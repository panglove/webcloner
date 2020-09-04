package ui

import (
	"webcloner/myapp"
	"fyne.io/fyne/dialog"
)

func Alert(msg string) {
	dialog.ShowInformation("Tips", msg, myapp.WindowInstall)

}
func Comfirm(msg string,callback func()) {

	newD := dialog.NewInformation("Tips", msg, myapp.WindowInstall)

	newD.SetOnClosed(callback)

	newD.Show()
}

func ComfirmOK(msg string,callback func(bool)) {

	newD := dialog.NewConfirm("Tips", msg,callback, myapp.WindowInstall)
	newD.Show()
}
