package gocov

import (
	"encoding/json"
)

func marshalJson(packages []*Package) ([]byte, error) {
	return json.Marshal(struct{ Packages []*Package }{packages})
}

func unmarshalJson(data []byte) (packages []*Package, err error) {
	result := &struct{ Packages []*Package }{}
	err = json.Unmarshal(data, result)
	if err == nil {
		packages = result.Packages
	}
	return
}
