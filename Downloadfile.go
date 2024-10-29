package main

import (
	"io"
	"net/http"
	"os"

	"github.com/cheggaaa/pb/v3"
)

func downloadFile(url string, filepath string) error {
	// Anfrage an die URL senden
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// Datei erstellen, in die wir herunterladen
	outfile, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer outfile.Close()

	// Dateigröße für den Ladebalken ermitteln
	size := response.ContentLength
	bar := pb.Full.Start64(size) // Ladebalken initialisieren
	defer bar.Finish()           // Ladebalken abschließen, wenn der Download beendet ist

	// Fortschritt verfolgen, indem wir den Response-Body durch einen TeeReader leiten
	reader := bar.NewProxyReader(response.Body)

	// Datei schreiben und Fortschritt anzeigen
	_, err = io.Copy(outfile, reader)
	if err != nil {
		return err

	}

	return nil // Null zurückgeben, wenn alles erfolgreich war
} // Das Null war notwenig weil die Funktion error als rückgabe wert definiert hatt. 
