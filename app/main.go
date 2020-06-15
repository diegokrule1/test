package main

import (
	c "github.com/diegokrule1/test/app/controller"
)

func main() {
	//fmt.Println("hello world")
	r := c.SetupRouter()
	r.Run(":8085")
}
