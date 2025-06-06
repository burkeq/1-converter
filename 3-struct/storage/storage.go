package storage

import (
	"converter/app/3-struct/bins"
	"converter/app/3-struct/files"
	"encoding/json"
	"errors"
	"os"
)


func SaveBinToFile(bin bins.Bin, filename string) error {
	data, err := bin.ToBytes()
	if err != nil {
		return err
	}
	err = writeDataToFile(data, filename)
	if err != nil {
		return err
	}
	return nil
}
func SaveBinListToFile(binList bins.BinList, filename string) error {
	data, err := binList.ToBytes()
	if err != nil {
		return err
	}
	err = writeDataToFile(data, filename)
	if err != nil {
		return err
	}
	return nil
}

func writeDataToFile(data []byte, filename string) error {
	file, err := os.OpenFile(filename, os.O_TRUNC|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil

}

func ReadBinFromFile(filename string) (*bins.Bin, error) {
	if !files.CheckJsonExtension(filename) {
		return nil, errors.New("Файл не json")
	}
	data, err := files.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var bin bins.Bin
	err = json.Unmarshal(data, &bin)
	if err != nil {
		return nil, err
	}
	return &bin, nil
}

func ReadBinListFromFile(filename string) (*bins.BinList, error) {
	if !files.CheckJsonExtension(filename) {
		return nil, errors.New("Файл не json")
	}
	data, err := files.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var bins bins.BinList
	err = json.Unmarshal(data, &bins)
	if err != nil {
		return nil, err
	}
	return &bins, nil
}
