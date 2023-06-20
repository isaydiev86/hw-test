package main

import (
	"errors"
	"io"

	"github.com/cheggaaa/pb/v3"
)

func copyFile(src io.Reader, dt io.Writer, limit int64) error {
	pBar := pb.Start64(limit)
	defer pBar.Finish()

	var lr io.Reader

	if limit > 0 {
		lr = io.LimitReader(src, limit)
	} else {
		lr = src
	}
	_, err := io.Copy(dt, lr)
	if err != nil {
		if errors.Is(err, io.EOF) {
			return nil
		}
		return err
	}

	return nil
}
