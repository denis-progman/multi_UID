package models

type Service struct {
	Id    int64    `column:"id"`
	CreatedAt string `column:"created_at"`
	UpdatedAt string `column:"updated_at"`
	UserId  int64 `column:"user_id"`
	Data  string `column:"data"`
}