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

var repoPath string

func CheckRepo(path string) {
	repoPath = path
	checkPath(repoPath)
}

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

func checkPath(path string) {
	path = strings.Replace(path, "\\", "/", -1)
	ex, err := exists(path)
	if err != nil {
		panic(err)
	}
	if !ex {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			log.Println(err)
			panic(err)
		}
	}
}

func WriteResource(data []byte) {
	var fileNameLen = ConvertByteArrayToInt32(data[0:4])

	var skipLength = fileNameLen + 4
	var path = string(data[4:skipLength])

	var fullPath = repoPath + path

	directoryOfFile := filepath.Dir(fullPath)
	checkPath(directoryOfFile)

	fmt.Printf("Writing file: %s\n", fullPath)

	os.WriteFile(fullPath, data[skipLength:], os.ModePerm)
}
