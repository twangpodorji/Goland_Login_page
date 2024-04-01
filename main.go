package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {

    fmt.Printf("Sever listening at port 1024\n")

    // ? Create a home page handler
    http.HandleFunc("/", homeHandler)
    // ? Create a hello handler
    http.HandleFunc("/hello", HelloHandler)
    // ? Create a login handler
    http.HandleFunc("/login", loginPageHandler)
    // ? Create a authentication handler
    http.HandleFunc("/auth", AuthHandler)

    // ? Start the server, if any error, console log it
    if err := http.ListenAndServe(":1024", nil); err != nil {
        log.Fatal(err)
    }
}

// // ? Handler for the hello route
// func helloHandler(w http.ResponseWriter, r *http.Request) {
//     fmt.Fprintf(w, "Hello Page!")
// }

// // ? Handler for the home route
// func homePageHandler(w http.ResponseWriter, r *http.Request) {
//     fmt.Fprintf(w, "Home Page!")
// }

// ? Handler for the login route
// func loginPageHandler(w http.ResponseWriter, r *http.Request) {
//     fmt.Fprintf(w, "login page")
// }

// ? Handler for the login route
func loginPageHandler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "./login.html")
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "./about.html")
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "./home.html")
}

func AuthHandler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "./home.html")
    // http.ServeFile(w, r, "./auth.html")
}