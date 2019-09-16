package main

import (
	"fmt"
	"math/rand"
	"time"
)

var deckMaster = [][]int{
	{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13},
	{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13},
	{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13},
	{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13},
}
var deckFlag = [4][13]bool{}

type Deck struct {
}

func (d *Deck) Drawcard() int {
	rand.Seed(time.Now().UnixNano())
	var suit, num int
	for {
		suit = rand.Intn(4)
		num = rand.Intn(13)
		if !deckFlag[suit][num] {
			deckFlag[suit][num] = true
			break
		}
	}
	return deckMaster[suit][num]
}

func ShowDeckFlag() {
	fmt.Printf("%#v\n", deckFlag)
}
