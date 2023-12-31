package models

type Uid struct {
	_meta string `table:"uids"`

	Id    int64    `column:"id"`
	CreatedAt string `column:"created_at"`
	UserId  int64 `column:"user_id"`
	Access  string `column:"access"`
	Hidden  bool `column:"hidden"`
}