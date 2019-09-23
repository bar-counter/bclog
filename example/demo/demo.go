package main

import (
	"fmt"
	bclog "github.com/bar-counter/template"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	addr := ":31000"
	bclog.Config()
	bclog.Infof("run addr: %v", addr)

	err := r.Run(addr)
	if err != nil {
		fmt.Printf("run err %v\n", err)
		return
	}
}
