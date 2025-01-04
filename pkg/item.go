package pkg

type Item struct {
	id        int
	name      string
	itype     string // e.g., "Weapon", "Potion", "Tool"
	rarity    string // e.g., "Common", "Rare", "Epic"
	minLevel  int    // Minimum character level required
	spawnRate float64
	effects   []Effect
}
