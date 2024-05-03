package main

import (
	. "main/binary"
	. "main/request"
)

func main() {
	LogInit()
	RedisInit()
	DBInit()
	ServeInit()
	defer F.Close()
	//go func() {
	R.Run(":7234")
	//}()
}
