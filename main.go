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
	White = iota
	Black

	GamePaused
	GameReset
	GameOver
)

// Hold clock time in seconds.
type ClockTime uint64
type GameState uint64


type ChessClock struct {
	whiteTime ClockTime
	blackTime ClockTime

	active uint64

	// State of game (over, paused, active, etc)
	gameState uint64
}

// toggleActive toggles between the time left for the two players. 
// Use this function to toggle between the two players so the correct
// players clock decrements.
func (c *ChessClock) toggleActive() {
        switch c.active {
        case White:
            c.active = Black
        case Black:
            c.active = White
        }
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

        // test struct.. 
        // clk := &ChessClock{whiteTime: 100, blackTime: 100, active:White}
	tick := time.Tick(1 * time.Second)

loopend:
	for {
		select {
		case in := <-userInput:

			if in == "q" {
				break loopend
			}
			fmt.Printf("Received user input: %s\n", in)

		case <-tick:
			fmt.Println("tickkk")

		}
	}

	fmt.Println("Quitting...")
}
