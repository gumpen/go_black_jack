package main

import "fmt"

var allowedPlayerResponse = []string{
	"y",
	"n",
}

func main() {
	player := &Player{}
	dealer := &Player{dealer: true}

	dealer.DrawCard()
	player.DrawCard()
	dealer.DrawCard()
	player.DrawCard()
	fmt.Printf("%#v\n", player)

	for {
		if dealer.score < 17 {
			dealer.DrawCard()
		} else {
			dealer.stand = true
		}
		if !player.stand {
			fmt.Println("ヒットしますか？(y or n)")
			var playerRes string
			for {
				_, err := fmt.Scanln(&playerRes)
				if err != nil {
					panic(err)
				}
				if contain(allowedPlayerResponse, playerRes) {
					break
				}
				fmt.Println("y または n を入力してください")
			}
			if playerRes == "y" {
				player.DrawCard()
				fmt.Printf("%#v\n", player)
			} else {
				player.stand = true
			}
		}

		if (dealer.burst || player.burst) || (dealer.stand && player.stand) {
			break
		}
	}

	fmt.Printf("ディーラーの点数:%#v\n", dealer.score)
	fmt.Printf("プレイヤーの点数:%#v\n", player.score)

	if !player.burst && (player.score > dealer.score) {
		fmt.Println("プレイヤーの勝ち")
	} else {
		fmt.Println("ディーラーの勝ち")
	}
}

func ScoringByRules(hands []int) int {
	var sum int
	for _, n := range hands {
		if n >= 10 {
			sum += 10
		} else {
			sum += n
		}
	}
	return sum
}
