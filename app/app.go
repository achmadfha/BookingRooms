package app

import (
	"BookingRoom/config"
	"BookingRoom/model/dto"
	"BookingRoom/router"
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
	"os"
	"strconv"
	"time"
)

func initEnv() (dto.ConfigData, error) {
	// load env data
	var configData dto.ConfigData
	if err := godotenv.Load(".env"); err != nil {
		return configData, err
	}

	if port := os.Getenv("PORT"); port != "" {
		configData.AppConfig.Port = port
	}

	dbHost := os.Getenv("DB_PORT")
	dbPort := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_DATABASE")
	dbMaxIdle := os.Getenv("MAX_IDLE")
	dbMaxConn := os.Getenv("MAX_CONN")
	dbMaxLifeTime := os.Getenv("MAX_LIFE_TIME")
	dbLogMode := os.Getenv("LOG_MODE")
	saltInt := os.Getenv("SALT")
	secretToken := os.Getenv("SECRET_TOKEN")
	tokenExpired := os.Getenv("TOKEN_EXPIRED")

	if dbHost == "" || dbPort == "" || dbUser == "" || dbPassword == "" || dbName == "" || dbMaxIdle == "" || dbMaxConn == "" || dbMaxLifeTime == "" || dbLogMode == "" || saltInt == "" || secretToken == "" || tokenExpired == "" {
		return configData, errors.New("DB Config not set")
	}

	var err error
	configData.DbConfig.MaxConn, err = strconv.Atoi(dbMaxConn)
	if err != nil {
		return configData, err
	}

	configData.DbConfig.MaxIdle, err = strconv.Atoi(dbMaxIdle)
	if err != nil {
		return configData, err
	}
	configData.DbConfig.Host = dbHost
	configData.DbConfig.DbPort = dbPort
	configData.DbConfig.User = dbUser
	configData.DbConfig.Pass = dbPassword
	configData.DbConfig.Database = dbName
	configData.DbConfig.MaxLifeTime = dbMaxLifeTime
	configData.DbConfig.LogMode, err = strconv.Atoi(dbLogMode)
	if err != nil {
		return configData, err
	}

	return configData, nil
}

func RunService() {
	//add log
	zerolog.TimeFieldFormat = "02-01-2006 15:04:05"
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	//load env config
	configData, err := initEnv()
	if err != nil {
		log.Error().Msg(err.Error())
		return
	}

	conn, err := config.ConnectToDB(configData, log.Logger)
	if err != nil {
		log.Error().Msg("service.ConnectToDB.Err: " + err.Error())
		return
	}

	duration, err := time.ParseDuration(configData.DbConfig.MaxLifeTime)
	if err != nil {
		log.Error().Msg("service.Duration.Err: " + err.Error())
		return
	}

	conn.SetConnMaxLifetime(duration)
	conn.SetMaxIdleConns(configData.DbConfig.MaxIdle)
	conn.SetMaxOpenConns(configData.DbConfig.MaxConn)

	defer func() {
		errClose := conn.Close()
		if errClose != nil {
			log.Error().Msg(errClose.Error())
		}
	}()

	// set up time zone
	time.Local = time.FixedZone("Asia/Jakarta", 7*60*60)
	r := gin.New()
	r.Use(cors.New(cors.Config{
		AllowAllOrigins: false,
		AllowOrigins:    []string{"*"},
		AllowMethods:    []string{"POST", "GET", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{
			"Origin", "Content-type", "Authorization",
		},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           120 * time.Second,
	}))

	log.Logger = log.With().Caller().Logger()

	r.Use(logger.SetLogger(
		logger.WithLogger(func(_ *gin.Context, l zerolog.Logger) zerolog.Logger {
			return l.Output(os.Stdout).With().Logger()
		}),
	))

	//gin recovery for handle panic
	r.Use(gin.Recovery())
	initializeDomainModule(r, conn)

	version := "0.0.1"
	log.Info().Msg(fmt.Sprintf("Service running versions %s", version))
	addr := flag.String("Port: ", ":"+configData.AppConfig.Port, "Address to listen and serve")
	err = r.Run(*addr)
	if err != nil {
		log.Error().Msg(err.Error())
	}

}

func initializeDomainModule(r *gin.Engine, db *sql.DB) {
	apiGroup := r.Group("/api")
	v1Group := apiGroup.Group("/v1")

	router.InitRouter(v1Group, db)
}
