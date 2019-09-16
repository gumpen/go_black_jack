package main

import (
	"fmt"
)

func main() {
	deck := &Deck{}
	for i := 0; i < 52; i++ {
		fmt.Println(deck.Drawcard())
	}

}
