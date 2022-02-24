package asset

import (
	"embed"
	"fmt"
	"io/fs"
)

var (
	//go:embed assets/*
	FS embed.FS
)

func NewFS() (fs.FS, error) {
	s, err := fs.Sub(FS, "assets")
	if err != nil {
		return nil, fmt.Errorf("connect to contents fs: %w", err)
	}

	return s, nil
}
