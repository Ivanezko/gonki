package main

/*
https://www.upwork.com/ab/proposals/1242538704290312193
https://www.upwork.com/messages/rooms/room_9303e63a3a997453c17d8b94235d92f9

Create a small backend:

Login to a website with anti-captcha API when necessary.
Check a string on the website at regular intervals (set in config). Send email (mailgun API) or SMS (Clockwork SMS API) based on the string.
Use a .env file for remote API credentials, interval setting, destination SMS number/s and email address/es.
Create a readme file.
*/

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"main/internal/config"
	"main/internal/myhttp"
	"os"
	"sync"
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime | log.Lmicroseconds)
	if _, err := os.Stat("config.yml"); err == nil {
		err := cleanenv.ReadConfig("config.yml", &config.Server)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Print(err)
		log.Print("config.yml not found. hope everything is in env?")
	}
}

func main() {

	var wg sync.WaitGroup
	wg.Add(1)
	go myhttp.Server(&wg)
	wg.Wait()
	log.Println("app end")
}
