package main

import (
	"task_21/internal/adapter"
	"task_21/internal/animal"
)

func main() {
	checkTheFish := adapter.NewCheckFish(&animal.Fish{})
	checkTheBird := adapter.NewCheckBird(&animal.Bird{})
	checkTheCat := adapter.NewCheckCat(&animal.Cat{})
	checkTheFish.CheckYourAnimals()
	checkTheBird.CheckYourAnimals()
	checkTheCat.CheckYourAnimals()
}
