package services

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/mitchellh/go-homedir"
)

func Manager(path string) {
	var home string
	var err error
	if strings.Contains(path, "~") {
		home, err = homedir.Expand("~")
		path = strings.Replace(path, "~", home, 1)
	}

	if err != nil {
		log.Fatal(err)
	}

	files, err := getPath(path)

	extensions := extractExtensions(files)

	for ffEx := range FolderFileExtensions {
		ffExSlice := FolderFileExtensions[ffEx]

		for ext := range extensions {
			key, ok := mapkey(FolderFileExtensions, ffExSlice)
			if !ok {
				log.Fatalf("could not extract key")
			}
			if slices.Contains(ffExSlice, extensions[ext]) {
				fmt.Printf("Key: %s Extension: %s\n", key, extensions[ext])
				cool := arrangeFiles(key, path, extensions[ext])
				if !cool {
					panic("")
				}
			}
		}
	}
}

func getPath(path string) ([]string, error) {
	files := []string{}
	err := filepath.Walk(path,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			files = append(files, info.Name())
			return nil
		})
	if err != nil {
		log.Fatal(err)
	}

	return files, err
}

func arrangeFiles(key string, path string, extension string) bool {
	curPath := path + "/" + key
	_ = os.Mkdir(curPath, os.ModePerm)
	filesToMove, err := getPath(path)
	if err != nil {
		log.Fatalf("Error on getting path ~ %v", err)
		return false
	}

	for file := range filesToMove {
		ftm := filesToMove[file]
		if strings.Contains(ftm, extension) {
			if strings.Contains(ftm, "Images") {
				continue
			}
			originalPath := path + "/" + ftm
			newPath := curPath + "/" + ftm
			if err = os.Rename(originalPath, newPath); err != nil {
				// log.Fatalf("Path Err:: %v", err)
				continue
			}
		}
	}

	return true
}

func extractExtensions(files []string) []string {
	extensions := []string{}

	for file := range files {
		parts := strings.Split(files[file], ".")
		if len(parts) < 2 || parts[1] == "" {
			continue
		}

		extensions = append(extensions, "."+parts[1])

	}

	return extensions
}

func mapkey(m map[string][]string, value []string) (key string, ok bool) {
	for k, v := range m {
		if slices.Equal(v, value) {
			key = k
			ok = true
			return
		}
	}
	return
}
