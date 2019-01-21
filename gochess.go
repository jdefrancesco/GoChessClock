package main

import (
	"fmt"
	"log"
	_ "os"
	"time"

	"github.com/eiannone/keyboard"
	_ "github.com/jroimartin/gocui"
)

const (
	// Players, identified by color.
	WHITE = iota
	BLACK

	// Game states.
	GAMEPAUSED
	GAMERESET
	GAMEOVER
)

type ClockTime uint64
type GameState uint64

// ChessClock is a structue containing elements
type ChessClock struct {
	WhiteTime ClockTime
	BlackTime ClockTime

	// Specify which clock is currently active (ticking down)
	ActiveClock uint64
	// State of game (over, paused, active, etc)
	State uint64
}

func main() {

	log.Println("Started chess clock main...")

	userInput := make(chan string)

	// Grab user keyboard input (cbreak is 0)
	go func() {
		err := keyboard.Open()
		if err != nil {
			panic(err)
		}
		defer keyboard.Close()

		for {
			c, _, err := keyboard.GetKey()
			if err != nil {
				panic(err)
			}
			userInput <- string(c)
		}

	}()

loop:
	for {
		select {
		case in := <-userInput:
			if in == "q" {
				break loop
			}
			fmt.Printf("Received user input: %s\n", in)
		default:
			log.Println("...")
			time.Sleep(1 * time.Second)
		}
	}

	fmt.Println("Quitting...")
}
