package models

import "time"

type Essence struct {
	Id      string
	Name    string
	Content string
	Time    time.Time
}

func NewEssence(id, name, content string, time time.Time) *Essence {
	return &Essence{id, name, content, time}
}
