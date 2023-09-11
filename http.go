package main

import (
"encoding/json"
"fmt"
"net/http"
)

type Post struct {
UserID int `json:"userId"`
ID int `json:"id"`
Title string `json:"title"`
Body string `json:"body"`
}

func main() {
	url := "https://jsonplaceholder.typicode.com/posts/1"
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("HTTP GET request failed: %s\n", err)
		return
	}
	defer response.Body.Close()

	var post Post

	if err := json.NewDecoder(response.Body).Decode(&post); err != nil {
		fmt.Printf("Failed to parse JSON: %s\n", err)
		return
	}

	fmt.Printf("UserID: %d\nID: %d\nTitle: %s\nBody: %s\n", post.UserID,
	post.ID, post.Title, post.Body)
}
