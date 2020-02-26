package main

import (
	"encoding/json"
	"flag"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

var (
	zipLocation  = flag.String("modzip", "", "Curseforge modpack zip file containing a manifest.json and overrides")
	installDir   = flag.String("installdir", "", "The directory to create the mods and config director")
	logLevel     = flag.String("loglevel", "debug", "")
	workerAmount = flag.Int("workers", 15, "amount of goroutines to use to download mod files")
)

func main() {
	startTime  := time.Now()


	flag.Parse()

	level, err := log.ParseLevel(*logLevel)
	if err != nil {
		level = log.DebugLevel
	}

	log.SetLevel(level)
	log.SetFormatter(&log.TextFormatter{ForceColors: true})
	log.SetOutput(os.Stdout)

	tempPath      := filepath.Join(*installDir, "temp", "cursedownloader")
	manifestPath  := filepath.Join(tempPath, "manifest.json")
	overridesPath := filepath.Join(tempPath, "overrides")
	modsPath      := filepath.Join(*installDir, "mods")

	log.Info("Unzipping")
	err = unzipFile(*zipLocation, tempPath)

	if err != nil {
		log.Fatal(err)
	}

	err = copyDirectory(overridesPath, *installDir)
	if err != nil {
		log.Fatal(err)
	}

	manifest, err := loadManifest(manifestPath)
	if err != nil { log.Fatal(err) }

	log.Infof("Downloading mods from manifest.json %v", manifest.Name)
	log.Infof("Using %v goroutines", *workerAmount)

	m := ModDownloader{
		Manifest: manifest,
	}

	downloadUrlChannel := make(chan string)

	var done = make(chan bool)
	go func() {
		defer func() { close(done) }()
		m.FetchDownloadUrls(downloadUrlChannel)
	}()

	DownloadFilesFromChannel(downloadUrlChannel, modsPath, *workerAmount)
	<- done
	log.Infof("Finished installing modpack, took %v", time.Since(startTime))
}

func loadManifest(fpath string) (m *Manifest, err error) {
	data, err := ioutil.ReadFile(fpath)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &m)
	return m, err
}
