package safe

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
)

var modDir string

func init() {
	stdout, err := exec.Command("go", "env", "GOMOD").Output()
	if err != nil {
		return
	}
	fp := string(bytes.TrimSpace(stdout))
	if fp == os.DevNull || fp == "" {
		return
	}
	modDir = filepath.Dir(fp)
}

// Path 兼容的返回相对执行目录
func Path(fp string) string {
	if modDir == "" || fp == "" || filepath.IsAbs(fp) {
		return fp
	}
	return filepath.Join(modDir, fp)
}
