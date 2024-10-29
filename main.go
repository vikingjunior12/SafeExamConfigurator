package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {

	fmt.Println("Starting the application...")
	fmt.Println("Willkommen beim Mac SafeExam Konfigurator")

	var SafeExamBrowserLinkMAC string = "https://github.com/SafeExamBrowser/seb-mac/releases/download/3.4/SafeExamBrowser-3.4.dmg"

	homeDir := os.Getenv("HOME") // Holt das Home-Verzeichnis
	var downloadPath string = homeDir + "/Downloads/SafeExamBrowser-3.4.dmg"

	fmt.Println("Downloading SafeExamBrowser for Mac...") //Funktion downloadFile wird aufgerufen

	// Funktion downloadFile wird aufgerufen
	err := downloadFile(SafeExamBrowserLinkMAC, downloadPath) // hier wird die Funktion aufgerufen
	if err != nil {
		fmt.Println("Fehler beim Herunterladen:", err)
		return
	}
	fmt.Println("Download erfolgreich abgeschlossen")

	// 1. Mounten der .dmg-Datei und Output erhalten
	cmd := exec.Command("hdiutil", "attach", downloadPath)
	var out bytes.Buffer // Output wird in einen Buffer geschrieben und kann später ausgegeben werden
	cmd.Stdout = &out    // Output wird in den Buffer geschrieben
	if err := cmd.Run(); err != nil {
		fmt.Println("Fehler beim Mounten der .dmg-Datei:", err)
		return
	}


// Pfad des gemounteten Volumes ermitteln
scanner := bufio.NewScanner(&out)
var volumePath string
for scanner.Scan() {
    line := scanner.Text()
    if strings.Contains(line, "/Volumes/SafeExamBrowser") {
        // Extrahiere den Pfad bis zum Tabulator oder dem Ende der Zeile
        fields := strings.Fields(line)
        if len(fields) > 0 {
            volumePath = fields[0]
        }
        break
    }
}

if volumePath == "" {
    fmt.Printf("Das Volume 'SafeExamBrowser' konnte nicht gefunden werden.")
} else {
    fmt.Printf("Gefundener Pfad: %s\n", volumePath)
}



	// Kopiere die App in den Applications-Ordner
	copyCmd := exec.Command("sudo", "cp", "-r", "/Volumes/SafeExamBrowser-3.4/Safe Exam Browser.app", "/Applications/")
	output, err := copyCmd.CombinedOutput()
	if err != nil {
	fmt.Printf("Error: %v\n", err)
		fmt.Printf("Output: %s\n", string(output))
		return
	}
	fmt.Printf("Output: %s\n", string(output))

	// 3. Unmounten des Volumes
	unmountCmd := exec.Command("hdiutil", "detach", volumePath)
	if err := unmountCmd.Run(); err != nil {
		fmt.Println("Fehler beim Unmounten des Volumes:", err)
		return
	}
	fmt.Println("Volume erfolgreich unmounted")
}
