package main

type Player struct {
	rawHands []int
	score    int
	stand    bool
	burst    bool
	dealer   bool
}

func (p *Player) DrawCard() {
	deck := &Deck{}
	p.rawHands = append(p.rawHands, deck.Drawcard())
	p.Scoring()
	p.BurstCheck()
	return
}

func (p *Player) Scoring() {
	p.score = ScoringByRules(p.rawHands)
}

func (p *Player) BurstCheck() {
	if p.score > 21 {
		p.burst = true
	}
	return
}
