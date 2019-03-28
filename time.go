package main

import (
	"fmt"
	"time"
)

func main() {
	today := time.Now()
	time1 := today.Format("2006/01/02")
	time2 := today.Format("15:04")
	time3 := today.Format(time.RFC1123)
	fmt.Println(time1)
	fmt.Printf("%T %v\n", today, today)
	fmt.Printf("%T %v\n", time1, time1)
	fmt.Printf("%T %v\n", time2, time2)
	fmt.Printf("%T %v\n", time3, time3)

	t, err := time.Parse("2006/1/2", "2019/3/28")
	if err != nil {
		panic(err)
	}
	fmt.Println(t)

}
