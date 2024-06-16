package types

import (
	"errors"

	cstaskmanager "github.com/connect3world/c3-data-avs/contracts/bindings/Connect3DATaskManager"
)

type TaskResponseData struct {
	TaskResponse              cstaskmanager.IConnect3DATaskManagerTaskResponse
	TaskResponseMetadata      cstaskmanager.IConnect3DATaskManagerTaskResponseMetadata
	NonSigningOperatorPubKeys []cstaskmanager.BN254G1Point
}

var (
	NoErrorInTaskResponse = errors.New("100. Task response is valid")
)
