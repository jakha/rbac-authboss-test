package config

import (
	"context"
	"github.com/volatiletech/authboss/v3"
)

type Storer struct {
	Users map[authboss.User]bool
}

func GetDb() authboss.ServerStorer {
	b := Storer{}


	return b
}

func (s Storer) Load(ctx context.Context, key string) (authboss.User, error) {

	return nil, nil
}

func (s Storer) Save(ctx context.Context, user authboss.User) error {

	return nil
}
