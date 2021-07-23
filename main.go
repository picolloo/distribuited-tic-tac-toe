package main

import (
	"fmt"

	"github.com/picolloo/distribuited-tic-tac-toe/cmd"
	"github.com/picolloo/distribuited-tic-tac-toe/game"
)

func main() {
	cmd.Execute()
  g := game.NewGame()


  pos, _ := g.NewPosition(1,1)
  err := g.NextTurn(pos)
  if err != nil {
    fmt.Println(err.Error())
  }
  pos, _ = g.NewPosition(2,1)
  err = g.NextTurn(pos)
  if err != nil {
    fmt.Println(err.Error())
  }
  g.Draw()
}
