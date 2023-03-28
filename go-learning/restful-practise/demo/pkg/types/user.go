package types

type User struct {
	Id     string `gorm:"Id"`
	Name   string `gorm:"name"`
	Age    int    `gorm:"age"`
	School string `gorm:"school"`
}
