//go:build !jsoniter
// +build !jsoniter

package json

import "encoding/json"

// RawMessage is exported by component-base/pkg/json package.
type RawMessage = json.RawMessage

var (
	Marshal       = json.Marshal
	Unmarshal     = json.Unmarshal
	MarshalIndent = json.MarshalIndent
	NewDecoder    = json.NewDecoder
	NewEncoder    = json.NewEncoder
)
