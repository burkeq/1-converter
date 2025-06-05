package storage

import (
	"converter/app/3-struct/bins"
	"converter/app/3-struct/files"
	"encoding/json"
	"os"
)
const storageFileName = "storage.json"
func SaveBinToFile(bin bins.Bin) error{
	data, err := bin.ToBytes()
	if err != nil{
		return err
	}
	err = writeDataToFile(data)
	if err != nil{
		return err
	}
	return nil
}

func writeDataToFile(data []byte) error{
	file, err := os.Open(storageFileName)
	if err != nil{
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil{
		return err
	}

	return nil

}

func ReadBinFromFile()(*bins.Bin, error){
	data, err := files.ReadFile(storageFileName)
	if err != nil{
		return nil, err
	}
	var bin bins.Bin
	err = json.Unmarshal(data, &bin)
	if err != nil{
		return nil, err
	}
	return &bin, nil
}