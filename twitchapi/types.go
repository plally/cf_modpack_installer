package twitchapi

import "time"

type Dependency struct {
	Type    int `json:"type"`
	AddonID int `json:"addonId"`
	FileID int `json:"fileId"`
}

type Module struct {
	Foldername  string `json:"Foldername"`
	Fingerprint int    `json:"Fingerprint"`
}

type FileInfo struct {
	ID              int       `json:"id"`
	DisplayName     string    `json:"displayName"`
	FileName        string    `json:"fileName"`
	FileDate        time.Time `json:"fileDate"`
	FileLength      int       `json:"fileLength"`
	ReleaseType     int       `json:"releaseType"`
	FileStatus      int       `json:"fileStatus"`
	DownloadURL     string    `json:"downloadUrl"`
	IsAlternate     bool      `json:"isAlternate"`
	AlternateFileID int       `json:"alternateFileId"`
	Dependencies    []Dependency `json:"dependencies"`
	IsAvailable bool `json:"isAvailable"`
	Modules     []Module`json:"modules"`
	PackageFingerprint      int64       `json:"packageFingerprint"`
	GameVersion             []string    `json:"gameVersion"`
	HasInstallScript        bool        `json:"hasInstallScript"`
	GameVersionDateReleased time.Time   `json:"gameVersionDateReleased"`
}

type AddonInfo struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	WebsiteURL    string  `json:"websiteUrl"`
	GameID        int     `json:"gameId"`
	Summary       string  `json:"summary"`
	DefaultFileID int     `json:"defaultFileId"`
	DownloadCount float64 `json:"downloadCount"`
	LatestFiles   []FileInfo`json:"latestFiles"`
	Status            int `json:"status"`
	PrimaryCategoryID int `json:"primaryCategoryId"`
	Slug                   string `json:"slug"`
	IsFeatured         bool      `json:"isFeatured"`
	PopularityScore    float64   `json:"popularityScore"`
	GamePopularityRank int       `json:"gamePopularityRank"`
	PrimaryLanguage    string    `json:"primaryLanguage"`
	GameSlug           string    `json:"gameSlug"`
	GameName           string    `json:"gameName"`
	PortalName         string    `json:"portalName"`
	DateModified       time.Time `json:"dateModified"`
	DateCreated        time.Time `json:"dateCreated"`
	DateReleased       time.Time `json:"dateReleased"`
	IsAvailable        bool      `json:"isAvailable"`
	IsExperiemental    bool      `json:"isExperiemental"`
}