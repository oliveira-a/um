package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

type Color string
type Value string

type Card struct {
	Color Color
	Value Value
}

func NewCard(value, color string) *Card {
	return &Card{
		Color: strToColor[color],
		Value: strToValue[value],
	}
}

const (
	Red    Color = "Red"
	Blue   Color = "Blue"
	Green  Color = "Green"
	Yellow Color = "Yellow"
	Wild   Color = "Wild"

	One   Value = "1"
	Two   Value = "2"
	Three Value = "3"
	Four  Value = "4"
	Five  Value = "5"
	Six   Value = "6"
	Seven Value = "7"
	Eight Value = "8"
	Nine  Value = "9"

	DrawTwo  Value = "Draw Two"
	Skip     Value = "Skip"
	Reverse  Value = "Reverse"
	WildCard Value = "Wild"
)

var strToColor = map[string]Color{
	"red":    Red,
	"green":  Green,
	"yellow": Yellow,
	"blue":   Blue,
	"wild":   Wild,
}

var strToValue = map[string]Value{
	"drawTwo":  DrawTwo,
	"skip":     Skip,
	"reverse":  Reverse,
	"wildCard": WildCard,
}

type Strategy interface {
	Choose(hand []Card, topCard Card) *Card
}

func parse(cards ...string) []*Card {
	buf := make([]*Card, len(cards))
	for _, card := range cards {
		var v Value
		var c Color
		if isDigit(card[0]) {
			v = strToValue[string(card[0])]
			c = strToColor[card[1:]]
		}
		if match, key := isSpecial(card); match {
			remaining := strings.Replace(card, key, "", 1)
			remaining = strings.TrimSpace(remaining)
			v = strToValue[key]
			c = strToColor[remaining]
		}
		buf = append(
			buf,
			&Card{
				Color: c,
				Value: v,
			},
		)
	}

	return buf
}

func isSpecial(s string) (bool, string) {
	for key := range strToValue {
		if strings.Contains(s, key) {
			return true, key
		}
	}
	return false, ""
}

func isDigit(c byte) bool {
	return unicode.IsDigit(rune(c))
}

func main() {
	if len(os.Args) < 3 {
		help()
		os.Exit(0)
	}

	cards := parse(os.Args[1])
	for _, c := range cards {
		fmt.Println(c)
	}
}

func help() {
	fmt.Println("You must provide the played card and the available cards. i.e.,")
	fmt.Println("\tum 7red 1blue 7green wild")
}
