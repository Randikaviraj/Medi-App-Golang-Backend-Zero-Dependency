package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/joho/godotenv"

	"github.com/randika/Medi-App-Golang-Backend/database"
	"github.com/randika/Medi-App-Golang-Backend/routes"
)

func main() {
	// loading enviroment variables from .env file
	if err := godotenv.Load(".env"); err!=nil {
		log.Panicf(err.Error())
	}

	
	// creating database connection
	database.CreateDbConnection(os.Getenv("DataSourceName"))

	portnum:=8080

	// check for the command line argument of port number
	if  len(os.Args) >1{
		num, err := strconv.Atoi(os.Args[1])
		if err!=nil{
			portnum=num
		}
	}
	log.Printf("Going to listen on port %d\n", portnum)

	// routes
	routes.Setup()

	// safe server close function
	serverCloseHandler()
	

	err := http.ListenAndServe(":"+strconv.Itoa(portnum), nil)
	if err != nil {
		log.Panicf(err.Error())
	}
}

func serverCloseHandler() {
	close := make(chan os.Signal)
	signal.Notify(close, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-close
		log.Printf("<<<<<<<<<<-----------------------Server Closing----------------------->>>>>>>>")
		fmt.Println("- Good bye!")
		os.Exit(0)
	}()

}
