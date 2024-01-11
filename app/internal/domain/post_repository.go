package domain

type PostRepository interface {
	GetAllPosts() ([]Post, error)
}
