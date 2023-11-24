package external_service

import (
	"context"
	"fmt"
	"github.com/shivansh98/kahawatein/internal/network"
	"github.com/spf13/viper"
	"log"
)

func SearchUnsplash(ctx context.Context, query string) map[string]interface{} {
	url := "https://api.unsplash.com/search/photos"
	headers := map[string]string{
		"content-type":  "application/json",
		"Authorization": fmt.Sprintf("Client-ID %s", viper.GetString("UNSPLASH_KEY")),
	}
	resp, err := network.Get(url, headers, map[string]string{"query": query})
	if err != nil {
		log.Println("got an error in unsplash call", err)
	}
	return resp
}
