package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

var CurrentTime int = 0

func TimeCheck() {
	CurrentTime := time.Now().Format("15")
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

	// Toon het juiste bericht op basis van of het kenteken aanwezig is of niet
	if found {
		fmt.Println("Welkom op het parkeerterrein.")
		//TimeCheck staat hier vanwege het moment dat "Goedendag" wordt aangeroepen. Dit moet pas gebeuren als het kenteken is gecontroleerd.
		TimeCheck()
	} else {
		fmt.Println("Je hebt helaas geen toegang tot het parkeerterrein.")
	}
}

func main() {
	KentekenCheck()
	//TimeCheck staat in de KentekenCheck omdat het moment dat "Goedendag" wordt aangeroepen afhankelijk is van het kenteken.
}
