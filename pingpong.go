package main

import "fmt"
import "time"

type Ball struct {
	hits int
}

func Player(name string, table chan *Ball) {
	for {
		ball := <-table
		ball.hits++
		fmt.Println(name, ball.hits)
		time.Sleep(100 * time.Millisecond)
		table <- ball
	}
}

func main() {
	table := make(chan *Ball)
	go Player("Ping", table)
	go Player("Pong", table)
	table <- new(Ball) //game on, toss ball
	time.Sleep(1 * time.Second)
	<-table //game over, grab the ball
	panic("Show me the stacks")
}
