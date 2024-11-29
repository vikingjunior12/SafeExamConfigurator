package main

import (
	"log"
	"os/exec"
)

func uninstallSEB(d string) {

	cmd := exec.Command(d, "/uninstall", "/silent")
	if err := cmd.Start(); err != nil {
		log.Fatalf("Fehler beim Deinstallieren der Datei: %v", err)
	}
	err := cmd.Wait()
	if err != nil {
		log.Fatalf("Fehler beim deinstalieren auf den Abschluss der Datei: %v", err)
	}

}
