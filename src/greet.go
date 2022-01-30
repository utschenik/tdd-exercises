package src

import (
	"fmt"
	"strings"
)

type Greeter interface {
	Greet() string
	checkForEmptyName()
	checkNameForUppercase()
	getRightGreetings() string
	getRightNames() string
}

type Greetings struct {
	Name            string
	Names           []string
	isNameUppercase bool
}

func (g *Greetings) Greet() string {
	g.checkForEmptyName()
	g.checkNameForUppercase()
	return fmt.Sprintf("%s, %s.", g.getRightGreetings(), g.getRightNames())
}

func (g *Greetings) checkForEmptyName() {
	if g.Name == "" {
		g.Name = "my friend"
	}
}

func (g *Greetings) checkNameForUppercase() {
	if strings.ToUpper(g.Name) == g.Name {
		g.isNameUppercase = true
	} else {
		g.isNameUppercase = false
	}
}

func (g *Greetings) getRightGreetings() string {
	if g.isNameUppercase {
		return "HELLO"
	} else {
		return "Hello"
	}
}

func (g *Greetings) getRightNames() string {
	if len(g.Names) == 0 {
		return g.Name
	} else {
		return strings.Join(g.Names, " and ")
	}
}
