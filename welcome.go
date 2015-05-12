package main

import (
       //"fmt"
        "net/http"
        "os"
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

func index2(w http.ResponseWriter, r *http.Request) {

    //fmt.Fprint(w, "Welcome from the language of the coming decade - Go!",r.URL.Path[1:] + ".html")
    http.ServeFile(w, r, "index_2.html")

}

// func (h *handler) ServeHTTP(w http.ResponseWriter,r *http.Request) {

//     w.Header().Set("Content-Type", "application/json")
//     w.WriteHeader(http.StatusOK)
//     enc := json.NewEncoder(w)
//     if err := enc.Encode(&MyResponse{}); nil != err {
//         fmt.Fprintf(w, `{"error":"%s"}`, err)
//     }
// }

func main() {

            http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
            //fmt.Fprint(w,r.URL.Path[1:])
            http.ServeFile(w, r, r.URL.Path[1:])
            }) 
            http.HandleFunc("/index2", index2)
   
    //http.ListenAndServe(":8080", nil)
    err := http.ListenAndServe(":" + os.Getenv("PORT"), nil)
    if err != nil {
      panic(err)
    }    
}

