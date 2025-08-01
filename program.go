package program

import "encoding/json"

type Program struct {
	Name       string      `json:"name"`
	Assets     []Asset     `json:"assets"`
	Challenges []Challenge `json:"challenges"`
}

func (p *Program) MarshalJSON() ([]byte, error) {
	return json.Marshal(p)
}

func (p *Program) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, p)
}
