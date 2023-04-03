package main

import (
	"fmt"
	flag "github.com/spf13/pflag"
)

var strFlag = flag.String("s", "", "Description")

type User struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	School string `json:"school"`
}

func main() {

	println(*strFlag)

	var user User
	FlagSet(&user)

	fmt.Println(user)
}

func FlagSet(user *User) {
	flag.StringVar(&user.Name, "name", user.Name, "please input your name")
	flag.IntVar(&user.Age, "age", user.Age, "please input your age")
	flag.StringVar(&user.School, "school", user.School, "please input your school")
	flag.Parse()
}
