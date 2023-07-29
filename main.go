package main

import (
	connectDB "bootcamp-api-hmsi/connectDB"
	"bootcamp-api-hmsi/query"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Error().Msg("Error loading .env file")
		os.Exit(1)
	}

	DB_HOST := os.Getenv("DB_HOST")
	DB_NAME := os.Getenv("DB_NAME")
	DB_PORT := os.Getenv("DB_PORT")
	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("POSTGRES_PASSWORD")
	DB_DRIVER := os.Getenv("DB_DRIVER")
	PORT := os.Getenv("PORT")

	log.Info().Msg("DB_HOST: " + DB_HOST)
	log.Info().Msg("DB_USER:" + DB_USER)
	log.Info().Msg("DB_PASSWORDD:" + DB_PASSWORD)
	log.Info().Msg("DB_NAME:" + DB_NAME)
	log.Info().Msg("PORT:" + DB_PORT)
	log.Info().Msg("DB_DRIVER:" + DB_DRIVER)
	log.Info().Msg("DB_DRIVER:" + PORT)

	db, errConn := connectDB.GetConnPostgres(DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME, DB_DRIVER)

	if errConn != nil {
		log.Error().Msg(errConn.Error())
		os.Exit(1)
	}

	fmt.Println("Successfully connected")
	// DB STRUCT INITIALIZE

	DB := query.DB{Conn: db}

	err := DB.Create(&query.Customers{
		Name:  "Tiar",
		Phone: "085770854766",
		Email: "muhamaadrahmatt@gmail.com",
		Age:   21,
	})

	if err != nil {
		log.Error().Msg(errConn.Error())
		os.Exit(1)
	}

	fmt.Println("Insert Data Berhasil")

	// read
	result, err := DB.Read()

	if err != nil {
		log.Error().Msg(errConn.Error())
		os.Exit(1)
	}

	fmt.Println(result)

	// update
	err = DB.Update(&query.Customers{
		Id:    1,
		Name:  "Bagas",
		Phone: "085781544293",
		Email: "bagas@gmail.com",
		Age:   23,
	})

	if err != nil {
		log.Error().Msg(errConn.Error())
		os.Exit(1)
	}

	fmt.Println("Update Berhasil")

	// delete
	err = DB.Delete(2)

	if err != nil {
		log.Error().Msg(errConn.Error())
		os.Exit(1)
	}

	fmt.Println("Update Berhasil")

}
