package registry

type Registration struct {
	ServiceName ServiceName
	ServiceURL  string

	//拓展，使其支持所依赖的其他服务
	RequiredServices []ServiceName

	ServiceUpdateURL string

	HeartbeatURL string
}

type ServiceName string

const (
	LogService     = ServiceName("LogService")
	GradingService = ServiceName("GradingService")
	PortalService  = ServiceName("Portald")
)

type patchEntry struct {
	Name ServiceName
	URL  string
}

type patch struct {
	Added   []patchEntry
	Removed []patchEntry
}
