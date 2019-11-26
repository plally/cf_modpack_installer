package main

import (
	"errors"
	"github.com/plally/curseforge_modpack_downloader/twitchapi"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"
)

type ModDownloader struct {
	*Manifest
}

func (m *ModDownloader) FetchDownloadUrls(ch chan string) {
	for _, manifestFile := range m.Files {
		url, err := m.GetFileURL(manifestFile)
		if err != nil {
			log.Fatal(err)
		}
		log.Debugf("fetched download url %v", url)
		ch <- url
	}
	log.Info("Finished fetching download urls ")
	time.Sleep(time.Second * 1)
	close(ch)

}

func (m *ModDownloader) GetFileURL(f *CurseForgeFile) (string, error) {
	url, err := twitchapi.GetDownloadUrl(f.ProjectID, f.FileID)
	return url, err
}

func DownloadFromFilesChannel(ch chan string, directory string) {
	os.MkdirAll(directory, os.ModePerm)
	for url := range ch {
		fpath := filepath.Join(directory, path.Base(url))
		log.Debugf("Downloading file %v to %v", url, fpath)
		DownloadFile(url, fpath)
	}
}

func DownloadFile(url, filepath string) (err error) {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
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
