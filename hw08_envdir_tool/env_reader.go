package main

import (
	"bufio"
	"errors"
	"io"
	"io/fs"
	"log"
	"os"
	"path"
	"strings"
)

var ErrorInvalidFilename = errors.New("invalid filename")

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	fileInfos, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	env := make(Environment, len(fileInfos))

	for _, fInfo := range fileInfos {
		if fInfo.IsDir() {
			continue
		}

		if strings.Contains(fInfo.Name(), "=") {
			return nil, ErrorInvalidFilename
		}

		val, err := getValueFromFile(path.Join(dir, fInfo.Name()))
		if err != nil {
			return nil, err
		}

		var info fs.FileInfo
		info, err = fInfo.Info()
		if err != nil {
			return nil, err
		}

		env[fInfo.Name()] = EnvValue{
			Value:      val,
			NeedRemove: info.Size() == 0,
		}
	}

	return env, nil
}

func getValueFromFile(fullName string) (string, error) {
	f, err := os.Open(fullName)
	if err != nil {
		return "", err
	}
	defer func() {
		err = f.Close()
		if err != nil {
			log.Println("failed to close file: %w", err)
		}
	}()

	buf := bufio.NewReader(f)
	line, err := buf.ReadBytes('\n')
	if err != nil && !errors.Is(err, io.EOF) {
		return "", err
	}

	val := strings.ReplaceAll(string(line), "\x00", "\n")
	val = strings.TrimRight(val, " \t\n")

	return val, nil
}
