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

var RepoPath string

func CheckRepo(path string) {
	RepoPath = path
	checkPath(RepoPath)
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
	result := strings.Replace(fullPath[len(RepoPath):], "\\", "/", -1)
	if !strings.HasPrefix(result, "/") {
		result = "/" + result
	}
	return result
}

func getCheckSum(fullPath string) string {
	f, err := os.OpenFile(fullPath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatal("getCheckSum ", err)
	}
	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal("getCheckSum md5 ", err)
	}
	res := fmt.Sprintf("%x", h.Sum(nil))
	res = strings.Replace(res, "-", "", -1)
	return strings.ToUpper(res)
}

func GetResourceByFullPath(path string, name string) resources.Resource {
	return resources.Resource{
		FullPath:     path,
		RelativePath: removeExtraFromPath(path),
		Checksum:     getCheckSum(path),
		ResourceName: name,
	}
}

func GetResourceWithoutChecksum(path string, name string) resources.Resource {
	return resources.Resource{
		FullPath:     path,
		RelativePath: removeExtraFromPath(path),
		ResourceName: name,
	}
}

func GetAllFilesWithChecksum() []resources.Resource {
	var result []resources.Resource

	err := filepath.Walk(RepoPath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if info.IsDir() {
				return nil
			}
			resource := GetResourceByFullPath(path, info.Name())
			result = append(result, resource)
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

func WriteResource(data []byte) string {
	var fileNameLen = ConvertByteArrayToInt32(data[0:4])

	var skipLength = fileNameLen + 4
	var path = string(data[4:skipLength])

	var fullPath = RepoPath + path

	directoryOfFile := filepath.Dir(fullPath)
	checkPath(directoryOfFile)

	log.Printf("Writing file: %s\n", fullPath)

	os.WriteFile(fullPath, data[skipLength:], os.ModePerm)

	return fullPath
}
