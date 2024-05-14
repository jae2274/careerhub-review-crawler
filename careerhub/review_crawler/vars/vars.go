package vars

import (
	"fmt"
	"os"
)

type Vars struct {
	ReviewGrpcEndpoint string
}

type ErrNotExistedVar struct {
	VarName string
}

func NotExistedVar(varName string) *ErrNotExistedVar {
	return &ErrNotExistedVar{VarName: varName}
}

func (e *ErrNotExistedVar) Error() string {
	return fmt.Sprintf("%s is not existed", e.VarName)
}

func Variables() (*Vars, error) {
	reviewGrpcEndpoint, err := getFromEnv("REVIEW_GRPC_ENDPOINT")
	if err != nil {
		return nil, err
	}

	return &Vars{
		ReviewGrpcEndpoint: reviewGrpcEndpoint,
	}, nil
}

func getFromEnv(envVar string) (string, error) {
	ev := os.Getenv(envVar)

	if ev == "" {
		return "", fmt.Errorf("%s is not existed", envVar)
	}

	return ev, nil
}
