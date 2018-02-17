package lib

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
)

func CreateNewFile(filePath, writeStr string) string {
	if !Exists(filePath) {
		file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			log.Println(err)
		}
		defer file.Close()
		fmt.Fprint(file, writeStr)
		log.Printf("Create new file %v\n", filePath)
	}
	return filePath
}

func RemoveFile(filePath string) {
	if err := os.Remove(filePath); err != nil {
		fmt.Println(err)
	}
	log.Printf("Remove success %v\n", filePath)
}

func Mkdir(filePath string) {
	if err := os.Mkdir(filePath, 0777); err != nil {
		fmt.Println(err)
	}
	log.Printf("Create new Dir %v\n", filePath)
}

func MkdirAll(filePath string) string {
	if !Exists(filePath) {
		if err := os.MkdirAll(filePath, 0777); err != nil {
			log.Println(err)
		}
		log.Printf("Create new Dir %v\n", filePath)
	}
	return filePath
}

func GetHomePath() string {
	usr, err := user.Current()
	if err != nil {
		log.Println(err)
	}
	return usr.HomeDir
}

func ReadFile(filePath string) string {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Println(err)
	}
	return string(data)
}

func Exists(fileName string) bool {
	_, err := os.Stat(fileName)
	return err == nil
}

func Getwd() string {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return wd
}
