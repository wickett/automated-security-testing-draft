package main

import (
    "io/ioutil"
    "log"
    "math/rand"
    "net/http"
    "path/filepath"
    "time"
)

func main() {
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
