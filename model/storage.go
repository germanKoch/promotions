package model

import "time"

type FileLiner interface {
	ReadNext() bool
	NextLine() string
	Close() error
}

type FileData interface {
	Path() string
	ModificationDate() time.Time
	Content() (FileLiner, error)
}

type Storage interface {
	Walk(process func(file FileData))
}
