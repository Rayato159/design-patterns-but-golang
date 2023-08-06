package main

import "fmt"

type playerStatus struct {
	Strength     int
	Intelligence int
	Agility      int
}

type PlayerClass interface {
	GetWeapon() string
	GetPlayerStatus() playerStatus
}

type playerClassInit struct {
	Weapon string
	Status playerStatus
}

func PlayerFactory(playerClass string) PlayerClass {
	switch playerClass {
	case "warrior":
		return newWarrior()
	case "mage":
		return newMage()
	default:
		panic("Error: player class not found!")
	}
}

type warrior struct {
	*playerClassInit
}

func newWarrior() PlayerClass {
	return &warrior{
		playerClassInit: &playerClassInit{
			Weapon: "Sword",
			Status: playerStatus{
				Strength:     10,
				Intelligence: 3,
				Agility:      5,
			},
		},
	}
}

func (w *warrior) GetWeapon() string             { return w.Weapon }
func (w *warrior) GetPlayerStatus() playerStatus { return w.Status }

type mage struct {
	*playerClassInit
}

func newMage() PlayerClass {
	return &warrior{
		playerClassInit: &playerClassInit{
			Weapon: "Staff",
			Status: playerStatus{
				Strength:     3,
				Intelligence: 10,
				Agility:      5,
			},
		},
	}
}

func (m *mage) GetWeapon() string             { return m.Weapon }
func (m *mage) GetPlayerStatus() playerStatus { return m.Status }

func main() {
	warrior := PlayerFactory("warrior")
	fmt.Println(warrior.GetWeapon())
	fmt.Println(warrior.GetPlayerStatus())

	mage := PlayerFactory("mage")
	fmt.Println(mage.GetWeapon())
	fmt.Println(mage.GetPlayerStatus())
}
