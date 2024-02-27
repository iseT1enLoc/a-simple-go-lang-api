package database

import "simpleAPI/models"

var WordList = []models.Word{
	{WordID: "1", TopicID: "1", WordName: "Táo", WordType: "n", IsLearned: false},
	{WordID: "2", TopicID: "1", WordName: "Lê", WordType: "n", IsLearned: true},
	{WordID: "3", TopicID: "1", WordName: "Xoài", WordType: "n", IsLearned: false},
	{WordID: "4", TopicID: "2", WordName: "Xe đạp", WordType: "n", IsLearned: true},
	{WordID: "5", TopicID: "2", WordName: "Xe tải", WordType: "n", IsLearned: true},
	{WordID: "6", TopicID: "2", WordName: "Xe xích lô", WordType: "n", IsLearned: false},
	{WordID: "7", TopicID: "3", WordName: "Mắt", WordType: "n", IsLearned: false},
	{WordID: "8", TopicID: "3", WordName: "Mũi", WordType: "n", IsLearned: false},
	{WordID: "9", TopicID: "3", WordName: "Chạy", WordType: "v", IsLearned: false},
}
