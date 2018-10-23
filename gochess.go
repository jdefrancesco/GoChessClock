package main

import (
    "fmt"
    "time"
    "github.com/jroimartin/gocui"
)



func main() {
    log.Println("Started chess clock...")

    // Init gocui 
    g, err := gocui.NewGui(gocui.OutputNormal)
    if err != nil {
	log.Panicln(err)
    }
    defer g.Close()

    active := make(chan bool)
    // Two go funcs each with a select checking active flag


}

