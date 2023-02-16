package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

var CurrentTime int = 0
var status bool = false
var welcomeMessage string = ""

func Welcome(welcomeMessage string) string {
	CurrentTime := time.Now().Format("15")
	if CurrentTime >= "7" && CurrentTime < "12" {
		//fmt.Println("Goede morgen!")
		welcomeMessage = "Goede morgen!"
		return welcomeMessage
	} else if CurrentTime > "12" && CurrentTime < "18" {
		//fmt.Println("Goede middag!")
		welcomeMessage = "Goede middag!"
		return welcomeMessage
	} else if CurrentTime > "18" && CurrentTime < "23" {
		//fmt.Println("Goede avond!")
		welcomeMessage = "Goede avond!"
		return welcomeMessage
	} else if CurrentTime > "23" && CurrentTime < "7" {
		//fmt.Println("Sorry, de parkeerplaats is snachts gesloten")
		welcomeMessage = "Sorry, de parkeerplaats is snachts gesloten"
		return welcomeMessage
	}
	return ""
}

func KentekenCheck() {
	// Lijst met toegestane kentekens
	allowedLicenses := []string{"AB-123-CD", "EF-456-GH", "IJ-789-KL"}

	// Vraag het kenteken aan de gebruiker
	fmt.Print("Voer het kenteken in: ")
	reader := bufio.NewReader(os.Stdin)
	license, _ := reader.ReadString('\n')
	license = strings.TrimSpace(license)

	// Controleer of het kenteken aanwezig is in de lijst met toegestane kentekens
	found := false
	for _, l := range allowedLicenses {
		if l == license {
			found = true
			break
		}
	}
	if found {
		status = true
	}
}

func main() {
	KentekenCheck()
	if status == true {
		fmt.Println("Welkom op het parkeerterrein.")
		welcomeMessage = Welcome(welcomeMessage)
		fmt.Println(welcomeMessage)
	} else {
		fmt.Println("Je hebt helaas geen toegang tot het parkeerterrein.")
	}
}
