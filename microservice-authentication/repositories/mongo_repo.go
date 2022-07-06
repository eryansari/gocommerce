package repositories

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gocommerce/constanta"
)

type MongoRepository struct {
	DB       *mongo.Database
	authRepo *UserAuthRepo
}

func BuildMongoRepository(db *mongo.Database) *MongoRepository {
	logger := logrus.StandardLogger()
	return &MongoRepository{
		DB:       db,
		authRepo: NewUserAuthRepo(db, logger),
	}
}

func (m *MongoRepository) UserRepository() UserRepository {
	return m.authRepo
}

// Open Connection MongoDB
func InitMongoDB(logger *logrus.Logger) *mongo.Database {
	logger.Println("trying connect to database")

	client, err := mongo.NewClient(options.Client().ApplyURI(viper.GetString(constanta.ConfigMongoURI)))
	if err != nil {
		logger.Fatalf("%v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	db := client.Database(viper.GetString(constanta.ConfigMongoDatabase))
	err = client.Connect(ctx)
	if err != nil {
		logger.Fatalf("%v", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		logger.Fatalf("%v", err)
	}

	logger.Println("successfully connect to database")

	return db
}
