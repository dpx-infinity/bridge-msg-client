package main

import (
    "github.com/mattn/go-gtk/gtk"
    "io"
    "io/ioutil"
    "net/http"
    "os"
)

func handleEdit(w http.ResponseWriter, r *http.Request) {
    r, err := os.Open("pages/edit.html")
    if err != nil {
    }
    io.Copy(w, src)
}

func handleSend(w http.ResponseWriter, r *http.Request) {
}

func main() {
    // gtk.Init(&os.Args)

    // builder := gtk.Builder()
    // builder.AddFromFile(locateUiFile())
    // builder.ConnectSignals(nil)
    // ui.Init(builder)
    
    http.RedirectHandler(url, code)
    http.Handler("/", http.RedirectHandler("/edit", 301))
    http.HandleFunc("/edit", handleEdit)
    http.HandleFunc("/send", handleSend)
    http.ListenAndServe(":8080", nil)
}