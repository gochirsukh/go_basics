package main

import (
	"fmt"

	"doggystore.com/dog_breed_app/dog_module"
)

func main() {
	list := dog_module.ListOfDogs{
		Dog1: "Shepperd\n",
		Dog2: "PitBull\n",
		Dog3: "Boston Terrier",
	}
	fmt.Println(list.Dog1, list.Dog2, list.Dog3)
}
