package authboss_components

import (
	"context"
	"github.com/volatiletech/authboss/v3"
)

type ServerStorage struct {
}

func (s ServerStorage) Load(ctx context.Context, key string) (authboss.User, error) {

	return nil, nil
}

func (s ServerStorage) Save(ctx context.Context, user authboss.User) error {

	return nil
}
