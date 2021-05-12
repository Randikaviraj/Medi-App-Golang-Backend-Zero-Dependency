package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"log"
	"strconv"
)

func main() {
	serverCloseHandler()
	portnum:=8080
	if  num, err := strconv.Atoi(os.Args[1]); err == nil{
		portnum=num
	}
	log.Printf("Going to listen on port %d\n", portnum)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("Error occured in server----->>" + err.Error())
	}
}

func serverCloseHandler() {
	close := make(chan os.Signal)
	signal.Notify(close, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-close
		fmt.Println("<<<<<<<<<<-----------------------Server Closing----------------------->>>>>>>>")
		fmt.Println("- Good bye!")
		os.Exit(0)
	}()

}
