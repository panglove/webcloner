package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/theme"
	"webcloner/config"
	"webcloner/myapp"
	"webcloner/ui"
)

func main() {

	OnStart()
}
func OnStart() {

	appInstall := app.NewWithID(config.AppID)
	appInstall.Settings().SetTheme(theme.LightTheme())
	winInstall := appInstall.NewWindow(config.AppName)
	myapp.AppInstall = appInstall
	myapp.WindowInstall = winInstall

	LoadIndex()

}
func LoadIndex() {
	myapp.ReadSetting()

	if len(myapp.AppSetting.OutPutDir) == 0 {
		myapp.WriteSetting()
	}

	myapp.WindowInstall.SetContent(ui.GetIndexLayout())
	myapp.WindowInstall.SetMaster()
	myapp.WindowInstall.SetFixedSize(true)
	myapp.WindowInstall.ShowAndRun()

}
