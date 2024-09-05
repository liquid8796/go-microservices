package dto

type LogEntry struct {
	Name string `bson:"name" json:"name"`
	Data string `bson:"data" json:"data"`
}
