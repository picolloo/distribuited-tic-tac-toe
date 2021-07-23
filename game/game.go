package game

import (
	"fmt"
	"strings"
)

type Token string
type Position map[string]int

func (pos Position) X() int {
  return pos["x"]
}

func (pos Position) Y() int {
  return pos["y"]
}

const DEFAULT_TOKEN = Token("_")

func (t Token) String() string {
  return string(t)
}

func NewToken(value string) (Token, error) {
  if value != "X" || value != "O" {
    return "", fmt.Errorf("Invalid token value. The only possible values are X and O.")
  }
  return Token(value), nil
}

func (t Token) DefaultToken() bool {
  return t == DEFAULT_TOKEN
}

type Game struct {
   board [3][3]Token
   currentToken Token
}


func NewGame() *Game {
  return &Game{
    board: [3][3]Token{
      {Token("_"),Token("_"),Token("_")},
      {Token("_"),Token("_"),Token("_")},
      {Token("_"),Token("_"),Token("_")},
    },
    currentToken: Token("O"),
  }
}

func (g *Game) getToken(pos Position) Token {
  return g.board[pos.Y()][pos.X()]
}

func (g *Game) Draw() {
  for _, i := range(g.board) {
    var string_row []string

    for _,j := range(i) {
      string_row = append(string_row, j.String())
    }

    fmt.Println(strings.Join(string_row, "|"))
  }
}


func (g *Game) NewPosition(x, y int) (Position, error) {
  if x < 0 || x > len(g.board[0]) -1 || y < 0 || y > len(g.board) {
    return nil, fmt.Errorf("Invalid position values")
  }

  return Position{"x": x, "y": y}, nil
}

func (g *Game) commitPosition(pos Position) {
  g.board[pos.Y()][pos.X()] = g.currentToken
}

func (g *Game) nextToken() Token {
  if g.currentToken == Token("X") {
    return Token("Y")
  }
  return Token("X")
}

func (g *Game) NextTurn(pos Position) error {
  if !g.getToken(pos).DefaultToken() {
    return fmt.Errorf("Invalid position")
  }

  g.commitPosition(pos)
  g.currentToken = g.nextToken()

  // TODO check winner

  return nil
}

