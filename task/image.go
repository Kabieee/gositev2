package task

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gositev2/app/models"
	"myrepo"
	"myrepo/myclient"
	"strconv"
	"time"
)

type ImageTask struct {
	images []interface{}
}

func (i *ImageTask) GetImage() {
	request := myclient.NewRequest()
	reqUrl := "http://tt.makemake.in/bing-images/?format=js&idx=0&n=8&mkt=zh-CN"
	get, err := request.Get(reqUrl, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	var m map[string]interface{}
	decode := json.NewDecoder(bytes.NewReader(get.Body))
	err = decode.Decode(&m)
	if err != nil {
		fmt.Println(err)
		return
	}
	if mapImages, ok := m["images"].([]interface{}); ok {
		i.images = mapImages
	}
}

func (i *ImageTask) SaveImage() {
	i.GetImage()
	//fmt.Println(i.images)
	b := new(models.Bing)
	now := time.Now()
	today, _ := strconv.Atoi(now.Format("20060102"))
	lastDay, _ := strconv.Atoi(now.AddDate(0, 0, -15).Format("20060102"))
	//fmt.Println(today, lastDay)
	date := b.GetDate(lastDay, today)
	fmt.Println(date)
	//return

	type BingTemp struct {
		models.Bing
		FullstartdateString string `json:"fullstartdate" gorm:"-"`
		StartdateString     string `json:"startdate" gorm:"-"`
		EnddateString       string `json:"enddate" gorm:"-"`
	}
	baseUrl := "https://www.bing.com"
	insertData := make([]*models.Bing, 0)
	for _, image := range i.images {
		switch image.(type) {
		case map[string]interface{}:
			var temp BingTemp
			buf := bytes.NewBuffer(nil)
			_ = json.NewEncoder(buf).Encode(image)
			err := json.NewDecoder(buf).Decode(&temp)
			if err != nil {
				fmt.Println(err)
				continue
			}
			temp.Bing.Startdate, _ = strconv.Atoi(temp.StartdateString)
			temp.Bing.Enddate, _ = strconv.Atoi(temp.EnddateString)
			temp.Bing.Fullstartdate, _ = strconv.Atoi(temp.FullstartdateString)

			if myrepo.InArray(temp.Bing.Startdate, date) {
				//fmt.Println(temp.Bing.Startdate)
				continue
			} else {
				fmt.Println(temp.Bing.Startdate)
			}
			temp.Bing.FullImageUrl = baseUrl + temp.Bing.Url
			temp.Bing.QuizUrl = baseUrl + temp.Bing.Quiz

			//fmt.Printf("%#v\n", temp)
			insertData = append(insertData, &temp.Bing)
		default:
			fmt.Println("not data", image)
			continue
		}
	}
	if len(insertData) > 0 {
		fmt.Printf("%#v\n", insertData)
		b.BatchSave(insertData)
	}
}
