package main

import (
	"fmt"
	"math/rand"
	"time"
)

var suits = []string{"♣", "♦", "♥", "♠"}
var values = []string{"Ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King"}

type Card struct {
	value string
	sign  string
}

type Player struct {
	hand     []Card
	wonCards []Card
	numCards int
}

type Talon struct {
	card   []Card
	number int
}

type Deck struct {
	cards []Card
	num   int
}

func (d *Deck) createDeck() {
	for _, e := range suits {
		for _, i := range values {
			c := Card{value: i, sign: e}
			d.cards = append(d.cards, c)
			d.num++
		}
	}
}

func (d *Deck) shuffleDeck() {
	rand.Seed(time.Now().UnixNano())
	for i := d.num - 1; i > 0; i-- {
		j := rand.Intn(d.num)
		t := d.cards[i]
		d.cards[i] = d.cards[j]
		d.cards[j] = t
	}
}

func (d *Deck) drawCard(h *Player) {
	h.hand = append(h.hand, d.cards[0])
	d.cards = d.cards[1:]
	d.num--
	h.numCards++

}

func (d *Deck) newRound(p1 *Player, p2 *Player) {
	han1 := d.cards[:4]
	p1.hand = append(p1.hand, han1...)
	d.cards = d.cards[4:]
	p1.numCards = len(p1.hand)
	han2 := d.cards[:4]
	p2.hand = append(p2.hand, han2...)
	p2.numCards = len(p2.hand)
	d.cards = d.cards[4:]
	d.num = len(d.cards)
}

func (p *Player) removeCard(i int) {
	p.hand[i] = p.hand[len(p.hand)-1]
	p.hand = p.hand[:len(p.hand)-1]
	p.numCards--
}

func (p *Player) move(t *Talon) {
	var c int
	fmt.Println("Your move, which card are you throwing, here is what you have", p.hand)

	fmt.Scanln(&c)
	if t.number != 0 {
		if p.hand[c].value == t.card[t.number-1].value || string(p.hand[c].value) == "Jack" {
			p.wonCards = append(p.wonCards, t.card...)
			t.card = t.card[:0]
			t.number = 0
			p.removeCard(c)
		} else {
			t.card = append(t.card, p.hand[c])
			t.number++
			p.removeCard(c)

		}
	} else {
		t.card = append(t.card, p.hand[c])
		t.number++
		p.removeCard(c)
	}

}
func (p *Player) evaluate() int {
	sum := 0
	for _, i := range p.wonCards {
		if string(i.value) == "Queen" || string(i.value) == "King" || string(i.value) == "Jack" || string(i.value) == "Ace" || string(i.value) == "10" {
			sum++
		}
	}
	return sum
}

func startGame(d *Deck, p1 *Player, p2 *Player, t *Talon) {
	talon := d.cards[:12]
	t.card = append(t.card, talon...)
	d.cards = d.cards[12:]
	t.number = len(t.card)
	han1 := d.cards[:4]
	p1.hand = append(p1.hand, han1...)
	d.cards = d.cards[4:]
	p1.numCards = len(p1.hand)
	han2 := d.cards[:4]
	p2.hand = append(p2.hand, han2...)
	p2.numCards = len(p2.hand)
	d.cards = d.cards[4:]
	d.num = len(d.cards)
}

func main() {
	deck := &Deck{}
	t := &Talon{}
	p1 := &Player{}
	p2 := &Player{}
	deck.createDeck()
	deck.shuffleDeck()
	fmt.Println("schuffled", deck)
	startGame(deck, p1, p2, t)
	i := 0
	for deck.num != 0 {
		if t.number != 0 {
			fmt.Println(t.card[len(t.card)-1])
		}
		if p1.numCards == 0 && p2.numCards == 0 {
			deck.newRound(p1, p2)
			continue
		}
		if i%2 == 0 {
			p1.move(t)
			fmt.Println(t)
			i++
		} else {
			p2.move(t)
			fmt.Println(t)
			i++
		}
	}
	p1Score := p1.evaluate()
	p2Score := p2.evaluate()
	if p1Score > p2Score {
		fmt.Println("Player1 won")
	} else {
		fmt.Println("Player2 won")
	}

}
