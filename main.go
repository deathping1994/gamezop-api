package main
 
import (
    // "encoding/json"
    "log"
    "net/http"
    "flag"
)


func main() {
    hostPtr := flag.String("host", "0.0.0.0", "Host (Default is 0.0.0.0)")
    portPtr := flag.String("port", "9090", "Port (Default is 9090)")
    flag.Parse()
    *hostPtr = *hostPtr + ":" + *portPtr
    log.Println("Listening on ",*hostPtr)
    router := NewRouter()
    InitDb()
    log.Fatal(http.ListenAndServe(*hostPtr, router))
}
