package twitchapi

import (
	"strings"
	"testing"
)

func TestGetDownloadUrl(t *testing.T) {
	fileInfo := []struct{
		addonId int
		fileid int
	}{
		{261251, 2638317},
		{248453, 2727070},
		{228404, 2774057},
	}

	for _, info := range fileInfo {
		fileUrl, err  := GetDownloadUrl(info.addonId, info.fileid)
		if err != nil {t.Error(err)}

		if !strings.HasSuffix(fileUrl, ".jar") {
			t.Errorf("%v, %v: file url (%v) has no .jar suffix", info.addonId, info.fileid, fileUrl)
		}
	}
}
