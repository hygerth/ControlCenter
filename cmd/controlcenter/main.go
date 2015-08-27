package main

import (
    "flag"
    "fmt"
    "github.com/hygerth/controlcenter"
    "os"
)

var exit = os.Exit

var (
    usage      = "Usage: controlcenter [OPTIONS]"
    options    = "Options:\n-h, -help \t Print this help text and exit \n-v, -version \t Print program version and exit"
    version    = "2015.08.11"
    help       = fmt.Sprintf("%s\nVersion: %s\n%s", usage, version, options)
    cliVersion = flag.Bool("version", false, version)
    cliHelp    = flag.Bool("help", false, help)
)

func init() {
    flag.BoolVar(cliVersion, "v", false, version)
    flag.BoolVar(cliHelp, "h", false, help)
}

func main() {
    flag.Parse()

    if *cliVersion {
        fmt.Println(flag.Lookup("version").Usage)
        exit(0)
        return
    }
    if *cliHelp {
        fmt.Println(flag.Lookup("help").Usage)
        exit(0)
        return
    }
    controlcenter.Start()
    exit(0)
    return
}
