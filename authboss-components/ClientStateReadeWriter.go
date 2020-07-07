package authboss_components

import (
	"github.com/volatiletech/authboss/v3"
	"net/http"
)

type ClientStateReadWriter struct {
}

func (cs ClientStateReadWriter) ReadState(r *http.Request) (authboss.ClientState, error) {
	return nil, nil
}

func (cs ClientStateReadWriter) WriteState(
	w http.ResponseWriter,
	state authboss.ClientState,
	stateEvent []authboss.ClientStateEvent) error {

	return nil
}
