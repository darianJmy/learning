package types

type MenuUri struct {
	ID int64 `uri:"id" binding:"required"`
}

type Menu struct {
	Id         int64  `json:"id"`
	Name       string `json:"username"`
	UserNameCn string `json:"username_cn"`
	Nick       string `json:"nickname"`
	Password   string `json:"password"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
}
