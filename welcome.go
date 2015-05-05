package main

import (
       //"fmt"
        "net/http"
        //"os"
)

func init() {
        //http.HandleFunc("/", handler)
        
        // http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        //     //fmt.Fprint(w,r.URL.Path[1:])
        //     http.ServeFile(w, r, r.URL.Path[1:])
        // })        
}

func handler(w http.ResponseWriter, r *http.Request) {

        http.ServeFile(w, r, r.URL.Path[1:])
        // fmt.Fprint(w, "Welcome from the language of the coming decade - Go!")
        //http.Handle("/", http.FileServer(http.Dir("./static")))
        //http.ListenAndServe(":8080", nil)
}

func main() {

            http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
            //fmt.Fprint(w,r.URL.Path[1:])
            http.ServeFile(w, r, r.URL.Path[1:])
        }) 
   
    http.ListenAndServe(":8080", nil)
}

