package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
)

type Article struct {
    Title string `json:"Title"`
    Description string `json:"description"`
    Body string `json:"body"`
}

type Articles []Article

// GET /
func topPage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome to the Top")
    fmt.Println("[Log] Get the top")
}

// GET /articles
func allArticles(w http.ResponseWriter, r *http.Request) {
    articles := Articles{}
    for i := 0; i < 10; i++ {
        title := "Hello_%d"
        description := "World_%d"
        articles = append(
            articles,
            Article{Title: fmt.Sprintf(title, i), Description: fmt.Sprintf(description, i), Body: "Article Body"})
    }
    fmt.Println("[Log] Get returnAllArticles")
    json.NewEncoder(w).Encode(articles)
}

func handleRequests() {
    http.HandleFunc("/", topPage)
    http.HandleFunc("/articles", allArticles)
    log.Fatal(http.ListenAndServe(":9000", nil))
}

func main() {
    handleRequests()
}
