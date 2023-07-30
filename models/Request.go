package models

type Request struct {
	_meta string `table:"requests"`

	Id    int64    `column:"id" json:"id"`
	CreatedAt string `column:"created_at" json:"created_at"`
	UId  int64 `column:"u_id" json:"u_id"`
	Hash  string `column:"hash" json:"hash"`
	ServiceFamily  string `column:"service_family" json:"service_family"`
	Data  string `column:"data" json:"data"`
}