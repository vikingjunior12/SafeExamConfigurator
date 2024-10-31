package main

import (
	"fmt"
	"log"
	"os/exec"
)

func seteupexe(path string) {

	// Startet die Installation
	cmd := exec.Command(path, "/silent")
	if err := cmd.Start(); err != nil {
		log.Fatalf("Fehler beim Starten der Datei: %v", err)
	}
	fmt.Println("Installation wird ausgeführt, bitte warten...")

	err := cmd.Wait()
	if err != nil {
		log.Fatalf("Fehler beim Warten auf den Abschluss der Datei: %v", err)
	}
	fmt.Println("Installation abgeschlossen.")

}
