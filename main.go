package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Hello, World!")

	// 通过启动参数指定端口，默认为 80
	port := flag.Int("port", 80, "Specify the port to use, e.g., 8080")
	flag.Parse()

	wd, _ := os.Getwd()

	fmt.Printf("Serving HTTP on http://0.0.0.0:%d", *port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), http.FileServer(http.Dir(wd))))
}
