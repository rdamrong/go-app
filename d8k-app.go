package main

import (
    "fmt"
    "net/http"
    "os"
    "time"
)

var t0 = time.Now()
func main() {
    fmt.Println("Starting r20221002, ready to service !!")
    http.HandleFunc("/", HelloServer)
    http.ListenAndServe(":4000", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
        t1 := time.Now()
        servTime := t1.Sub(t0).Seconds()
        fmt.Printf("%s,%0.2f\n",r.URL.Path[1:],servTime)
        hname, err := os.Hostname()
        if err != nil {
              panic(err)
        }
        fmt.Fprintf(w, "[Hello, %s][Process ID: %d][service time = %0.2f sec][hostname: %s]\n", r.URL.Path[1:],os.Getpid(),servTime,hname)
}
