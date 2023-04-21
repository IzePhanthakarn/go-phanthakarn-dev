package fileutil

import (
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path/filepath"
)

const (
	// ContentTypeExcel excel
	ContentTypeExcel = "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
	// ContentTypePDF pdf
	ContentTypePDF = "application/pdf"
)

// File file
type File struct {
	filename    string
	basePath    string
	contentType string
}

// New new file
func New(mf *multipart.FileHeader) (*File, error) {
	src, err := mf.Open()
	if err != nil {
		return nil, err
	}
	defer func() { _ = src.Close() }()

	td, err := ioutil.TempDir("", "com7-")
	if err != nil {
		return nil, err
	}

	dst, err := os.Create(fmt.Sprintf("%s/%s", td, mf.Filename))
	if err != nil {
		return nil, err
	}
	defer func() { _ = dst.Close() }()

	if _, err = io.Copy(dst, src); err != nil {
		return nil, err
	}
	return &File{
		filename: mf.Filename,
		basePath: td,
	}, nil
}

// NewWithFilename new file
func NewWithFilename(name string) (*File, error) {
	td, err := ioutil.TempDir("", "smh-")
	if err != nil {
		return nil, err
	}
	var ct string
	switch filepath.Ext(name) {
	case ".xlsx":
		ct = ContentTypeExcel
	case ".pdf":
		ct = ContentTypePDF
	}
	return &File{
		filename:    name,
		basePath:    td,
		contentType: ct,
	}, nil
}

// Path file path
func (f *File) Path() string {
	return fmt.Sprintf("%s/%s", f.basePath, f.filename)
}

// Name file name
func (f *File) Name() string {
	return f.filename
}

// Close close
func (f *File) Close() error {
	return os.RemoveAll(f.basePath)
}

// ContentType content type
func (f *File) ContentType() string {
	return f.contentType
}

// Ext ext
func (f *File) Ext() string {
	return filepath.Ext(f.filename)
}
