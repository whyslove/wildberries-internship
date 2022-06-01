package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	l0 "github.com/whyslove/wildberries-l0"
	"github.com/whyslove/wildberries-l0/core/handler"
	"github.com/whyslove/wildberries-l0/core/repository"
)

// @title Wildberries L0
// @version 1.0
// @description API for Getting purchases

// @host localhost:8000
// @BasePath /api

func main() {
	logrus.SetLevel(logrus.DebugLevel)

	//Configs
	if err := initConfig(); err != nil {
		logrus.Fatalf("error in config init %s", err.Error())
	}
	if err := godotenv.Load("./configs/.env"); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	//Postgres configuration
	db, err := repository.NewPostgresDb(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		logrus.Fatalf("error in connectinf to DB %s", err.Error())
	}

	repository := repository.NewRepository(db)
	handlers := handler.NewHandler(repository)

	p, err := repository.ReturnAllPurchasesDTO()
	if err != nil {
		logrus.Fatalf("error while returning all purchases %s", err)
	}
	repository.RecoverAllPurchasesCache(p)

	logrus.Infof("End loading purchases into cache, loaded %s", p)

	//Stan configuration
	sc, err := stan.Connect(viper.GetString("nats.clusterID"), viper.GetString("nats.clientID"))
	if err != nil {
		logrus.Fatalf("error in connecting to nats %s", err.Error())
	}

	_, err = sc.Subscribe(viper.GetString("nats.channelName"), handlers.ListenChannelPurchases,
		stan.DeliverAllAvailable(), stan.DurableName(viper.GetString("nats.durableName")))
	if err != nil {
		logrus.Fatalf("error in subscribing to channel in nats %s", err.Error())
	}

	//Http server configuration
	srv := new(l0.Server)
	go func() {
		if err := srv.Run(viper.GetString("server.port"), handlers.InitRoutes()); err != nil {
			if err != http.ErrServerClosed {
				logrus.Fatalf("Error in running server, error is %s", err.Error())
			}

		}
	}()

	logrus.Info("done with setting up")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Info("now shutting down")
	err = srv.Shutdown(context.Background())
	if err != nil {
		logrus.Errorf("Error while shutdown server")
	}

	err = db.Close()
	if err != nil {
		logrus.Errorf("Error while disconnect db")
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
