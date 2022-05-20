package main

import (
	"betprophet1.com/wagers/internal/domains"
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

func main() {
	dsn := "wager:123456@tcp(127.0.0.1:3306)/wager?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
	}
	err = db.AutoMigrate(&domains.Wager{})
	if err != nil {
		fmt.Println(err.Error())
	}
	r := mux.NewRouter()
	srv := &http.Server{
		Addr:              "0.0.0.0:8080",
		Handler:           r,
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      15 * time.Second,
	}
	log.Fatalln(srv.ListenAndServe())
}