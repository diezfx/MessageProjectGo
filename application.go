package main

import (
	"fmt"
	"net/http"
	"time"
	"google.golang.org/appengine"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func greeting(w http.ResponseWriter, r *http.Request) {

	// utc time?
	var hour = time.Now().Hour()

	print(hour)

	var greeting = ""
	if (hour >= 0 && hour < 6) {
		greeting = "Good night"
	} else if (hour >= 6 && hour < 10) {
		greeting = "Good morning"
	} else if (hour >= 10 && hour < 14) {
		greeting = "Good mid-day"
	} else if (hour >= 14 && hour < 18) {
		greeting = "Good afternoon"
	} else {
		greeting = "Good evening"
	}

	// slashes are allowed
	var message = greeting + " " + r.URL.Path[len("/welcome/"):] + "."

	fmt.Fprint(w, message)

}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/welcome/", greeting)

	appengine.Main() // Starts the server to receive requests
}
