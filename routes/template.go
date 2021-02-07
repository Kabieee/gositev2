package routes

import (
	"html/template"
	"time"
)

func TemplateFunc() template.FuncMap {
	return template.FuncMap{
		"Inc": Inc,
		"Sub": Sub,
		"FormatDate": FormatDate,
	}
}

func Inc(i int) int {
	return i + 1
}

func Sub(i int) int {
	return i - 1
}

func FormatDate(layout string, timestamp... int64) string {
	if len(timestamp) > 0 {
		return  time.Unix(timestamp[0], 0).Format(time.RFC850)
	} else {
		return time.Now().Format(layout)

	}
}