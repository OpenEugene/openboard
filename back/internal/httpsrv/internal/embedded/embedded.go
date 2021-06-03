package embedded

import (
	"embed"
	"fmt"
	"io/fs"
)

var (
	//go:embed contents/*
	FS embed.FS
)

func NewFS() (fs.FS, error) {
	s, err := fs.Sub(FS, "contents")
	if err != nil {
		return nil, fmt.Errorf("connect to contents fs: %w", err)
	}

	return s, nil
}
