package tesla

import "encoding/json"

// The vehicle struct
type Vehicle struct {
	ID                     json.Number `json:"id" gqlgen:"id"`
	VehicleID              int64       `json:"vehicle_id" gqlgen:"vehicleID"`
	Vin                    string      `json:"vin"`
	DisplayName            string      `json:"display_name"`
	OptionCodes            string      `json:"option_codes"`
	Color                  interface{} `json:"color"`
	Tokens                 []string    `json:"tokens"`
	State                  string      `json:"state"`
	InService              bool        `json:"in_service"`
	IDS                    string      `json:"id_s"`
	CalendarEnabled        bool        `json:"calendar_enabled"`
	ApiVersion             int         `json:"api_version"`
	BackseatToken          interface{} `json:"backseat_token"`
	BackseatTokenUpdatedAt interface{} `json:"backseat_token_updated_at"`
}

type VehiclesResult struct {
	Count    int        `json:"count"`
	Response []*Vehicle `json:"reponse"`
}

func (c *Client) Vehicles() ([]*Vehicle, error) {
	body, err := c.get("/api/1/vehicles")
	if err != nil {
		return nil, err
	}
	response := &VehiclesResult{}
	err = json.Unmarshal(body, response)

	if err != nil {
		return nil, err
	}
	return response.Response, nil
}

func (c *Client) Vehicle(id string) (*Vehicle, error) {
	body, err := c.get("/api/1/vehicle/" + id)
	if err != nil {
		return nil, err
	}
	response := &(struct {
		Reponse *Vehicle `json:"response"`
	}{})
	err = json.Unmarshal(body, response)
	if err != nil {
		return nil, err
	}
	return response.Reponse, nil
}
