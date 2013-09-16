package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// signals := make(chan os.Signal)
	// signal.Notify(signals)

	// go func() {
	// 	signal := <-signals
	// 	fmt.Printf("Got signal %v", signal)
	// 	panic(signal)
	// }()

	portNumber := os.Getenv("PORT")
	if portNumber == "" {
		portNumber = "5000"
	}
	listenOn := fmt.Sprintf(":%v", portNumber)
	log.Fatal(http.ListenAndServe(listenOn, http.FileServer(http.Dir("public"))))
}
