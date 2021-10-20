package main

import "fmt"
import "net/http"
import "os"
import "log"

func main() {
    fmt.Println("Hello, World!")

    wd, _ := os.Getwd()
    port := 80

    fmt.Printf("Serving HTTP on http://0.0.0.0:%d", port)

    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), http.FileServer(http.Dir(wd))))
}
