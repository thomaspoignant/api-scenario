package model

type Variable struct {
	Source   Source `json:"Source"`
	Property string `json:"property"`
	Name     string `json:"name"`
}
