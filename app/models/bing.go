package models

import (
	"fmt"
	"gositev2/common"
	"gositev2/database"
)

type Bing struct {
	ID            int    `json:"id"`
	Fullstartdate int    `json:"fullstartdate"`
	Startdate     int    `json:"startdate"`
	Enddate       int    `json:"enddate"`
	Url           string `json:"url"`
	Urlbase       string `json:"urlbase"`
	Copyright     string `json:"copyright"`
	Copyrightlink string `json:"copyrightlink"`
	Title         string `json:"title"`
	Quiz          string `json:"quiz"`
	FullImageUrl  string `json:"full_image_url"`
	QuizUrl       string `json:"quiz_url"`
	CreatedAt     int64  `json:"created_at"`
}

func (b *Bing) TableName() string {
	return "bing"
}

func (b *Bing) GetList() []*Bing {
	var result = make([]*Bing, 0)
	find := database.DB.Where("?", 1).Find(&result)
	if find.Error != nil {
		common.SendPanic(500, find.Error.Error())
	}
	return result
}


func (b *Bing) GetDate(start, end int) []int {
	var result = make([]int, 0)
	find := database.DB.Table("bing").Select("startdate").
		Where("startdate BETWEEN ? AND ?", start, end).
		Pluck("startdate",&result)
	if find.Error != nil {
		fmt.Println(find.Error)
	}
	return result
}

func (* Bing) BatchSave(data []*Bing) {
	create := database.DB.Model(Bing{}).Omit("ID").Create(&data)
	if create.Error != nil {
		fmt.Println(create.Error)
		return
	}
}
