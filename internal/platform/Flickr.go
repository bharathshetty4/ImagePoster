package platform

import (
	"context"
	"fmt"
)

type Flickr struct {
	username string
	password string
}

func (flickr *Flickr) AddAccount(ctx context.Context) error {
	fmt.Println("add acount of FLICKR")

	return nil
}
