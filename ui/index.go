package ui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
	"net/url"
	"webcloner/config"
	"webcloner/myapp"
	"webcloner/util/web"
)

const (
	Width  = 600
	Height = 400
)

var downloadChan chan bool
var urlEdit, dirEdit *widget.Entry
var cloneBt *widget.Button

func GetIndexLayout() fyne.CanvasObject {
	downloadChan = make(chan bool)
	viewSize := fyne.NewSize(Width, Height)

	welcomeTip := widget.NewLabel("Welcome to " + config.AppName)
	welcomeTip.Resize(welcomeTip.MinSize())
	SetWidgetHCenter(welcomeTip, viewSize)
	SetWidgetY(welcomeTip, 40)

	urlEdit = widget.NewEntry()
	urlEdit.SetPlaceHolder("Input Website Url")
	urlEdit.Resize(fyne.NewSize(400, 30))
	SetWidgetHCenter(urlEdit, viewSize)
	SetWidgetY(urlEdit, 120)

	dirEdit = widget.NewEntry()
	dirEdit.SetText("./dist")
	dirEdit.SetPlaceHolder("Input Output Dir")
	dirEdit.Resize(fyne.NewSize(400, 30))
	SetWidgetHCenter(dirEdit, viewSize)
	SetWidgetY(dirEdit, 180)

	cloneBt = widget.NewButton("Get Start", cloneBtClick)

	cloneBt.Resize(fyne.NewSize(200, 30))
	SetWidgetHCenter(cloneBt, viewSize)
	SetWidgetY(cloneBt, 300)

	lay := fyne.NewContainerWithLayout(&AbLayout{Width, Height}, welcomeTip, urlEdit, dirEdit, cloneBt)

	return lay

}
func cloneBtClick() {

	urlStr := urlEdit.Text

	dirStr := dirEdit.Text

	if len(urlStr) <= 0 {

		Alert("Please input website url")
		return
	}

	if len(dirStr) <= 0 {

		Alert("Please input the output url")
		return
	}

	cloneBt.SetText("Dowloading...")
	cloneBt.Disable()

	go func() {
		downloadChan <- web.DownloadWebsite(urlStr, dirStr)

	}()
	isOk := <-downloadChan
	cloneBt.SetText("Get Start")
	cloneBt.Enable()
	if isOk {
		localPath := web.GetWebsitePath(urlStr, dirStr)

		web.HttpFileServer(8082, localPath)

		ComfirmOK("Download Success,Do you want open it?", func(b bool) {
			if b {
				url ,_:=url.Parse("http://127.0.0.1:8082")

				myapp.AppInstall.OpenURL(url)

			}

		})
	} else {
		Alert("Download Error!")
	}

}
