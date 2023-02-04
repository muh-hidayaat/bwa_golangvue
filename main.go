package main

import (
	"bwa_golangvue/user"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:root@tcp(127.0.0.1:8889)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("koneksi datbase berhasil")

	userRepository := user.NewRepository(db)
	user := user.User{
		Name: "Test",
	}
	userRepository.Save(user)

}
