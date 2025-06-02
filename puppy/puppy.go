package puppy

import (
	"github.com/Devanshu1610/dog"
)

func Barks() string {
	return "woof!woofwoof"
}

func Bark() string {
	return "woof"
}

func BigBarks() string {

	return dog.WhenGrownUp(Barks())
}
func BigBark() string {

	return dog.WhenGrownUp(Bark())
}
