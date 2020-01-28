package main

import "github.com/guygubaby/dytt-server/server"

func main() {
	r := server.InitServer()
	r.Run(":5000")
}
