package documents

import "time"

type EssenceDocument struct {
	Id      string `bson:"_id,omitempty"`
	Name    string
	Content string
	Time    time.Time
}
