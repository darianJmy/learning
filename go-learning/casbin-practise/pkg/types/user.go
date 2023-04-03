package types

type User struct {
	UserName   string `json:"username"`
	UserNameCn string `json:"username_cn"`
	NickName   string `json:"nickname"`
	Password   string `json:"password"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
}
