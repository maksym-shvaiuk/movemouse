package main

import (
	"fmt"
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {
	for {
		randomY := int(time.Now().Unix()%300) // Random Y between 0 and 300
		randomX := int(time.Now().Unix()%300) // Random X between 0 and 300
		randomSeconds1 := time.Duration(1 + time.Now().Unix()%10) // Random seconds between 1 and 10
		randomSeconds2 := time.Duration(1 + time.Now().Unix()%10) // Random seconds between 1 and 10

		robotgo.MoveSmoothRelative(randomX, randomY, 1.0, 2.0)
		fmt.Printf("Moved mouse %d pixels down, %d pixels right, now sleeping for %d seconds\n", randomX, randomY, randomSeconds1)
		time.Sleep(randomSeconds1 * time.Second)
		robotgo.MoveSmoothRelative(-randomX, -randomY, 1.0, 2.0)
		fmt.Printf("Moved mouse %d pixels up, %d pixels left, now sleeping for %d seconds\n", randomX, randomY, randomSeconds2)
		time.Sleep(randomSeconds2 * time.Second)
	}
}