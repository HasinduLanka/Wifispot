package main

import (
	"net"
	"regexp"

	"github.com/HasinduLanka/Spindly/spindlyapp"
)

var EndCh chan bool = nil

func main() {

	println(" --- Spindly Server --- ")

	spindlyapp.Configure()

	println(spindlyapp.Global.GetAppName())

	spindlyapp.Global.Refresh.OnChange(OnRefresh)
	spindlyapp.Global.Start.OnChange(OnStart)

	spindlyapp.Serve()
}

func OnRefresh(hub interface{}) {
	l, err := net.Interfaces()
	if err != nil {
		spindlyapp.Global.NetworkInterfaces.Set([]string{})
		return
	}

	NIs := make([]string, 0, len(l))

	for _, ni := range l {
		if (ni.Flags&net.FlagUp != 0) && (ni.Flags&net.FlagLoopback == 0) {
			NIs = append(NIs, ni.Name)
		}
	}

	spindlyapp.Global.NetworkInterfaces.Set(NIs)
}

func OnStart(hub interface{}) {

	if EndCh == nil {

		println("Starting Hotspot")

		src := spindlyapp.Global.SrcInterface.Get().(string)
		dst := spindlyapp.Global.DstInterface.Get().(string)
		name := spindlyapp.Global.Name.Get().(string)
		pass := spindlyapp.Global.Pass.Get().(string)

		if src == "" || dst == "" {
			spindlyapp.Global.Msg.Set("Please select source and destination interfaces")
			return
		}

		if name == "" || pass == "" {
			spindlyapp.Global.Msg.Set("Please enter name and password")
			return
		}

		re := regexp.MustCompile("^[a-zA-Z0-9_]*$")

		if !re.MatchString(name) {
			spindlyapp.Global.Msg.Set("Name can only contain alphanumeric characters and underscore")
			return
		}

		if !re.MatchString(pass) {
			spindlyapp.Global.Msg.Set("Password can only contain alphanumeric characters and underscore")
			return
		}

		if !re.MatchString(src) {
			spindlyapp.Global.Msg.Set("Source interface can only contain alphanumeric characters and underscore")
			return
		}

		if !re.MatchString(dst) {
			spindlyapp.Global.Msg.Set("Destination interface can only contain alphanumeric characters and underscore")
			return
		}

		println("Starting Hotspot on " + src + " to " + dst)

		EndCh = make(chan bool)

		go func() {
			_, exErr := ExcecCmdTask("sudo pkill -9 hostapd ; sudo pkill -9 dnsmasq ; sudo create_ap "+src+" "+dst+" "+name+" "+pass, EndCh)

			if exErr != nil {
				spindlyapp.Global.Msg.Set("Error starting hotspot" + exErr.Error())
				EndCh <- true
				EndCh = nil
				return
			}
		}()

		spindlyapp.Global.IsRunning.Set(true)
		spindlyapp.Global.Msg.Set("Hotspot started")

	} else {
		println("Stopping Hotspot")

		EndCh <- true
		EndCh = nil

		spindlyapp.Global.IsRunning.Set(false)
		spindlyapp.Global.Msg.Set("Hotspot stopped")

	}

}
