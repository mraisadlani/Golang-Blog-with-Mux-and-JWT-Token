package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

type Configuration struct {
	DBUSER string `json:"DB_USER"`
	DBPASS string `json:"DB_PASS"`
	DBHOST string `json:"DB_HOST"`
	DBPORT string `json:"DB_PORT"`
	DBNAME string `json:"DB_NAME"`
	LOGFILENAME string `json:"LOG_FILE_NAME"`
	LOGMAXSIZE string `json:"LOG_MAX_SIZE"`
	LOGMAXBACKUPS string `json:"LOG_MAX_BACKUPS"`
	LOGMAXAGE string `json:"LOG_MAX_AGE"`
	PORT string `json:"PORT"`
	READTIMEOUT time.Duration `json:"READ_TIME_OUT"`
	WRITETIMEOUT time.Duration `json:"WRITE_TIME_OUT"`
	SECRETKEY string `json:"SECRET_KEY"`
}

var (
	Config *Configuration
)

func SetupConfiguration() {
	if err := godotenv.Load();err != nil {
		log.Fatal("File .env tidak ditemukan")
		os.Exit(1)
	}

	Config = new(Configuration)

	read, _ := time.ParseDuration(os.Getenv("READ_TIME_OUT"))
	write, _ := time.ParseDuration(os.Getenv("WRITE_TIME_OUT"))

	Config.DBUSER = os.Getenv("DB_USER")
	Config.DBPASS = os.Getenv("DB_PASS")
	Config.DBHOST = os.Getenv("DB_HOST")
	Config.DBPORT = os.Getenv("DB_PORT")
	Config.DBNAME = os.Getenv("DB_NAME")
	Config.LOGFILENAME = os.Getenv("LOG_FILE_NAME")
	Config.LOGMAXSIZE = os.Getenv("LOG_MAX_SIZE")
	Config.LOGMAXBACKUPS = os.Getenv("LOG_MAX_BACKUPS")
	Config.LOGMAXAGE = os.Getenv("LOG_MAX_AGE")
	Config.PORT = os.Getenv("PORT")
	Config.READTIMEOUT = read
	Config.WRITETIMEOUT = write
	Config.SECRETKEY = os.Getenv("SECRET_KEY")
}