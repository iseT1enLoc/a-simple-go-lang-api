package database

import "simpleAPI/models"

var TopicList = []models.Topic{
	{Id: "1", TopicName: "Trái cây", NumberOfLearnedLesson: 1, TotalLessons: 3},
	{Id: "2", TopicName: "Phương tiện", NumberOfLearnedLesson: 2, TotalLessons: 3},
	{Id: "3", TopicName: "Con người", NumberOfLearnedLesson: 0, TotalLessons: 3},
}
