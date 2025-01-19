package models



type 	SetKeyValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type DeleteKeys struct {
	Keys []string `json:"keys"`
}
