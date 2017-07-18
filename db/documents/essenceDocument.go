package documents

type EssenceDocument struct {
	Id      string `bson:"_id,omitempty"`
	Name    string
	Content string
	Time    string
}
