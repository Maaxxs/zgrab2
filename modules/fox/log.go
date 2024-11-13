package fox

import (
	"github.com/zmap/zgrab2"
)

// FoxLog is the struct returned to the caller.
type FoxLog struct {
	// IsFox should always be true (otherwise, the result should have been nil).
	IsFox bool `json:"is_fox"`

	// Version corresponds to the "fox.version" response field.
	Version string `json:"version"`

	// Id corresponds to the "id" response field, which is decoded as a decimal integer.
	Id uint32 `json:"id"`

	// Hostname corresponds to the "hostName" field.
	Hostname string `json:"hostname"`

	// HostAddress corresponds to the "hostAddress" field.
	HostAddress string `json:"host_address"`

	// AppName corresponds to the "app.name" field.
	AppName string `json:"app_name"`

	// AppVersion corresponds to the "app.version" field.
	AppVersion string `json:"app_version"`

	// VMName corresponds to the "vm.name" field.
	VMName string `json:"vm_name"`

	// VMVersion corresponds to the "vm.version" field.
	VMVersion string `json:"vm_version"`

	// OSName corresponds to the "os.name" field.
	OSName string `json:"os_name"`

	// OSVersion corresponds to the "os.version" field.
	OSVersion string `json:"os_version"`

	// StationName corresponds to the "station.name" field.
	StationName string `json:"station_name"`

	// Language corresponds to the "lang" field.
	Language string `json:"language"`

	// TimeZone corresponds to the "timeZone" field (or, that portion of it before the first semicolon).
	TimeZone string `json:"time_zone"`

	// HostId corresponds to the "hostId" field.
	HostId string `json:"host_id"`

	// VMUuid corresponds to the "vmUuid" field.
	VMUuid string `json:"vm_uuid"`

	// BrandId corresponds to the "brandId" field.
	BrandId string `json:"brand_id"`

	// SysInfo corresponds to the "sysInfo" field.
	SysInfo string `json:"sys_info"`

	// AuthAgentType corresponds to the "authAgentTypeSpecs" field.
	AuthAgentType string `json:"auth_agent_type"`

	TLSLog *zgrab2.TLSLog `json:"tls"`
}
