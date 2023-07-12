package types

import (
	"github.com/asians-cloud/crowdsec/pkg/models"
)

type Backend interface {
	Init() error
	ShutDown() error
	Add(*models.Decision) error
	Delete(*models.Decision) error
	Commit() error
	CollectMetrics()
}

func BoolPtr(b bool) *bool {
	return &b
}
