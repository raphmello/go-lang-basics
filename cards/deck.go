package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}

func (d deck) saveToFile(fileName string) error {
	file, error := os.Create(fileName)
	file.Write([]byte(d.toString()))
	defer file.Close()
	return error
}

func newDeckFromFile(fileName string) deck {
	byte, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	cards := strings.Split(string(byte), ",")
	d := deck(cards)
	return d
}

func (d deck) shuffle() {
	randSource := rand.NewSource(int64(time.Now().Nanosecond()))
	randomGenerator := rand.New(randSource)

	for i := range d {
		newPosition := randomGenerator.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
