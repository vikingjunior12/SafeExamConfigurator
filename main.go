package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {

	fmt.Println("\033[32mStarting the application...\033[0m")
	fmt.Println("\033[32mWillkommen beim Windows SafeExam Konfigurator\033[0m")

	SafeExamBrowserLinkWindows, version, err := getLatestSafeExamBrowserURL()
	if err != nil {
		fmt.Println("\033[31mFehler beim Abrufen des Download-Links:\033[0m", err)
		return
	}

	modifiedVersion := strings.ReplaceAll(version, ".", "-")

	homeDir := os.Getenv("USERPROFILE")
	var downloadPath string = fmt.Sprintf("%s/Downloads/SafeExamBrowser-%s.exe", homeDir, modifiedVersion)

	fmt.Println("\033[32mDownloading SafeExamBrowser for Windows...\033[0m")

	err = downloadFile(SafeExamBrowserLinkWindows, downloadPath)
	if err != nil {
		fmt.Println("\033[31mFehler beim Herunterladen:\033[0m", err)
		return
	}
	fmt.Println("\033[32mDownload erfolgreich abgeschlossen\033[0m")

	seteupexe(downloadPath) // Installation, wird gewartet

	fmt.Println("Starte SEB Careum Konfiguration...")
	err = setupSEB() // Seb Datei wird eingerichtet
	if err != nil {
		fmt.Println("Fehler beim Einrichten von SEB:", err)
	} else {
		fmt.Println("SEB Careum Konfiguration erfolgreich abgeschlossen.")
	}
	//auräumen
	err = os.Remove(downloadPath)
	if err != nil {
		fmt.Println("Fehler beim Löschen der Datei:", err)
	} else {
		fmt.Println("Datei erfolgreich gelöscht")
	}

}

// Funktion zum Abrufen des neuesten Download-Links für SafeExamBrowser und gibt zwei Werte zurück
func getLatestSafeExamBrowserURL() (string, string, error) {
	url := "https://api.github.com/repos/SafeExamBrowser/seb-win-refactoring/releases/latest"
	resp, err := http.Get(url)
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", "", fmt.Errorf("failed to get latest release: %s", resp.Status)
	}

	var releaseData struct {
		TagName string `json:"tag_name"`
		Assets  []struct {
			BrowserDownloadURL string `json:"browser_download_url"`
			Name               string `json:"name"`
		} `json:"assets"`
	}

	err = json.NewDecoder(resp.Body).Decode(&releaseData)
	if err != nil {
		return "", "", err
	}

	for _, asset := range releaseData.Assets {
		if strings.HasSuffix(asset.Name, ".exe") {
			return asset.BrowserDownloadURL, releaseData.TagName, nil
		}
	}

	return "", "", fmt.Errorf("no .exe asset found for SafeExamBrowser")
}
