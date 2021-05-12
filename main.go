package main;
import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)
func main()  {
	serverCloseHandler()
	err:=http.ListenAndServe(":8080",nil)
	if err!=nil {
		panic("Error occured in server----->>"+err.Error())
	}
}

func serverCloseHandler()  {
	close:=make(chan os.Signal)
	signal.Notify(close, os.Interrupt, syscall.SIGTERM)
	<- close
	fmt.Println("<<<<<<<<<<-----------------------Server Closing----------------------->>>>>>>>")
	fmt.Println("- Good bye!")
	os.Exit(0)
}