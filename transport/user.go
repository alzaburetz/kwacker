package transport

type User struct {
	Name            string `json:"name"`
	Surname         string `json:"surname"`
	NickNameHandler string `json:"nick_name"`
	Posts           int    `json:"posts"`
}
