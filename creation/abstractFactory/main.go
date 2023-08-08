package main

import "fmt"

type Mage interface {
	GetWeapon() string
	GetSkill() string
}

type mage struct {
	weapon string
	skill  string
}

type MageFactory struct {
	Factory Mage
}

func NewMageFactory(factory Mage) *MageFactory {
	return &MageFactory{
		Factory: factory,
	}
}

type pyromancer struct {
	*mage
}

func NewPyromancer() Mage {
	return &pyromancer{
		&mage{
			weapon: "Fist",
			skill:  "Fire ball",
		},
	}
}

type preist struct {
	*mage
}

func NewPreist() Mage {
	return &preist{
		&mage{
			weapon: "Staff of healing",
			skill:  "Healing",
		},
	}
}

func (m *mage) GetWeapon() string { return m.weapon }
func (m *mage) GetSkill() string  { return m.skill }

func main() {
	mage := NewMageFactory(NewPyromancer())
	fmt.Println(mage.Factory.GetWeapon())
	fmt.Println(mage.Factory.GetSkill())

	fmt.Println()

	mage.Factory = NewPreist()
	fmt.Println(mage.Factory.GetWeapon())
	fmt.Println(mage.Factory.GetSkill())
}
