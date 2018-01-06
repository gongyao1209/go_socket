package server

import (
	"fmt"
	"os"
	"net"
	"go_socket/database"

	"github.com/garyburd/redigo/redis"
)

var redis_ redis.Conn

func checkError(err error)  {
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		os.Exit(1)
	}
}

func recvConnMsg(conn net.Conn)  {
	buf := make([]byte, 50)

	//redis_ := database.GetRedis()

	defer conn.Close()

	for {
		n, err := conn.Read(buf)

		if err != nil {
			fmt.Println("conn cloesed")
			return
		}

		str := string(buf[0:n])

		_, err1 := redis_.Do("HSET", "go_socket", str, 1);

		if err1 != nil {
			fmt.Println("redis error")
			return
		}

		fmt.Println("recv msg: ", string(buf[0:n]))
	}
}

func init ()  {
	redis_ = database.GetRedis()
}

func SocketRun(port string)  {
	host := "127.0.0.1:" + port
	listen_sock, err := net.Listen("tcp", host)

	checkError(err)

	defer listen_sock.Close()
	defer redis_.Close()

	for {
		new_conn, err := listen_sock.Accept()

		if err != nil {
			continue
		}

		go recvConnMsg(new_conn)
	}
}