package files

import (
	"errors"
	"os"
	"strings"
)

func ReadAllFile(filename string)([]byte, error){
	data, err := os.ReadFile(filename)
	if err != nil{
		return nil, err
	}
	return data, nil
}
func CheckJsonExtension(filename string)([]byte, error){
	if !strings.HasSuffix(filename, ".json"){
		return nil, errors.New("ФАйл не json формата!")
	}
	data, err := ReadAllFile(filename)
	if err != nil{
		return nil, err
	}
	return data, nil
}