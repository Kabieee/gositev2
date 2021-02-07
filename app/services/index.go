package services

import (
	"fmt"
	"gositev2/app/models"
	"gositev2/common"
	"gositev2/database"
	"math"
	"net/url"
	"strconv"
)

type IndexServices struct {
}

func (i *IndexServices) Index() []*models.Bing {
	m := &models.Bing{}
	return m.GetList()
}

func (i *IndexServices) Search(params SearchParams) map[string]interface{} {
	var list = make([]*models.Bing, 0)
	var int64Count = new(int64)
	query := database.DB.Model(models.Bing{}).Order("startdate DESC")
	limit := 7

	u := url.URL{}
	values := u.Query()
	if params.Search != "" {
		values.Set("s", params.Search)
		query.Where("CONCAT(title,copyright) LIKE ?", "%"+params.Search+"%")
	}

	query.Count(int64Count)
	totalCount := int(*int64Count)
	totalPage := int(math.Ceil(float64(totalCount) / float64(limit)))

	if params.Page >= totalPage {
		params.Page = totalPage
	}

	offset := (params.Page - 1) * limit
	query.Limit(limit).Offset(offset)

	find := query.Find(&list)

	if find.Error != nil {
		common.SendPanic(404, find.Error.Error())
	}
	fmt.Println()

	m := map[string]interface{}{
		"List":        list,
		"TotalCount":  totalCount,
		"Page":        params.Page,
		"Search":      params.Search,
		"Pagination":  Pagination(params, totalPage),
	}
	return m
}

type Pager struct {
	PageUrls    []map[string]string
	Previous    bool
	Next        bool
	PreviousUrl string
	NextUrl     string
}

func Pagination(params SearchParams, totalPage int) Pager {
	currentPage := params.Page
	fmt.Println(totalPage)
	show := 5
	offset := show / 2
	s := make([]int, 0)
	if totalPage <= show {
		// 页码过小
		fmt.Println("min")
		s = getRandInt(1, totalPage)
	} else {
		if currentPage < show {
			// 靠近开头
			fmt.Println("start")
			s = getRandInt(1, show)
		} else if currentPage > totalPage-show+1 {
			// 靠近结尾
			fmt.Println("end")
			s = getRandInt(totalPage-show+1, totalPage)
		} else {
			// 在中间
			fmt.Println("middle")
			s = getRandInt(currentPage-offset, currentPage+offset)
		}
	}

	var pageUrls = make([]map[string]string, 0)
	u := url.URL{}
	values := u.Query()
	if params.Search != "" {
		values.Set("s", params.Search)
	}
	for _, v := range s {
		var m = make(map[string]string, 2)
		if v == currentPage {
			m["class"] = "active"
		} else {
			m["class"] = ""
		}
		page := strconv.Itoa(v)
		values.Set("p", page)
		m["url"] = "?" + values.Encode()
		m["page"] = page
		pageUrls = append(pageUrls, m)
	}

	p := Pager{
		PageUrls:    nil,
		Previous:    true,
		Next:        true,
		PreviousUrl: "",
		NextUrl:     "",
	}
	p.PageUrls = pageUrls

	if 1 == currentPage {
		p.Previous = false
		p.PreviousUrl = "#"
	}

	if totalPage == currentPage {
		p.Next = false
		p.NextUrl = "#"
	}

	if p.Previous {
		values.Set("p", strconv.Itoa(currentPage-1))
		p.PreviousUrl = "?" + values.Encode()
	}
	if p.Next {
		values.Set("p", strconv.Itoa(currentPage+1))
		p.NextUrl = "?" + values.Encode()
	}

	return p
}

func getRandInt(start, end int) []int {
	var s = make([]int, 0)
	for end >= start {
		s = append(s, start)
		start++
	}
	return s
}
