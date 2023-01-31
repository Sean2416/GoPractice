package main

import (
	"os"
	"testing"
)

func TestNewDecK(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Number Error should have 52 cards but have %v", len(d))
	}

	if d[0] != "黑桃A" || d[len(d)-1] != "梅花K" {

		t.Errorf("The first card should be 黑桃A But get " + d[len(d)-1])

	}
}

func TestSaveAndGetFile(t *testing.T) {
	os.Remove("test.txt")
	d := newDeck()
	d.shuffle()
	_, hand, _ := d.deal(13)

	hand.saveTofile("test.txt")

	cards := readFile("test.txt")

	if len(cards) != 13 {
		t.Errorf("Number Error should have 13 cards but have %v", len(cards))
	}
	os.Remove("test.txt")
}
