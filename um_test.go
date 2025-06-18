package main

import "testing"

func TestParsesSingleCard(t *testing.T) {
	card := "7red"

	rslt := parse(card)[0]

	if rslt.Color != Red {
		t.Errorf("Color does not match expected 'red', got: %s", rslt.Color)
	}
	if rslt.Value != Seven {
		t.Errorf("Value does not match expected 'seven', got: %s", rslt.Color)
	}
}

func TestParsesMultiple(t *testing.T) {
	cards := parse("7red", "drawTwoRed")
	expected := []*Card{
		&Card{
			Color: Red,
			Value: Seven,
		},
		&Card{
			Color: Red,
			Value: DrawTwo,
		},
	}

	for i, c := range cards {
		if c.Color != expected[i].Color {
			t.Errorf(
				"Color does not match expected '%s', got: %s",
				expected[i].Color,
				c.Color,
			)
		}
		if c.Value != expected[i].Value {
			t.Errorf("Value does not match expected '%s', got: %s",
				expected[i].Value,
				c.Value,
			)
		}
	}
}

func TestIsSpecial(t *testing.T) {
	card := "drawTwoRed"

	match, rslt := isSpecial(card)

	if !match {
		t.Error("Did not match.")
	}
	if rslt != "drawTwo" {
		t.Error("Could not parse.")
	}
}
