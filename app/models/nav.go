package models

import (
	"fmt"
	"gositev2/database"
)

type Nav struct {
	Name string
	Route string
	Level int
	CreatedAt int
}

func (n *Nav) TableName() string {
	return "nav"
}

func (n *Nav) ListNavs() []*Nav {
	var result = make([]*Nav, 0)
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover", err)
		}
	}()
	find := database.DB.Find(&result)
	if find.Error != nil {
		panic(find.Error)
	}
	return result
}
