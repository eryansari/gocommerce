package services

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"strings"

	jwt "github.com/golang-jwt/jwt"
	log "github.com/sirupsen/logrus"
	"gocommerce/constanta"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
)

var insecuredPaths = [...]string{
	"/pb.AuthApi/Login",
	"/pb.AuthApi/Register",
	"/pb.ProductApi/GetProductByID",
	"/pb.ProductApi/ListProduct",
}

func insecuredPath(path string) bool {
	for _, p := range insecuredPaths {
		if p == path {
			return true
		}
	}
	return false
}

func AuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	userAgent, _ := md["user-agent"]
	peer, _ := peer.FromContext(ctx)
	token, tokenFound := md["authorization"]
	var tkn string
	if len(token) <= 0 {
		tkn = "n/a"
	} else {
		tkn = token[0]
		tkn = fmt.Sprintf("%s...%s", tkn[0:17], tkn[len(tkn)-10:])
	}
	insecured := insecuredPath(info.FullMethod)
	var secured = "Y"
	if insecured {
		secured = "N"
	}
	log.Debugf("rpc secured: %s, method: %s, token: %v, ipAddress: %s, userAgent: %v",
		secured, info.FullMethod, tkn, peer.Addr.String(), userAgent)
	// allowed method without authorization
	if insecured {
		return handler(ctx, req)
	}
	// bearer token not found
	if !tokenFound {
		return nil, status.Errorf(codes.Unauthenticated, "Authorization not found")
	}

	tokenStr := strings.Replace(strings.Join(token, ""), "Bearer ", "", -1)
	var signkey = []byte(viper.GetString(constanta.ConfigJWTSign))
	jwtToken, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return signkey, nil
	})
	if jwtToken == nil {
		return nil, status.Errorf(codes.Unauthenticated, "Invalid Token")
	}
	if claims, ok := jwtToken.Claims.(jwt.MapClaims); ok && jwtToken.Valid {
		ctx = context.WithValue(ctx, constanta.CtxKeyUserID, claims[constanta.CtxKeyUserID])
		ctx = context.WithValue(ctx, constanta.CtxKeyDeviceID, claims[constanta.CtxKeyDeviceID])
		ctx = context.WithValue(ctx, constanta.CtxKeyUsername, claims[constanta.CtxKeyUsername])
		ctx = context.WithValue(ctx, constanta.CtxKeyUserEmail, claims[constanta.CtxKeyUserEmail])
		return handler(ctx, req)
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return nil, status.Errorf(codes.Unauthenticated, "Invalid Token")
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			return nil, status.Errorf(codes.Unauthenticated, "Token Expired")
		} else {
			return nil, status.Errorf(codes.Unauthenticated, "Invalid Token")
		}
	} else {
		return nil, status.Errorf(codes.Unauthenticated, "Invalid Token")
	}

	//jwtToken, _ := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
	//	if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
	//		return nil, fmt.Errorf("Signing method invalid")
	//	} else if method != jwt.SigningMethodHS256 {
	//		return nil, fmt.Errorf("Signing method invalid")
	//	}
	//
	//	return jwt.SigningMethodHS256, nil
	//})

	//claims, tokenFound := jwtToken.Claims.(jwt.MapClaims)
}
