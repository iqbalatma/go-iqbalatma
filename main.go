package main

import (
	"fmt"
	"iqbalatma/go-iqbalatma/app/enum"
	"iqbalatma/go-iqbalatma/cmd"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Running server")
	} else {
		switch os.Args[1] {
		case "server":
		case string(enum.CommandSeeder):
			cmd.RunningSeeder()
		default:
			fmt.Println("❌ Unknown command:", os.Args[1])
		}
	}
}
