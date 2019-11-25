package main

import (
	"net/http"
	"errors"
	"os"
	"github.com/plally/curseforge/twitchapi"
	"path"
	"io"
	log "github.com/sirupsen/logrus"
)

type ModDownloader struct {
	*Manifest
	FileUrlCache map[int]string
}

func (m *ModDownloader) getCachedUrl(addonID int) (string, bool) {
	url, ok := m.FileUrlCache[addonID]
	return url, ok
}

func (m *ModDownloader) setCachedUrl(addonID int, url string) {
	m.FileUrlCache[addonID] = url
}

func (m *ModDownloader) FetchDownloadUrls(ch chan string) {
	for _, file := range m.Files {
		info, err := twitchapi.GetAddonInfo(file.ProjectID)
		if err != nil {
			log.Error(err)
			continue
		}
		if _, ok := m.getCachedUrl(file.ProjectID); ok {
			continue
		}

		for _, f := range info.LatestFiles {
			if f.ID == file.FileID {
				m.setCachedUrl(file.ProjectID, f.DownloadURL)
				ch <- f.DownloadURL
				log.Debugf("Download url retrieved %v", f.DownloadURL)
				break
			}
		}
	}
	close(ch)
	log.Info("Finished fetching download urls")
}

func DownloadFromFilesChannel(ch chan string, directory string) {
	for url := range ch {
		DownloadFile(url, directory + path.Base(url))
	}
}

func DownloadFile(url, filepath string) (err error) {
	resp, err := http.Get(url)
	if err != nil { return err }
	if resp.StatusCode != http.StatusOK {
		return errors.New(resp.Status)
	}

	f, err := os.Create(filepath)

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	io.Copy(f, resp.Body)

	return nil
}
