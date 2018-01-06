package main

import (
	//"go_socket/server"
	"go_socket/database"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main()  {
	//server.SocketRun("10001")
	redis_ := database.GetRedis()

	defer redis_.Close()

	str, err := redis.String(redis_.Do("GET", "bcd"))

	if err != nil {
		fmt.Println("Error %s", err.Error())
	}


	fmt.Printf("redis: %s", str)
}