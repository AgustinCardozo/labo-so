package models

type User struct {
	Id   int    `json:"id"`
	User string `json:"usuario"`
	//Pass string `json:"password"`
	Mail string `json:"mail"`
	Name string `json:"nombre"`
	//Role     string    `json:"rol"`
	//CreateAt time.Time `json:"createAt"`
}
