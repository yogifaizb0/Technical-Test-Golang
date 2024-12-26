package services

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

const uploadDir = "./uploads"

func UploadFile(file multipart.File, header *multipart.FileHeader) error {
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		return err
	}

	dest, err := os.Create(filepath.Join(uploadDir, header.Filename))
	if err != nil {
		return err
	}

	if _, err := io.Copy(dest, file); err != nil {
		return err
	}
	return nil
}

func GetAllUploadedFiles() ([]map[string]string ,error) {
	dirEntries, err := os.ReadDir(uploadDir)
	if err != nil {
		return []map[string]string{}, err
	}

	files := []map[string]string{}
	for _, dirEntry := range dirEntries {
		if !dirEntry.IsDir() {
			files = append(files, map[string]string{
				"filename": dirEntry.Name(),
			})
		}
	}
	return files, nil
}

func CheckFile(filename string) (string, error) {
	filepath := filepath.Join(uploadDir, filename)
	_, err := os.Open(filepath)
	if err != nil {
		return "", err
	}
	return filepath, nil
}