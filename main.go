package main

import (
	"fmt"
	"os"

	"github.com/kiranS4/task/parse"
)

func main() {
	args := os.Args[1:]
	d := parse.NewData(args[0])
	d.Parse()
	fmt.Println("minutes        ", d.GetMin())
	fmt.Println("hour           ", d.GetHour())
	fmt.Println("day of month   ", d.GetDayOfMonth())
	fmt.Println("month          ", d.GetMonth())
	fmt.Println("day of week    ", d.GetDayofWeek())
	fmt.Println("command        ", d.GetCommand())
}
