package handlers

import "context"

type DeployRequest struct {
}

type DeployResponse struct {
}

func Deploy(ctx context.Context, req DeployRequest) (DeployResponse, *Error) {
	return DeployResponse{}, nil
}
