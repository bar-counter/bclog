package main

import (
	"fmt"
	"github.com/bar-counter/bclog"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	addr := ":31000"
	log.Config()
	log.Infof("run addr: %v", addr)

	err := r.Run(addr)
	if err != nil {
		fmt.Printf("run err %v\n", err)
		return
	}
}
