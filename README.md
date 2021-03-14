# cf_modpack_installer
downloads all content given a modpack zip from curseforge 

You can download compiled binaries from the releases tab for OS X, Windows, or Linux. Alternatively you can build the project yourself with the go tool `go build .` in the root directory of this project.

#### command usage:
```
-installdir string
  The directory to create the mods and config director
-loglevel string
  (default "debug")
-modzip string
  Curseforge modpack zip file containing a manifest.json and overrides
-workers int
  amount of goroutines to use to download mod files (default 15)
```

#### examples
* `cf_modpack_installer --modzip "mc_modpack.zip"`
* `cf_modpack_installer --modzip "mc_modpack.zip" --installdir "/home/steve/.minecraft/" --workers 20`



