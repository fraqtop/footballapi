package server

type competitionResponse struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

type errorResponse struct {
	Reason string `json:"reason"`
}
