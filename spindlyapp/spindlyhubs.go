package spindlyapp

import "github.com/HasinduLanka/Spindly/Spindly"

var HubManager *Spindly.HubManager = Spindly.NewHubManager()

type GlobalHub struct {
	Instance          *Spindly.HubInstance
	AppName           Spindly.SpindlyStore
	Msg               Spindly.SpindlyStore
	SrcInterface      Spindly.SpindlyStore
	DstInterface      Spindly.SpindlyStore
	Pass              Spindly.SpindlyStore
	Name              Spindly.SpindlyStore
	NetworkInterfaces Spindly.SpindlyStore
	Refresh           Spindly.SpindlyStore
	Start             Spindly.SpindlyStore
	IsRunning         Spindly.SpindlyStore
}

var GlobalHub_OnInstanciate func(*GlobalHub)

var Global *GlobalHub

func (hub GlobalHub) New(InstanceID string) *GlobalHub {
	hub.Instanciate(InstanceID)
	return &hub
}

func (hub *GlobalHub) Instanciate(InstanceID string) *Spindly.HubInstance {
	hub.Instance = &Spindly.HubInstance{
		HubClass:   "GlobalHub",
		InstanceID: InstanceID,
		Stores:     make(map[string]*Spindly.SpindlyStore),
	}

	hub.AppName = Spindly.NewSpindlyStore(
		"AppName",
		func() interface{} {
			return ``
		},
		`Wifispot`,
	)
	hub.Instance.Register(&hub.AppName)

	hub.Msg = Spindly.NewSpindlyStore(
		"Msg",
		func() interface{} {
			return ""
		},
		nil,
	)
	hub.Instance.Register(&hub.Msg)

	hub.SrcInterface = Spindly.NewSpindlyStore(
		"SrcInterface",
		func() interface{} {
			return ""
		},
		nil,
	)
	hub.Instance.Register(&hub.SrcInterface)

	hub.DstInterface = Spindly.NewSpindlyStore(
		"DstInterface",
		func() interface{} {
			return ""
		},
		nil,
	)
	hub.Instance.Register(&hub.DstInterface)

	hub.Pass = Spindly.NewSpindlyStore(
		"Pass",
		func() interface{} {
			return ""
		},
		nil,
	)
	hub.Instance.Register(&hub.Pass)

	hub.Name = Spindly.NewSpindlyStore(
		"Name",
		func() interface{} {
			return ""
		},
		nil,
	)
	hub.Instance.Register(&hub.Name)

	hub.NetworkInterfaces = Spindly.NewSpindlyStore(
		"NetworkInterfaces",
		func() interface{} {
			return []string{}
		},
		nil,
	)
	hub.Instance.Register(&hub.NetworkInterfaces)

	hub.Refresh = Spindly.NewSpindlyStore(
		"Refresh",
		func() interface{} {
			return false
		},
		nil,
	)
	hub.Instance.Register(&hub.Refresh)

	hub.Start = Spindly.NewSpindlyStore(
		"Start",
		func() interface{} {
			return false
		},
		nil,
	)
	hub.Instance.Register(&hub.Start)

	hub.IsRunning = Spindly.NewSpindlyStore(
		"IsRunning",
		func() interface{} {
			return false
		},
		nil,
	)
	hub.Instance.Register(&hub.IsRunning)

	HubManager.Register(hub.Instance)
	if GlobalHub_OnInstanciate != nil {
		go GlobalHub_OnInstanciate(hub)
	}
	return hub.Instance
}

func (hub *GlobalHub) GetAppName() string {
	return hub.AppName.Get().(string)
}
func (hub *GlobalHub) GetMsg() string {
	return hub.Msg.Get().(string)
}
func (hub *GlobalHub) GetSrcInterface() string {
	return hub.SrcInterface.Get().(string)
}
func (hub *GlobalHub) GetDstInterface() string {
	return hub.DstInterface.Get().(string)
}
func (hub *GlobalHub) GetPass() string {
	return hub.Pass.Get().(string)
}
func (hub *GlobalHub) GetName() string {
	return hub.Name.Get().(string)
}
func (hub *GlobalHub) GetNetworkInterfaces() []string {
	return hub.NetworkInterfaces.Get().([]string)
}
func (hub *GlobalHub) GetRefresh() bool {
	return hub.Refresh.Get().(bool)
}
func (hub *GlobalHub) GetStart() bool {
	return hub.Start.Get().(bool)
}
func (hub *GlobalHub) GetIsRunning() bool {
	return hub.IsRunning.Get().(bool)
}

func InitializeHubs() {
	HubManager.RegisterClass("GlobalHub", func() Spindly.HubClass { return &GlobalHub{} })
	Global = GlobalHub{}.New("Global")
}
