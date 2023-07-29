package main

type Service struct {
	Id    int    `column:"id"`
	CreatedAt string `column:"created_at"`
	UpdatedAt string `column:"updated_at"`
	UserId  int `column:"user_id"`
	Data  string `column:"data"`
}