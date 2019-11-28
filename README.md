# cf_modpack_installer
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
* `cf_modpack_installer --modzip "mc_modpack.zip" --installdir "/home/name/.minecraft/" --workers 20`



