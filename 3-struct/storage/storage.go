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
func SaveBinListToFile(binList bins.BinList) error{
	data, err := binList.ToBytes()
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
	file, err := os.OpenFile(storageFileName, os.O_WRONLY, 0644)
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

func ReadBinListFromFile()(*bins.BinList, error){
	data, err := files.ReadFile(storageFileName)
	if err != nil{
		return nil, err
	}
	var bins bins.BinList
	err = json.Unmarshal(data, &bins)
	if err != nil{
		return nil, err
	}
	return &bins, nil
}