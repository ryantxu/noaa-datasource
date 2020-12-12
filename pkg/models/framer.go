package models

import "github.com/grafana/grafana-plugin-sdk-go/data"

type ErrorFramer struct {
	Error error
}

func (f *ErrorFramer) Frames() (data.Frames, error) {
	return nil, f.Error
}
