package pkg

type Character struct {
	Id         int
	Health     int
	Experience int
	Level      int
	Food       int
	Water      int
	Energy     int
	Items      []Item
}

func (c *Character) setHealth(Health int) {
	c.Health = Health
}

func (c *Character) getHealth() int {
	return c.Health
}

func (c *Character) setExperience(Experience int) {
	c.Experience = Experience
}

func (c *Character) getExperience() int {
	return c.Experience
}

func (c *Character) setLevel(Level int) {
	c.Level = Level
}

func (c *Character) getLevel() int {
	return c.Level
}

func (c *Character) addItem(item Item) {
	c.Items = append(c.Items, item)
}

func (c *Character) getItems() []Item {
	return c.Items
}

func (c *Character) setFood(Food int) {
	c.Food = Food
}

func (c *Character) getFood() int {
	return c.Food
}

func (c *Character) setWater(Water int) {
	c.Water = Water
}

func (c *Character) getWater() int {
	return c.Water
}

func (c *Character) setEnergy(Energy int) {
	c.Energy = Energy
}

func (c *Character) getEnergy() int {
	return c.Energy
}
