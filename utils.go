package controlcenter

import (
    "go/build"
    "log"
)

func getPath() string {
    p, _ := build.Default.Import("github.com/hygerth/controlcenter", "", build.FindOnly)
    return p.Dir
}

func checkerr(err error) {
    if err != nil {
        log.Fatal(err)
    }
}
