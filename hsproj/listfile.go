package hsproj

import (
	"encoding/json"
	"path/filepath"
	"strings"
	"io/ioutil"
)

type ListFile struct {
	MpqName string
	Files   []ListFilePath
}

type ListFilePath struct{
	DisplayName string
	FilePath    string
}

func getListFilePath(folderpath string, mpqname string) string {
	return filepath.Join(folderpath, strings.ReplaceAll(mpqname, ".mpq", "_mpq_listfile.json"))
}

func (v *ListFile) Save(folderpath string) error {
	lf, err := json.MarshalIndent(v, "", " ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(getListFilePath(folderpath, v.MpqName), lf, 0644)
	if err != nil {
		return err
	}
	return nil
}

func LoadListFile(folderpath string, mpqname string) (*ListFile, error) {
	lfJSON, err := ioutil.ReadFile(getListFilePath(folderpath, mpqname))
	if err != nil {
		return nil, err
	}

	var plf ListFile
	err = json.Unmarshal(lfJSON, &plf)
	if err != nil {
		return nil, err
	}

	return &plf, nil
}

func CreateListFileFromMpq(mpq *MpqInfo) *ListFile {
	// should return a blank one (as such) if none is found internally or if loading it errors
	result := ListFile{}
	result.MpqName = mpq.Data.FileName
	result.Files = make([]ListFilePath, 0)

	// if the mpq has an existing lifefile internally, convert it to our listfile format
	names, err := mpq.Data.GetFileList()
	if err == nil {
		// found the listfile, load it in
		for _, name := range names {
			newlfp := ListFilePath{
				name,
				name,
			}
			result.Files = append(result.Files, newlfp)
		}
	}

	return &result
}