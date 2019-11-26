package main

import (
	"archive/zip"
	"path/filepath"
	"os"
	"io"
	"io/ioutil"
	"fmt"
)

func unzipFile(location, dest string) (err error) {
	fmt.Println(location)
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

func copyDirectory(src, dest string) error {
	entries, err := ioutil.ReadDir(src)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		srcpath := filepath.Join(src, entry.Name())
		destpath := filepath.Join(dest, entry.Name())

		if entry.IsDir() {
			err = copyDirectory(srcpath, destpath)
			if err != nil { return err }
		} else {
			err = copy(srcpath, destpath)
			if err != nil { return err }
		}

	}
	return nil
}

func copy(src, dest string) error{
	os.MkdirAll(filepath.Dir(dest), os.ModePerm)
	newFile, err  := os.Create(dest)
	if err != nil {
		return err
	}
	original, err := os.Open(src)
	defer original.Close()
	_, err = io.Copy(newFile, original)

	return err

}