package models

import (
	"fmt"
	"gositev2/database"
)

type Title struct {
	Name string `json:"name"`
	CreateAt  int64  `json:"create_at"`
}

func (t *Title) TableName() string {
	return "title"
}

func (t *Title) GetMainTitle() *Title {
	result := &Title{}
	find := database.DB.Model(&Title{}).Where(map[string]interface{}{
		"type": "main",
	}).First(result)

	if find.Error != nil {
		fmt.Println(find.Error)
	}
	return result
}
