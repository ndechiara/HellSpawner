package hsproj

import (
	"strings"
	"io/ioutil"
	"path/filepath"
	"log"
	"sync"
	"errors"

	"github.com/OpenDiablo2/HellSpawner/hsutil"

	"github.com/OpenDiablo2/OpenDiablo2/d2data/d2mpq"
)

type MpqList struct {
	Mpqs []MpqInfo
}

type MpqInfo struct {
	Data     *d2mpq.MPQ
	ListFile *ListFile
}

var errMpqInfoNotFound = errors.New("mpq not found in mpq list")

var mpqMutex = sync.Mutex{}

func LoadMpqList(folderpath string) (*MpqList, error) {
	result := MpqList{}
	err := result.Populate(folderpath)
	if err != nil {
		return nil, err
	}

	return &result, err
}

func (v *MpqList) Populate(folderpath string) error {
	v.Mpqs = make([]MpqInfo, 0)
	// search the folder for Mpqs
	files, err := ioutil.ReadDir(folderpath)
    if err != nil {
        return err
	}
	
	mpqMutex.Lock()
	defer mpqMutex.Unlock()

    for _, file := range files {
        if filepath.Ext(file.Name()) == ".mpq" {
			newinfo := MpqInfo{}
			archive, archiveErr := d2mpq.Load(filepath.Join(folderpath, file.Name()))
			if archiveErr != nil {
				log.Printf("Could not load MPQ '%s'", file.Name())
				log.Println(archiveErr)
				continue
			} 
			log.Printf("Loaded MPQ '%s'", file.Name())
			newinfo.Data = archive

			// now try to load the listfile
			lf, lfErr := LoadListFile(folderpath, file.Name())
			if lfErr != nil {
				// couldn't load the listfile
				lf = CreateListFileFromMpq(&newinfo)
			}
			newinfo.ListFile = lf

			v.Mpqs = append(v.Mpqs, newinfo)
		}
	}
	return nil
}

func (v *MpqList) FindMpq(mpqname string) *MpqInfo {
	for _, m := range v.Mpqs {
		if m.Data.FileName == mpqname {
			return &m
		}
	}
	return nil
}

func (v *MpqList) LoadFile(mpqpath hsutil.MpqPath) ([]byte, error) {
	minfo := v.FindMpq(mpqpath.MpqName)
	if minfo == nil {
		return nil, errMpqInfoNotFound
	}
	return minfo.LoadFile(mpqpath.FilePath)
}

func (v *MpqInfo) LoadFile(fileName string) ([]byte, error) {
	//fileName = strings.ReplaceAll(fileName, "{LANG}", d2resource.LanguageCode)
	fileName = strings.ToLower(fileName)
	fileName = strings.ReplaceAll(fileName, `/`, "\\")
	if fileName[0] == '\\' {
		fileName = fileName[1:]
	}
	mpqMutex.Lock()
	defer mpqMutex.Unlock()

	result, err := v.Data.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (v *MpqList) Save(folderpath string) error {
	for _, m := range v.Mpqs {
		err := m.Save(folderpath)
		if err != nil {
			return err
		}
	}
	return nil
}

func (v *MpqInfo) Save(folderpath string) error {
	// save the listfile
	err := v.ListFile.Save(folderpath)
	if err != nil {
		return err
	}

	// TODO: save the mpq data itself
	return nil
}