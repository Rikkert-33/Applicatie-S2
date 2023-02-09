package main

import (
	"fmt"
	"time"
)

var CurrentTime int = 0

func TimeCheck() {
	CurrentTime := time.Now().Format("15")

	//fmt.Println(CurrentTime)

	if CurrentTime >= "7" && CurrentTime < "12" {
		fmt.Println("Goede morgen!")
	} else if CurrentTime > "12" && CurrentTime < "18" {
		fmt.Println("Goede middag!")
	} else if CurrentTime > "18" && CurrentTime < "23" {
		fmt.Println("Goede avond!")
	} else if CurrentTime > "23" && CurrentTime < "7" {
		fmt.Println("Sorry, de parkeerplaats is snachts gesloten")
	}
}

func main() {
	TimeCheck()
}
