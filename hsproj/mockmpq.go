package hsproj

import (
	"path/filepath"
)

type MockMPQ struct {
	FileName string
}

func LoadMockMPQ(path string) (*MockMPQ, error) {
	//folderpath := filepath.Dir(path)
	filename := filepath.Base(path)
	return &MockMPQ{filename}, nil
}

func (v *MockMPQ) ReadFile(name string) ([]byte, error) {
	b := make([]byte, 10)
	return b, nil
}

func (v *MockMPQ) GetFileList() ([]string, error) {
	names := make([]string, 0)
	names = append(names, "armor.txt")
	names = append(names, "weapons.txt")
	names = append(names, "misc.txt")
	return names, nil
}