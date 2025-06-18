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

	"1": One,
	"2": Two,
	"3": Three,
	"4": Four,
	"5": Five,
	"6": Six,
	"7": Seven,
	"8": Eight,
	"9": Nine,
}

type Strategy interface {
	Choose(topCard *Card, hand []*Card) *Card
}

type OffensiveStrategy struct{}

func (OffensiveStrategy) Choose(topCard *Card, hand []*Card) *Card {
	for _, i := range hand {
		fmt.Println(i.Color)
	}
	if c, match := colorMatch(topCard, hand); match {
		return c
	}
	return &Card{
		Color: Red,
		Value: Seven,
	}
}

func colorMatch(card *Card, cards []*Card) (*Card, bool) {
	for _, c := range cards {
		if card.Color == c.Color {
			return c, true
		}
	}
	return nil, false
}

type DefensiveStrategy struct{}

func (DefensiveStrategy) Choose(topCard *Card, hand []*Card) *Card {
	// todo:
	// 1. Play duplicated card
	// 2. Play wild card
	// wild cards last
	// play a duplicated card
	return &Card{
		Color: Blue,
		Value: Eight,
	}
}

func parse(cards ...string) []*Card {
	var buf []*Card
	for _, card := range cards {
		var v Value
		var c Color
		if isDigit(card[0]) {
			v = strToValue[string(card[0])]
			c = strToColor[card[1:]]
		} else if match, key := isSpecial(card); match {
			remaining := strings.Replace(card, key, "", 1)
			remaining = strings.TrimSpace(remaining)
			v = strToValue[key]
			c = strToColor[strings.ToLower(remaining)]
			fmt.Println(c)
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

// todo: rename this to something better
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

	offStrat := &OffensiveStrategy{}
	defStrat := &DefensiveStrategy{}

	cards := parse(os.Args[1:]...)

	oc := offStrat.Choose(cards[0], cards[1:])
	dc := defStrat.Choose(cards[0], cards[1:])

	fmt.Printf("Offesnively -> Play '%s %s'\n", oc.Color, oc.Value)
	fmt.Printf("Defensively -> Play '%s %s'\n", dc.Color, dc.Value)
}

func help() {
	fmt.Println("You must provide the played card and the available cards. i.e.,")
	fmt.Println("\tum 7red 1blue 7green wild")
}
