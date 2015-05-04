package main

import (
       //"fmt"
        "net/http"
)

func init() {
        //http.HandleFunc("/", handler)
 		
 		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
 			//fmt.Fprint(w,r.URL.Path[1:])
        	http.ServeFile(w, r, r.URL.Path[1:])
    	})        
}

func handler(w http.ResponseWriter, r *http.Request) {

        http.ServeFile(w, r, r.URL.Path[1:])
        // fmt.Fprint(w, "Welcome from the language of the coming decade - Go!")
		//http.Handle("/", http.FileServer(http.Dir("./static")))
        //http.ListenAndServe(":3000", nil)
}

func main() {
    // http.HandleFunc("/", handler)

    // http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    //     http.ServeFile(w, r, r.URL.Path[1:])
    // })
    // //panic(http.ListenAndServe(":8080", nil))
    
    //  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
 			// //fmt.Fprint(w,r.URL.Path[1:])
    //     http.ServeFile(w, r, r.URL.Path[1:])
    // })        
     // http.ListenAndServe(":8080", nil)
}

// func homeHandler(w http.ResponseWriter, r *http.Request) {
	
// 	//fmt.Fprint(w, "Welcome from the language of the coming decade - Go!")
//     fmt.Fprintf(w, "<html><head></head><body><h1>Welcome Home!</h1><a href=\"/static/img/test.gif\">Show Image!</a></body></html>")
// }