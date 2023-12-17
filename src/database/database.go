package database

import (
	"fmt"
	"log"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"os"
)

type DB struct {
	conn *gorm.DB
}

func Connect() *DB {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	time.Sleep(5 * time.Second)
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslMode := os.Getenv("SSL_MODE")

	//dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, dbname)
	fmt.Printf("DB_HOST: %s\n", host)
	fmt.Printf("DB_PORT: %s\n", port)
	fmt.Printf("DB_USER: %s\n", user)
	fmt.Printf("DB_PASSWORD: %s\n", password)
	fmt.Printf("DB_NAME: %s\n", dbname)

	// dsn := "root:mysql@tcp(172.16.238.10:3306)/db_users?parseTime=true"
	//dsn := "root:mysql@tcp(localhost:3306)/db_users?parseTime=true"
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", user, password, dbname, sslMode)
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error al conectar a la base de datos postgres: %v", err)
	}
	return &DB{conn}
}

// Función para obtener la conexión de GORM
func (db *DB) GetConn() *gorm.DB {
	return db.conn
}
