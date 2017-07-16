package models

type Essence struct {
	id      string
	name    string
	content string
	time    string
}

func NewEssence(id, name, content, time string) *Essence {
	return &Essence{id, name, content, time}
}
