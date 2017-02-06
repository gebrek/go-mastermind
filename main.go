package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

// dictFileLength could be derived from manually counting the lines
// once, but ehh whatever
const dictFileLength = 45403

type game struct {
	code    string
	guesses int
	running bool
}

func (g *game) guess(s string) {
	g.guesses++
	if strings.Compare(g.code, s) == 0 {
		g.win()
	} else {
		g.hint(s)
	}
}

func (g *game) win() {
	g.running = false
	fmt.Println("You made", g.guesses, "guesses.")
	fmt.Println("You win!")
}

func (g *game) hint(s string) {
	var h string
	for i, c := range g.code {
		if i >= len(s) {
			h += "*"
		} else if s[i] == byte(c) {
			h += string(c)
		} else {
			h += "*"
		}
	}
	fmt.Println(h)
}

func newGame() *game {
	return &game{
		code:    randomWord(),
		guesses: 0,
		running: true,
	}
}

func randomWord() (word string) {
	file, err := os.Open("linuxwords")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	i := 0
	goal := int(rand.Float64() * dictFileLength)

	for scanner.Scan() {
		i++
		if i == goal {
			word = strings.ToLower(scanner.Text())
		}
	}
	return
}

func main() {
	// a seed based on time is different every run
	rand.Seed(time.Now().UTC().UnixNano())

	g := newGame()
	reader := bufio.NewReader(os.Stdin)
	log.Println(g.code)
	for g.running {
		fmt.Print("Enter your guess: ")
		input, err := reader.ReadString('\n')
		input = input[:len(input)-1] // remove newline char
		if err != nil {
			log.Fatal(err)
		}
		g.guess(input)
	}
}
