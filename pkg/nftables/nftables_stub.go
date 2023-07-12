//go:build !linux
// +build !linux

package nftables

import (
	"github.com/asians-cloud/firewall-bouncer/pkg/cfg"
	"github.com/asians-cloud/firewall-bouncer/pkg/types"
)

func NewNFTables(config *cfg.BouncerConfig) (types.Backend, error) {
	return nil, nil
}
