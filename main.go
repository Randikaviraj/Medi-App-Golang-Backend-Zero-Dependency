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
	// safe server close function
	serverCloseHandler()

	portnum:=8080
	// check for the command line argument of port number
	if  num, err := strconv.Atoi(os.Args[1]); err == nil{
		portnum=num
	}
	log.Printf("Going to listen on port %d\n", portnum)

	// routes
	http.HandleFunc("/",func(w http.ResponseWriter,r *http.Request){
		w.Write([]byte("Hello"))
	})
	err := http.ListenAndServe(":"+strconv.Itoa(portnum), nil)
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
