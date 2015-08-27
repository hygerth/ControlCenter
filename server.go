package controlcenter

import (
    "github.com/googollee/go-socket.io"
    "log"
    "net/http"
    "time"
)

// Start starts the server
func Start() {
    server, err := socketio.NewServer(nil)
    checkerr(err)
    err = server.On("connection", func(so socketio.Socket) {
        log.Println("on connection")
        err := so.Join("room")
        checkerr(err)
        err = so.Emit("list", listDevices())
        checkerr(err)
        err = so.On("toggle", func(data Device) {
            go setDevice(data)
            err := so.BroadcastTo("room", "changed", data)
            checkerr(err)
        })
        checkerr(err)
        err = so.On("disconnection", func() {
            log.Println("on disconnect")
            so.Leave("room")
        })
        checkerr(err)
    })
    checkerr(err)
    err = server.On("error", func(so socketio.Socket, err error) {
        log.Println("error:", err)
    })
    checkerr(err)
    go func() {
        timer := time.Tick(60 * time.Second)
        for _ = range timer {
            server.BroadcastTo("room", "list", listDevices())
        }
    }()
    http.HandleFunc("/", indexHandler)
    http.Handle("/socket.io/", server)
    http.HandleFunc("/static/", staticFileHandler)
    log.Println("Serving at localhost:5000...")
    panic(http.ListenAndServe(":5000", nil))
}
