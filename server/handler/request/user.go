package request

type ReqUpdateUser struct {
	ID    uint    `param:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type ReqLoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ReqUpdateUserGenre struct {
	GenreID []uint `json:"genre_ids"`
}

type ReqCreateUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
