package main

import (
	"context"
	"net"
	"net/http"
	"net/url"
	"os"
	"runtime"
	"strings"
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func createServer() *http.Server {
	home_, _ := os.UserHomeDir()
	svr := &http.Server{
		Addr:    ":8080",
		Handler: http.FileServer(http.Dir(home_)),
	}
	go func() {
		svr.ListenAndServe()
	}()
	time.Sleep(time.Second)
	return svr
}

func main() {

	var localIP string
	addrs, _ := net.InterfaceAddrs()
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil && strings.HasPrefix(ipnet.IP.String(), "192.168") {
				localIP = ipnet.IP.String()
				break
			}
		}
	}

	var isStarted = true
	svr := createServer()
	svrURL, _ := url.Parse("http://" + localIP + ":8080")

	app := app.New()
	window := app.NewWindow("File Server")

	name_ := widget.NewLabel("Welcome to File Server!\n")
	status := widget.NewLabel("Status: Started")
	info_ := widget.NewLabel("Os: " + runtime.GOOS + " | Arch: " + runtime.GOARCH + "\n")
	ip := widget.NewLabel("URL: " + localIP + ":8080\n")
	mainbtn := widget.NewButton("Stop", func() {})
	secbtn := widget.NewButton("Open in Browser", func() {
		app.OpenURL(svrURL)
	})
	mainbtn.OnTapped = func() {
		if isStarted {
			isStarted = false
			status.SetText("Status: Stopped")
			mainbtn.SetText("Start")
			svr.Shutdown(context.Background())
		} else {
			isStarted = true
			status.SetText("Status: Started")
			mainbtn.SetText("Stop")
			svr = createServer()
		}
	}

	window.SetContent(container.NewVBox(
		name_,
		status,
		info_,
		ip,
		mainbtn,
		secbtn,
	))
	window.SetTitle("ShareiT")
	window.ShowAndRun()
}
