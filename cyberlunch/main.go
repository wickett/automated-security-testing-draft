package main

import (
    "html/template"
    "io/ioutil"
    "log"
    "math/rand"
    "net/http"
    "path/filepath"
    "time"
)

func main() {
    // Parse the HTML template file
    welcomeTemplate, err := template.ParseFiles("welcome.html")
    if err != nil {
        log.Fatalf("Error parsing template file: %v", err)
    }

    // Handle the root path
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if err := welcomeTemplate.Execute(w, nil); err != nil {
            log.Printf("Error executing template: %v", err)
            http.Error(w, "Internal server error", http.StatusInternalServerError)
        }
    })

    // Handle the /fun path
    http.HandleFunc("/fun", func(w http.ResponseWriter, r *http.Request) {
        images, err := ioutil.ReadDir("images")
        if err != nil {
            http.Error(w, "Internal server error", http.StatusInternalServerError)
            return
        }

        if len(images) == 0 {
            http.Error(w, "No images found", http.StatusNotFound)
            return
        }

        rand.Seed(time.Now().UnixNano())
        randomImage := images[rand.Intn(len(images))]
        imagePath := filepath.Join("images", randomImage.Name())

        http.ServeFile(w, r, imagePath)
    })

    log.Println("Server started on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
