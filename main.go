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
type ClockTime time.Duration
type GameState time.Duration

type ChessClock struct {
	whiteTime ClockTime
	blackTime ClockTime

	active uint64

	// State of game (over, paused, active, etc)
	gameState uint64
}

// toggleActive toggles between the time left for the two players. 
func (c *ChessClock) toggleActive() {
        switch c.active {
        case White:
            c.active = Black
        case Black:
            c.active = White
        }
}

// decrementCurrentTimer will decrement the current active players 
// time by one (in this case one second).
func (c *ChessClock) decrementCurrentTimer() {

    // NOTE: Add logic for timer running out (White Lose) and draws
    switch c.active {
    case White:
        if c.whiteTime > 0 {
            c.whiteTime--
        } else {
            log.Println("White is out of time..")
        }
    case Black:
        if c.blackTime > 0 {
            c.blackTime--
        } else {
            log.Println("Black is out of time")
        }
    }
}

// printTime displays time as seconds to its minutes equivalent 
func displayTimeMinutes(s ClockTime) (t time.Time) {
    mins := s / time.Minute
    secs := s % time.Minute

    fmt.Println("%d : %d ", mins, secs);
}

func main() {

	userInput := make(chan string)

	// Grab user keyboard input, (cbreak is 0) replace with termbox
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
        clk := &ChessClock{whiteTime: 100, blackTime: 100, active:White}
	tick := time.Tick(1 * time.Second)

Loopend:
	for {
		select {
		case in := <-userInput:
                    switch in {
                    // Spacebar toggles the timer 
                    case "\x00":
                        clk.toggleActive()
                    case "q":
                        break Loopend
                    }


		case <-tick:
                    // Stop the chess clock application. Someone won, lost, or quit..
                    if (clk.whiteTime == 0) || (clk.blackTime == 0) {
                        fmt.Println("Clock stopping...")
                    }
                    clk.decrementCurrentTimer()
                    // Display current timers
                    log.Printf("Clock object: %+v \n", clk)

		}
	}

	fmt.Println("Quitting...")
}
