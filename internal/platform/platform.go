package platform

import (
	"context"
)

// interface for the internal operations related to platform. All internal call should go through this interface.
// i.e all functions should be part of this interface or private to this package.
type PlatformInt interface {
	AddAccount(ctx context.Context) error
	// UpdateAccount(ctx context.Context) error
	// DeleteAccount(ctx context.Context) error
}

func NewPlatform(platformName string) PlatformInt {
	if platformName == "flickr" {
		return &Flickr{}
	}
	return nil
}
