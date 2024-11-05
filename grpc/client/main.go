package client

import (
	"context"
)

type client struct {
}

type OperationRequest struct {
	Name      string `json:"Name"`
	Age       int    `json:"Age"`
	Operation string `json:"Operation"`
}

func SendRequest(ctx context.Context, or *OperationRequest) {

}
