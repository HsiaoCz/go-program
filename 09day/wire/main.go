package main

import "fmt"

// wire use
// there is a monster
type Monster struct {
	Name string
}

func NewMonster() Monster {
	return Monster{
		Name: "hello kitty",
	}
}

// and there also has a brove man
type Player struct {
	Name string
}

func NewPlayer(name string) Player {
	return Player{Name: name}
}

// yuong man beat the monster
type Mission struct {
	Player  Player
	Monster Monster
}

func NewMission(p Player, m Monster) Mission {
	return Mission{Player: p, Monster: m}
}
// finally young man beat the monster hello kitty
// function about that 
func (m Mission) Start() {
	fmt.Printf("%s defeats %s, whole world peace now!\n", m.Player.Name, m.Monster.Name)
}

func main() {
	monster := NewMonster()
	player := NewPlayer("zhangsan")
	mission := NewMission(player, monster)
	mission.Start()
}
