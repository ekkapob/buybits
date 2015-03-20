package model

type Status struct {
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
}

type LoginResponse struct {
	Status
	User `json:"user,omitempty"`
}

type PostsResponse struct {
	Status
	Posts []Post `json:"posts"`
}
