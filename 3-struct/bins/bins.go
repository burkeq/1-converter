package bins

import (
	"time"
	"encoding/json"
)

type Bin struct{
	Id string `json:"id"`
	Private bool `json:"private"`
	CreatedAt time.Time `json:"createdAt"`
	Name string `json:"name"`
}

func NewBin(id, name string, private bool, createdAt time.Time) *Bin{
	return &Bin{
		Id: id,
		Name: name,
		Private: private,
		CreatedAt: createdAt,
	}
}
func (bin *Bin) ToBytes()([]byte, error){
	data, err := json.Marshal(bin)
	if err != nil{
		return nil, err
	}
	return data, nil
}

type BinList struct{
	Bins []Bin `json:"bins"`
	UpdatedAt time.Time `json:"updatedAt"`
}
func NewBinList() *BinList{
	return &BinList{
		Bins: []Bin{},
		UpdatedAt: time.Now(),
	}
}
func (bins *BinList) ToBytes()([]byte, error){
	data, err := json.Marshal(bins)
	if err != nil{
		return nil, err
	}
	return data, nil
}
