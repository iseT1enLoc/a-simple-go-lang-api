package models

type Topic struct {
	Id                    string `json:"id"`
	TopicName             string `json:"topic_name"`
	NumberOfLearnedLesson int    `json:"number_of_learned_lesson"`
	TotalLessons          int    `json:"total_lessons"`
}
