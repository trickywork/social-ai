package service

import (
	"encoding/base64"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"reflect"
	"strings"

	"socialai/backend"
	"socialai/constants"
	"socialai/model"

	"github.com/olivere/elastic/v7"
)

func SearchPostsByUser(user string) ([]model.Post, error) {
	if backend.MemoryStore != nil {
		return backend.MemoryStore.SearchPostsByUser(user), nil
	}

	query := elastic.NewTermQuery("user", user)
	searchResult, err := backend.ESBackend.ReadFromES(query, constants.POST_INDEX)
	if err != nil {
		return nil, err
	}
	return getPostFromSearchResult(searchResult), nil
}

func SearchPostsByKeywords(keywords string) ([]model.Post, error) {
	if backend.MemoryStore != nil {
		return backend.MemoryStore.SearchPostsByKeywords(keywords), nil
	}

	query := elastic.NewMatchQuery("message", keywords)
	query.Operator("AND")
	if keywords == "" {
		query.ZeroTermsQuery("all")
	}
	searchResult, err := backend.ESBackend.ReadFromES(query, constants.POST_INDEX)
	if err != nil {
		return nil, err
	}
	return getPostFromSearchResult(searchResult), nil
}

func getPostFromSearchResult(searchResult *elastic.SearchResult) []model.Post {
	var ptype model.Post
	var posts []model.Post

	for _, item := range searchResult.Each(reflect.TypeOf(ptype)) {
		p := item.(model.Post)
		posts = append(posts, p)
	}
	return posts
}

func SavePost(post *model.Post, file multipart.File) error {
	if backend.MemoryStore != nil {
		data, err := io.ReadAll(file)
		if err != nil {
			return err
		}
		contentType := http.DetectContentType(data)
		if strings.HasPrefix(contentType, "application/octet-stream") && post.Type == "video" {
			contentType = "video/mp4"
		}
		post.Url = fmt.Sprintf("data:%s;base64,%s", contentType, base64.StdEncoding.EncodeToString(data))
		backend.MemoryStore.SavePost(post)
		return nil
	}

	medialink, err := backend.GCSBackend.SaveToGCS(file, post.Id)
	if err != nil {
		return err
	}
	post.Url = medialink

	err = backend.ESBackend.SaveToES(post, constants.POST_INDEX, post.Id)
	if err != nil {
		// Consider rolling back the GCS upload by deleting the object
		_ = backend.GCSBackend.DeleteFromGCS(post.Id) // Example of a rollback
		return err
	}

	return nil
}

func DeletePost(id string, user string) error {
	if backend.MemoryStore != nil {
		return backend.MemoryStore.DeletePost(id, user)
	}

	query := elastic.NewBoolQuery()
	query.Must(elastic.NewTermQuery("id", id))
	query.Must(elastic.NewTermQuery("user", user))

	return backend.ESBackend.DeleteFromES(query, constants.POST_INDEX)
}
