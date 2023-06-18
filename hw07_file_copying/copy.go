package main

import (
	"errors"
	"io"
	"os"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	params, err := validate(fromPath, offset, limit)
	if err != nil {
		return err
	}
	file, err := os.OpenFile(fromPath, os.O_RDONLY, 0o666)
	if err != nil {
		return err
	}
	defer file.Close()

	if offset > 0 {
		_, err = file.Seek(params.offset, io.SeekStart)
		if err != nil {
			return err
		}
	}

	tmpFile, err := os.CreateTemp("", "temp.*")
	if err != nil {
		return err
	}

	err = copyFile(file, tmpFile, params.limit)
	if err != nil {
		return err
	}
	err = os.Rename(tmpFile.Name(), toPath)
	if err != nil {
		return err
	}
	return nil
}
