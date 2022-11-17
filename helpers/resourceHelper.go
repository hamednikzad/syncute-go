package helpers

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"syncute-go/messages/resources"
)

var repoPath = "C:/SyncRepo/"

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func removeExtraFromPath(fullPath string) string {
	result := strings.Replace(fullPath[len(repoPath):], "\\", "/", -1)
	if !strings.HasPrefix(result, "/") {
		result = "/" + result
	}
	return result
}

func getCheckSum(fullPath string) string {
	f, err := os.Open(fullPath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}
	res := fmt.Sprintf("%x", h.Sum(nil))
	res = strings.Replace(res, "-", "", -1)
	return strings.ToUpper(res)
}

func GetAllFilesWithChecksum() []resources.Resource {
	var result []resources.Resource

	err := filepath.Walk(repoPath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			//fmt.Println("Walk ", info.IsDir(), path)
			if info.IsDir() {
				return nil
			}
			resource := resources.Resource{
				FullPath:     path,
				RelativePath: removeExtraFromPath(path),
				Checksum:     getCheckSum(path),
				ResourceName: info.Name(),
			}
			result = append(result, resource)
			fmt.Println(resource)
			return nil
		})
	if err != nil {
		log.Println(err)
	}

	return result
}

func CheckRepository() {
	ex, err := exists(repoPath)
	if err != nil {
		panic(err)
	}
	if !ex {
		err := os.Mkdir(repoPath, os.ModePerm)
		if err != nil {
			log.Println(err)
			panic(err)
		}
	}
}
