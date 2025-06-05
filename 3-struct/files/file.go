package files

import (
	"errors"
	"os"
	"strings"
)

func ReadFile(filename string)([]byte, error){
	if !strings.Contains(filename, ".json"){
		return nil, errors.New("ФАйл не json формата!")
	}
	data, err := os.ReadFile(filename)
	if err != nil{
		return nil, err
	}
	return data, nil
}