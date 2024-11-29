package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/charmbracelet/huh/spinner"
)

var downloadlink string = "https://api.github.com/repos/SafeExamBrowser/seb-win-refactoring/releases/latest"

func main() {

	uninstall := flag.Bool("u", false, "Uninstall SafeExamBrowser")
	flag.Parse()

	if *uninstall == true {
		fmt.Println("\033[31mDeinstallation von SafeExamBrowser wird gestartet...\033[0m")
		fmt.Println("\033[31mFür die Deinstaltion wird das Exe Runtergeladen und wieder gelöscht...\033[0m")
	} else {
		fmt.Println("\033[32mWillkommen beim Careum Windows SafeExam Konfigurator\033[0m")
	}

	SafeExamBrowserLinkWindows, version, err := getLatestSafeExamBrowserURL()
	if err != nil {
		fmt.Println("\033[31mFehler beim Abrufen des Download-Links:\033[0m", err)
		return
	}

	modifiedVersion := strings.ReplaceAll(version, ".", "-")

	homeDir := os.Getenv("USERPROFILE")
	var downloadPath string = fmt.Sprintf("%s/Downloads/SafeExamBrowser-%s.exe", homeDir, modifiedVersion)
	fmt.Println("")
	fmt.Println("\033[32mDownloading SafeExamBrowser for Windows...\033[0m")

	err = downloadFile(SafeExamBrowserLinkWindows, downloadPath)
	if err != nil {
		fmt.Println("\033[31mFehler beim Herunterladen:\033[0m", err)
		return
	}
	fmt.Println("\033[32mDownload erfolgreich abgeschlossen\033[0m")
	fmt.Println("")

	if *uninstall == true {

		action := func() {
			uninstallSEB(downloadPath) // Installation, wird gewartet
			cleanup(downloadPath)
		}
		_ = spinner.New().Title("Deinstallation von SafeExamBrowser läuft...").Action(action).Run()
		return
	}

	action := func() {
		seteupexe(downloadPath) // Installation, wird gewartet

	}
	_ = spinner.New().Title("Installation von SafeExamBrowser läuft...").Action(action).Run()

	fmt.Println("Starte SEB Careum Konfiguration...")
	err = setupSEB() // Seb Datei wird eingerichtet
	if err != nil {
		fmt.Println("Fehler beim Einrichten von SEB:", err)
	} else {
		fmt.Println("SEB Careum Konfiguration erfolgreich abgeschlossen.")
	}
	err = cleanup(downloadPath) // Datei wird gelöscht
	if err != nil {
		fmt.Println("Fehler beim Löschen der Datei:", err)
	}

}

// error muss zurückgegeben werden, kein error, nu
func cleanup(d string) error {

	err := os.Remove(d)
	if err != nil {
		return err
	}
	return nil
}

// Funktion zum Abrufen des neuesten Download-Links für SafeExamBrowser und gibt zwei Werte zurück
func getLatestSafeExamBrowserURL() (string, string, error) {
	url := downloadlink
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
