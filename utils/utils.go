package utils

import (
	"errors"
	"github.com/mitchellh/go-homedir"
	"github.com/paulvollmer/go-concatenate"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type Validate func(dir string) bool

func FindDir(possibles []string, data *Context, validate Validate) (string, error) {
	var err error
	for _, dir := range possibles {
		if data != nil {
			dir, err = Execute(dir, data)
			if err != nil {
				continue
			}
		}
		if Exists(dir) && validate(dir) {
			return dir, nil
		}
	}
	return "", errors.New("not found files or directories from [" + strings.Join(possibles, ", ") + "]")
}

func Exists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func ConcatFiles(dest string, del string, files ...string) error {
	return concatenate.FilesToFile(dest, 0644, del, files...)
}

func ConcatCerts(dest string, files ...string) error {
	return ConcatFiles(dest, "", files...)
}

func CopyFile(srcFile, destFile string) error {
	file, err := os.Open(srcFile)
	if err != nil {
		return err
	}
	defer file.Close()

	dest, err := os.Create(destFile)
	if err != nil {
		return err
	}
	defer dest.Close()

	io.Copy(dest, file)

	return nil
}

func ResolvePath(path string) (string, error) {
	answer, err := homedir.Expand(path)
	if err != nil {
		return "", nil
	}
	if answer == "" {
		return answer, nil
	}
	answer, err = filepath.Abs(answer)
	if err != nil {
		return "", nil
	}
	return answer, nil
}
