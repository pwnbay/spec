package program

type Program struct {
	Name       string      `json:"name"`
	Assets     []Asset     `json:"assets"`
	Challenges []Challenge `json:"challenges"`
	// Rewards    []Reward    `json:"rewards"`
	// Resources  Resources   `json:"resources"`
}

//func (p *Program) MarshalJSON() ([]byte, error) {
//	return json.Marshal(p)
//}
//
//func (p *Program) UnmarshalJSON(data []byte) error {
//	return json.Unmarshal(data, p)
//}
