package aggregator

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"
	"github.com/connect3world/c3-data-avs/aggregator/types"
	cstaskmanager "github.com/connect3world/c3-data-avs/contracts/bindings/Connect3DATaskManager"
	"github.com/connect3world/c3-data-avs/core"
)

func TestProcessSignedTaskResponse(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	var TASK_INDEX = uint32(0)
	var BLOCK_NUMBER = uint32(100)
	var ADDRESS_TO_FETCH_C3DATA = common.HexToAddress("0x523D3bd434b4E1F4F0268F6cC1f735C02f08C998")

	MOCK_OPERATOR_BLS_PRIVATE_KEY, err := bls.NewPrivateKey(MOCK_OPERATOR_BLS_PRIVATE_KEY_STRING)
	assert.Nil(t, err)
	MOCK_OPERATOR_KEYPAIR := bls.NewKeyPair(MOCK_OPERATOR_BLS_PRIVATE_KEY)
	MOCK_OPERATOR_G1PUBKEY := MOCK_OPERATOR_KEYPAIR.GetPubKeyG1()
	MOCK_OPERATOR_G2PUBKEY := MOCK_OPERATOR_KEYPAIR.GetPubKeyG2()

	operatorPubkeyDict := map[sdktypes.OperatorId]types.OperatorInfo{
		MOCK_OPERATOR_ID: {
			OperatorPubkeys: sdktypes.OperatorPubkeys{
				G1Pubkey: MOCK_OPERATOR_G1PUBKEY,
				G2Pubkey: MOCK_OPERATOR_G2PUBKEY,
			},
			OperatorAddr: common.Address{},
		},
	}
	aggregator, _, mockBlsAggServ, err := createMockAggregator(mockCtrl, operatorPubkeyDict)
	assert.Nil(t, err)

	signedTaskResponse, err := createMockSignedTaskResponse(MockTask{
		TaskNum:              TASK_INDEX,
		BlockNumber:          BLOCK_NUMBER,
		AddressToFetchC3Data: ADDRESS_TO_FETCH_C3DATA,
	}, *MOCK_OPERATOR_KEYPAIR)
	assert.Nil(t, err)
	signedTaskResponseDigest, err := core.GetTaskResponseDigest(&signedTaskResponse.TaskResponse)
	assert.Nil(t, err)

	// TODO(samlaf): is this the right way to test writing to external service?
	// or is there some wisdom to "don't mock 3rd party code"?
	// see https://hynek.me/articles/what-to-mock-in-5-mins/
	mockBlsAggServ.EXPECT().ProcessNewSignature(context.Background(), TASK_INDEX, signedTaskResponseDigest,
		&signedTaskResponse.BlsSignature, signedTaskResponse.OperatorId)
	err = aggregator.ProcessSignedTaskResponse(signedTaskResponse, nil)
	assert.Nil(t, err)
}

// mocks an operator signing on a task response
func createMockSignedTaskResponse(mockTask MockTask, keypair bls.KeyPair) (*SignedTaskResponse, error) {
	addressInResponse := common.HexToAddress("0x523D3bd434b4E1F4F0268F6cC1f735C02f08C998")
	taskResponse := &cstaskmanager.IConnect3DATaskManagerTaskResponse{
		ReferenceTaskIndex: mockTask.TaskNum,
		AddressInResponse:  addressInResponse,
	}
	taskResponseHash, err := core.GetTaskResponseDigest(taskResponse)
	if err != nil {
		return nil, err
	}
	blsSignature := keypair.SignMessage(taskResponseHash)
	signedTaskResponse := &SignedTaskResponse{
		TaskResponse: *taskResponse,
		BlsSignature: *blsSignature,
		OperatorId:   MOCK_OPERATOR_ID,
	}
	return signedTaskResponse, nil
}
