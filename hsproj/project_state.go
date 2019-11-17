package hsproj

import (
	"encoding/json"
	"io/ioutil"
	"errors"
	"path/filepath"
	"log"
)

var ActiveProject *ProjectState

func SetDefaultActiveProject() {
	ActiveProject = GetEmptyProjectState()
}

type ProjectState struct {
	// identification
	Name       string
	Version    int
	FolderPath string

	// state
	Loaded         bool
	UnsavedChanges bool

	// data
	MpqList *MpqList
}

func GetEmptyProjectState() *ProjectState {
	result := ProjectState{}
	result.Loaded = false
	result.Name = "Untitled"
	result.UnsavedChanges = false
	return &result
}

// the savable state is just the info that gets saved to the odproj file
type projectStateSavable struct {
	Name    string
	Version int
}

func (s *ProjectState) getSavable() projectStateSavable {
	return projectStateSavable{
		s.Name,
		s.Version,
	}
}

func (s *projectStateSavable) getStateFromSavable() *ProjectState {
	result := GetEmptyProjectState()
	result.Name = s.Name
	result.Version = s.Version
	return result;
}

// errors
var errProjectStateNotLoaded = errors.New("no project folder is loaded")


func LoadProjectStateFromFolder(folderpath string) (*ProjectState, error) {
	// check if the folderpath contains an odproj file, and if it does
	// load from that instead 
	testState, err := LoadProjectStateFromProj(filepath.Join(folderpath, "odproj.json"))
	if err == nil {
		return testState, nil
	}

	result := GetEmptyProjectState()
	result.Loaded = true
	result.FolderPath = folderpath
	err = result.postLoad()
	if err != nil {
		return nil, err
	}

	log.Printf("Project '%s' loaded from '%s'", result.Name, result.FolderPath);

	return result, nil
}

// load from a odproj file
func LoadProjectStateFromProj(projpath string) (*ProjectState, error) {
	odprojJSON, err := ioutil.ReadFile(projpath)
	if err != nil {
		return nil, err
	}

	var psav projectStateSavable
	err = json.Unmarshal(odprojJSON, &psav)
	if err != nil {
		return nil, err
	}

	result := psav.getStateFromSavable()
	result.Loaded = true
	result.FolderPath = filepath.Dir(projpath)
	err = result.postLoad()
	if err != nil {
		return nil, err
	}

	log.Printf("Project '%s' loaded from '%s'", result.Name, result.FolderPath);

	return result, nil
}

func (s *ProjectState) postLoad() error {
	// use this to trigger load for other parts of the project

	// load mpq list
	mpqlist, err := LoadMpqList(s.FolderPath)
	if err != nil {
		return err
	}
	s.MpqList = mpqlist

	return nil
}

func (s *ProjectState) PromptUnsavedChanges() {
	if !s.UnsavedChanges {
		return
	}

	// TODO: prompt the unsaved changes dialog
}

func (s *ProjectState) Save() error {
	if !s.Loaded {
		return errProjectStateNotLoaded
	}

	// TODO: check if folderpath is valid

	s.Version += 1

	// save the .odproj file
	odproj, err := json.MarshalIndent(s.getSavable(), "", " ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filepath.Join(s.FolderPath, "odproj.json"), odproj, 0644)
	if err != nil {
		return err
	}

	// TODO: call actual saving code for subcomponents...

	// save the mpqs
	s.MpqList.Save(s.FolderPath)

	log.Printf("Project '%s' saved in '%s'", s.Name, s.FolderPath);
	s.UnsavedChanges = false
	return nil
}

func (s *ProjectState) SaveAs(newpath string) error {
	// TODO: check for invalid paths

	s.FolderPath = newpath
	return s.Save()
}

func (s *ProjectState) Rename(newname string) error {
	// TOOD: check for invalid names

	s.Name = newname
	s.UnsavedChanges = true
	return nil
}

func (s *ProjectState) Close() {
	// discard the project
}