package model

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

type FileItem struct {
	Name string
	Path string
	IsDir bool
}

type FileModel struct {
	CurrentPath string
	Files       []FileItem
}

func NewFileModel(start string) (*FileModel, error) {
	if start == "" {
		start = os.Getenv("HOME")
	}
	files, err := readDir(start)
	if err != nil {
		return nil, err
	}
	return &FileModel{CurrentPath: start, Files: files}, nil
}

func (m *FileModel) EnterDir(subdir string) error {
	newPath := filepath.Join(m.CurrentPath, subdir)
	files, err := readDir(newPath)
	if err != nil {
		return err
	}
	m.CurrentPath = newPath
	m.Files = files
	return nil
}

func (m *FileModel) GoUp() error {
	parent := filepath.Dir(m.CurrentPath)
	files, err := readDir(parent)
	if err != nil {
		return err
	}
	m.CurrentPath = parent
	m.Files = files
	return nil
}

func (m *FileModel) Delete(target string) error {
	return os.RemoveAll(filepath.Join(m.CurrentPath, target))
}

func (m *FileModel) Create(name string, dir bool) error {
	path := filepath.Join(m.CurrentPath, name)
	if dir {
		return os.Mkdir(path, 0755)
	}
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return nil
}

func readDir(path string) ([]FileItem, error) {
	list, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	files := make([]FileItem, 0, len(list))
	for _, f := range list {
		files = append(files, FileItem{
			Name:  f.Name(),
			Path:  filepath.Join(path, f.Name()),
			IsDir: f.IsDir(),
		})
	}
	return files, nil
}
