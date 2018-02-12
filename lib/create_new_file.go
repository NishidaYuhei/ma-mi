package lib

import (
	"fmt"
	"os"
)

func CreateNewFile(filePath, writeStr string) {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	fmt.Fprintln(file, writeStr)
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
