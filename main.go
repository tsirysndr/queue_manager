package main

import (
	"log"
	"time"

	"github.com/labstack/gommon/color"
	natsd "github.com/nats-io/nats-server/server"
	"github.com/nats-io/nats.go"
)

func main() {
	const banner = `
   ____                          __  __                                   
  / __ \                        |  \/  |                                  
 | |  | |_   _  ___ _   _  ___  | \  / | __ _ _ __   __ _  __ _  ___ _ __ 
 | |  | | | | |/ _ \ | | |/ _ \ | |\/| |/ _' | '_ \ / _' |/ _' |/ _ \ '__|
 | |__| | |_| |  __/ |_| |  __/ | |  | | (_| | | | | (_| | (_| |  __/ |   
  \___\_\\__,_|\___|\__,_|\___| |_|  |_|\__,_|_| |_|\__,_|\__, |\___|_|   
                                                           __/ |          
                                                          |___/           
	`

	nopts := &natsd.Options{}
	nopts.HTTPPort = 8222
	nopts.Port = 4223

	color.Println(color.Magenta(banner))

	color.Printf(color.Cyan("🚀 NATS Server ready at: nats://localhost:%d ⭐️ \n"), nopts.Port)

	// Create the NATS Server
	ns := natsd.New(nopts)

	go ns.Start()

	nc, err := nats.Connect("localhost:4223")

	if err != nil {
		log.Fatal(err)
	}

	defer nc.Close()

	// Subscribe
	sub, err := nc.SubscribeSync("articles")
	if err != nil {
		log.Fatal(err)
	}

	for {
		// Wait for a message
		msg, err := sub.NextMsg(10 * time.Second)
		if err != nil {
			log.Println(err)
		}

		if msg != nil {
			// Use the response
			res, _ := nc.Request("save_article", msg.Data, 10*time.Second)
			if res != nil {
				log.Printf("Reply: %s %s\n", msg.Data, res.Data)
			}
		}

	}

}
