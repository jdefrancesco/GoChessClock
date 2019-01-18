package main

import (
	"fmt"
	"log"
	"time"

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

	// Grab user keyboard input
	go func() { userInput <- os.Stdin.Read(make([]byte, 1)) }()

	for {
		select {
		case in := <-userInput:
			log.Println("Received user input: %c", in)
		default:
			log.Println("...")
			time.Sleep(1 * time.Second)
		}
	}
}
