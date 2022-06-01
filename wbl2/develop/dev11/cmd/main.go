package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"whyslove/wbl2/dev11/core/handler"
	"whyslove/wbl2/dev11/core/repository"
	"whyslove/wbl2/dev11/core/server"
	"whyslove/wbl2/dev11/core/usecase"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

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
	usecase := usecase.NewUseCase(repository)
	handlers := handler.NewHandler(usecase)

	srv := new(server.Server)
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
