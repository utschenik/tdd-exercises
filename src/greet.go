package src

import (
	"fmt"
	"strings"
)

type Greetings struct {
	Name               string
	Names              []string
	isAnyNameUppercase bool
}

func (g *Greetings) Greet() string {
	g.checkForEmptyName()
	g.checkNameForUppercase()
	g.checkNamesForUppercase()
	return fmt.Sprintf("%s, %s", g.getRightGreetings(), g.getRightNames())
}

func (g *Greetings) checkForEmptyName() {
	if g.Name == "" {
		g.Name = "my friend"
	}
}

func (g *Greetings) checkNameForUppercase() {
	checkIfNameIsUpperCase(g.Name, g)
}

func (g *Greetings) checkNamesForUppercase() {
	for _, nameIn := range g.Names {
		checkIfNameIsUpperCase(nameIn, g)
		if g.isAnyNameUppercase {
			return
		}
	}
}

func checkIfNameIsUpperCase(nameIn string, g *Greetings) {
	if strings.ToUpper(nameIn) == nameIn {
		g.isAnyNameUppercase = true
	} else {
		g.isAnyNameUppercase = false
	}
}

func (g *Greetings) getRightGreetings() string {
	if g.isAnyNameUppercase && len(g.Names) == 0 {
		return "HELLO"
	} else {
		return "Hello"
	}
}

func (g *Greetings) getRightNames() string {
	lengthOfNames := len(g.Names)

	if lengthOfNames == 0 {
		return g.Name + "."
	} else if lengthOfNames == 2 {
		return strings.Join(g.Names, " and ") + "."
	} else {
		lowercaseNames := []string{}
		uppercaseNamesGreetings := []string{}

		for _, nameIn := range g.Names {
			if strings.ToUpper(nameIn) != nameIn {
				lowercaseNames = append(lowercaseNames, nameIn)
			} else {
				uppercaseNamesGreetings = append(uppercaseNamesGreetings, fmt.Sprintf("AND HELLO %s!", nameIn))
			}
		}

		lastNameInSlice := lowercaseNames[len(lowercaseNames)-1]
		g.Names = lowercaseNames[:len(lowercaseNames)-1]

		namesForGreeting := strings.Join(g.Names, ", ")
		uppercaseGreeting := strings.Join(uppercaseNamesGreetings, " ")

		if g.isAnyNameUppercase {
			return namesForGreeting + " and " + lastNameInSlice + ". " + uppercaseGreeting
		} else {
			return namesForGreeting + " and " + lastNameInSlice + "."
		}
	}
}
