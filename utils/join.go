package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func JoinPathUtil(path string, others ...string) string {
	wd, _ := os.Getwd()
	SIMP_PRODUCTION := os.Getenv("SIMP_PRODUCTION")
	var staticPath string
	if SIMP_PRODUCTION == "Yes" {
		SIMP_SERVER_PATH := os.Getenv("SIMP_SERVER_PATH")
		arr := []string{}
		arr = append(arr, SIMP_SERVER_PATH)
		arr = append(arr, path)
		arr = append(arr, others...)
		staticPath = filepath.Join(arr...)
	} else {
		arr := []string{}
		arr = append(arr, wd)
		arr = append(arr, path)
		arr = append(arr, others...)
		staticPath = filepath.Join(arr...)
	}
	fmt.Println("JoinPathUtil", staticPath)
	return staticPath
}
