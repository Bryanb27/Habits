package main

type World struct {
	id         int
	kind       string
	npcs       []NPC
	items      []Item
	mainEvents []string
}
