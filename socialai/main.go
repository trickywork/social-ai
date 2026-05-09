package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"socialai/backend"
	"socialai/handler"
	"socialai/util"
	"strings"
)

func main() {
	fmt.Println("started-service")

	config, err := util.LoadApplicationConfig("conf", "deploy.yml")
	if err != nil {
		panic(err)
	}

	if secret := os.Getenv("TOKEN_SECRET"); secret != "" {
		config.TokenConfig.Secret = secret
	}

	mode := strings.ToLower(os.Getenv("SOCIALAI_MODE"))
	if mode == "" {
		mode = "elastic"
	}

	if mode == "demo" || mode == "memory" {
		backend.InitMemoryStore()
		log.Println("SocialAI is running with the in-memory demo backend")
	} else {
		backend.InitElasticsearchBackend(config.ElasticsearchConfig)
		backend.InitGCSBackend(config.GCSConfig)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(http.ListenAndServe(":"+port, handler.InitRouter(config.TokenConfig)))
}
