package storage

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
)

// Storage store a file to the storage
type Storage struct {
	config *Config
}

// PutFile to the storage
func (s *Storage) PutFile(dir string, file *multipart.FileHeader) error {
	return s.PutFileAs(dir, file, file.Filename)
}

// PutFileAs with costume filename to the storage
func (s *Storage) PutFileAs(dir string, file *multipart.FileHeader, filename string) error {
	if !s.Exists(dir) {
		err := s.MakeDir(dir)
		if err != nil {
			return err
		}
	}

	dirPath := fmt.Sprintf("%s/%s", s.config.Root, dir)

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(fmt.Sprintf("%s/%s", dirPath, filename))
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		return err
	}

	return nil
}

// Exists check dir or file is exists or not
func (s *Storage) Exists(path string) bool {
	fullPath := fmt.Sprintf("%s/%s", s.config.Root, path)
	_, err := os.Stat(fullPath)

	return !os.IsNotExist(err)
}

// URL to get url of file
func (s *Storage) URL(path string) string {
	return fmt.Sprintf("%s/%s", s.config.URL, path)
}

// MakeDir  make a new directory based on dir and config root dir
func (s *Storage) MakeDir(dir string) error {
	dirPath := fmt.Sprintf("%s/%s", s.config.Root, dir)

	return os.MkdirAll(dirPath, 0755)
}

// NewStorage to make instance of Storage
func NewStorage(config *Config) *Storage {
	return &Storage{config}
}
