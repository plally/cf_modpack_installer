package twitchapi

import (
	"encoding/json"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"time"
)

var baseUrl = "https://addons-ecs.forgesvc.net/api"
var MAX_TRIES = 5

func get(url string) (resp *http.Response, err error) {
	url = baseUrl + url
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}
	req.Header.Set("User-Agent", "Simple_mc_Modpack_downloader/1.0")
	for i := 1; i < MAX_TRIES; i++ {
		resp, err = client.Do(req)
		if err != nil {
			return nil, err
		}
		if resp.StatusCode == http.StatusOK {
			break
		}
		time.Sleep(time.Second * 1)
		log.Warnf("http request GET %v failed retrying", url)
	}
	return resp, err
}

func GetDownloadUrl(addonID, fileID int) (fileUrl string, err error) {
	url := fmt.Sprintf("/v2/addon/%v/file/%v/download-url", addonID, fileID)
	resp, err := get(url)
	defer resp.Body.Close()
	if err != nil {
		return
	}
	if resp == nil || resp.StatusCode != http.StatusOK {
		return "", errors.New(resp.Status + "when fetching " + url)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	return string(data), nil
}

func GetAddonInfo(addonID int) (info *AddonInfo, err error) {
	url := fmt.Sprintf("/v2/addon/%v", addonID)
	resp, err := get(url)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	if resp == nil || resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status + "when fetching " + url)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	json.Unmarshal(data, &info)

	return
}
