package main

import (
	//"go_socket/server"
	"go_socket/database"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"reflect"
)

type User  struct{
	Title string `redis:"title"`
	Author string `redis:"author"`
	Body string `redis:"body"`
}

var p1,p2 User

func main()  {
	//server.SocketRun("10001")
	redis_ := database.GetRedis()

	defer redis_.Close()

	p1.Title = "gongyao book"
	p1.Author = "gyao"
	p1.Body = "gongyao graduate from yanshan university"

	_, err := redis_.Do("HMSET", redis.Args{}.Add("gongyao").AddFlat(p1)...)
	if err != nil {
		fmt.Printf("error %s", err.Error())
	}

	//str, err := redis.String(redis_.Do("GET", "bcd"))
	str,err := redis.Values(redis_.Do("HGETALL", "gongyao"))

	if err != nil {
		fmt.Printf("Error %s", err.Error())
	}

	redis.ScanStruct(str, &p2)

	fmt.Printf("redis: %v\n", p2)

	//遍历结构体
	t := reflect.TypeOf(p2)
	value := reflect.ValueOf(p2)
	for i := 0; i < value.NumField(); i++ {
		fmt.Printf("key:%s, value:%s\n", t.Field(i).Name, value.Field(i))
	}
}