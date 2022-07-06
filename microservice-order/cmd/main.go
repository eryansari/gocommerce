package main

import (
	"fmt"
	"log"
	"net"
	"path"
	"runtime"
	"time"

	"github.com/sirupsen/logrus"

	logRotate "github.com/lestrrat/go-file-rotatelogs"

	"github.com/spf13/viper"
	"gocommerce/constanta"
	"gocommerce/microservice-order/repositories"
	"gocommerce/microservice-order/services"
	"google.golang.org/grpc"
)

func init() {
	configName := "config"
	viper.SetConfigName(configName)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("conf")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("error load config: %v", err))
	}

	// set log
	logpath := viper.GetString(constanta.ConfigLogPath)
	log.Printf("get path path to %s", logpath)
	writer, err := logRotate.New(
		fmt.Sprintf("./logs/%s.log", "%y%m%d"),
		logRotate.WithRotationTime(time.Hour*24),
	)
	if err != nil {
		panic(fmt.Errorf("error initialize log file: %v", err))
	}

	logrus.SetOutput(writer)
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetReportCaller(true)
	logfmtr := &logrus.TextFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			filename := path.Base(f.File)
			return "", fmt.Sprintf("%s %s:%d", time.Now().Format(time.RFC3339), filename, f.Line)
		},
	}
	logrus.SetFormatter(logfmtr)
}

func main() {
	logger := logrus.StandardLogger()

	initmongo := repositories.InitMongoDB(logger)
	mongorepo := repositories.BuildMongoRepository(initmongo)

	server := grpc.NewServer(
		grpc.UnaryInterceptor(services.AuthInterceptor),
	)
	services.BuildGRPCService(server, mongorepo)

	port := viper.GetString(constanta.ConfigGrpcHost) + ":" + viper.GetString(constanta.ConfigGrpcPort)
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("error listening to %s: %v", port, err)
	}

	panic(server.Serve(listener))
}
