package main

import (
	"fmt"

	"github.com/osamikoyo/hrm-notify/internal/server"
)

func main() {
	server, err := server.New()
	if err != nil{
		fmt.Println(err)
	}

	server.Run()
}