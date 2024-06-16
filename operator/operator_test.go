package operator

import (
	"context"
	"fmt"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"

	"github.com/connect3world/c3-data-avs/aggregator"
	aggtypes "github.com/connect3world/c3-data-avs/aggregator/types"
	cstaskmanager "github.com/connect3world/c3-data-avs/contracts/bindings/Connect3DATaskManager"
	chainiomocks "github.com/connect3world/c3-data-avs/core/chainio/mocks"
	operatormocks "github.com/connect3world/c3-data-avs/operator/mocks"
)

func TestOperator(t *testing.T) {
	operator, err := createMockOperator()
	assert.Nil(t, err)
	const taskIndex = 1

	t.Run("ProcessNewTaskCreatedLog", func(t *testing.T) {
		var addressToFetchC3Data = gethcommon.HexToAddress("0x523D3bd434b4E1F4F0268F6cC1f735C02f08C998")
		newTaskCreatedLog := &cstaskmanager.ContractConnect3DATaskManagerNewTaskCreated{
			TaskIndex: taskIndex,
			Task: cstaskmanager.IConnect3DATaskManagerTask{
				AddressToFetchC3Data:      addressToFetchC3Data,
				TaskCreatedBlock:          1000,
				QuorumNumbers:             aggtypes.QUORUM_NUMBERS.UnderlyingType(),
				QuorumThresholdPercentage: uint32(aggtypes.QUORUM_THRESHOLD_NUMERATOR),
			},
			Raw: types.Log{},
		}
		got := operator.ProcessNewTaskCreatedLog(newTaskCreatedLog)
		want := &cstaskmanager.IConnect3DATaskManagerTaskResponse{
			ReferenceTaskIndex: taskIndex,
			AddressInResponse:  addressToFetchC3Data,
		}
		assert.Equal(t, got, want)
	})

	t.Run("Start", func(t *testing.T) {
		var addressToFetchC3Data = gethcommon.HexToAddress("0x523D3bd434b4E1F4F0268F6cC1f735C02f08C998")

		// new task event
		newTaskCreatedEvent := &cstaskmanager.ContractConnect3DATaskManagerNewTaskCreated{
			TaskIndex: taskIndex,
			Task: cstaskmanager.IConnect3DATaskManagerTask{
				AddressToFetchC3Data:      addressToFetchC3Data,
				TaskCreatedBlock:          1000,
				QuorumNumbers:             aggtypes.QUORUM_NUMBERS.UnderlyingType(),
				QuorumThresholdPercentage: uint32(aggtypes.QUORUM_THRESHOLD_NUMERATOR),
			},
			Raw: types.Log{},
		}
		fmt.Println("newTaskCreatedEvent", newTaskCreatedEvent)
		X, ok := big.NewInt(0).SetString("7926134832136282318561896451042374984489965925674521194255549259381336496956", 10)
		assert.True(t, ok)
		Y, ok := big.NewInt(0).SetString("15243507701692917330954619280683582177901049846125926696838777109165913318327", 10)
		assert.True(t, ok)

		signedTaskResponse := &aggregator.SignedTaskResponse{
			TaskResponse: cstaskmanager.IConnect3DATaskManagerTaskResponse{
				ReferenceTaskIndex: taskIndex,
				AddressInResponse:  addressToFetchC3Data,
			},
			BlsSignature: bls.Signature{
				G1Point: bls.NewG1Point(X, Y),
			},
			OperatorId: operator.operatorId,
		}
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mockAggregatorRpcClient := operatormocks.NewMockAggregatorRpcClienter(mockCtrl)
		mockAggregatorRpcClient.EXPECT().SendSignedTaskResponseToAggregator(signedTaskResponse)
		operator.aggregatorRpcClient = mockAggregatorRpcClient

		mockSubscriber := chainiomocks.NewMockAvsSubscriberer(mockCtrl)
		mockSubscriber.EXPECT().SubscribeToNewTasks(operator.newTaskCreatedChan).Return(event.NewSubscription(func(quit <-chan struct{}) error {
			// loop forever
			<-quit
			return nil
		}))
		operator.avsSubscriber = mockSubscriber

		mockReader := chainiomocks.NewMockAvsReaderer(mockCtrl)
		mockReader.EXPECT().IsOperatorRegistered(gomock.Any(), operator.operatorAddr).Return(true, nil)
		operator.avsReader = mockReader

		ctx, cancel := context.WithCancel(context.Background())
		go func() {
			err := operator.Start(ctx)
			assert.Nil(t, err)
		}()
		operator.newTaskCreatedChan <- newTaskCreatedEvent
		time.Sleep(1 * time.Second)

		cancel()
	})

}
