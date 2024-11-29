package main

import (
	"log"
	"os/exec"
)

func seteupexe(path string) {

	// Startet die Installation
	cmd := exec.Command(path, "/silent")
	if err := cmd.Start(); err != nil {
		log.Fatalf("Fehler beim Starten der Datei: %v", err)
	}

	err := cmd.Wait()
	if err != nil {
		log.Fatalf("Fehler beim Warten auf den Abschluss der Datei: %v", err)
	}
}
