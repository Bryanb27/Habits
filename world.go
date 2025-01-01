package main

type World struct {
	id                    int
	kind                  string
	nonplayablecharacters []NonPlayableCharacter
	items                 []Item
	mainEvents            []string
}
