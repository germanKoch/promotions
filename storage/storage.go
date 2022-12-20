package storage

import (
	"bufio"
	"os"
	"path/filepath"
	"promotions/config"
	"promotions/model"
	"time"
)

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

func (fileLiner LocalFileLiner) ReadNext() bool {
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

func (fileData LocalFileData) Content() (model.FileLiner, error) {
	file, err := os.Open(fileData.path)
	if err != nil {
		return LocalFileLiner{}, err
	}
	scanner := bufio.NewScanner(file)
	return LocalFileLiner{
		file:    file,
		scanner: scanner,
	}, nil
}

func GetLocalStorage(localStorageConfig config.LocalStorageConfig) model.Storage {
	return LocalStorage{
		monitoredDirectory: localStorageConfig.MonitoredDirectory,
	}
}

func (storage LocalStorage) Walk(process func(fileData model.FileData)) {
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
