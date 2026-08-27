package internal

import (
	"context"

	gomt5 "github.com/mukbeast4/go-mt5"
)

type MT5 struct {
	Client *gomt5.Client
	Ctx context.Context
}

// func Connect(id string, pipeName string) (*MT5, error) {
// 	gomt5.WithPipeName()
// }