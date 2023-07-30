package models

type RequestAPIRequest struct {
	Id    int64    `json:"id"`
	UserID  int64 `json:"user_id"`
	Access string `json:"access"`
	Hidden  bool `json:"hidden"`
	ServiceFamily  string `json:"service_family"`
	Data  string `json:"data"`
	Token string `json:"token"`
}