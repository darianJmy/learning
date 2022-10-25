package main

import "github.com/darianJmy/learning/go-learning/casbin-practise/cmd/options"

func main() {
	o := options.NewOptions()

	if err := o.Complete(); err != nil {
		panic(err)
	}

}
