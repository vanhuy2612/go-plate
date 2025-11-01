package service

import (
	"encoding/json"
	"gorm.io/gorm"
	"io"
	"net/http"
	"root/src/models"
	"sync"
)

type IPostService interface {
	Index()
}

type PostService struct {
	DB *gorm.DB
}

func fetch(wg *sync.WaitGroup, ch chan<- models.Post) {
	defer wg.Done()
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var post models.Post
	if err := json.Unmarshal(body, &post); err != nil {
		panic(err)
	}
	ch <- post
}

func Index() []models.Post {
	batch := 100
	var wg sync.WaitGroup
	ch := make(chan models.Post, batch)
	for i := 0; i < batch; i++ {
		wg.Add(1)
		go fetch(&wg, ch)
	}
	wg.Wait()
	close(ch)

	// Collect all posts
	var posts []models.Post
	for post := range ch {
		posts = append(posts, post)
	}
	return posts
}
