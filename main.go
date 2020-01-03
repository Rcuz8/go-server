package main

import (
    "net/http"
    "sync"
    "strconv"
    "fmt"
    "log"
)

var total_added int
var PORT = 8000;
var mutex = &sync.Mutex{}

func porter(portnum int) (string) {
	return ":" + strconv.Itoa(portnum)
}

func running(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "running..")
}

func check_param(r *http.Request, param string) (int) {
    res, contains := r.URL.Query()[param]
    
    if !contains || len(res[0]) < 1 {
        log.Println("Param is missing")
        return 0;
    }

    first := res[0]
    var ret int
    ret =  0
    ret, e := strconv.Atoi(first)
    if e != nil {
        fmt.Println("Error in conversion !")
        return 0
    }
    return ret;
}

func perform_add(r *http.Request) {
    var a = check_param(r, "a");
    var b = check_param(r, "b");

    mutex.Lock()
    total_added+=a;
    total_added+=b;
    mutex.Unlock()
    
}

func writeResult(w http.ResponseWriter, res string) {
    fmt.Fprintf(w, res);
}

func handle_add(w http.ResponseWriter, r *http.Request) {
    perform_add(r);
    writeResult(w, strconv.Itoa(total_added));
}

func main() {

    http.HandleFunc("/", running)
    http.HandleFunc("/addr", running)
    http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "Hi")
    })
    http.HandleFunc("/hii", func(w http.ResponseWriter, r *http.Request){
        handle_add(w,r);
    })

    http.HandleFunc("/add", handle_add)
    log.Fatal(http.ListenAndServe(porter(PORT), nil))

}