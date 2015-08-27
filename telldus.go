package controlcenter

import (
    "bytes"
    "os/exec"
    "strings"
)

// Device is the struct describing an registered device in the telldus program
type Device struct {
    ID     string
    Name   string
    Status string
}

func setDevice(data Device) {
    args := []string{"--" + data.Status, data.ID}
    cmd := exec.Command("tdtool", args...)
    err := cmd.Start()
    checkerr(err)
    err = cmd.Wait()
    checkerr(err)
}

func listDevices() []Device {
    args := []string{"--list"}
    cmd := exec.Command("tdtool", args...)
    var buff bytes.Buffer
    cmd.Stdout = &buff
    err := cmd.Start()
    checkerr(err)
    err = cmd.Wait()
    checkerr(err)
    total := buff.String()
    total = strings.Replace(total, "\r", "", -1)
    s := strings.Split(total, "\n")
    var list []Device
    for i := 1; i < len(s)-2; i++ {
        d := strings.Split(s[i], "\t")
        list = append(list, Device{ID: d[0], Name: d[1], Status: d[2]})
    }
    return list
}
