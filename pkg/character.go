package main

/*
The character:
Health bar
Experience bar
items

Additional features for character:
Hunger
Thirst
Energy
*/
type Character struct {
	//id string
	health     int
	experience int
	level      int
	items      []Item
	food       int
	water      int
	energy     int
}

func (c *Character) setHealth(health int) {
	c.health = health
}

func (c *Character) getHealth() int {
	return c.health
}

func (c *Character) setExperience(experience int) {
	c.experience = experience
}

func (c *Character) getExperience() int {
	return c.experience
}

func (c *Character) setLevel(level int) {
	c.level = level
}

func (c *Character) getLevel() int {
	return c.level
}

func (c *Character) addItem(item Item) {
	c.items = append(c.items, item)
}

func (c *Character) getItems() []Item {
	return c.items
}

func (c *Character) setFood(food int) {
	c.food = food
}

func (c *Character) getFood() int {
	return c.food
}

func (c *Character) setWater(water int) {
	c.water = water
}

func (c *Character) getWater() int {
	return c.water
}

func (c *Character) setEnergy(energy int) {
	c.energy = energy
}

func (c *Character) getEnergy() int {
	return c.energy
}
