package ui

import (
    "github.com/mattn/go-gtk/gtk"
    "unsafe"
)

var (
    messageNameEntry *gtk.GtkEntry
    messagePartsView *gtk.GtkTreeView
    messageTreeStore *gtk.GtkTreeStore
)

func Assign(to unsafe.Pointer, from unsafe.Pointer) {
    fptr := (*unsafe.Pointer)(from)
    tptr := (*unsafe.Pointer)(to)
    *tptr = *fptr
}

func Init(builder *gtk.GtkBuilder) {
    messageNameEntryWidget := gtk.GtkWidget{}
    Assign(unsafe.Pointer(&messageNameEntryWidget.Widget),
           unsafe.Pointer(&builder.GetObject("messageNameEntry").Object))
    // p := (*unsafe.Pointer)(unsafe.Pointer(&messageNameEntryWidget.Widget))
    // *p = builder.GetObject("messageNameEntry").Object
    // messageNameEntry = &gtk.GtkEntry{messageNameEntryWidget, gtk.To_GtkEditable(messageNameEntryWidget.Widget)}
}

//export on_addHeaderButton_clicked
func on_addHeaderButton_clicked() {

}

//export on_addBodyPartButton_clicked
func on_addBodyPartButton_clicked() {

}

//export on_deleteElementButton_clicked
func on_deleteElementButton_clicked() {

}

//export on_sendButton_clicked
func on_sendButton_clicked() {

}