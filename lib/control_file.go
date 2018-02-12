package lib

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
)

func CreateNewFile(filePath, writeStr string) {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	fmt.Fprint(file, writeStr)
	fmt.Printf("config write %v\n", writeStr)

}

func RemoveFile(filePath string) {
	if err := os.Remove(filePath); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v remove success\n", filePath)
}

func Mkdir(filePath string) {
	if err := os.Mkdir(filePath, 0777); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v create success\n", filePath)
}

func MkdirAll(filePath string) {
	if err := os.MkdirAll(filePath, 0777); err != nil {
		log.Println(err)
	}
	log.Printf("%v create success\n", filePath)
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

func MakeDateDir() {
	MkdirAll(ReadFile(GetHomePath()+"/.ma-mi/config") + "/make_minutes/" + GetNowYearAndMonth())
}
