package aoc2021day4

import (
	"context"
	"fmt"
	"strconv"
	"strings"
)

type Game struct {
	Cards []Card
}
type Card struct {
	Rows [5]Row
}
type Row struct {
	Squares [5]Square
}
type Square struct {
	Number string
	Called bool
}

// Part One: 41668
func Part1(ctx context.Context, instructions []string) int {
	var (
		calls         []string
		game          Game
		winningCardID []int
		winningCall   string
		debug         = ctx.Value("debug").(bool)
	)

	calls = append(calls, strings.Split(instructions[0], ",")...)

	for j := 2; j < len(instructions); j += 6 {
		card := Card{}
		for k := 0; k <= 4; k++ {
			row := Row{}
			instruction := strings.Split(strings.ReplaceAll(strings.Trim(instructions[j+k], " "), "  ", " "), " ")
			for l, ins := range instruction {
				row.Squares[l] = Square{ins, false}
			}
			card.Rows[k] = row
		}
		game.Cards = append(game.Cards, card)
	}

	// Check the rows and columns for a winner.
	for _, call := range calls {
		game = markCard(call, game)

		winningCardID = checkRows(game)
		if len(winningCardID) > 0 {
			winningCall = call
			break
		}

		winningCardID = checkCols(game)
		if len(winningCardID) > 0 {
			winningCall = call
			break
		}
	}

	if debug {
		fmt.Printf("Card %d on call %s is a winner!\n", winningCardID, winningCall)
	}

	sumOfCard := countUncalledNumbersOnCard(game.Cards[winningCardID[0]])

	if debug {
		fmt.Println("Sum of remaining numbers on card:", sumOfCard)
	}

	winningCallInt, _ := strconv.Atoi(winningCall)
	return sumOfCard * winningCallInt
}

// Part Two: 10478
func Part2(ctx context.Context, instructions []string) int {
	var (
		calls         []string
		game          Game
		winningCardID []int
		winningCall   string
		debug         = ctx.Value("debug").(bool)
	)

	calls = append(calls, strings.Split(instructions[0], ",")...)

	for j := 2; j < len(instructions); j += 6 {
		card := Card{}
		for k := 0; k <= 4; k++ {
			row := Row{}
			instruction := strings.Split(strings.ReplaceAll(strings.Trim(instructions[j+k], " "), "  ", " "), " ")
			for l, ins := range instruction {
				row.Squares[l] = Square{ins, false}
			}
			card.Rows[k] = row
		}
		game.Cards = append(game.Cards, card)
	}

	// Check the rows and columns for a winner.
	for i, call := range calls {

		if debug {
			fmt.Printf("Call %d is %s.\n", i, call)
		}

		game = markCard(call, game)

		winningCardID = checkRows(game)
		if len(winningCardID) > 0 {
			if len(game.Cards) == 1 {
				winningCall = call
				break
			}

			for _, id := range winningCardID {
				game = removeCard(game, id)
			}
		}

		winningCardID = checkCols(game)
		if len(winningCardID) > 0 {
			if len(game.Cards) == 1 {
				winningCall = call
				break
			}

			for _, id := range winningCardID {
				game = removeCard(game, id)
			}
		}

		if debug {
			drawCards(game)
		}
	}

	if debug {
		fmt.Println("Winning call:", winningCall)
		drawCards(game)
	}

	sumOfCard := countUncalledNumbersOnCard(game.Cards[0])
	if debug {
		fmt.Println("Sum of remaining numbers on card:", sumOfCard)
	}

	winningCallInt, _ := strconv.Atoi(winningCall)
	return sumOfCard * winningCallInt
}

// Loop through each square on a give card, and count the uncallled ones.
func countUncalledNumbersOnCard(card Card) int {
	var output int

	for _, row := range card.Rows {
		for _, square := range row.Squares {
			if !square.Called {
				int, _ := strconv.Atoi(square.Number)
				output += int
			}
		}
	}

	return output
}

// Check all squares in all cards for a match against the called number, and mark it as called.
func markCard(call string, game Game) Game {
	for i1, g := range game.Cards {
		for i2, cards := range g.Rows {
			for i3, c := range cards.Squares {
				if c.Number == call {
					game.Cards[i1].Rows[i2].Squares[i3].Called = true
				}
			}
		}
	}

	return game
}

// Draw the cards as a grid.
func drawCards(game Game) {
	fmt.Println()
	for i, c := range game.Cards {
		fmt.Println("Card", i)
		for i, r := range c.Rows {
			fmt.Printf("row %d: ", i)
			for _, s := range r.Squares {
				if s.Called {
					// Format a 'called' number as another colour, or bold, or something.
					fmt.Printf("%s%2s%s ", "\u001b[32m", s.Number, "\u001b[0m")
				} else {
					fmt.Printf("%2s ", s.Number)
				}
			}
			fmt.Println()
		}
		fmt.Println()
	}
	fmt.Println()
}

func removeCard(game Game, cardID int) Game {
	var newGame Game

	for i := range game.Cards {
		if i != cardID {
			newGame.Cards = append(newGame.Cards, game.Cards[i])
		}
	}

	return newGame
}

// Check all rows in all cards for a full / winning row.
func checkRows(game Game) []int {
	var out []int

	for i, card := range game.Cards {
		for _, row := range card.Rows {
			if row.Squares[0].Called && row.Squares[1].Called && row.Squares[2].Called && row.Squares[3].Called && row.Squares[4].Called {
				out = append(out, i)
			}
		}
	}

	return out
}

// Check all columns in all cards for a full / winning column.
func checkCols(game Game) []int {
	var out []int

	for i, card := range game.Cards {
		for j := 0; j <= 4; j++ {
			if card.Rows[0].Squares[j].Called && card.Rows[1].Squares[j].Called && card.Rows[2].Squares[j].Called && card.Rows[3].Squares[j].Called && card.Rows[4].Squares[j].Called {
				out = append(out, i)
			}
		}
	}

	return out
}
