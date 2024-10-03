package request

type CreateNewAuthorRequest struct {
	Name string `json:"name"`
	Bio  string `json:"bio"`
}
