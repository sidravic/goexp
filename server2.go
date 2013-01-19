package main

import "fmt"
import "net/http"
import "time"


func HomeHandler(w http.ResponseWriter, r *http.Request){
	msg := fmt.Sprintf("%v %v", time.Now(), " Fuck You Again.")
	fmt.Printf("Request is %v", r.URL)
	fmt.Println("")
	fmt.Fprint(w, msg)
}

func main(){
	http.HandleFunc("/home", HomeHandler)
	http.ListenAndServe("localhost:9900", nil)
}


