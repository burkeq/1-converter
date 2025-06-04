package bins

import "time"

type Bin struct{
	id string
	private bool
	createdAt time.Time
	name string
}
func newBin(id, name string, private bool, createdAt time.Time) *Bin{
	return &Bin{
		id: id,
		name: name,
		private: private,
		createdAt: createdAt,
	}
}
var BinList = make([]Bin, 0, 20)