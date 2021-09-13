/*
2019 © Postgres.ai
*/

package models

// FileSystem describes state of a file system.
type FileSystem struct {
	Mode            string  `json:"mode"`
	Size            uint64  `json:"size"`
	Free            uint64  `json:"free"`
	Used            uint64  `json:"used"`
	DataSize        uint64  `json:"dataSize"`
	UsedBySnapshots uint64  `json:"usedBySnapshots"`
	UsedByClones    uint64  `json:"usedByClones"`
	CompressRatio   float64 `json:"compressRatio"`
}

// FileSystemView describes a view of file system state.
type FileSystemView struct {
	*FileSystem
	Size            Size `json:"size"`
	Free            Size `json:"free"`
	Used            Size `json:"used"`
	DataSize        Size `json:"dataSize"`
	UsedBySnapshots Size `json:"usedBySnapshots"`
	UsedByClones    Size `json:"usedByClones"`
}
