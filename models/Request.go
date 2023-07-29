package models

type Request struct {
	_meta string `table:"requests"`

	Id    int    `column:"id"`
	CreatedAt string `column:"created_at"`
	UId  int `column:"u_id"`
	Hash  string `column:"hash"`
	Data  string `column:"data"`
	ServiceFamily  string `column:"service_family"`
}