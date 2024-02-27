package models

type Topic struct {
	Id                    string `json:"id"`
	TopicName             string `json:"topic_name"`
	NumberOfLearnedLesson int    `json:"number_of_learned_lesson"`
	TotalLessons          int    `json:"total_lessons"`
}

// var TopicList = []Topic{
// 	{Id: "1", TopicName: "Trái cây", NumberOfLearnedLesson: 1, TotalLessons: 3},
// 	{Id: "2", TopicName: "Phương tiện", NumberOfLearnedLesson: 2, TotalLessons: 3},
// 	{Id: "3", TopicName: "Con người", NumberOfLearnedLesson: 0, TotalLessons: 3},
// }
