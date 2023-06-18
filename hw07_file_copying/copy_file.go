package main

import (
	"errors"
	"io"

	"github.com/cheggaaa/pb/v3"
)

func copyFile(src io.Reader, dt io.Writer, limit int64) error {
	pBar := pb.Start64(limit)
	defer pBar.Finish()

	var total int64
	for {
		n, err := io.CopyN(dt, src, 1)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			return err
		}

		pBar.Increment()

		total += n
		if total == limit {
			break
		}
	}

	return nil
}
