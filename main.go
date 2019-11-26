package main

import (
	"encoding/json"
	"flag"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	zipLocation := flag.String("modzip", "", "Curseforge modpack zip file containing a manifest.json and override")
	installDir := flag.String("installdir", "", "The directory to create the mods and config director")
	logLevel := flag.String("loglevel", "debug", "")
	flag.Parse()

	level, err := log.ParseLevel(*logLevel)
	if err != nil {
		level = log.DebugLevel
	}
	log.SetLevel(level)
	log.SetFormatter(&log.TextFormatter{ForceColors: true})
	log.SetOutput(os.Stdout)

	tempPath := filepath.Join(*installDir, "temp/cursedownloader/")
	manifestPath := filepath.Join(tempPath, "manifest.json")

	log.Info("Unzipping")
	err = unzipFile(*zipLocation, tempPath)

	if err != nil {
		log.Fatal(err)
	}

	err = copyDirectory(filepath.Join(tempPath, "overrides/"), *installDir)
	if err != nil {
		log.Fatal(err)
	}
	manifest, err := loadManifest(manifestPath)
	log.Infof("Downloading mods from manifest.json %v", manifest.Name)

	m := ModDownloader{
		Manifest: manifest,
	}
	downloadUrlChannel := make(chan string)

	go m.FetchDownloadUrls(downloadUrlChannel)
	DownloadFromFilesChannel(downloadUrlChannel, filepath.Join(*installDir, "mods"))
}

func loadManifest(fpath string) (m *Manifest, err error) {
	data, err := ioutil.ReadFile(fpath)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &m)
	return m, err
}
