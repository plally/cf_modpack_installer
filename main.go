package main

import (
	log "github.com/sirupsen/logrus"
	"os"
	"archive/zip"
	"path/filepath"
	"io"
	"io/ioutil"
	"encoding/json"
	"fmt"
)
var modsDirectory = "mods/"

func main() {
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.TextFormatter{ForceColors: true})
	log.SetOutput(os.Stdout)

	err := unzipFile("test/cfc_mc.zip", "temp/cursedownloader")
	if err != nil {
		log.Fatal(err)
	}

	manifest, err := loadManifest("temp/cursedownloader/manifest.json")
	log.Infof("Downloading mods from manifest.json %v", manifest.Name)
	fmt.Println(len(manifest.Files))
	m := ModDownloader{
		Manifest: manifest,
	}
	downloadUrlChannel := make(chan string)

	go m.FetchDownloadUrls(downloadUrlChannel)
	DownloadFromFilesChannel(downloadUrlChannel, "mods/")
}

func loadManifest(fpath string) (m *Manifest, err error){
	data, err := ioutil.ReadFile(fpath)
	if err != nil {return nil, err}
	err = json.Unmarshal(data, &m)
	return m, err
}

func unzipFile(location, dest string) (err error) {
	r, err := zip.OpenReader(location)
	if err != nil {return}
	defer r.Close()

	for _, f := range r.File {
		fpath := filepath.Join(dest, f.Name)
		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}
		err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm)
		if err != nil {
			return
		}
		rc, err := f.Open()
		if err != nil { return err }


		if err != nil { return err }

		out, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil { return err }
		io.Copy(out, rc)
		out.Close()
		rc.Close()
	}
	return
}
