package output

type competitionResponse struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}
