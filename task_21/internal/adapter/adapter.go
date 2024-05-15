package adapter

import (
	"task_21/internal/animal"
)

type Check interface {
	CheckYourAnimals()
}

type FishAdapter struct {
	*animal.Fish
}

type BirdAdapter struct {
	*animal.Bird
}

type CatAdapter struct {
	*animal.Cat
}

func (f *FishAdapter) CheckYourAnimals() {
	f.Status()
}

func (b *BirdAdapter) CheckYourAnimals() {
	b.Status()
}

func (c *CatAdapter) CheckYourAnimals() {
	c.Status()
}

func NewCheckFish(f *animal.Fish) Check {
	return &FishAdapter{f}
}

func NewCheckBird(b *animal.Bird) Check {
	return &BirdAdapter{b}
}

func NewCheckCat(c *animal.Cat) Check {
	return &CatAdapter{c}
}
