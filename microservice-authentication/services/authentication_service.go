package services

import (
	"context"
	"encoding/hex"
	"fmt"
	"regexp"
	"strings"
	"time"

	jwt "github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gocommerce/constanta"
	"gocommerce/microservice-authentication/pkg/model"
	"gocommerce/microservice-authentication/pkg/pbuff"
	"gocommerce/microservice-authentication/repositories"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// AuthSvc Service
type AuthSvc struct {
	urepo repositories.UserRepository
	log   *logrus.Logger
}

// NewAuthSvc
func NewAuthSvc(urepo repositories.UserRepository, log *logrus.Logger) *AuthSvc {
	return &AuthSvc{
		urepo: urepo,
		log:   log,
	}
}

// Register Request
func (s *AuthSvc) Register(ctx context.Context, udto *pbuff.RegisterDto) (*pbuff.UserAuthResDto, error) {
	var (
		reqUname = strings.ToLower(udto.GetUsername())
		reqEmail = strings.ToLower(udto.GetEmail())
	)

	if len(reqUname) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "username required")
	}

	if len(reqEmail) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "email required")
	}

	if len(udto.GetPassword()) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "password required")
	}

	if len(udto.GetPhone()) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "phone required")
	}

	hasPwd, err := hashPassword(udto.GetPassword())
	if err != nil {
		s.log.Errorln(err)
		return nil, status.Errorf(codes.DataLoss, "password is not hashing")
	}

	// check username exist
	checkUser, _ := s.urepo.FindByUsername(ctx, reqUname)
	if checkUser.Username == reqUname {
		s.log.Errorf("username exist: %s, err: %v", reqUname, err)
		return nil, status.Errorf(codes.NotFound, "username is exist: %s", reqUname)
	}

	// check email exist
	checkEmail, _ := s.urepo.FindByEmail(ctx, reqEmail)
	if checkEmail.Email == reqEmail {
		s.log.Errorf("email exist: %s, err: %v", reqEmail, err)
		return nil, status.Errorf(codes.NotFound, "email is exist: %s", reqEmail)
	}

	// check phone exist
	checkPhone, _ := s.urepo.FindByPhone(ctx, udto.GetPhone())
	if checkPhone.Phone == udto.GetPhone() {
		s.log.Errorf("phone exist: %s, err: %v", udto.GetPhone(), err)
		return nil, status.Errorf(codes.NotFound, "phone is exist: %s", udto.GetPhone())
	}

	u := &model.UserAuth{
		Username: reqUname,
		Email:    reqEmail,
		Phone:    udto.GetPhone(),
		Password: hasPwd,
		Status:   model.UserStatusActive,
	}

	uauth, err := s.urepo.Create(ctx, u)
	fmt.Println(err)
	if err != nil {
		s.log.Errorln(err)
		return nil, status.Errorf(codes.Internal, "Create user failed")
	}
	user := &pbuff.UserAuthResDto{
		Username: uauth.Username,
		Email:    uauth.Email,
		Phone:    uauth.Phone,
		Status:   "active",
	}

	s.log.Infof("new user (%s) created", user.Username)
	return user, nil
}

// Login Request
func (s *AuthSvc) Login(ctx context.Context, login *pbuff.LoginDto) (*pbuff.AccessTokenDto, error) {
	username := strings.ToLower(login.GetUsername())
	deviceID := login.GetDeviceId()
	s.log.Debugf("Login(Username : %s, DeviceId : %s)", username, deviceID)

	if len(username) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "username required")
	}
	if !validUsername(username) {
		return nil, status.Errorf(codes.InvalidArgument, "invalid username")
	}
	if len(deviceID) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "deviceID required")
	}

	// find username
	user, err := s.urepo.FindByUsername(ctx, username)
	if err != nil {
		s.log.Errorf("username : %s not found, err: %v", username, err)
		return nil, status.Errorf(codes.NotFound, "username: %s not found", username)
	}
	// password is not match
	passOk := checkPasswordHash(login.GetPassword(), user.Password)
	if !passOk {
		s.log.Errorln("invalid password")
		return nil, status.Errorf(codes.InvalidArgument, "invalid password")
	}
	// user is not active
	if user.Status != model.UserStatusActive {
		s.log.Errorln("user is not active")
		return nil, status.Errorf(codes.FailedPrecondition, "user is not active")
	}

	// find existing token
	curToken, err := s.urepo.FindTokenByUserAndDevice(ctx, username, deviceID)
	if err == nil {

		expIn := curToken.ExpiredDate.Unix() - time.Now().Unix()
		s.log.Debugf("token will expired in %v seconds", expIn)
		if expIn >= 420 { //420 secs = 7 minutes
			accesToken := &pbuff.AccessTokenDto{
				AccessToken: curToken.Token,
				ExpiresIn:   expIn,
				Scope:       "all",
				TokenType:   "Bearer",
				Username:    user.Username,
				Email:       user.Email,
				Phone:       user.Phone,
			}
			return accesToken, nil
		} else {
			td, _ := s.urepo.DeleteTokenByUserAndDevice(ctx, username, deviceID)
			if td {
				s.log.Debugf("token removed for %s on %s", username, deviceID)
			}
		}

	}

	var issuer = viper.GetString(constanta.ConfigJWTIssuer)
	var issTime = time.Now().Unix()
	var expTime = issTime + viper.GetInt64(constanta.ConfigJWTExpired)
	var signkey = []byte(viper.GetString(constanta.ConfigJWTSign))

	uid, _ := hex.DecodeString(user.ID.Hex())
	claims := &JwtClaims{
		StandardClaims: jwt.StandardClaims{
			Id:        uuid.New().String(),
			Audience:  username,
			Issuer:    issuer,
			IssuedAt:  issTime,
			ExpiresAt: expTime,
		},
		UserID:   string(uid),
		DeviceID: deviceID,
		Username: user.Username,
		Email:    user.Email,
		Phone:    user.Phone,
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	signedToken, err := token.SignedString(signkey)
	if err != nil {
		s.log.Errorf("failed to sign token: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to sign token: %v", err)
	}

	logDate := time.Now()
	expDate := time.Unix(expTime, 0)
	userToken := &model.UserToken{
		UserID:      user.ID,
		Username:    user.Username,
		DeviceID:    deviceID,
		LoginDate:   &logDate,
		ExpiredDate: &expDate,
		Token:       signedToken,
	}

	utoken, err := s.urepo.CreateToken(ctx, userToken)
	if err != nil {
		s.log.Errorf("unable to create token: %v", err)
		return nil, status.Errorf(codes.Internal, "unable to create token: %v", err)
	}

	accesToken := &pbuff.AccessTokenDto{
		AccessToken: utoken.Token,
		ExpiresIn:   utoken.ExpiredDate.Unix() - time.Now().Unix(),
		Scope:       "all",
		TokenType:   "Bearer",
		Username:    user.Username,
		Email:       user.Email,
		Phone:       user.Phone,
	}

	return accesToken, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// username regex
func validUsername(username string) bool {
	matched, _ := regexp.MatchString(`^\w*$`, username)
	return matched
}
