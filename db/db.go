package db

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	var flag_env = flag.String("GO_ENV", "", "開発環境フラグ")
	flag.Parse()
	if *flag_env == "dev" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalln(err)
		}
	}

	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PW"), os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_DB"))

	/* url = fmt.Sprintf("postgres://%s:%s@%s:%s/%s", "shouta",
	"shouta", "localhost",
	"5434", "shouta") */
	//なぜかos.Gotenvが働かない

	db, err := gorm.Open(postgres.Open(url), &gorm.Config{}) //第二引数の空の構造体で初期化

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected")
	return db
}

func CloseDB(db *gorm.DB) {
	sqlDB, _ := db.DB()
	if err := sqlDB.Close(); err != nil {
		log.Fatalln(err)
	}
}
