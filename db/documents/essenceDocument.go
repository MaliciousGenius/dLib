package documents

type EssenceDocument struct {
	id      string `bson:"_id,omitempty"`
	name    string
	content string
	time    string
}
