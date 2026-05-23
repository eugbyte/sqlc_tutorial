package model

type AuthorRequest struct {
	Name string  `json:"name"`
	Bio  *string `json:"bio"`
}

type AuthorResponse struct {
	ID   int64   `json:"id"`
	Name string  `json:"name"`
	Bio  *string `json:"bio"`
}
