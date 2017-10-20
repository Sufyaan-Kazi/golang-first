package main

import (
    "fmt"
    "log"
    "net/http"
    "strconv"
)

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
    size := r.URL.Query()["size"]
    if sizeR, err := strconv.Atoi(size[0]); err == nil {
      for i := 0;  i<sizeR - 1; i++ {
        log.Println("log    ",i)
      }
    }
}

func handleRequests() {
    http.HandleFunc("/", homePage)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
    handleRequests()
}
