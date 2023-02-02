package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/eiannone/keyboard"
	"github.com/pterm/pterm"
)

const (
	// Players, identified by color.
	White = iota
	Black

	GameActive
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

	// Active players time.
	active uint64
	// State of game (over, paused, active, etc)
	gameState uint64
}

func NewChessClock(gameTime ClockTime) *ChessClock {

	// Create new clock with corresponding times. White activtes first as
	// in chess white always makes the first move.
	clk := &ChessClock{whiteTime: gameTime, blackTime: gameTime, active: White}
	return clk
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

// Take ClockTime type (seconds), return mins and secs values.
func secToMins(t ClockTime) (mins, secs uint) {
	mins = uint(t) / 60
	secs = uint(t) % 60

	return mins, secs
}

func main() {

	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stdin, "Usage: chessclk <TIME_SECONDS>")
		os.Exit(1)
	}

	var gameTime ClockTime = 0

	// If user supplied times, usse for both openents.
	// TODO: Refactor with proper flags interface.
	if os.Args[1] != "" {
		initTime := os.Args[1]
		m, err := time.ParseDuration((initTime))
		if err != nil {
			log.Fatal("error parsing time argument")
		}
		// Convert to proper type
		gameTime = ClockTime(m.Round(time.Second).Seconds())
	} else {
		log.Println("No argument passed, will use default time value of 15 mins")
		gameTime = ClockTime(15 * 60)
	}

	// Make sure our clock isn't set to a useless zero.
	if gameTime == 0 {
		pterm.Warning.Println("[+] No time set on game clock.")
		os.Exit(1)
	}

	// Grab user keyboard input, (cbreak is 0) replace with termbox
	userInput := make(chan string)
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

	// Create our chess clock.
	clk := NewChessClock(gameTime)
	tick := time.Tick(1 * time.Second)

	fmt.Println("")

	// Start area for our TUI output
	area, _ := pterm.DefaultArea.Start()
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
				fmt.Println("\nClock stopping...")
				break Loopend
			}

			clk.decrementCurrentTimer()
			wMins, wSecs := secToMins(clk.whiteTime)
			bMins, bSecs := secToMins(clk.blackTime)

			// Display current clock times for each player.
			// Aside: Formatting needs refactor. Four spaces seperate the respective times for each player.
			clksStr, _ := pterm.DefaultBigText.
				WithLetters(pterm.NewLettersFromString(fmt.Sprintf("%02d:%02d    ", wMins, wSecs) +
					fmt.Sprintf("%02d:%02d", bMins, bSecs))).Srender()
			clksStr = pterm.DefaultCenter.Sprint(clksStr)

			area.Update(clksStr)
		}
	}
	area.Stop()

	fmt.Println("\nQuitting...")
}
