package puppy

import (
	"github.com/Harsh00067/dog"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof! Woof!"
}

func BigBark() string {
	dog.WhenGrownUp(Bark())
}

func BigBarks() string {
	dog.WhenGrownUp(Barks())
}
