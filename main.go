package main

import (
	"fmt"
	"gositev2/database"
	"gositev2/routes"
	"gositev2/task"
)

func main()  {
	routes.InitRoute()
	database.InitDB()
	go task.Run()
	_ = routes.Engine.Run(":9900")
	//pager(5)
}

const PSIZE = 6

func pager(pageNo int) {
	var start int = pageNo/(PSIZE-1)*(PSIZE-1) + 1
	fmt.Println(start)
	if pageNo%(PSIZE-1) == 0 {
		start = start - (PSIZE - 1)
	}
	fmt.Println(pageNo%(PSIZE-1), start)

	for i := start; i < start+PSIZE; i++ {
		if i == pageNo {
			fmt.Printf("(%d) ", i)
		} else {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Print("\n")
}