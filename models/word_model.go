package models

type Word struct {
	WordID    string `json:"word_id"`
	TopicID   string `json:"topic_id"`
	WordName  string `json:"word_name"`
	WordType  string `json:"word_type"`
	IsLearned bool   `json:"is_learned"`
}
