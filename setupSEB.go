package main

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
)

//go:embed assets/KonfigurationsDateiETutorCareum.seb
var sebFile []byte

func setupSEB() error {
	// Zielpfad: %APPDATA%/SafeExamBrowser/KonfigurationsDateiETutorCareum.seb
	appDataDir, err := os.UserConfigDir()
	if err != nil {
		return fmt.Errorf("APPDATA-Verzeichnis konnte nicht ermittelt werden: %w", err)
	}
	targetPath := filepath.Join(appDataDir, "SafeExamBrowser", "KonfigurationsDateiETutorCareum.seb")

	// Zielverzeichnis sicherstellen
	err = os.MkdirAll(filepath.Dir(targetPath), 0755)
	if err != nil {
		return fmt.Errorf("Fehler beim Erstellen des Zielverzeichnisses: %w", err)
	}

	// Datei an den Zielort schreiben
	err = os.WriteFile(targetPath, sebFile, 0644)
	if err != nil {
		return fmt.Errorf("Fehler beim Schreiben der Datei: %w", err)
	}

	fmt.Printf("Datei erfolgreich nach %s kopiert\n", targetPath)
	return nil
}
