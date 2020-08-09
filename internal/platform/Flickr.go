package platform

import (
	"context"
	"fmt"

	"github.com/mayflyman4/image-poster/model"
)

type Flickr struct {
	username string
	password string
}

func (flickr *Flickr) AddAccount(ctx context.Context, pf *model.Platform) error {
	fmt.Println("add acount of Flickr with token", pf.AccessToken)

	return nil
}

func (flickr *Flickr) UpdateAccount(ctx context.Context, pf *model.Platform) error {
	fmt.Println("UpdateAccount acount of Flickr")

	return nil
}

func (flickr *Flickr) DeleteAccount(ctx context.Context, pf *model.Platform) error {
	fmt.Println("DeleteAccount acount of Flickr")

	return nil
}
