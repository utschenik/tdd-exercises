package tests

import (
	"testing"

	"github.com/utschenik/tdd-exercises/src"
)

func TestGreetGiveNameReturnNameToMatch(t *testing.T) {
	greeter := src.Greetings{Name: "Amy"}
	greetings := greeter.Greet()

	if greetings != "Hello, Amy." {
		t.Errorf("Greeting was incorrect, got: %s, want %s", greetings, "Hello, Amy.")
	}
}

func TestGreetGiveNameEmptyReturnHelloFriend(t *testing.T) {
	greeter := src.Greetings{}
	greetings := greeter.Greet()

	if greetings != "Hello, my friend." {
		t.Errorf("Greeting was incorrect, got: %s, want %s", greetings, "Hello, my friend.")
	}
}

func TestGreetGiveNameAllUppercaseReturnGreetingsUppercaseAndReverse(t *testing.T) {
	greeter := src.Greetings{Name: "AMY"}
	greetings := greeter.Greet()

	if greetings != "HELLO, AMY." {
		t.Errorf("Greeting was incorrect, got: %s, want %s", greetings, "HELLO, AMY.")
	}

	greeter.Name = "Amy"
	greetings = greeter.Greet()

	if greetings != "Hello, Amy." {
		t.Errorf("Greeting was incorrect, got: %s, want %s", greetings, "Hello, Amy.")
	}
}

func TestGreetGiveMultipleNamesReturnMultipleNames(t *testing.T) {
	greeter := src.Greetings{Names: []string{"Amy", "John"}}
	greetings := greeter.Greet()

	if greetings != "Hello, Amy and John." {
		t.Errorf("Greeting was incorrect, got: %s, want %s", greetings, "Hello, Amy and John.")
	}

	greeter.Names = []string{"Bill", "John"}
	greetings = greeter.Greet()

	if greetings != "Hello, Bill and John." {
		t.Errorf("Greeting was incorrect, got: %s, want %s", greetings, "Hello, Bill and John.")
	}
}
