package backend

import (
	"fmt"
	"sort"
	"strings"
	"sync"

	"socialai/model"
)

var MemoryStore *InMemoryStore

type InMemoryStore struct {
	mu    sync.RWMutex
	users map[string]model.User
	posts map[string]model.Post
}

func InitMemoryStore() {
	MemoryStore = &InMemoryStore{
		users: make(map[string]model.User),
		posts: make(map[string]model.Post),
	}
}

func (s *InMemoryStore) AddUser(user *model.User) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.users[user.Username]; exists {
		return false
	}
	s.users[user.Username] = *user
	return true
}

func (s *InMemoryStore) CheckUser(username, password string) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	user, exists := s.users[username]
	return exists && user.Password == password
}

func (s *InMemoryStore) SavePost(post *model.Post) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.posts[post.Id] = *post
}

func (s *InMemoryStore) SearchPostsByUser(user string) []model.Post {
	s.mu.RLock()
	defer s.mu.RUnlock()

	posts := make([]model.Post, 0)
	for _, post := range s.posts {
		if post.User == user {
			posts = append(posts, post)
		}
	}
	sortPosts(posts)
	return posts
}

func (s *InMemoryStore) SearchPostsByKeywords(keywords string) []model.Post {
	s.mu.RLock()
	defer s.mu.RUnlock()

	needle := strings.ToLower(strings.TrimSpace(keywords))
	posts := make([]model.Post, 0)
	for _, post := range s.posts {
		if needle == "" || strings.Contains(strings.ToLower(post.Message), needle) {
			posts = append(posts, post)
		}
	}
	sortPosts(posts)
	return posts
}

func (s *InMemoryStore) DeletePost(id string, user string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	post, exists := s.posts[id]
	if !exists {
		return fmt.Errorf("no such post exists")
	}
	if post.User != user {
		return fmt.Errorf("post belongs to another user")
	}
	delete(s.posts, id)
	return nil
}

func sortPosts(posts []model.Post) {
	sort.Slice(posts, func(i, j int) bool {
		return posts[i].Id > posts[j].Id
	})
}
