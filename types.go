package main

type ModLoader struct {
	ID      string `json:"id"`
	Primary bool   `json:"primary"`
}

type MinecraftInfo struct {
	Version    string      `json:"version"`
	ModLoaders []ModLoader `json:"modLoaders"`
}
type CurseForgeFile struct {
	ProjectID int  `json:"projectID"`
	FileID    int  `json:"fileID"`
	Required  bool `json:"required"`
}

type Manifest struct {
	Minecraft       MinecraftInfo     `json:"minecraft"`
	ManifestType    string            `json:"manifestType"`
	ManifestVersion int               `json:"manifestVersion"`
	Name            string            `json:"name"`
	Version         string            `json:"version"`
	Author          string            `json:"author"`
	Files           []*CurseForgeFile `json:"files"`
	Overrides       string            `json:"overrides"`
}
