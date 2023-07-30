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
	// Untuk memuat file .env sebagai konfigurasi
	err := godotenv.Load("config/.env")

	if err != nil {
		log.Error().Msg(err.Error())
		os.Exit(1)
	}
	DB_HOST := os.Getenv("DB_HOST")
	DB_NAME := os.Getenv("DB_NAME")
	DB_PORT := os.Getenv("DB_PORT")
	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_DRIVER := os.Getenv("DB_DRIVER")
	PORT := os.Getenv("PORT")

	log.Info().Msg(DB_HOST)
	log.Info().Msg(DB_NAME)
	log.Info().Msg(DB_PORT)
	log.Info().Msg(DB_USER)
	log.Info().Msg(DB_PASSWORD)
	log.Info().Msg(DB_DRIVER)
	log.Info().Msg(PORT)

	db, errConn := connectDB.GetConnPostgres(DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME, DB_DRIVER)
	if errConn != nil {
		log.Error().Msg(errConn.Error())
		os.Exit(1)
	}
	fmt.Println("Successfully connected!")

	// DB struct initialize
	DB := query.DB{Conn: db}

	// Create customer
	err = DB.Create(&query.Customers{
		Name:  "Muhammad Rahmat",
		Phone: "085770854766",
		Email: "rahmat@mail.com",
		Age:   21,
	})
	fmt.Println("err", err)
	if err != nil {
		log.Error().Msg(errConn.Error())
		os.Exit(1)
	}
	fmt.Println("Insert Data Berhasil")

	// Read customer
	result, err := DB.Read()
	if err != nil {
		log.Error().Msg(errConn.Error())
		os.Exit(1)
	}
	fmt.Println(result)

	// Update customer
	err = DB.Update(&query.Customers{
		Id:    1,
		Name:  "Budi",
		Phone: "012345564",
		Email: "budi@example.com",
		Age:   31,
	})
	if err != nil {
		log.Error().Msg(errConn.Error())
		os.Exit(1)
	}
	fmt.Println("Update Berhasil")

	// Delete customer
	err = DB.Delete(1)
	if err != nil {
		log.Error().Msg(errConn.Error())
		os.Exit(1)
	}
	fmt.Println("Delete Berhasil")

}
