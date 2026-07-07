package dto

type ArticleResponse struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	TypeName string `json:"typeName"`
}
