package domain

type Device struct {
	Name    string  `json:"name"`
	Sensors Sensors `json:"sensors"`
}
