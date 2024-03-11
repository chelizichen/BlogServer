package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func JoinPathUtil(path string) string {
	wd, _ := os.Getwd()
	SIMP_PRODUCTION := os.Getenv("SIMP_PRODUCTION")
	var staticPath string
	if SIMP_PRODUCTION == "Yes" {
		SIMP_SERVER_PATH := os.Getenv("SIMP_SERVER_PATH")
		staticPath = filepath.Join(SIMP_SERVER_PATH, path)
	} else {
		staticPath = filepath.Join(wd, path)
	}
	fmt.Println("JoinPathUtil", staticPath)
	return staticPath
}
