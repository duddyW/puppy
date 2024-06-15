package puppy

import (
	"fmt"

	"github.com/duddyW/dog"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof! Woof!"
}

func BigBurk() string {
	return dog.WhenGrowUp(Bark())
}

func BigBurks() string {
	return dog.WhenGrowUp(Barks())
}

func From11() {
	fmt.Println("I'm from version v1.1.0")
}

func From12() {
	fmt.Println("I'm from version v1.2.0")
}
