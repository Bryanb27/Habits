package pkg

type World struct {
	Id                    int
	Kind                  string
	Nonplayablecharacters []NonPlayableCharacter
}

func NewWorld(id int, kind string) World {
	return World{
		Id:                    id,
		Kind:                  kind,
		Nonplayablecharacters: []NonPlayableCharacter{},
	}
}
