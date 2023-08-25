package main

import (
	"path"
	"runtime"
)

func BasePath() (pwd string) {
	_, fullFilename, _, _ := runtime.Caller(0)
	pwd = path.Dir(fullFilename)
	return
}
