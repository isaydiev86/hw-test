package main

import "os"

type params struct {
	offset int64
	limit  int64
}

func validate(filePath string, offset int64, limit int64) (*params, error) {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return nil, err
	}

	if !fileInfo.Mode().IsRegular() {
		return nil, ErrUnsupportedFile
	}

	if fileInfo.Size() < offset {
		return nil, ErrOffsetExceedsFileSize
	}

	if limit > (fileInfo.Size() - offset) {
		limit = fileInfo.Size() - offset
	}

	return &params{
		offset: offset,
		limit:  limit,
	}, nil
}
