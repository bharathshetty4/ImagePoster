package platform

import (
	"context"

	"github.com/mayflyman4/image-poster/model"
)

// interface for the internal operations related to platform. All internal call should go through this interface.
// i.e all functions should be part of this interface or private to this package.
type PlatformInt interface {
	AddAccount(ctx context.Context, account *model.Platform) error
}
