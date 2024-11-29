package main

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/charmbracelet/huh/spinner"
)

func seteupexe(path string) {

	// Startet die Installation
	cmd := exec.Command(path, "/silent")
	if err := cmd.Start(); err != nil {
		log.Fatalf("Fehler beim Starten der Datei: %v", err)
	}
	fmt.Println("Installatwion wird ausgeführt, bitte warten...")

	action := func() {
		err := cmd.Wait()
		if err != nil {
			log.Fatalf("Fehler beim Warten auf den Abschluss der Datei: %v", err)
		}
	}
	_ = spinner.New().Title("Preparing your SafeExamBrowser...").Action(action).Run()
	fmt.Println("Installation abgeschlossen")
}
