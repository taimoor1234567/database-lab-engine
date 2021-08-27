/*
2019 © Postgres.ai
*/

package models

// FileSystem describes state of a file system.
type FileSystem struct {
	Mode              string  `json:"mode"`
	SizeHR            string  `json:"sizeHR"`
	FreeHR            string  `json:"freeHR"`
	UsedHR            string  `json:"usedHR"`
	DataSizeHR        string  `json:"dataSizeHR"`
	UsedBySnapshotsHR string  `json:"usedBySnapshotsHR"`
	UsedByClonesHR    string  `json:"usedByClonesHR"`
	CompressRatio     float64 `json:"compressRatio"`
}
