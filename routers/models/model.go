package models

type RegisterReq struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}
type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
