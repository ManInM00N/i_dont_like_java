package main

import (
	. "main/binary"
	. "main/request"
)

func main() {
	ConfigInit()
	LogInit()
	RedisInit()
	DBInit()
	ServeInit()
	defer F.Close()
	R.Run(":7234")
}
