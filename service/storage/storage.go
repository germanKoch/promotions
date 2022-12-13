package storage

import (
	"bufio"
	"os"

	"path/filepath"
	"time"
)

type FileLiner interface {
	HasNext() bool
	NextLine() string
	Close() error
}

type FileData interface {
	Path() string
	ModificationDate() time.Time
	Content() FileLiner
}

type Storage interface {
	Walk(process func(file FileData))
}

type LocalFileLiner struct {
	file    *os.File
	scanner *bufio.Scanner
}

type LocalFileData struct {
	path string
	file os.FileInfo
}

type LocalStorage struct {
	monitoredDirectory string
}

func (fileLiner LocalFileLiner) HasNext() bool {
	return fileLiner.scanner.Scan()
}

func (fileLiner LocalFileLiner) NextLine() string {
	s := fileLiner.scanner.Text()
	return s
}

func (fileLiner LocalFileLiner) Close() error {
	return fileLiner.file.Close()
}

func (fileData LocalFileData) Path() string {
	return fileData.path
}

func (fileData LocalFileData) ModificationDate() time.Time {
	return fileData.file.ModTime()
}

func (fileData LocalFileData) Content() FileLiner {
	file, err := os.Open(fileData.path)
	if err != nil {
		//TODO: error handling
	}
	scanner := bufio.NewScanner(file)
	return LocalFileLiner{
		file:    file,
		scanner: scanner,
	}
}

func GetLocalStorage(monitoredDirectory string) Storage {
	return LocalStorage{
		monitoredDirectory: monitoredDirectory,
	}
}

func (storage LocalStorage) Walk(process func(fileData FileData)) {
	filepath.Walk(storage.monitoredDirectory, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			process(
				LocalFileData{
					path: path,
					file: info,
				},
			)
		}
		return nil
	})
}
