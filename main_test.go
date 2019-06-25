package main

import "testing"
import "log"

func TestToggleActive(t *testing.T) {
    clk := &ChessClock{whiteTime: 100, blackTime:100, active:White}
    log.Printf("currently... %+v", clk)
    clk.toggleActive()
    if clk.active != Black {
        t.Error(
            "Current active value", clk.active,
            "Expected", Black)
    }

    log.Printf("After toggle: %+v", clk)
    // toggle back
    clk.toggleActive()
    if clk.active != White {
        t.Error("clock.active is ", clk.active, "Should be set to White")
    }
    log.Printf("Final toggle is: %+v", clk)
}
