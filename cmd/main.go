package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joseluis8906/cleanexample/internal/order"
	mymongo "github.com/joseluis8906/cleanexample/pkg/mongo"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx := context.Background()

	// check environment
	env := os.Getenv("ENVIRONMENT")
	if env == "" {
		env = "dev"
	}

	// load configs by environment
	viper.SetConfigName(env)
	viper.SetConfigType("json")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	// conect to mongodb
	cleandbHost := viper.GetString("mongo.databases.cleandb.host")
	cleandbPort := viper.GetInt("mongo.databases.cleandb.port")
	cleandbUser := viper.GetString("mongo.databases.cleandb.user")
	cleandbPasswd := viper.GetString("mongo.databases.cleandb.password")
	cleandbDB := viper.GetString("mongo.databases.cleandb.db")
	cleandbURI := fmt.Sprintf("mongodb://%s:%s@%s:%d/%s", cleandbUser, cleandbPasswd, cleandbHost, cleandbPort, cleandbDB)
	cleandb, err := mongo.NewClient(options.Client().ApplyURI(cleandbURI))
	if err != nil {
		panic(err)
	}

	err = cleandb.Connect(ctx)
	if err != nil {
		panic(err)
	}

	defer func() {
		err := cleandb.Disconnect(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}()

	err = cleandb.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	// driver to infra db
	mongoDriver := mymongo.MongoDriver{
		Client: cleandb.Database(cleandbDB).Collection("orders"),
	}

	// order
	orderRepo := order.Repository{
		Saver:  &mongoDriver,
		Finder: &mongoDriver,
	}

	odCreator := order.Creator{
		Saver: orderRepo,
	}
	err = odCreator.Exec(ctx, "Jhon Doe", "Cll 7 # 6 - 09", "Katy Heinz", 10000)
	if err != nil {
		log.Print(err)
	}

	orders, err := orderRepo.Find(ctx, bson.M{"customer": "Jhon Doe"})
	if err != nil {
		log.Print(err)
	}

	fmt.Printf("result from db: %+v\n", orders)

	// run
	fmt.Println("running")
	os.Exit(0)
}
