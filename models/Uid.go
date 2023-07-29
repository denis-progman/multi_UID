package models

type Uid struct {
	_meta string `table:"uids"`

	Id    int    `column:"id"`
	CreatedAt string `column:"created_at"`
	UserId  int `column:"user_id"`
	Access  string `column:"access"`
	Hidden  bool `column:"hidden"`
	Token  string `column:"token"`
}