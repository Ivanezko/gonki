package main

import (
	"log"
	"main/internal/app"
	"os"
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime | log.Lmicroseconds)
	log.Println("init...")
	app.Init()
	log.Println("init done")
}

func main() {
	log.Println("=====app start, welcome")
	app.Main()
	log.Println("=====app end, bye")
}
