package twitchapi

import (
	"fmt"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"errors"
	"time"
	log "github.com/sirupsen/logrus"
)

var baseUrl = "https://addons-ecs.forgesvc.net/api"
var MAX_TRIES = 5

func GetAddonInfo(addonID int) (info *AddonInfo, err error) {
	url := baseUrl + "/v2/addon/%v"
	url = fmt.Sprintf(url, addonID)
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}
	req.Header.Set("User-Agent", "Simple_mc_Modpack_downloader/1.0")
	var resp *http.Response

	for i:=1;i<MAX_TRIES;i++ {
		resp, err = client.Do(req)

		if err != nil {
			return nil, err
		}
		if resp.StatusCode == http.StatusOK {
			break
		}
		time.Sleep(time.Second*1)
		log.Warnf("Failed fetching info for %v retrying", addonID)
	}
	defer resp.Body.Close()
	if resp == nil || resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status + "when fetching "+url)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil { return }

	json.Unmarshal(data, &info)

	return
}
