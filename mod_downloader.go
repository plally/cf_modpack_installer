package main

import (
	"errors"
	"github.com/plally/cf_modpack_installer/twitchapi"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"sync"
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

func DownloadFilesFromChannel(ch chan string, directory string, workers int) {
	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go DownloadFilesWorker(ch, directory, &wg)
	}
	wg.Wait()
}

func DownloadFilesWorker(ch chan string, directory string, wg *sync.WaitGroup) {
	os.MkdirAll(directory, os.ModePerm)
	for url := range ch {
		fpath := filepath.Join(directory, path.Base(url))
		log.Debugf("Downloading file %v to %v", url, fpath)
		DownloadFile(url, fpath)
	}
	wg.Done()
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
