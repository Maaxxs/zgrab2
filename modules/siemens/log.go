package siemens

// S7Log is the output type for the Siemens S7 scan.
type S7Log struct {
	// IsS7 indicates that S7 was actually detected, so it should always be true.
	IsS7 bool `json:"is_s7"`

	// System is the first field returned in the component ID response.
	System string `json:"system"`

	// Module is the second field returned in the component ID response.
	Module string `json:"module"`

	// PlantId is the third field returned in the component ID response.
	PlantId string `json:"plant_id"`

	// Copyright is the fourth field returned in the component ID response.
	Copyright string `json:"copyright"`

	// SerialNumber is the fifth field returned in the component ID response.
	SerialNumber string `json:"serial_number"`

	// ModuleType is the sixth field returned in the component ID response.
	ModuleType string `json:"module_type"`

	// ReservedForOS is the seventh field returned in the component ID response.
	ReservedForOS string `json:"reserved_for_os"`

	// MemorySerialNumber is the eighth field returned in the component ID response.
	MemorySerialNumber string `json:"memory_serial_number"`

	// CpuProfile is the ninth field returned in the component ID response.
	CpuProfile string `json:"cpu_profile"`

	// OemId is the tenth field returned in the component ID response.
	OEMId string `json:"oem_id"`

	// Location is the eleventh field returned in the component ID response.
	Location string `json:"location"`

	// ModuleId is the first field returned in the module identification response.
	ModuleId string `json:"module_id"`

	// Hardware is the second field returned in the module identification response.
	Hardware string `json:"hardware"`

	// Fiirmware is the third field returned in the module identification response.
	Firmware string `json:"firmware"`
}
