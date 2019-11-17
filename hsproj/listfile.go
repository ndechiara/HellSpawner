package hsproj

import (
	"path/filepath"
	"log"
)

type ListFile struct {
	Files   []ListFilePath
}

type ListFilePath struct{
	Name string 
	Path string 
}

func CreateListFileFromMpq(mpq *MpqInfo) *ListFile {
	// should return a blank one (as such) if none is found internally or if loading it errors
	result := ListFile{}
	result.Files = make([]ListFilePath, 0)

	// if the mpq has an existing lifefile internally, convert it to our listfile format
	names, err := mpq.Data.GetFileList()
	if err == nil {
		log.Printf("List file found inside MPQ '%s'", mpq.Name)
		// found the listfile, load it in
		for _, name := range names {
			newlfp := ListFilePath{
				filepath.Base(name),
				name,
			}
			result.Files = append(result.Files, newlfp)
		}
	} else {
		log.Printf("List file not found inside MPQ '%s'", mpq.Name)
	}

	return &result
}