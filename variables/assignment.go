package main

import (
	"fmt"
	"path"
	"time"
)

func main(){ 
	var speed int
	fmt.Println(speed)

	speed = 111
	fmt.Println(speed)

	speed = speed - 1
	fmt.Println(speed)


	var running bool
	var force float64

	_, _ = running, force //Keep the compiler happy


	var (
		s int
		now time.Time
	)

	s, now = 7, time.Now	()

	fmt.Println(s, now)

	var (
		fname string = "John"
		lname string = "Doe"
	)

	fname, lname = lname, fname

	var dir, file string
	dir, file = path.Split("route/index.js")

	fmt.Println(dir, file)
}