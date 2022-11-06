package dto

const (
	OK   = "OK"
	HTTP = "Http"
)

type (
	Status struct {
		Name   string      `json:"name"`
		Status string      `json:"status"`
		Data   interface{} `json:"data"`
	}

	HCStatus struct {
		Status []Status `json:"status"`
	}

	HCData struct {
		HandlerCount uint32 `json:"handlerCount"`
	}
)
