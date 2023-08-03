package request

type ReqCreateGenre struct {
	Title string `json:"title"`
	Color string `json:"color"`
}

type ReqUpdateGenre struct {
	ID    uint   `param:"id"`
	Title string `json:"title"`
	Color string `json:"color"`
}
