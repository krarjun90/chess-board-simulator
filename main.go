package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	e "technogise/chess-board/board-engine"
	l "technogise/chess-board/board-layout"
	m "technogise/chess-board/messages"
	u "technogise/chess-board/utils"
)

func main() {
	fmt.Println(m.WelcomeMessage)
	fmt.Println(m.InputFormatMessage)
	reader := bufio.NewReader(os.Stdin)

	go simulate(reader)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	<-c

	fmt.Printf("\n%v", m.ExitMessage)
}

func simulate(reader *bufio.Reader) {
	layout := l.NewBoardLayout()
	for true {
		input, _ := reader.ReadString('\n')
		piece, cell, err := u.ParseInput(strings.TrimSpace(input))
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(e.SimulateMovement(layout, piece, *cell))
		}
	}
}
