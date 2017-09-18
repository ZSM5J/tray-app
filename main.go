package main

import (
   "github.com/getlantern/systray"
   "io/ioutil"
   "log"
   "github.com/skratchdot/open-golang/open"
	"./bash"
)

func main() {
	// Should be called at the very beginning of main().
	systray.Run(onReady, onExit)
}

func onReady() {
	content, err := ioutil.ReadFile("assets/sc.png")
	if err != nil {
		log.Fatal(err)
	}
	systray.SetIcon(content)
	systray.SetTitle("Tray App")
	systray.SetTooltip("Tray App")

	

	go func() {
		mURL := systray.AddMenuItem("cx sandbox", "cx-web")
		vs := systray.AddMenuItem("viscript", "cx-web")
		mQuit := systray.AddMenuItem("Quit", "Quit the whole app")
		systray.AddSeparator()

		for {
			select {
			case <-mURL.ClickedCh:
				open.Run("http://cx.skycoin.net")
			case <-vs.ClickedCh:
				go runViscript()
			case <-mQuit.ClickedCh:
				systray.Quit()
				log.Println("Quit now...")
				return
			}
		}
	}()
}

func onExit() {
	// clean up here
}


func runViscript() {
	bash.Execute("./apps/viscript/viscript")
}

