package lorawan

type PayloadRx struct {
	ApplicationID string `json:"applicationID"`
	ApplicationName string `json:"applicationName,omitempty"`
	DeviceName string `json:"deviceName,omitempty"`
	DevEUI string `json:"devEUI"`
	Data string `json:"data"`
	DataObj *map[string]interface{} `json:"object,omitempty"`

	FPort int `json:"fPort"`
}

type PayloadTx struct {
	FPort uint8 `json:"fPort"`
	Data string `json:"data"`
}