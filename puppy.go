package puppy

import (
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
