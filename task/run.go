package task

import (
	"fmt"
	"time"
)

func Run()  {
	ticker := time.NewTicker(time.Minute * 15)
	defer ticker.Stop()
	fmt.Println("run task")
	i := new(ImageTask)
	i.SaveImage()
	for  {
		select  {
		case t := <-ticker.C:
			fmt.Printf("start image ticker: %s\n", t.Format(time.RFC3339))
			go i.SaveImage()
		default:
		}
	}

}
