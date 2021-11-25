package databases

import (
	"database/sql"
	"log"
	"os"
	// _ "github.com/lib/pq"
)

const (
	host     = "ec2-3-217-68-126.compute-1.amazonaws.com"
	port     = 5432
	user     = "jqbgkdyymcxpll"
	password = "e3236ebb5eed5b81e8fe54c3a2faeec73932e8e343219b22080ac114c6539abf"
	dbname   = "der3dq4oku2f96"
)

var DB *sql.DB

func ConnectDatabase() {
	// psqlInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=require TimeZone=Asia/Shanghai", host, user, password, dbname, port)
	dsn := "host=ec2-3-217-68-126.compute-1.amazonaws.com user=jqbgkdyymcxpll password=e3236ebb5eed5b81e8fe54c3a2faeec73932e8e343219b22080ac114c6539abf dbname=der3dq4oku2f96 port=5432 sslmode=require TimeZone=Asia/Shanghai"
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db, err := sql.Open("postgres", dsn )
	// db.AutoMigrate(&models.User{})
	if err == nil {
		log.Printf("Failed to connect ")
		log.Println(err.Error())
		os.Exit(100)
	}
	log.Printf("Connected to db")

  	DB = db
}



// Create User Table
func CreateUserTable(db *sql.DB) error {
	// opts := &orm.CreateTableOptions{
	// 	IfNotExists: true,
	// }
	// createError := db.CreateTable(&models.User{}, opts)
	// if createError != nil {
	// 	log.Printf("Error while creating user table, Reason: %v\n", createError)
	// 	return createError
	// }
	// log.Printf("User table created")
	return nil
}