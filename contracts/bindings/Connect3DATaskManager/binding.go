// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractConnect3DATaskManager

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// BN254G1Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G1Point struct {
	X *big.Int
	Y *big.Int
}

// BN254G2Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G2Point struct {
	X [2]*big.Int
	Y [2]*big.Int
}

// IBLSSignatureCheckerNonSignerStakesAndSignature is an auto generated low-level Go binding around an user-defined struct.
type IBLSSignatureCheckerNonSignerStakesAndSignature struct {
	NonSignerQuorumBitmapIndices []uint32
	NonSignerPubkeys             []BN254G1Point
	QuorumApks                   []BN254G1Point
	ApkG2                        BN254G2Point
	Sigma                        BN254G1Point
	QuorumApkIndices             []uint32
	TotalStakeIndices            []uint32
	NonSignerStakeIndices        [][]uint32
}

// IBLSSignatureCheckerQuorumStakeTotals is an auto generated low-level Go binding around an user-defined struct.
type IBLSSignatureCheckerQuorumStakeTotals struct {
	SignedStakeForQuorum []*big.Int
	TotalStakeForQuorum  []*big.Int
}

// IConnect3DATaskManagerTask is an auto generated low-level Go binding around an user-defined struct.
type IConnect3DATaskManagerTask struct {
	AddressToFetchC3Data      common.Address
	TaskCreatedBlock          uint32
	QuorumNumbers             []byte
	QuorumThresholdPercentage uint32
}

// IConnect3DATaskManagerTaskResponse is an auto generated low-level Go binding around an user-defined struct.
type IConnect3DATaskManagerTaskResponse struct {
	ReferenceTaskIndex uint32
	AddressInResponse  common.Address
	C3DataHash         [32]byte
}

// IConnect3DATaskManagerTaskResponseMetadata is an auto generated low-level Go binding around an user-defined struct.
type IConnect3DATaskManagerTaskResponseMetadata struct {
	TaskResponsedBlock uint32
	HashOfNonSigners   [32]byte
}

// OperatorStateRetrieverCheckSignaturesIndices is an auto generated low-level Go binding around an user-defined struct.
type OperatorStateRetrieverCheckSignaturesIndices struct {
	NonSignerQuorumBitmapIndices []uint32
	QuorumApkIndices             []uint32
	TotalStakeIndices            []uint32
	NonSignerStakeIndices        [][]uint32
}

// OperatorStateRetrieverOperator is an auto generated low-level Go binding around an user-defined struct.
type OperatorStateRetrieverOperator struct {
	Operator   common.Address
	OperatorId [32]byte
	Stake      *big.Int
}

// ContractConnect3DATaskManagerMetaData contains all meta data concerning the ContractConnect3DATaskManager contract.
var ContractConnect3DATaskManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"_taskResponseWindowBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"TASK_CHALLENGE_WINDOW_BLOCK\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TASK_RESPONSE_WINDOW_BLOCK\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"aggregator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allTaskHashes\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allTaskResponses\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"blsApkRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBLSApkRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkSignatures\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.QuorumStakeTotals\",\"components\":[{\"name\":\"signedStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"},{\"name\":\"totalStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"}]},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createNewTask\",\"inputs\":[{\"name\":\"addressToFetchC3Data\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"generator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCheckSignaturesIndices\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"nonSignerOperatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOperatorStateRetriever.CheckSignaturesIndices\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuorumBitmapsAtBlockNumber\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTaskResponseWindowBlock\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_aggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_generator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"latestTaskNum\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"raiseAndResolveChallenge\",\"inputs\":[{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structIConnect3DATaskManager.Task\",\"components\":[{\"name\":\"addressToFetchC3Data\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"taskResponse\",\"type\":\"tuple\",\"internalType\":\"structIConnect3DATaskManager.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"addressInResponse\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"c3DataHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"taskResponseMetadata\",\"type\":\"tuple\",\"internalType\":\"structIConnect3DATaskManager.TaskResponseMetadata\",\"components\":[{\"name\":\"taskResponsedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"pubkeysOfNonSigningOperators\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registryCoordinator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"respondToTask\",\"inputs\":[{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structIConnect3DATaskManager.Task\",\"components\":[{\"name\":\"addressToFetchC3Data\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"taskResponse\",\"type\":\"tuple\",\"internalType\":\"structIConnect3DATaskManager.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"addressInResponse\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"c3DataHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"nonSignerStakesAndSignature\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPauserRegistry\",\"inputs\":[{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setStaleStakesForbidden\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakeRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"staleStakesForbidden\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"taskNumber\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"taskSuccesfullyChallenged\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"trySignatureAndApkVerification\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"apk\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"pairingSuccessful\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"siganatureIsValid\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NewTaskCreated\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"task\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIConnect3DATaskManager.Task\",\"components\":[{\"name\":\"addressToFetchC3Data\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PauserRegistrySet\",\"inputs\":[{\"name\":\"pauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StaleStakesForbiddenUpdate\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskChallengedSuccessfully\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"challenger\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskChallengedUnsuccessfully\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"challenger\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskCompleted\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskResponded\",\"inputs\":[{\"name\":\"taskResponse\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIConnect3DATaskManager.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"addressInResponse\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"c3DataHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"taskResponseMetadata\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIConnect3DATaskManager.TaskResponseMetadata\",\"components\":[{\"name\":\"taskResponsedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x6101206040523480156200001257600080fd5b5060405162005ede38038062005ede8339810160408190526200003591620001f7565b81806001600160a01b03166080816001600160a01b031681525050806001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200008f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620000b591906200023e565b6001600160a01b031660a0816001600160a01b031681525050806001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200010d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200013391906200023e565b6001600160a01b031660c0816001600160a01b03168152505060a0516001600160a01b031663df5cf7236040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200018d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001b391906200023e565b6001600160a01b031660e052506097805460ff1916600117905563ffffffff16610100525062000265565b6001600160a01b0381168114620001f457600080fd5b50565b600080604083850312156200020b57600080fd5b82516200021881620001de565b602084015190925063ffffffff811681146200023357600080fd5b809150509250929050565b6000602082840312156200025157600080fd5b81516200025e81620001de565b9392505050565b60805160a05160c05160e05161010051615be7620002f7600039600081816102900152818161059c01526109da0152600081816105650152612a660152600081816104310152818161173b0152612c4801526000818161045801528181612e1e0152612fe001526000818161047f0152818161187501528181612750015281816128c90152612b030152615be76000f3fe608060405234801561001057600080fd5b506004361061021c5760003560e01c8063683048351161012557806391085183116100ad578063f2fde38b1161007c578063f2fde38b14610587578063f5c9899d1461059a578063f63c5bab146105c0578063f8c8765e146105c8578063fabc1cbc146105db57600080fd5b8063910851831461051f578063b98d090814610532578063cefdc1d41461053f578063df5cf7231461056057600080fd5b806372d18e8d116100f457806372d18e8d146104ca5780637afa1eed146104d8578063886f1195146104eb5780638b00ce7c146104fe5780638da5cb5b1461050e57600080fd5b806368304835146104535780636d14a9871461047a5780636efb4636146104a1578063715018a6146104c257600080fd5b80633878e94b116101a85780635ac86ab7116101775780635ac86ab7146103ae5780635c155662146103e15780635c975abb146104015780635decc3f5146104095780635df459461461042c57600080fd5b80633878e94b14610360578063416c7e5e146103735780634f739f7414610386578063595c6a67146103a657600080fd5b80631ad43189116101ef5780631ad431891461028b578063245a7bfc146102c75780632cb223d5146102f25780632d89f6fc146103205780633563b0d11461034057600080fd5b806310d67a2f14610221578063136439dd1461023657806314e5e9ba14610249578063171f1d5b1461025c575b600080fd5b61023461022f3660046146c0565b6105ee565b005b6102346102443660046146dd565b6106aa565b610234610257366004614b46565b6107e9565b61026f61026a366004614bba565b610c68565b6040805192151583529015156020830152015b60405180910390f35b6102b27f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff9091168152602001610282565b60cd546102da906001600160a01b031681565b6040516001600160a01b039091168152602001610282565b610312610300366004614c0b565b60cb6020526000908152604090205481565b604051908152602001610282565b61031261032e366004614c0b565b60ca6020526000908152604090205481565b61035361034e366004614c28565b610df2565b6040516102829190614d80565b61023461036e366004614d9a565b611288565b610234610381366004614e39565b611873565b610399610394366004614e9e565b6119e8565b6040516102829190614fa2565b61023461210e565b6103d16103bc36600461506c565b606654600160ff9092169190911b9081161490565b6040519015158152602001610282565b6103f46103ef366004615089565b6121d5565b6040516102829190615135565b606654610312565b6103d1610417366004614c0b565b60cc6020526000908152604090205460ff1681565b6102da7f000000000000000000000000000000000000000000000000000000000000000081565b6102da7f000000000000000000000000000000000000000000000000000000000000000081565b6102da7f000000000000000000000000000000000000000000000000000000000000000081565b6104b46104af366004615179565b61239d565b604051610282929190615239565b610234613295565b60c95463ffffffff166102b2565b60ce546102da906001600160a01b031681565b6065546102da906001600160a01b031681565b60c9546102b29063ffffffff1681565b6033546001600160a01b03166102da565b61023461052d366004615282565b6132a9565b6097546103d19060ff1681565b61055261054d3660046152e6565b613430565b604051610282929190615328565b6102da7f000000000000000000000000000000000000000000000000000000000000000081565b6102346105953660046146c0565b6135c2565b7f00000000000000000000000000000000000000000000000000000000000000006102b2565b6102b2606481565b6102346105d6366004615349565b613638565b6102346105e93660046146dd565b613789565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610641573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061066591906153a5565b6001600160a01b0316336001600160a01b03161461069e5760405162461bcd60e51b8152600401610695906153c2565b60405180910390fd5b6106a7816138e5565b50565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa1580156106f2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610716919061540c565b6107325760405162461bcd60e51b815260040161069590615429565b606654818116146107ab5760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e70617573653a20696e76616c696420617474656d70742060448201527f746f20756e70617573652066756e6374696f6e616c69747900000000000000006064820152608401610695565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d906020015b60405180910390a250565b60cd546001600160a01b031633146108435760405162461bcd60e51b815260206004820152601d60248201527f41676772656761746f72206d757374206265207468652063616c6c65720000006044820152606401610695565b60006108556040850160208601614c0b565b90503660006108676040870187615471565b9092509050600061087e6080880160608901614c0b565b905060ca60006108916020890189614c0b565b63ffffffff1663ffffffff16815260200190815260200160002054876040516020016108bd91906154e0565b60405160208183030381529060405280519060200120146109465760405162461bcd60e51b815260206004820152603d60248201527f737570706c696564207461736b20646f6573206e6f74206d617463682074686560448201527f206f6e65207265636f7264656420696e2074686520636f6e74726163740000006064820152608401610695565b600060cb8161095860208a018a614c0b565b63ffffffff1663ffffffff16815260200190815260200160002054146109d55760405162461bcd60e51b815260206004820152602c60248201527f41676772656761746f722068617320616c726561647920726573706f6e64656460448201526b20746f20746865207461736b60a01b6064820152608401610695565b6109ff7f0000000000000000000000000000000000000000000000000000000000000000856155ab565b63ffffffff164363ffffffff161115610a705760405162461bcd60e51b815260206004820152602d60248201527f41676772656761746f722068617320726573706f6e64656420746f207468652060448201526c7461736b20746f6f206c61746560981b6064820152608401610695565b600086604051602001610a83919061560d565b604051602081830303815290604052805190602001209050600080610aab8387878a8c61239d565b9150915060005b85811015610baa578460ff1683602001518281518110610ad457610ad461561b565b6020026020010151610ae69190615631565b6001600160601b0316606484600001518381518110610b0757610b0761561b565b60200260200101516001600160601b0316610b229190615660565b1015610b98576040805162461bcd60e51b81526020600482015260248101919091527f5369676e61746f7269657320646f206e6f74206f776e206174206c656173742060448201527f7468726573686f6c642070657263656e74616765206f6620612071756f72756d6064820152608401610695565b80610ba28161567f565b915050610ab2565b5060408051808201825263ffffffff43168152602080820184905291519091610bd7918c9184910161569a565b6040516020818303038152906040528051906020012060cb60008c6000016020810190610c049190614c0b565b63ffffffff1663ffffffff168152602001908152602001600020819055507fb6a7079e92345de4742a12e32ffe70803bdaaac6f5417156288caefd356e4b028a82604051610c5392919061569a565b60405180910390a15050505050505050505050565b60008060007f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000187876000015188602001518860000151600060028110610cb057610cb061561b565b60200201518951600160200201518a60200151600060028110610cd557610cd561561b565b60200201518b60200151600160028110610cf157610cf161561b565b602090810291909101518c518d830151604051610d4e9a99989796959401988952602089019790975260408801959095526060870193909352608086019190915260a085015260c084015260e08301526101008201526101200190565b6040516020818303038152906040528051906020012060001c610d7191906156c6565b9050610de4610d8a610d8388846139dc565b8690613a73565b610d92613b07565b610dda610dcb85610dc5604080518082018252600080825260209182015281518083019092526001825260029082015290565b906139dc565b610dd48c613bc7565b90613a73565b886201d4c0613c57565b909890975095505050505050565b60606000846001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa158015610e34573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e5891906153a5565b90506000856001600160a01b0316639e9923c26040518163ffffffff1660e01b8152600401602060405180830381865afa158015610e9a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ebe91906153a5565b90506000866001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa158015610f00573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f2491906153a5565b9050600086516001600160401b03811115610f4157610f41614720565b604051908082528060200260200182016040528015610f7457816020015b6060815260200190600190039081610f5f5790505b50905060005b875181101561127c576000888281518110610f9757610f9761561b565b0160200151604051638902624560e01b815260f89190911c6004820181905263ffffffff8a16602483015291506000906001600160a01b03871690638902624590604401600060405180830381865afa158015610ff8573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261102091908101906156e8565b905080516001600160401b0381111561103b5761103b614720565b60405190808252806020026020018201604052801561108657816020015b60408051606081018252600080825260208083018290529282015282526000199092019101816110595790505b508484815181106110995761109961561b565b602002602001018190525060005b8151811015611266576040518060600160405280876001600160a01b03166347b314e88585815181106110dc576110dc61561b565b60200260200101516040518263ffffffff1660e01b815260040161110291815260200190565b602060405180830381865afa15801561111f573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061114391906153a5565b6001600160a01b031681526020018383815181106111635761116361561b565b60200260200101518152602001896001600160a01b031663fa28c6278585815181106111915761119161561b565b60209081029190910101516040516001600160e01b031960e084901b168152600481019190915260ff8816602482015263ffffffff8f166044820152606401602060405180830381865afa1580156111ed573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112119190615778565b6001600160601b031681525085858151811061122f5761122f61561b565b602002602001015182815181106112485761124861561b565b6020026020010181905250808061125e9061567f565b9150506110a7565b50505080806112749061567f565b915050610f7a565b50979650505050505050565b60006112976020850185614c0b565b905060006112a860208701876146c0565b63ffffffff8316600090815260cb60205260409020549091506113175760405162461bcd60e51b815260206004820152602160248201527f5461736b206861736e2774206265656e20726573706f6e64656420746f2079656044820152601d60fa1b6064820152608401610695565b848460405160200161132a9291906157a1565b60408051601f19818403018152918152815160209283012063ffffffff8516600090815260cb909352912054146113c95760405162461bcd60e51b815260206004820152603d60248201527f5461736b20726573706f6e736520646f6573206e6f74206d617463682074686560448201527f206f6e65207265636f7264656420696e2074686520636f6e74726163740000006064820152608401610695565b63ffffffff8216600090815260cc602052604090205460ff16156114615760405162461bcd60e51b815260206004820152604360248201527f54686520726573706f6e736520746f2074686973207461736b2068617320616c60448201527f7265616479206265656e206368616c6c656e676564207375636365737366756c606482015262363c9760e91b608482015260a401610695565b60646114706020860186614c0b565b61147a91906155ab565b63ffffffff164363ffffffff1611156114fb5760405162461bcd60e51b815260206004820152603760248201527f546865206368616c6c656e676520706572696f6420666f72207468697320746160448201527f736b2068617320616c726561647920657870697265642e0000000000000000006064820152608401610695565b600061150d60408701602088016146c0565b6001600160a01b038381169116149050600181141561156157604051339063ffffffff8516907ffd3e26beeb5967fc5a57a0446914eabc45b4aa474c67a51b4b5160cac60ddb0590600090a350505061186d565b600084516001600160401b0381111561157c5761157c614720565b6040519080825280602002602001820160405280156115a5578160200160208202803683370190505b50905060005b8551811015611617576115e88682815181106115c9576115c961561b565b6020026020010151805160009081526020918201519091526040902090565b8282815181106115fa576115fa61561b565b60209081029190910101528061160f8161567f565b9150506115ab565b50600061162a60408a0160208b01614c0b565b8260405160200161163c9291906157d7565b604051602081830303815290604052805190602001209050866020013581146116e65760405162461bcd60e51b815260206004820152605060248201527f546865207075626b657973206f66206e6f6e2d7369676e696e67206f7065726160448201527f746f727320737570706c69656420627920746865206368616c6c656e6765722060648201526f30b932903737ba1031b7b93932b1ba1760811b608482015260a401610695565b600086516001600160401b0381111561170157611701614720565b60405190808252806020026020018201604052801561172a578160200160208202803683370190505b50905060005b875181101561181d577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663e8bb9ae685838151811061177a5761177a61561b565b60200260200101516040518263ffffffff1660e01b81526004016117a091815260200190565b602060405180830381865afa1580156117bd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117e191906153a5565b8282815181106117f3576117f361561b565b6001600160a01b0390921660209283029190910190910152806118158161567f565b915050611730565b5063ffffffff8616600081815260cc6020526040808220805460ff19166001179055513392917fc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec91a35050505050505b50505050565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156118d1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118f591906153a5565b6001600160a01b0316336001600160a01b0316146119a15760405162461bcd60e51b815260206004820152605c60248201527f424c535369676e6174757265436865636b65722e6f6e6c79436f6f7264696e6160448201527f746f724f776e65723a2063616c6c6572206973206e6f7420746865206f776e6560648201527f72206f6620746865207265676973747279436f6f7264696e61746f7200000000608482015260a401610695565b6097805460ff19168215159081179091556040519081527f40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc9060200160405180910390a150565b611a136040518060800160405280606081526020016060815260200160608152602001606081525090565b6000876001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa158015611a53573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611a7791906153a5565b9050611aa46040518060800160405280606081526020016060815260200160608152602001606081525090565b6040516361c8a12f60e11b81526001600160a01b038a169063c391425e90611ad4908b908990899060040161581f565b600060405180830381865afa158015611af1573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611b199190810190615869565b81526040516340e03a8160e11b81526001600160a01b038316906381c0750290611b4b908b908b908b906004016158f7565b600060405180830381865afa158015611b68573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611b909190810190615869565b6040820152856001600160401b03811115611bad57611bad614720565b604051908082528060200260200182016040528015611be057816020015b6060815260200190600190039081611bcb5790505b50606082015260005b60ff811687111561201f576000856001600160401b03811115611c0e57611c0e614720565b604051908082528060200260200182016040528015611c37578160200160208202803683370190505b5083606001518360ff1681518110611c5157611c5161561b565b602002602001018190525060005b86811015611f1f5760008c6001600160a01b03166304ec63518a8a85818110611c8a57611c8a61561b565b905060200201358e88600001518681518110611ca857611ca861561b565b60200260200101516040518463ffffffff1660e01b8152600401611ce59392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa158015611d02573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611d269190615920565b90506001600160c01b038116611dca5760405162461bcd60e51b815260206004820152605c60248201527f4f70657261746f7253746174655265747269657665722e676574436865636b5360448201527f69676e617475726573496e64696365733a206f70657261746f72206d7573742060648201527f6265207265676973746572656420617420626c6f636b6e756d62657200000000608482015260a401610695565b8a8a8560ff16818110611ddf57611ddf61561b565b6001600160c01b03841692013560f81c9190911c600190811614159050611f0c57856001600160a01b031663dd9846b98a8a85818110611e2157611e2161561b565b905060200201358d8d8860ff16818110611e3d57611e3d61561b565b6040516001600160e01b031960e087901b1681526004810194909452919091013560f81c60248301525063ffffffff8f166044820152606401602060405180830381865afa158015611e93573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611eb79190615949565b85606001518560ff1681518110611ed057611ed061561b565b60200260200101518481518110611ee957611ee961561b565b63ffffffff9092166020928302919091019091015282611f088161567f565b9350505b5080611f178161567f565b915050611c5f565b506000816001600160401b03811115611f3a57611f3a614720565b604051908082528060200260200182016040528015611f63578160200160208202803683370190505b50905060005b82811015611fe45784606001518460ff1681518110611f8a57611f8a61561b565b60200260200101518181518110611fa357611fa361561b565b6020026020010151828281518110611fbd57611fbd61561b565b63ffffffff9092166020928302919091019091015280611fdc8161567f565b915050611f69565b508084606001518460ff1681518110611fff57611fff61561b565b60200260200101819052505050808061201790615966565b915050611be9565b506000896001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa158015612060573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061208491906153a5565b60405163354952a360e21b81529091506001600160a01b0382169063d5254a8c906120b7908b908b908e90600401615986565b600060405180830381865afa1580156120d4573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526120fc9190810190615869565b60208301525098975050505050505050565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa158015612156573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061217a919061540c565b6121965760405162461bcd60e51b815260040161069590615429565b600019606681905560405190815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2565b60606000846001600160a01b031663c391425e84866040518363ffffffff1660e01b81526004016122079291906159b0565b600060405180830381865afa158015612224573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261224c9190810190615869565b9050600084516001600160401b0381111561226957612269614720565b604051908082528060200260200182016040528015612292578160200160208202803683370190505b50905060005b855181101561239357866001600160a01b03166304ec63518783815181106122c2576122c261561b565b6020026020010151878685815181106122dd576122dd61561b565b60200260200101516040518463ffffffff1660e01b815260040161231a9392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa158015612337573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061235b9190615920565b6001600160c01b03168282815181106123765761237661561b565b60209081029190910101528061238b8161567f565b915050612298565b5095945050505050565b60408051808201909152606080825260208201526000846124145760405162461bcd60e51b81526020600482015260376024820152600080516020615b9283398151915260448201527f7265733a20656d7074792071756f72756d20696e7075740000000000000000006064820152608401610695565b6040830151518514801561242c575060a08301515185145b801561243c575060c08301515185145b801561244c575060e08301515185145b6124b65760405162461bcd60e51b81526020600482015260416024820152600080516020615b9283398151915260448201527f7265733a20696e7075742071756f72756d206c656e677468206d69736d6174636064820152600d60fb1b608482015260a401610695565b8251516020840151511461252e5760405162461bcd60e51b815260206004820152604460248201819052600080516020615b92833981519152908201527f7265733a20696e707574206e6f6e7369676e6572206c656e677468206d69736d6064820152630c2e8c6d60e31b608482015260a401610695565b4363ffffffff168463ffffffff161061259d5760405162461bcd60e51b815260206004820152603c6024820152600080516020615b9283398151915260448201527f7265733a20696e76616c6964207265666572656e636520626c6f636b000000006064820152608401610695565b6040805180820182526000808252602080830191909152825180840190935260608084529083015290866001600160401b038111156125de576125de614720565b604051908082528060200260200182016040528015612607578160200160208202803683370190505b506020820152866001600160401b0381111561262557612625614720565b60405190808252806020026020018201604052801561264e578160200160208202803683370190505b50815260408051808201909152606080825260208201528560200151516001600160401b0381111561268257612682614720565b6040519080825280602002602001820160405280156126ab578160200160208202803683370190505b5081526020860151516001600160401b038111156126cb576126cb614720565b6040519080825280602002602001820160405280156126f4578160200160208202803683370190505b50816020018190525060006127c68a8a8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505060408051639aa1653d60e01b815290516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169350639aa1653d925060048083019260209291908290030181865afa15801561279d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906127c19190615a04565b613e7b565b905060005b876020015151811015612a42576127f1886020015182815181106115c9576115c961561b565b836020015182815181106128075761280761561b565b602090810291909101015280156128c7576020830151612828600183615a21565b815181106128385761283861561b565b602002602001015160001c836020015182815181106128595761285961561b565b602002602001015160001c116128c7576040805162461bcd60e51b8152602060048201526024810191909152600080516020615b9283398151915260448201527f7265733a206e6f6e5369676e65725075626b657973206e6f7420736f727465646064820152608401610695565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166304ec63518460200151838151811061290c5761290c61561b565b60200260200101518b8b60000151858151811061292b5761292b61561b565b60200260200101516040518463ffffffff1660e01b81526004016129689392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa158015612985573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906129a99190615920565b6001600160c01b0316836000015182815181106129c8576129c861561b565b602002602001018181525050612a2e610d83612a0284866000015185815181106129f4576129f461561b565b602002602001015116613f0e565b8a602001518481518110612a1857612a1861561b565b6020026020010151613f3990919063ffffffff16565b945080612a3a8161567f565b9150506127cb565b5050612a4d8361401d565b60975490935060ff16600081612a64576000612ae6565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c448feb86040518163ffffffff1660e01b8152600401602060405180830381865afa158015612ac2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612ae69190615a38565b905060005b8a811015613164578215612c46578963ffffffff16827f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663249a0c428f8f86818110612b4257612b4261561b565b60405160e085901b6001600160e01b031916815292013560f81c600483015250602401602060405180830381865afa158015612b82573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612ba69190615a38565b612bb09190615a51565b11612c465760405162461bcd60e51b81526020600482015260666024820152600080516020615b9283398151915260448201527f7265733a205374616b6552656769737472792075706461746573206d7573742060648201527f62652077697468696e207769746864726177616c44656c6179426c6f636b732060848201526577696e646f7760d01b60a482015260c401610695565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166368bccaac8d8d84818110612c8757612c8761561b565b9050013560f81c60f81b60f81c8c8c60a001518581518110612cab57612cab61561b565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa158015612d07573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612d2b9190615a69565b6001600160401b031916612d4e8a6040015183815181106115c9576115c961561b565b67ffffffffffffffff191614612dea5760405162461bcd60e51b81526020600482015260616024820152600080516020615b9283398151915260448201527f7265733a2071756f72756d41706b206861736820696e2073746f72616765206460648201527f6f6573206e6f74206d617463682070726f76696465642071756f72756d2061706084820152606b60f81b60a482015260c401610695565b612e1a89604001518281518110612e0357612e0361561b565b602002602001015187613a7390919063ffffffff16565b95507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c8294c568d8d84818110612e5d57612e5d61561b565b9050013560f81c60f81b60f81c8c8c60c001518581518110612e8157612e8161561b565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa158015612edd573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612f019190615778565b85602001518281518110612f1757612f1761561b565b6001600160601b03909216602092830291909101820152850151805182908110612f4357612f4361561b565b602002602001015185600001518281518110612f6157612f6161561b565b60200260200101906001600160601b031690816001600160601b0316815250506000805b8a602001515181101561314f57612fd986600001518281518110612fab57612fab61561b565b60200260200101518f8f86818110612fc557612fc561561b565b600192013560f81c9290921c811614919050565b1561313d577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663f2be94ae8f8f8681811061301f5761301f61561b565b9050013560f81c60f81b60f81c8e896020015185815181106130435761304361561b565b60200260200101518f60e0015188815181106130615761306161561b565b6020026020010151878151811061307a5761307a61561b565b60209081029190910101516040516001600160e01b031960e087901b16815260ff909416600485015263ffffffff92831660248501526044840191909152166064820152608401602060405180830381865afa1580156130de573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906131029190615778565b87518051859081106131165761311661561b565b6020026020010181815161312a9190615a94565b6001600160601b03169052506001909101905b806131478161567f565b915050612f85565b5050808061315c9061567f565b915050612aeb565b50505060008061317e8c868a606001518b60800151610c68565b91509150816131ef5760405162461bcd60e51b81526020600482015260436024820152600080516020615b9283398151915260448201527f7265733a2070616972696e6720707265636f6d70696c652063616c6c206661696064820152621b195960ea1b608482015260a401610695565b806132505760405162461bcd60e51b81526020600482015260396024820152600080516020615b9283398151915260448201527f7265733a207369676e617475726520697320696e76616c6964000000000000006064820152608401610695565b5050600087826020015160405160200161326b9291906157d7565b60408051808303601f190181529190528051602090910120929b929a509198505050505050505050565b61329d6140b8565b6132a76000614112565b565b60ce546001600160a01b0316331461330d5760405162461bcd60e51b815260206004820152602160248201527f5461736b2067656e657261746f72206d757374206265207468652063616c6c656044820152603960f91b6064820152608401610695565b6040805160808101825260608183018190526001600160a01b038716825263ffffffff438116602080850191909152908716918301919091528251601f8501829004820281018201909352838352909190849084908190840183828082843760009201919091525050505060408083019190915251613390908290602001615abc565b60408051601f19818403018152828252805160209182012060c9805463ffffffff908116600090815260ca90945293909220555416907f024dec6bd1d353599633784df48acfe0f307837dd3da13643a805fc55c6eb8f9906133f3908490615abc565b60405180910390a260c95461340f9063ffffffff1660016155ab565b60c9805463ffffffff191663ffffffff929092169190911790555050505050565b604080516001808252818301909252600091606091839160208083019080368337019050509050848160008151811061346b5761346b61561b565b60209081029190910101526040516361c8a12f60e11b81526000906001600160a01b0388169063c391425e906134a790889086906004016159b0565b600060405180830381865afa1580156134c4573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526134ec9190810190615869565b6000815181106134fe576134fe61561b565b60209081029190910101516040516304ec635160e01b81526004810188905263ffffffff87811660248301529091166044820181905291506000906001600160a01b038916906304ec635190606401602060405180830381865afa15801561356a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061358e9190615920565b6001600160c01b0316905060006135a482614164565b9050816135b28a838a610df2565b9550955050505050935093915050565b6135ca6140b8565b6001600160a01b03811661362f5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610695565b6106a781614112565b600054610100900460ff16158080156136585750600054600160ff909116105b806136725750303b158015613672575060005460ff166001145b6136d55760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610695565b6000805460ff1916600117905580156136f8576000805461ff0019166101001790555b613703856000614230565b61370c84614112565b60cd80546001600160a01b038086166001600160a01b03199283161790925560ce8054928516929091169190911790558015613782576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050505050565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156137dc573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061380091906153a5565b6001600160a01b0316336001600160a01b0316146138305760405162461bcd60e51b8152600401610695906153c2565b6066541981196066541916146138ae5760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e756e70617573653a20696e76616c696420617474656d7060448201527f7420746f2070617573652066756e6374696f6e616c69747900000000000000006064820152608401610695565b606681905560405181815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c906020016107de565b6001600160a01b0381166139735760405162461bcd60e51b815260206004820152604960248201527f5061757361626c652e5f73657450617573657252656769737472793a206e657760448201527f50617573657252656769737472792063616e6e6f7420626520746865207a65726064820152686f206164647265737360b81b608482015260a401610695565b606554604080516001600160a01b03928316815291831660208301527f6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6910160405180910390a1606580546001600160a01b0319166001600160a01b0392909216919091179055565b60408051808201909152600080825260208201526139f86145d1565b835181526020808501519082015260408082018490526000908360608460076107d05a03fa9050808015613a2b57613a2d565bfe5b5080613a6b5760405162461bcd60e51b815260206004820152600d60248201526c1958cb5b5d5b0b59985a5b1959609a1b6044820152606401610695565b505092915050565b6040805180820190915260008082526020820152613a8f6145ef565b835181526020808501518183015283516040808401919091529084015160608301526000908360808460066107d05a03fa9050808015613a2b575080613a6b5760405162461bcd60e51b815260206004820152600d60248201526c1958cb5859190b59985a5b1959609a1b6044820152606401610695565b613b0f61460d565b50604080516080810182527f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c28183019081527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed6060830152815281518083019092527f275dc4a288d1afb3cbb1ac09187524c7db36395df7be3b99e673b13a075a65ec82527f1d9befcd05a5323e6da4d435f3b617cdb3af83285c2df711ef39c01571827f9d60208381019190915281019190915290565b604080518082019091526000808252602082015260008080613bf7600080516020615b72833981519152866156c6565b90505b613c038161431a565b9093509150600080516020615b72833981519152828309831415613c3d576040805180820190915290815260208101919091529392505050565b600080516020615b72833981519152600182089050613bfa565b604080518082018252868152602080820186905282518084019093528683528201849052600091829190613c89614632565b60005b6002811015613e4e576000613ca2826006615660565b9050848260028110613cb657613cb661561b565b60200201515183613cc8836000615a51565b600c8110613cd857613cd861561b565b6020020152848260028110613cef57613cef61561b565b60200201516020015183826001613d069190615a51565b600c8110613d1657613d1661561b565b6020020152838260028110613d2d57613d2d61561b565b6020020151515183613d40836002615a51565b600c8110613d5057613d5061561b565b6020020152838260028110613d6757613d6761561b565b6020020151516001602002015183613d80836003615a51565b600c8110613d9057613d9061561b565b6020020152838260028110613da757613da761561b565b602002015160200151600060028110613dc257613dc261561b565b602002015183613dd3836004615a51565b600c8110613de357613de361561b565b6020020152838260028110613dfa57613dfa61561b565b602002015160200151600160028110613e1557613e1561561b565b602002015183613e26836005615a51565b600c8110613e3657613e3661561b565b60200201525080613e468161567f565b915050613c8c565b50613e57614651565b60006020826101808560088cfa9151919c9115159b50909950505050505050505050565b600080613e878461439c565b9050808360ff166001901b11613f055760405162461bcd60e51b815260206004820152603f60248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206269746d61702065786365656473206d61782076616c7565006064820152608401610695565b90505b92915050565b6000805b8215613f0857613f23600184615a21565b9092169180613f3181615b4f565b915050613f12565b60408051808201909152600080825260208201526102008261ffff1610613f955760405162461bcd60e51b815260206004820152601060248201526f7363616c61722d746f6f2d6c6172676560801b6044820152606401610695565b8161ffff1660011415613fa9575081613f08565b6040805180820190915260008082526020820181905284906001905b8161ffff168661ffff161061401257600161ffff871660ff83161c81161415613ff557613ff28484613a73565b93505b613fff8384613a73565b92506201fffe600192831b169101613fc5565b509195945050505050565b6040805180820190915260008082526020820152815115801561404257506020820151155b15614060575050604080518082019091526000808252602082015290565b604051806040016040528083600001518152602001600080516020615b72833981519152846020015161409391906156c6565b6140ab90600080516020615b72833981519152615a21565b905292915050565b919050565b6033546001600160a01b031633146132a75760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610695565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b606060008061417284613f0e565b61ffff166001600160401b0381111561418d5761418d614720565b6040519080825280601f01601f1916602001820160405280156141b7576020820181803683370190505b5090506000805b8251821080156141cf575061010081105b15614226576001811b935085841615614216578060f81b8383815181106141f8576141f861561b565b60200101906001600160f81b031916908160001a9053508160010191505b61421f8161567f565b90506141be565b5090949350505050565b6065546001600160a01b031615801561425157506001600160a01b03821615155b6142d35760405162461bcd60e51b815260206004820152604760248201527f5061757361626c652e5f696e697469616c697a655061757365723a205f696e6960448201527f7469616c697a6550617573657228292063616e206f6e6c792062652063616c6c6064820152666564206f6e636560c81b608482015260a401610695565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2614316826138e5565b5050565b60008080600080516020615b728339815191526003600080516020615b7283398151915286600080516020615b72833981519152888909090890506000614390827f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f52600080516020615b72833981519152614529565b91959194509092505050565b6000610100825111156144255760405162461bcd60e51b8152602060048201526044602482018190527f4269746d61705574696c732e6f72646572656442797465734172726179546f42908201527f69746d61703a206f7264657265644279746573417272617920697320746f6f206064820152636c6f6e6760e01b608482015260a401610695565b815161443357506000919050565b600080836000815181106144495761444961561b565b0160200151600160f89190911c81901b92505b8451811015614520578481815181106144775761447761561b565b0160200151600160f89190911c1b915082821161450c5760405162461bcd60e51b815260206004820152604760248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206f72646572656442797465734172726179206973206e6f74206064820152661bdc99195c995960ca1b608482015260a401610695565b918117916145198161567f565b905061445c565b50909392505050565b600080614534614651565b61453c61466f565b602080825281810181905260408201819052606082018890526080820187905260a082018690528260c08360056107d05a03fa9250828015613a2b5750826145c65760405162461bcd60e51b815260206004820152601a60248201527f424e3235342e6578704d6f643a2063616c6c206661696c7572650000000000006044820152606401610695565b505195945050505050565b60405180606001604052806003906020820280368337509192915050565b60405180608001604052806004906020820280368337509192915050565b604051806040016040528061462061468d565b815260200161462d61468d565b905290565b604051806101800160405280600c906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b6040518060c001604052806006906020820280368337509192915050565b60405180604001604052806002906020820280368337509192915050565b6001600160a01b03811681146106a757600080fd5b6000602082840312156146d257600080fd5b8135613f05816146ab565b6000602082840312156146ef57600080fd5b5035919050565b60006080828403121561470857600080fd5b50919050565b60006060828403121561470857600080fd5b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b038111828210171561475857614758614720565b60405290565b60405161010081016001600160401b038111828210171561475857614758614720565b604051601f8201601f191681016001600160401b03811182821017156147a9576147a9614720565b604052919050565b60006001600160401b038211156147ca576147ca614720565b5060051b60200190565b63ffffffff811681146106a757600080fd5b80356140b3816147d4565b600082601f83011261480257600080fd5b81356020614817614812836147b1565b614781565b82815260059290921b8401810191818101908684111561483657600080fd5b8286015b8481101561485a57803561484d816147d4565b835291830191830161483a565b509695505050505050565b60006040828403121561487757600080fd5b61487f614736565b9050813581526020820135602082015292915050565b600082601f8301126148a657600080fd5b813560206148b6614812836147b1565b82815260069290921b840181019181810190868411156148d557600080fd5b8286015b8481101561485a576148eb8882614865565b8352918301916040016148d9565b600082601f83011261490a57600080fd5b604051604081018181106001600160401b038211171561492c5761492c614720565b806040525080604084018581111561494357600080fd5b845b81811015614012578035835260209283019201614945565b60006080828403121561496f57600080fd5b614977614736565b905061498383836148f9565b815261499283604084016148f9565b602082015292915050565b600082601f8301126149ae57600080fd5b813560206149be614812836147b1565b82815260059290921b840181019181810190868411156149dd57600080fd5b8286015b8481101561485a5780356001600160401b03811115614a005760008081fd5b614a0e8986838b01016147f1565b8452509183019183016149e1565b60006101808284031215614a2f57600080fd5b614a3761475e565b905081356001600160401b0380821115614a5057600080fd5b614a5c858386016147f1565b83526020840135915080821115614a7257600080fd5b614a7e85838601614895565b60208401526040840135915080821115614a9757600080fd5b614aa385838601614895565b6040840152614ab5856060860161495d565b6060840152614ac78560e08601614865565b6080840152610120840135915080821115614ae157600080fd5b614aed858386016147f1565b60a0840152610140840135915080821115614b0757600080fd5b614b13858386016147f1565b60c0840152610160840135915080821115614b2d57600080fd5b50614b3a8482850161499d565b60e08301525092915050565b600080600060a08486031215614b5b57600080fd5b83356001600160401b0380821115614b7257600080fd5b614b7e878388016146f6565b9450614b8d876020880161470e565b93506080860135915080821115614ba357600080fd5b50614bb086828701614a1c565b9150509250925092565b6000806000806101208587031215614bd157600080fd5b84359350614be28660208701614865565b9250614bf1866060870161495d565b9150614c008660e08701614865565b905092959194509250565b600060208284031215614c1d57600080fd5b8135613f05816147d4565b600080600060608486031215614c3d57600080fd5b8335614c48816146ab565b92506020848101356001600160401b0380821115614c6557600080fd5b818701915087601f830112614c7957600080fd5b813581811115614c8b57614c8b614720565b614c9d601f8201601f19168501614781565b91508082528884828501011115614cb357600080fd5b8084840185840137600084828401015250809450505050614cd6604085016147e6565b90509250925092565b6000815180845260208085019450848260051b86018286016000805b86811015614d72578484038a52825180518086529087019087860190845b81811015614d5d57835180516001600160a01b031684528a8101518b8501526040908101516001600160601b03169084015292890192606090920191600101614d19565b50509a87019a94505091850191600101614cfb565b509198975050505050505050565b602081526000614d936020830184614cdf565b9392505050565b60008060008084860360e0811215614db157600080fd5b85356001600160401b0380821115614dc857600080fd5b614dd489838a016146f6565b9650614de38960208a0161470e565b95506040607f1984011215614df757600080fd5b60808801945060c0880135925080831115614e1157600080fd5b5050614e1f87828801614895565b91505092959194509250565b80151581146106a757600080fd5b600060208284031215614e4b57600080fd5b8135613f0581614e2b565b60008083601f840112614e6857600080fd5b5081356001600160401b03811115614e7f57600080fd5b602083019150836020828501011115614e9757600080fd5b9250929050565b60008060008060008060808789031215614eb757600080fd5b8635614ec2816146ab565b95506020870135614ed2816147d4565b945060408701356001600160401b0380821115614eee57600080fd5b614efa8a838b01614e56565b90965094506060890135915080821115614f1357600080fd5b818901915089601f830112614f2757600080fd5b813581811115614f3657600080fd5b8a60208260051b8501011115614f4b57600080fd5b6020830194508093505050509295509295509295565b600081518084526020808501945080840160005b83811015614f9757815163ffffffff1687529582019590820190600101614f75565b509495945050505050565b600060208083528351608082850152614fbe60a0850182614f61565b905081850151601f1980868403016040870152614fdb8383614f61565b92506040870151915080868403016060870152614ff88383614f61565b60608801518782038301608089015280518083529194508501925084840190600581901b8501860160005b8281101561504f578487830301845261503d828751614f61565b95880195938801939150600101615023565b509998505050505050505050565b60ff811681146106a757600080fd5b60006020828403121561507e57600080fd5b8135613f058161505d565b60008060006060848603121561509e57600080fd5b83356150a9816146ab565b92506020848101356001600160401b038111156150c557600080fd5b8501601f810187136150d657600080fd5b80356150e4614812826147b1565b81815260059190911b8201830190838101908983111561510357600080fd5b928401925b8284101561512157833582529284019290840190615108565b8096505050505050614cd6604085016147e6565b6020808252825182820181905260009190848201906040850190845b8181101561516d57835183529284019291840191600101615151565b50909695505050505050565b60008060008060006080868803121561519157600080fd5b8535945060208601356001600160401b03808211156151af57600080fd5b6151bb89838a01614e56565b9096509450604088013591506151d0826147d4565b909250606087013590808211156151e657600080fd5b506151f388828901614a1c565b9150509295509295909350565b600081518084526020808501945080840160005b83811015614f975781516001600160601b031687529582019590820190600101615214565b60408152600083516040808401526152546080840182615200565b90506020850151603f198483030160608501526152718282615200565b925050508260208301529392505050565b6000806000806060858703121561529857600080fd5b84356152a3816146ab565b935060208501356152b3816147d4565b925060408501356001600160401b038111156152ce57600080fd5b6152da87828801614e56565b95989497509550505050565b6000806000606084860312156152fb57600080fd5b8335615306816146ab565b925060208401359150604084013561531d816147d4565b809150509250925092565b8281526040602082015260006153416040830184614cdf565b949350505050565b6000806000806080858703121561535f57600080fd5b843561536a816146ab565b9350602085013561537a816146ab565b9250604085013561538a816146ab565b9150606085013561539a816146ab565b939692955090935050565b6000602082840312156153b757600080fd5b8151613f05816146ab565b6020808252602a908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526939903ab73830bab9b2b960b11b606082015260800190565b60006020828403121561541e57600080fd5b8151613f0581614e2b565b60208082526028908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526739903830bab9b2b960c11b606082015260800190565b6000808335601e1984360301811261548857600080fd5b8301803591506001600160401b038211156154a257600080fd5b602001915036819003821315614e9757600080fd5b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b60208152600082356154f1816146ab565b6001600160a01b031660208381019190915283013561550f816147d4565b63ffffffff81166040840152506040830135601e1984360301811261553357600080fd5b830180356001600160401b0381111561554b57600080fd5b80360385131561555a57600080fd5b6080606085015261557260a0850182602085016154b7565b915050615581606085016147e6565b63ffffffff81166080850152509392505050565b634e487b7160e01b600052601160045260246000fd5b600063ffffffff8083168185168083038211156155ca576155ca615595565b01949350505050565b80356155de816147d4565b63ffffffff16825260208101356155f4816146ab565b6001600160a01b03166020830152604090810135910152565b60608101613f0882846155d3565b634e487b7160e01b600052603260045260246000fd5b60006001600160601b038083168185168183048111821515161561565757615657615595565b02949350505050565b600081600019048311821515161561567a5761567a615595565b500290565b600060001982141561569357615693615595565b5060010190565b60a081016156a882856155d3565b63ffffffff8351166060830152602083015160808301529392505050565b6000826156e357634e487b7160e01b600052601260045260246000fd5b500690565b600060208083850312156156fb57600080fd5b82516001600160401b0381111561571157600080fd5b8301601f8101851361572257600080fd5b8051615730614812826147b1565b81815260059190911b8201830190838101908783111561574f57600080fd5b928401925b8284101561576d57835182529284019290840190615754565b979650505050505050565b60006020828403121561578a57600080fd5b81516001600160601b0381168114613f0557600080fd5b60a081016157af82856155d3565b82356157ba816147d4565b63ffffffff16606083015260209290920135608090910152919050565b63ffffffff60e01b8360e01b1681526000600482018351602080860160005b83811015615812578151855293820193908201906001016157f6565b5092979650505050505050565b63ffffffff84168152604060208201819052810182905260006001600160fb1b0383111561584c57600080fd5b8260051b8085606085013760009201606001918252509392505050565b6000602080838503121561587c57600080fd5b82516001600160401b0381111561589257600080fd5b8301601f810185136158a357600080fd5b80516158b1614812826147b1565b81815260059190911b820183019083810190878311156158d057600080fd5b928401925b8284101561576d5783516158e8816147d4565b825292840192908401906158d5565b63ffffffff841681526040602082015260006159176040830184866154b7565b95945050505050565b60006020828403121561593257600080fd5b81516001600160c01b0381168114613f0557600080fd5b60006020828403121561595b57600080fd5b8151613f05816147d4565b600060ff821660ff81141561597d5761597d615595565b60010192915050565b60408152600061599a6040830185876154b7565b905063ffffffff83166020830152949350505050565b60006040820163ffffffff851683526020604081850152818551808452606086019150828701935060005b818110156159f7578451835293830193918301916001016159db565b5090979650505050505050565b600060208284031215615a1657600080fd5b8151613f058161505d565b600082821015615a3357615a33615595565b500390565b600060208284031215615a4a57600080fd5b5051919050565b60008219821115615a6457615a64615595565b500190565b600060208284031215615a7b57600080fd5b815167ffffffffffffffff1981168114613f0557600080fd5b60006001600160601b0383811690831681811015615ab457615ab4615595565b039392505050565b6000602080835260018060a01b038451168184015263ffffffff8185015116604084015260408401516080606085015280518060a086015260005b81811015615b135782810184015186820160c001528301615af7565b81811115615b2557600060c083880101525b50606086015163ffffffff811660808701529250601f01601f19169390930160c001949350505050565b600061ffff80831681811415615b6757615b67615595565b600101939250505056fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47424c535369676e6174757265436865636b65722e636865636b5369676e617475a2646970667358221220fa55d8e3409db9797d324212e2b705719959a42c629a039285738cb8be3eabaf64736f6c634300080c0033",
}

// ContractConnect3DATaskManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractConnect3DATaskManagerMetaData.ABI instead.
var ContractConnect3DATaskManagerABI = ContractConnect3DATaskManagerMetaData.ABI

// ContractConnect3DATaskManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractConnect3DATaskManagerMetaData.Bin instead.
var ContractConnect3DATaskManagerBin = ContractConnect3DATaskManagerMetaData.Bin

// DeployContractConnect3DATaskManager deploys a new Ethereum contract, binding an instance of ContractConnect3DATaskManager to it.
func DeployContractConnect3DATaskManager(auth *bind.TransactOpts, backend bind.ContractBackend, _registryCoordinator common.Address, _taskResponseWindowBlock uint32) (common.Address, *types.Transaction, *ContractConnect3DATaskManager, error) {
	parsed, err := ContractConnect3DATaskManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractConnect3DATaskManagerBin), backend, _registryCoordinator, _taskResponseWindowBlock)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractConnect3DATaskManager{ContractConnect3DATaskManagerCaller: ContractConnect3DATaskManagerCaller{contract: contract}, ContractConnect3DATaskManagerTransactor: ContractConnect3DATaskManagerTransactor{contract: contract}, ContractConnect3DATaskManagerFilterer: ContractConnect3DATaskManagerFilterer{contract: contract}}, nil
}

// ContractConnect3DATaskManager is an auto generated Go binding around an Ethereum contract.
type ContractConnect3DATaskManager struct {
	ContractConnect3DATaskManagerCaller     // Read-only binding to the contract
	ContractConnect3DATaskManagerTransactor // Write-only binding to the contract
	ContractConnect3DATaskManagerFilterer   // Log filterer for contract events
}

// ContractConnect3DATaskManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractConnect3DATaskManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractConnect3DATaskManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractConnect3DATaskManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractConnect3DATaskManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractConnect3DATaskManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractConnect3DATaskManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractConnect3DATaskManagerSession struct {
	Contract     *ContractConnect3DATaskManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                  // Call options to use throughout this session
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// ContractConnect3DATaskManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractConnect3DATaskManagerCallerSession struct {
	Contract *ContractConnect3DATaskManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                        // Call options to use throughout this session
}

// ContractConnect3DATaskManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractConnect3DATaskManagerTransactorSession struct {
	Contract     *ContractConnect3DATaskManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                        // Transaction auth options to use throughout this session
}

// ContractConnect3DATaskManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractConnect3DATaskManagerRaw struct {
	Contract *ContractConnect3DATaskManager // Generic contract binding to access the raw methods on
}

// ContractConnect3DATaskManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractConnect3DATaskManagerCallerRaw struct {
	Contract *ContractConnect3DATaskManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ContractConnect3DATaskManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractConnect3DATaskManagerTransactorRaw struct {
	Contract *ContractConnect3DATaskManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractConnect3DATaskManager creates a new instance of ContractConnect3DATaskManager, bound to a specific deployed contract.
func NewContractConnect3DATaskManager(address common.Address, backend bind.ContractBackend) (*ContractConnect3DATaskManager, error) {
	contract, err := bindContractConnect3DATaskManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractConnect3DATaskManager{ContractConnect3DATaskManagerCaller: ContractConnect3DATaskManagerCaller{contract: contract}, ContractConnect3DATaskManagerTransactor: ContractConnect3DATaskManagerTransactor{contract: contract}, ContractConnect3DATaskManagerFilterer: ContractConnect3DATaskManagerFilterer{contract: contract}}, nil
}

// NewContractConnect3DATaskManagerCaller creates a new read-only instance of ContractConnect3DATaskManager, bound to a specific deployed contract.
func NewContractConnect3DATaskManagerCaller(address common.Address, caller bind.ContractCaller) (*ContractConnect3DATaskManagerCaller, error) {
	contract, err := bindContractConnect3DATaskManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractConnect3DATaskManagerCaller{contract: contract}, nil
}

// NewContractConnect3DATaskManagerTransactor creates a new write-only instance of ContractConnect3DATaskManager, bound to a specific deployed contract.
func NewContractConnect3DATaskManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractConnect3DATaskManagerTransactor, error) {
	contract, err := bindContractConnect3DATaskManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractConnect3DATaskManagerTransactor{contract: contract}, nil
}

// NewContractConnect3DATaskManagerFilterer creates a new log filterer instance of ContractConnect3DATaskManager, bound to a specific deployed contract.
func NewContractConnect3DATaskManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractConnect3DATaskManagerFilterer, error) {
	contract, err := bindContractConnect3DATaskManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractConnect3DATaskManagerFilterer{contract: contract}, nil
}

// bindContractConnect3DATaskManager binds a generic wrapper to an already deployed contract.
func bindContractConnect3DATaskManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractConnect3DATaskManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractConnect3DATaskManager.Contract.ContractConnect3DATaskManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractConnect3DATaskManager.Contract.ContractConnect3DATaskManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractConnect3DATaskManager.Contract.ContractConnect3DATaskManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractConnect3DATaskManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractConnect3DATaskManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractConnect3DATaskManager.Contract.contract.Transact(opts, method, params...)
}

// TASKCHALLENGEWINDOWBLOCK is a free data retrieval call binding the contract method 0xf63c5bab.
//
// Solidity: function TASK_CHALLENGE_WINDOW_BLOCK() view returns(uint32)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerCaller) TASKCHALLENGEWINDOWBLOCK(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractConnect3DATaskManager.contract.Call(opts, &out, "TASK_CHALLENGE_WINDOW_BLOCK")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TASKCHALLENGEWINDOWBLOCK is a free data retrieval call binding the contract method 0xf63c5bab.
//
// Solidity: function TASK_CHALLENGE_WINDOW_BLOCK() view returns(uint32)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerSession) TASKCHALLENGEWINDOWBLOCK() (uint32, error) {
	return _ContractConnect3DATaskManager.Contract.TASKCHALLENGEWINDOWBLOCK(&_ContractConnect3DATaskManager.CallOpts)
}

// TASKCHALLENGEWINDOWBLOCK is a free data retrieval call binding the contract method 0xf63c5bab.
//
// Solidity: function TASK_CHALLENGE_WINDOW_BLOCK() view returns(uint32)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerCallerSession) TASKCHALLENGEWINDOWBLOCK() (uint32, error) {
	return _ContractConnect3DATaskManager.Contract.TASKCHALLENGEWINDOWBLOCK(&_ContractConnect3DATaskManager.CallOpts)
}

// TASKRESPONSEWINDOWBLOCK is a free data retrieval call binding the contract method 0x1ad43189.
//
// Solidity: function TASK_RESPONSE_WINDOW_BLOCK() view returns(uint32)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerCaller) TASKRESPONSEWINDOWBLOCK(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractConnect3DATaskManager.contract.Call(opts, &out, "TASK_RESPONSE_WINDOW_BLOCK")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TASKRESPONSEWINDOWBLOCK is a free data retrieval call binding the contract method 0x1ad43189.
//
// Solidity: function TASK_RESPONSE_WINDOW_BLOCK() view returns(uint32)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerSession) TASKRESPONSEWINDOWBLOCK() (uint32, error) {
	return _ContractConnect3DATaskManager.Contract.TASKRESPONSEWINDOWBLOCK(&_ContractConnect3DATaskManager.CallOpts)
}

// TASKRESPONSEWINDOWBLOCK is a free data retrieval call binding the contract method 0x1ad43189.
//
// Solidity: function TASK_RESPONSE_WINDOW_BLOCK() view returns(uint32)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerCallerSession) TASKRESPONSEWINDOWBLOCK() (uint32, error) {
	return _ContractConnect3DATaskManager.Contract.TASKRESPONSEWINDOWBLOCK(&_ContractConnect3DATaskManager.CallOpts)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerCaller) Aggregator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractConnect3DATaskManager.contract.Call(opts, &out, "aggregator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerSession) Aggregator() (common.Address, error) {
	return _ContractConnect3DATaskManager.Contract.Aggregator(&_ContractConnect3DATaskManager.CallOpts)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerCallerSession) Aggregator() (common.Address, error) {
	return _ContractConnect3DATaskManager.Contract.Aggregator(&_ContractConnect3DATaskManager.CallOpts)
}

// AllTaskHashes is a free data retrieval call binding the contract method 0x2d89f6fc.
//
// Solidity: function allTaskHashes(uint32 ) view returns(bytes32)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerCaller) AllTaskHashes(opts *bind.CallOpts, arg0 uint32) ([32]byte, error) {
	var out []interface{}
	err := _ContractConnect3DATaskManager.contract.Call(opts, &out, "allTaskHashes", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AllTaskHashes is a free data retrieval call binding the contract method 0x2d89f6fc.
//
// Solidity: function allTaskHashes(uint32 ) view returns(bytes32)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerSession) AllTaskHashes(arg0 uint32) ([32]byte, error) {
	return _ContractConnect3DATaskManager.Contract.AllTaskHashes(&_ContractConnect3DATaskManager.CallOpts, arg0)
}

// AllTaskHashes is a free data retrieval call binding the contract method 0x2d89f6fc.
//
// Solidity: function allTaskHashes(uint32 ) view returns(bytes32)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerCallerSession) AllTaskHashes(arg0 uint32) ([32]byte, error) {
	return _ContractConnect3DATaskManager.Contract.AllTaskHashes(&_ContractConnect3DATaskManager.CallOpts, arg0)
}

// AllTaskResponses is a free data retrieval call binding the contract method 0x2cb223d5.
//
// Solidity: function allTaskResponses(uint32 ) view returns(bytes32)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerCaller) AllTaskResponses(opts *bind.CallOpts, arg0 uint32) ([32]byte, error) {
	var out []interface{}
	err := _ContractConnect3DATaskManager.contract.Call(opts, &out, "allTaskResponses", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AllTaskResponses is a free data retrieval call binding the contract method 0x2cb223d5.
//
// Solidity: function allTaskResponses(uint32 ) view returns(bytes32)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerSession) AllTaskResponses(arg0 uint32) ([32]byte, error) {
	return _ContractConnect3DATaskManager.Contract.AllTaskResponses(&_ContractConnect3DATaskManager.CallOpts, arg0)
}

// AllTaskResponses is a free data retrieval call binding the contract method 0x2cb223d5.
//
// Solidity: function allTaskResponses(uint32 ) view returns(bytes32)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerCallerSession) AllTaskResponses(arg0 uint32) ([32]byte, error) {
	return _ContractConnect3DATaskManager.Contract.AllTaskResponses(&_ContractConnect3DATaskManager.CallOpts, arg0)
}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerCaller) BlsApkRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractConnect3DATaskManager.contract.Call(opts, &out, "blsApkRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerSession) BlsApkRegistry() (common.Address, error) {
	return _ContractConnect3DATaskManager.Contract.BlsApkRegistry(&_ContractConnect3DATaskManager.CallOpts)
}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerCallerSession) BlsApkRegistry() (common.Address, error) {
	return _ContractConnect3DATaskManager.Contract.BlsApkRegistry(&_ContractConnect3DATaskManager.CallOpts)
}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) params) view returns((uint96[],uint96[]), bytes32)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerCaller) CheckSignatures(opts *bind.CallOpts, msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	var out []interface{}
	err := _ContractConnect3DATaskManager.contract.Call(opts, &out, "checkSignatures", msgHash, quorumNumbers, referenceBlockNumber, params)

	if err != nil {
		return *new(IBLSSignatureCheckerQuorumStakeTotals), *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(IBLSSignatureCheckerQuorumStakeTotals)).(*IBLSSignatureCheckerQuorumStakeTotals)
	out1 := *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return out0, out1, err

}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) params) view returns((uint96[],uint96[]), bytes32)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerSession) CheckSignatures(msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	return _ContractConnect3DATaskManager.Contract.CheckSignatures(&_ContractConnect3DATaskManager.CallOpts, msgHash, quorumNumbers, referenceBlockNumber, params)
}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) params) view returns((uint96[],uint96[]), bytes32)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerCallerSession) CheckSignatures(msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	return _ContractConnect3DATaskManager.Contract.CheckSignatures(&_ContractConnect3DATaskManager.CallOpts, msgHash, quorumNumbers, referenceBlockNumber, params)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerCaller) Delegation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractConnect3DATaskManager.contract.Call(opts, &out, "delegation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerSession) Delegation() (common.Address, error) {
	return _ContractConnect3DATaskManager.Contract.Delegation(&_ContractConnect3DATaskManager.CallOpts)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerCallerSession) Delegation() (common.Address, error) {
	return _ContractConnect3DATaskManager.Contract.Delegation(&_ContractConnect3DATaskManager.CallOpts)
}

// Generator is a free data retrieval call binding the contract method 0x7afa1eed.
//
// Solidity: function generator() view returns(address)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerCaller) Generator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractConnect3DATaskManager.contract.Call(opts, &out, "generator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Generator is a free data retrieval call binding the contract method 0x7afa1eed.
//
// Solidity: function generator() view returns(address)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerSession) Generator() (common.Address, error) {
	return _ContractConnect3DATaskManager.Contract.Generator(&_ContractConnect3DATaskManager.CallOpts)
}

// Generator is a free data retrieval call binding the contract method 0x7afa1eed.
//
// Solidity: function generator() view returns(address)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerCallerSession) Generator() (common.Address, error) {
	return _ContractConnect3DATaskManager.Contract.Generator(&_ContractConnect3DATaskManager.CallOpts)
}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerCaller) GetCheckSignaturesIndices(opts *bind.CallOpts, registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	var out []interface{}
	err := _ContractConnect3DATaskManager.contract.Call(opts, &out, "getCheckSignaturesIndices", registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)

	if err != nil {
		return *new(OperatorStateRetrieverCheckSignaturesIndices), err
	}

	out0 := *abi.ConvertType(out[0], new(OperatorStateRetrieverCheckSignaturesIndices)).(*OperatorStateRetrieverCheckSignaturesIndices)

	return out0, err

}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerSession) GetCheckSignaturesIndices(registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	return _ContractConnect3DATaskManager.Contract.GetCheckSignaturesIndices(&_ContractConnect3DATaskManager.CallOpts, registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)
}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerCallerSession) GetCheckSignaturesIndices(registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	return _ContractConnect3DATaskManager.Contract.GetCheckSignaturesIndices(&_ContractConnect3DATaskManager.CallOpts, registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)
}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerCaller) GetOperatorState(opts *bind.CallOpts, registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	var out []interface{}
	err := _ContractConnect3DATaskManager.contract.Call(opts, &out, "getOperatorState", registryCoordinator, quorumNumbers, blockNumber)

	if err != nil {
		return *new([][]OperatorStateRetrieverOperator), err
	}

	out0 := *abi.ConvertType(out[0], new([][]OperatorStateRetrieverOperator)).(*[][]OperatorStateRetrieverOperator)

	return out0, err

}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerSession) GetOperatorState(registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	return _ContractConnect3DATaskManager.Contract.GetOperatorState(&_ContractConnect3DATaskManager.CallOpts, registryCoordinator, quorumNumbers, blockNumber)
}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerCallerSession) GetOperatorState(registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	return _ContractConnect3DATaskManager.Contract.GetOperatorState(&_ContractConnect3DATaskManager.CallOpts, registryCoordinator, quorumNumbers, blockNumber)
}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (address,bytes32,uint96)[][])
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerCaller) GetOperatorState0(opts *bind.CallOpts, registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	var out []interface{}
	err := _ContractConnect3DATaskManager.contract.Call(opts, &out, "getOperatorState0", registryCoordinator, operatorId, blockNumber)

	if err != nil {
		return *new(*big.Int), *new([][]OperatorStateRetrieverOperator), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new([][]OperatorStateRetrieverOperator)).(*[][]OperatorStateRetrieverOperator)

	return out0, out1, err

}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (address,bytes32,uint96)[][])
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerSession) GetOperatorState0(registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	return _ContractConnect3DATaskManager.Contract.GetOperatorState0(&_ContractConnect3DATaskManager.CallOpts, registryCoordinator, operatorId, blockNumber)
}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (address,bytes32,uint96)[][])
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerCallerSession) GetOperatorState0(registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	return _ContractConnect3DATaskManager.Contract.GetOperatorState0(&_ContractConnect3DATaskManager.CallOpts, registryCoordinator, operatorId, blockNumber)
}

// GetQuorumBitmapsAtBlockNumber is a free data retrieval call binding the contract method 0x5c155662.
//
// Solidity: function getQuorumBitmapsAtBlockNumber(address registryCoordinator, bytes32[] operatorIds, uint32 blockNumber) view returns(uint256[])
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerCaller) GetQuorumBitmapsAtBlockNumber(opts *bind.CallOpts, registryCoordinator common.Address, operatorIds [][32]byte, blockNumber uint32) ([]*big.Int, error) {
	var out []interface{}
	err := _ContractConnect3DATaskManager.contract.Call(opts, &out, "getQuorumBitmapsAtBlockNumber", registryCoordinator, operatorIds, blockNumber)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetQuorumBitmapsAtBlockNumber is a free data retrieval call binding the contract method 0x5c155662.
//
// Solidity: function getQuorumBitmapsAtBlockNumber(address registryCoordinator, bytes32[] operatorIds, uint32 blockNumber) view returns(uint256[])
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerSession) GetQuorumBitmapsAtBlockNumber(registryCoordinator common.Address, operatorIds [][32]byte, blockNumber uint32) ([]*big.Int, error) {
	return _ContractConnect3DATaskManager.Contract.GetQuorumBitmapsAtBlockNumber(&_ContractConnect3DATaskManager.CallOpts, registryCoordinator, operatorIds, blockNumber)
}

// GetQuorumBitmapsAtBlockNumber is a free data retrieval call binding the contract method 0x5c155662.
//
// Solidity: function getQuorumBitmapsAtBlockNumber(address registryCoordinator, bytes32[] operatorIds, uint32 blockNumber) view returns(uint256[])
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerCallerSession) GetQuorumBitmapsAtBlockNumber(registryCoordinator common.Address, operatorIds [][32]byte, blockNumber uint32) ([]*big.Int, error) {
	return _ContractConnect3DATaskManager.Contract.GetQuorumBitmapsAtBlockNumber(&_ContractConnect3DATaskManager.CallOpts, registryCoordinator, operatorIds, blockNumber)
}

// GetTaskResponseWindowBlock is a free data retrieval call binding the contract method 0xf5c9899d.
//
// Solidity: function getTaskResponseWindowBlock() view returns(uint32)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerCaller) GetTaskResponseWindowBlock(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractConnect3DATaskManager.contract.Call(opts, &out, "getTaskResponseWindowBlock")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetTaskResponseWindowBlock is a free data retrieval call binding the contract method 0xf5c9899d.
//
// Solidity: function getTaskResponseWindowBlock() view returns(uint32)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerSession) GetTaskResponseWindowBlock() (uint32, error) {
	return _ContractConnect3DATaskManager.Contract.GetTaskResponseWindowBlock(&_ContractConnect3DATaskManager.CallOpts)
}

// GetTaskResponseWindowBlock is a free data retrieval call binding the contract method 0xf5c9899d.
//
// Solidity: function getTaskResponseWindowBlock() view returns(uint32)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerCallerSession) GetTaskResponseWindowBlock() (uint32, error) {
	return _ContractConnect3DATaskManager.Contract.GetTaskResponseWindowBlock(&_ContractConnect3DATaskManager.CallOpts)
}

// LatestTaskNum is a free data retrieval call binding the contract method 0x8b00ce7c.
//
// Solidity: function latestTaskNum() view returns(uint32)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerCaller) LatestTaskNum(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractConnect3DATaskManager.contract.Call(opts, &out, "latestTaskNum")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LatestTaskNum is a free data retrieval call binding the contract method 0x8b00ce7c.
//
// Solidity: function latestTaskNum() view returns(uint32)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerSession) LatestTaskNum() (uint32, error) {
	return _ContractConnect3DATaskManager.Contract.LatestTaskNum(&_ContractConnect3DATaskManager.CallOpts)
}

// LatestTaskNum is a free data retrieval call binding the contract method 0x8b00ce7c.
//
// Solidity: function latestTaskNum() view returns(uint32)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerCallerSession) LatestTaskNum() (uint32, error) {
	return _ContractConnect3DATaskManager.Contract.LatestTaskNum(&_ContractConnect3DATaskManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractConnect3DATaskManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerSession) Owner() (common.Address, error) {
	return _ContractConnect3DATaskManager.Contract.Owner(&_ContractConnect3DATaskManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerCallerSession) Owner() (common.Address, error) {
	return _ContractConnect3DATaskManager.Contract.Owner(&_ContractConnect3DATaskManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerCaller) Paused(opts *bind.CallOpts, index uint8) (bool, error) {
	var out []interface{}
	err := _ContractConnect3DATaskManager.contract.Call(opts, &out, "paused", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerSession) Paused(index uint8) (bool, error) {
	return _ContractConnect3DATaskManager.Contract.Paused(&_ContractConnect3DATaskManager.CallOpts, index)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerCallerSession) Paused(index uint8) (bool, error) {
	return _ContractConnect3DATaskManager.Contract.Paused(&_ContractConnect3DATaskManager.CallOpts, index)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerCaller) Paused0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractConnect3DATaskManager.contract.Call(opts, &out, "paused0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerSession) Paused0() (*big.Int, error) {
	return _ContractConnect3DATaskManager.Contract.Paused0(&_ContractConnect3DATaskManager.CallOpts)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerCallerSession) Paused0() (*big.Int, error) {
	return _ContractConnect3DATaskManager.Contract.Paused0(&_ContractConnect3DATaskManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerCaller) PauserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractConnect3DATaskManager.contract.Call(opts, &out, "pauserRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerSession) PauserRegistry() (common.Address, error) {
	return _ContractConnect3DATaskManager.Contract.PauserRegistry(&_ContractConnect3DATaskManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerCallerSession) PauserRegistry() (common.Address, error) {
	return _ContractConnect3DATaskManager.Contract.PauserRegistry(&_ContractConnect3DATaskManager.CallOpts)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerCaller) RegistryCoordinator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractConnect3DATaskManager.contract.Call(opts, &out, "registryCoordinator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerSession) RegistryCoordinator() (common.Address, error) {
	return _ContractConnect3DATaskManager.Contract.RegistryCoordinator(&_ContractConnect3DATaskManager.CallOpts)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerCallerSession) RegistryCoordinator() (common.Address, error) {
	return _ContractConnect3DATaskManager.Contract.RegistryCoordinator(&_ContractConnect3DATaskManager.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerCaller) StakeRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractConnect3DATaskManager.contract.Call(opts, &out, "stakeRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerSession) StakeRegistry() (common.Address, error) {
	return _ContractConnect3DATaskManager.Contract.StakeRegistry(&_ContractConnect3DATaskManager.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerCallerSession) StakeRegistry() (common.Address, error) {
	return _ContractConnect3DATaskManager.Contract.StakeRegistry(&_ContractConnect3DATaskManager.CallOpts)
}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerCaller) StaleStakesForbidden(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ContractConnect3DATaskManager.contract.Call(opts, &out, "staleStakesForbidden")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerSession) StaleStakesForbidden() (bool, error) {
	return _ContractConnect3DATaskManager.Contract.StaleStakesForbidden(&_ContractConnect3DATaskManager.CallOpts)
}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerCallerSession) StaleStakesForbidden() (bool, error) {
	return _ContractConnect3DATaskManager.Contract.StaleStakesForbidden(&_ContractConnect3DATaskManager.CallOpts)
}

// TaskNumber is a free data retrieval call binding the contract method 0x72d18e8d.
//
// Solidity: function taskNumber() view returns(uint32)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerCaller) TaskNumber(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractConnect3DATaskManager.contract.Call(opts, &out, "taskNumber")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TaskNumber is a free data retrieval call binding the contract method 0x72d18e8d.
//
// Solidity: function taskNumber() view returns(uint32)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerSession) TaskNumber() (uint32, error) {
	return _ContractConnect3DATaskManager.Contract.TaskNumber(&_ContractConnect3DATaskManager.CallOpts)
}

// TaskNumber is a free data retrieval call binding the contract method 0x72d18e8d.
//
// Solidity: function taskNumber() view returns(uint32)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerCallerSession) TaskNumber() (uint32, error) {
	return _ContractConnect3DATaskManager.Contract.TaskNumber(&_ContractConnect3DATaskManager.CallOpts)
}

// TaskSuccesfullyChallenged is a free data retrieval call binding the contract method 0x5decc3f5.
//
// Solidity: function taskSuccesfullyChallenged(uint32 ) view returns(bool)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerCaller) TaskSuccesfullyChallenged(opts *bind.CallOpts, arg0 uint32) (bool, error) {
	var out []interface{}
	err := _ContractConnect3DATaskManager.contract.Call(opts, &out, "taskSuccesfullyChallenged", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TaskSuccesfullyChallenged is a free data retrieval call binding the contract method 0x5decc3f5.
//
// Solidity: function taskSuccesfullyChallenged(uint32 ) view returns(bool)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerSession) TaskSuccesfullyChallenged(arg0 uint32) (bool, error) {
	return _ContractConnect3DATaskManager.Contract.TaskSuccesfullyChallenged(&_ContractConnect3DATaskManager.CallOpts, arg0)
}

// TaskSuccesfullyChallenged is a free data retrieval call binding the contract method 0x5decc3f5.
//
// Solidity: function taskSuccesfullyChallenged(uint32 ) view returns(bool)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerCallerSession) TaskSuccesfullyChallenged(arg0 uint32) (bool, error) {
	return _ContractConnect3DATaskManager.Contract.TaskSuccesfullyChallenged(&_ContractConnect3DATaskManager.CallOpts, arg0)
}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerCaller) TrySignatureAndApkVerification(opts *bind.CallOpts, msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	var out []interface{}
	err := _ContractConnect3DATaskManager.contract.Call(opts, &out, "trySignatureAndApkVerification", msgHash, apk, apkG2, sigma)

	outstruct := new(struct {
		PairingSuccessful bool
		SiganatureIsValid bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PairingSuccessful = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.SiganatureIsValid = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerSession) TrySignatureAndApkVerification(msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	return _ContractConnect3DATaskManager.Contract.TrySignatureAndApkVerification(&_ContractConnect3DATaskManager.CallOpts, msgHash, apk, apkG2, sigma)
}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerCallerSession) TrySignatureAndApkVerification(msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	return _ContractConnect3DATaskManager.Contract.TrySignatureAndApkVerification(&_ContractConnect3DATaskManager.CallOpts, msgHash, apk, apkG2, sigma)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0x91085183.
//
// Solidity: function createNewTask(address addressToFetchC3Data, uint32 quorumThresholdPercentage, bytes quorumNumbers) returns()
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerTransactor) CreateNewTask(opts *bind.TransactOpts, addressToFetchC3Data common.Address, quorumThresholdPercentage uint32, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractConnect3DATaskManager.contract.Transact(opts, "createNewTask", addressToFetchC3Data, quorumThresholdPercentage, quorumNumbers)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0x91085183.
//
// Solidity: function createNewTask(address addressToFetchC3Data, uint32 quorumThresholdPercentage, bytes quorumNumbers) returns()
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerSession) CreateNewTask(addressToFetchC3Data common.Address, quorumThresholdPercentage uint32, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractConnect3DATaskManager.Contract.CreateNewTask(&_ContractConnect3DATaskManager.TransactOpts, addressToFetchC3Data, quorumThresholdPercentage, quorumNumbers)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0x91085183.
//
// Solidity: function createNewTask(address addressToFetchC3Data, uint32 quorumThresholdPercentage, bytes quorumNumbers) returns()
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerTransactorSession) CreateNewTask(addressToFetchC3Data common.Address, quorumThresholdPercentage uint32, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractConnect3DATaskManager.Contract.CreateNewTask(&_ContractConnect3DATaskManager.TransactOpts, addressToFetchC3Data, quorumThresholdPercentage, quorumNumbers)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _pauserRegistry, address initialOwner, address _aggregator, address _generator) returns()
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerTransactor) Initialize(opts *bind.TransactOpts, _pauserRegistry common.Address, initialOwner common.Address, _aggregator common.Address, _generator common.Address) (*types.Transaction, error) {
	return _ContractConnect3DATaskManager.contract.Transact(opts, "initialize", _pauserRegistry, initialOwner, _aggregator, _generator)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _pauserRegistry, address initialOwner, address _aggregator, address _generator) returns()
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerSession) Initialize(_pauserRegistry common.Address, initialOwner common.Address, _aggregator common.Address, _generator common.Address) (*types.Transaction, error) {
	return _ContractConnect3DATaskManager.Contract.Initialize(&_ContractConnect3DATaskManager.TransactOpts, _pauserRegistry, initialOwner, _aggregator, _generator)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _pauserRegistry, address initialOwner, address _aggregator, address _generator) returns()
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerTransactorSession) Initialize(_pauserRegistry common.Address, initialOwner common.Address, _aggregator common.Address, _generator common.Address) (*types.Transaction, error) {
	return _ContractConnect3DATaskManager.Contract.Initialize(&_ContractConnect3DATaskManager.TransactOpts, _pauserRegistry, initialOwner, _aggregator, _generator)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerTransactor) Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractConnect3DATaskManager.contract.Transact(opts, "pause", newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractConnect3DATaskManager.Contract.Pause(&_ContractConnect3DATaskManager.TransactOpts, newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerTransactorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractConnect3DATaskManager.Contract.Pause(&_ContractConnect3DATaskManager.TransactOpts, newPausedStatus)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractConnect3DATaskManager.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerSession) PauseAll() (*types.Transaction, error) {
	return _ContractConnect3DATaskManager.Contract.PauseAll(&_ContractConnect3DATaskManager.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerTransactorSession) PauseAll() (*types.Transaction, error) {
	return _ContractConnect3DATaskManager.Contract.PauseAll(&_ContractConnect3DATaskManager.TransactOpts)
}

// RaiseAndResolveChallenge is a paid mutator transaction binding the contract method 0x3878e94b.
//
// Solidity: function raiseAndResolveChallenge((address,uint32,bytes,uint32) task, (uint32,address,bytes32) taskResponse, (uint32,bytes32) taskResponseMetadata, (uint256,uint256)[] pubkeysOfNonSigningOperators) returns()
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerTransactor) RaiseAndResolveChallenge(opts *bind.TransactOpts, task IConnect3DATaskManagerTask, taskResponse IConnect3DATaskManagerTaskResponse, taskResponseMetadata IConnect3DATaskManagerTaskResponseMetadata, pubkeysOfNonSigningOperators []BN254G1Point) (*types.Transaction, error) {
	return _ContractConnect3DATaskManager.contract.Transact(opts, "raiseAndResolveChallenge", task, taskResponse, taskResponseMetadata, pubkeysOfNonSigningOperators)
}

// RaiseAndResolveChallenge is a paid mutator transaction binding the contract method 0x3878e94b.
//
// Solidity: function raiseAndResolveChallenge((address,uint32,bytes,uint32) task, (uint32,address,bytes32) taskResponse, (uint32,bytes32) taskResponseMetadata, (uint256,uint256)[] pubkeysOfNonSigningOperators) returns()
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerSession) RaiseAndResolveChallenge(task IConnect3DATaskManagerTask, taskResponse IConnect3DATaskManagerTaskResponse, taskResponseMetadata IConnect3DATaskManagerTaskResponseMetadata, pubkeysOfNonSigningOperators []BN254G1Point) (*types.Transaction, error) {
	return _ContractConnect3DATaskManager.Contract.RaiseAndResolveChallenge(&_ContractConnect3DATaskManager.TransactOpts, task, taskResponse, taskResponseMetadata, pubkeysOfNonSigningOperators)
}

// RaiseAndResolveChallenge is a paid mutator transaction binding the contract method 0x3878e94b.
//
// Solidity: function raiseAndResolveChallenge((address,uint32,bytes,uint32) task, (uint32,address,bytes32) taskResponse, (uint32,bytes32) taskResponseMetadata, (uint256,uint256)[] pubkeysOfNonSigningOperators) returns()
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerTransactorSession) RaiseAndResolveChallenge(task IConnect3DATaskManagerTask, taskResponse IConnect3DATaskManagerTaskResponse, taskResponseMetadata IConnect3DATaskManagerTaskResponseMetadata, pubkeysOfNonSigningOperators []BN254G1Point) (*types.Transaction, error) {
	return _ContractConnect3DATaskManager.Contract.RaiseAndResolveChallenge(&_ContractConnect3DATaskManager.TransactOpts, task, taskResponse, taskResponseMetadata, pubkeysOfNonSigningOperators)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractConnect3DATaskManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractConnect3DATaskManager.Contract.RenounceOwnership(&_ContractConnect3DATaskManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractConnect3DATaskManager.Contract.RenounceOwnership(&_ContractConnect3DATaskManager.TransactOpts)
}

// RespondToTask is a paid mutator transaction binding the contract method 0x14e5e9ba.
//
// Solidity: function respondToTask((address,uint32,bytes,uint32) task, (uint32,address,bytes32) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerTransactor) RespondToTask(opts *bind.TransactOpts, task IConnect3DATaskManagerTask, taskResponse IConnect3DATaskManagerTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractConnect3DATaskManager.contract.Transact(opts, "respondToTask", task, taskResponse, nonSignerStakesAndSignature)
}

// RespondToTask is a paid mutator transaction binding the contract method 0x14e5e9ba.
//
// Solidity: function respondToTask((address,uint32,bytes,uint32) task, (uint32,address,bytes32) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerSession) RespondToTask(task IConnect3DATaskManagerTask, taskResponse IConnect3DATaskManagerTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractConnect3DATaskManager.Contract.RespondToTask(&_ContractConnect3DATaskManager.TransactOpts, task, taskResponse, nonSignerStakesAndSignature)
}

// RespondToTask is a paid mutator transaction binding the contract method 0x14e5e9ba.
//
// Solidity: function respondToTask((address,uint32,bytes,uint32) task, (uint32,address,bytes32) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerTransactorSession) RespondToTask(task IConnect3DATaskManagerTask, taskResponse IConnect3DATaskManagerTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractConnect3DATaskManager.Contract.RespondToTask(&_ContractConnect3DATaskManager.TransactOpts, task, taskResponse, nonSignerStakesAndSignature)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerTransactor) SetPauserRegistry(opts *bind.TransactOpts, newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractConnect3DATaskManager.contract.Transact(opts, "setPauserRegistry", newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractConnect3DATaskManager.Contract.SetPauserRegistry(&_ContractConnect3DATaskManager.TransactOpts, newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerTransactorSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractConnect3DATaskManager.Contract.SetPauserRegistry(&_ContractConnect3DATaskManager.TransactOpts, newPauserRegistry)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerTransactor) SetStaleStakesForbidden(opts *bind.TransactOpts, value bool) (*types.Transaction, error) {
	return _ContractConnect3DATaskManager.contract.Transact(opts, "setStaleStakesForbidden", value)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerSession) SetStaleStakesForbidden(value bool) (*types.Transaction, error) {
	return _ContractConnect3DATaskManager.Contract.SetStaleStakesForbidden(&_ContractConnect3DATaskManager.TransactOpts, value)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerTransactorSession) SetStaleStakesForbidden(value bool) (*types.Transaction, error) {
	return _ContractConnect3DATaskManager.Contract.SetStaleStakesForbidden(&_ContractConnect3DATaskManager.TransactOpts, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ContractConnect3DATaskManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractConnect3DATaskManager.Contract.TransferOwnership(&_ContractConnect3DATaskManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractConnect3DATaskManager.Contract.TransferOwnership(&_ContractConnect3DATaskManager.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerTransactor) Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractConnect3DATaskManager.contract.Transact(opts, "unpause", newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractConnect3DATaskManager.Contract.Unpause(&_ContractConnect3DATaskManager.TransactOpts, newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerTransactorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractConnect3DATaskManager.Contract.Unpause(&_ContractConnect3DATaskManager.TransactOpts, newPausedStatus)
}

// ContractConnect3DATaskManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContractConnect3DATaskManager contract.
type ContractConnect3DATaskManagerInitializedIterator struct {
	Event *ContractConnect3DATaskManagerInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractConnect3DATaskManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractConnect3DATaskManagerInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractConnect3DATaskManagerInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractConnect3DATaskManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractConnect3DATaskManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractConnect3DATaskManagerInitialized represents a Initialized event raised by the ContractConnect3DATaskManager contract.
type ContractConnect3DATaskManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractConnect3DATaskManagerInitializedIterator, error) {

	logs, sub, err := _ContractConnect3DATaskManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractConnect3DATaskManagerInitializedIterator{contract: _ContractConnect3DATaskManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractConnect3DATaskManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _ContractConnect3DATaskManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractConnect3DATaskManagerInitialized)
				if err := _ContractConnect3DATaskManager.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerFilterer) ParseInitialized(log types.Log) (*ContractConnect3DATaskManagerInitialized, error) {
	event := new(ContractConnect3DATaskManagerInitialized)
	if err := _ContractConnect3DATaskManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractConnect3DATaskManagerNewTaskCreatedIterator is returned from FilterNewTaskCreated and is used to iterate over the raw logs and unpacked data for NewTaskCreated events raised by the ContractConnect3DATaskManager contract.
type ContractConnect3DATaskManagerNewTaskCreatedIterator struct {
	Event *ContractConnect3DATaskManagerNewTaskCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractConnect3DATaskManagerNewTaskCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractConnect3DATaskManagerNewTaskCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractConnect3DATaskManagerNewTaskCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractConnect3DATaskManagerNewTaskCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractConnect3DATaskManagerNewTaskCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractConnect3DATaskManagerNewTaskCreated represents a NewTaskCreated event raised by the ContractConnect3DATaskManager contract.
type ContractConnect3DATaskManagerNewTaskCreated struct {
	TaskIndex uint32
	Task      IConnect3DATaskManagerTask
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewTaskCreated is a free log retrieval operation binding the contract event 0x024dec6bd1d353599633784df48acfe0f307837dd3da13643a805fc55c6eb8f9.
//
// Solidity: event NewTaskCreated(uint32 indexed taskIndex, (address,uint32,bytes,uint32) task)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerFilterer) FilterNewTaskCreated(opts *bind.FilterOpts, taskIndex []uint32) (*ContractConnect3DATaskManagerNewTaskCreatedIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractConnect3DATaskManager.contract.FilterLogs(opts, "NewTaskCreated", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return &ContractConnect3DATaskManagerNewTaskCreatedIterator{contract: _ContractConnect3DATaskManager.contract, event: "NewTaskCreated", logs: logs, sub: sub}, nil
}

// WatchNewTaskCreated is a free log subscription operation binding the contract event 0x024dec6bd1d353599633784df48acfe0f307837dd3da13643a805fc55c6eb8f9.
//
// Solidity: event NewTaskCreated(uint32 indexed taskIndex, (address,uint32,bytes,uint32) task)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerFilterer) WatchNewTaskCreated(opts *bind.WatchOpts, sink chan<- *ContractConnect3DATaskManagerNewTaskCreated, taskIndex []uint32) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractConnect3DATaskManager.contract.WatchLogs(opts, "NewTaskCreated", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractConnect3DATaskManagerNewTaskCreated)
				if err := _ContractConnect3DATaskManager.contract.UnpackLog(event, "NewTaskCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewTaskCreated is a log parse operation binding the contract event 0x024dec6bd1d353599633784df48acfe0f307837dd3da13643a805fc55c6eb8f9.
//
// Solidity: event NewTaskCreated(uint32 indexed taskIndex, (address,uint32,bytes,uint32) task)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerFilterer) ParseNewTaskCreated(log types.Log) (*ContractConnect3DATaskManagerNewTaskCreated, error) {
	event := new(ContractConnect3DATaskManagerNewTaskCreated)
	if err := _ContractConnect3DATaskManager.contract.UnpackLog(event, "NewTaskCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractConnect3DATaskManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ContractConnect3DATaskManager contract.
type ContractConnect3DATaskManagerOwnershipTransferredIterator struct {
	Event *ContractConnect3DATaskManagerOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractConnect3DATaskManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractConnect3DATaskManagerOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractConnect3DATaskManagerOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractConnect3DATaskManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractConnect3DATaskManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractConnect3DATaskManagerOwnershipTransferred represents a OwnershipTransferred event raised by the ContractConnect3DATaskManager contract.
type ContractConnect3DATaskManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractConnect3DATaskManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractConnect3DATaskManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractConnect3DATaskManagerOwnershipTransferredIterator{contract: _ContractConnect3DATaskManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractConnect3DATaskManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractConnect3DATaskManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractConnect3DATaskManagerOwnershipTransferred)
				if err := _ContractConnect3DATaskManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerFilterer) ParseOwnershipTransferred(log types.Log) (*ContractConnect3DATaskManagerOwnershipTransferred, error) {
	event := new(ContractConnect3DATaskManagerOwnershipTransferred)
	if err := _ContractConnect3DATaskManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractConnect3DATaskManagerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ContractConnect3DATaskManager contract.
type ContractConnect3DATaskManagerPausedIterator struct {
	Event *ContractConnect3DATaskManagerPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractConnect3DATaskManagerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractConnect3DATaskManagerPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractConnect3DATaskManagerPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractConnect3DATaskManagerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractConnect3DATaskManagerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractConnect3DATaskManagerPaused represents a Paused event raised by the ContractConnect3DATaskManager contract.
type ContractConnect3DATaskManagerPaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*ContractConnect3DATaskManagerPausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractConnect3DATaskManager.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractConnect3DATaskManagerPausedIterator{contract: _ContractConnect3DATaskManager.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ContractConnect3DATaskManagerPaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractConnect3DATaskManager.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractConnect3DATaskManagerPaused)
				if err := _ContractConnect3DATaskManager.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerFilterer) ParsePaused(log types.Log) (*ContractConnect3DATaskManagerPaused, error) {
	event := new(ContractConnect3DATaskManagerPaused)
	if err := _ContractConnect3DATaskManager.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractConnect3DATaskManagerPauserRegistrySetIterator is returned from FilterPauserRegistrySet and is used to iterate over the raw logs and unpacked data for PauserRegistrySet events raised by the ContractConnect3DATaskManager contract.
type ContractConnect3DATaskManagerPauserRegistrySetIterator struct {
	Event *ContractConnect3DATaskManagerPauserRegistrySet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractConnect3DATaskManagerPauserRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractConnect3DATaskManagerPauserRegistrySet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractConnect3DATaskManagerPauserRegistrySet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractConnect3DATaskManagerPauserRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractConnect3DATaskManagerPauserRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractConnect3DATaskManagerPauserRegistrySet represents a PauserRegistrySet event raised by the ContractConnect3DATaskManager contract.
type ContractConnect3DATaskManagerPauserRegistrySet struct {
	PauserRegistry    common.Address
	NewPauserRegistry common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterPauserRegistrySet is a free log retrieval operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerFilterer) FilterPauserRegistrySet(opts *bind.FilterOpts) (*ContractConnect3DATaskManagerPauserRegistrySetIterator, error) {

	logs, sub, err := _ContractConnect3DATaskManager.contract.FilterLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return &ContractConnect3DATaskManagerPauserRegistrySetIterator{contract: _ContractConnect3DATaskManager.contract, event: "PauserRegistrySet", logs: logs, sub: sub}, nil
}

// WatchPauserRegistrySet is a free log subscription operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerFilterer) WatchPauserRegistrySet(opts *bind.WatchOpts, sink chan<- *ContractConnect3DATaskManagerPauserRegistrySet) (event.Subscription, error) {

	logs, sub, err := _ContractConnect3DATaskManager.contract.WatchLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractConnect3DATaskManagerPauserRegistrySet)
				if err := _ContractConnect3DATaskManager.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePauserRegistrySet is a log parse operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerFilterer) ParsePauserRegistrySet(log types.Log) (*ContractConnect3DATaskManagerPauserRegistrySet, error) {
	event := new(ContractConnect3DATaskManagerPauserRegistrySet)
	if err := _ContractConnect3DATaskManager.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractConnect3DATaskManagerStaleStakesForbiddenUpdateIterator is returned from FilterStaleStakesForbiddenUpdate and is used to iterate over the raw logs and unpacked data for StaleStakesForbiddenUpdate events raised by the ContractConnect3DATaskManager contract.
type ContractConnect3DATaskManagerStaleStakesForbiddenUpdateIterator struct {
	Event *ContractConnect3DATaskManagerStaleStakesForbiddenUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractConnect3DATaskManagerStaleStakesForbiddenUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractConnect3DATaskManagerStaleStakesForbiddenUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractConnect3DATaskManagerStaleStakesForbiddenUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractConnect3DATaskManagerStaleStakesForbiddenUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractConnect3DATaskManagerStaleStakesForbiddenUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractConnect3DATaskManagerStaleStakesForbiddenUpdate represents a StaleStakesForbiddenUpdate event raised by the ContractConnect3DATaskManager contract.
type ContractConnect3DATaskManagerStaleStakesForbiddenUpdate struct {
	Value bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterStaleStakesForbiddenUpdate is a free log retrieval operation binding the contract event 0x40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc.
//
// Solidity: event StaleStakesForbiddenUpdate(bool value)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerFilterer) FilterStaleStakesForbiddenUpdate(opts *bind.FilterOpts) (*ContractConnect3DATaskManagerStaleStakesForbiddenUpdateIterator, error) {

	logs, sub, err := _ContractConnect3DATaskManager.contract.FilterLogs(opts, "StaleStakesForbiddenUpdate")
	if err != nil {
		return nil, err
	}
	return &ContractConnect3DATaskManagerStaleStakesForbiddenUpdateIterator{contract: _ContractConnect3DATaskManager.contract, event: "StaleStakesForbiddenUpdate", logs: logs, sub: sub}, nil
}

// WatchStaleStakesForbiddenUpdate is a free log subscription operation binding the contract event 0x40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc.
//
// Solidity: event StaleStakesForbiddenUpdate(bool value)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerFilterer) WatchStaleStakesForbiddenUpdate(opts *bind.WatchOpts, sink chan<- *ContractConnect3DATaskManagerStaleStakesForbiddenUpdate) (event.Subscription, error) {

	logs, sub, err := _ContractConnect3DATaskManager.contract.WatchLogs(opts, "StaleStakesForbiddenUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractConnect3DATaskManagerStaleStakesForbiddenUpdate)
				if err := _ContractConnect3DATaskManager.contract.UnpackLog(event, "StaleStakesForbiddenUpdate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStaleStakesForbiddenUpdate is a log parse operation binding the contract event 0x40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc.
//
// Solidity: event StaleStakesForbiddenUpdate(bool value)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerFilterer) ParseStaleStakesForbiddenUpdate(log types.Log) (*ContractConnect3DATaskManagerStaleStakesForbiddenUpdate, error) {
	event := new(ContractConnect3DATaskManagerStaleStakesForbiddenUpdate)
	if err := _ContractConnect3DATaskManager.contract.UnpackLog(event, "StaleStakesForbiddenUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractConnect3DATaskManagerTaskChallengedSuccessfullyIterator is returned from FilterTaskChallengedSuccessfully and is used to iterate over the raw logs and unpacked data for TaskChallengedSuccessfully events raised by the ContractConnect3DATaskManager contract.
type ContractConnect3DATaskManagerTaskChallengedSuccessfullyIterator struct {
	Event *ContractConnect3DATaskManagerTaskChallengedSuccessfully // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractConnect3DATaskManagerTaskChallengedSuccessfullyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractConnect3DATaskManagerTaskChallengedSuccessfully)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractConnect3DATaskManagerTaskChallengedSuccessfully)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractConnect3DATaskManagerTaskChallengedSuccessfullyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractConnect3DATaskManagerTaskChallengedSuccessfullyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractConnect3DATaskManagerTaskChallengedSuccessfully represents a TaskChallengedSuccessfully event raised by the ContractConnect3DATaskManager contract.
type ContractConnect3DATaskManagerTaskChallengedSuccessfully struct {
	TaskIndex  uint32
	Challenger common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTaskChallengedSuccessfully is a free log retrieval operation binding the contract event 0xc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec.
//
// Solidity: event TaskChallengedSuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerFilterer) FilterTaskChallengedSuccessfully(opts *bind.FilterOpts, taskIndex []uint32, challenger []common.Address) (*ContractConnect3DATaskManagerTaskChallengedSuccessfullyIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractConnect3DATaskManager.contract.FilterLogs(opts, "TaskChallengedSuccessfully", taskIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return &ContractConnect3DATaskManagerTaskChallengedSuccessfullyIterator{contract: _ContractConnect3DATaskManager.contract, event: "TaskChallengedSuccessfully", logs: logs, sub: sub}, nil
}

// WatchTaskChallengedSuccessfully is a free log subscription operation binding the contract event 0xc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec.
//
// Solidity: event TaskChallengedSuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerFilterer) WatchTaskChallengedSuccessfully(opts *bind.WatchOpts, sink chan<- *ContractConnect3DATaskManagerTaskChallengedSuccessfully, taskIndex []uint32, challenger []common.Address) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractConnect3DATaskManager.contract.WatchLogs(opts, "TaskChallengedSuccessfully", taskIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractConnect3DATaskManagerTaskChallengedSuccessfully)
				if err := _ContractConnect3DATaskManager.contract.UnpackLog(event, "TaskChallengedSuccessfully", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTaskChallengedSuccessfully is a log parse operation binding the contract event 0xc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec.
//
// Solidity: event TaskChallengedSuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerFilterer) ParseTaskChallengedSuccessfully(log types.Log) (*ContractConnect3DATaskManagerTaskChallengedSuccessfully, error) {
	event := new(ContractConnect3DATaskManagerTaskChallengedSuccessfully)
	if err := _ContractConnect3DATaskManager.contract.UnpackLog(event, "TaskChallengedSuccessfully", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractConnect3DATaskManagerTaskChallengedUnsuccessfullyIterator is returned from FilterTaskChallengedUnsuccessfully and is used to iterate over the raw logs and unpacked data for TaskChallengedUnsuccessfully events raised by the ContractConnect3DATaskManager contract.
type ContractConnect3DATaskManagerTaskChallengedUnsuccessfullyIterator struct {
	Event *ContractConnect3DATaskManagerTaskChallengedUnsuccessfully // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractConnect3DATaskManagerTaskChallengedUnsuccessfullyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractConnect3DATaskManagerTaskChallengedUnsuccessfully)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractConnect3DATaskManagerTaskChallengedUnsuccessfully)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractConnect3DATaskManagerTaskChallengedUnsuccessfullyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractConnect3DATaskManagerTaskChallengedUnsuccessfullyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractConnect3DATaskManagerTaskChallengedUnsuccessfully represents a TaskChallengedUnsuccessfully event raised by the ContractConnect3DATaskManager contract.
type ContractConnect3DATaskManagerTaskChallengedUnsuccessfully struct {
	TaskIndex  uint32
	Challenger common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTaskChallengedUnsuccessfully is a free log retrieval operation binding the contract event 0xfd3e26beeb5967fc5a57a0446914eabc45b4aa474c67a51b4b5160cac60ddb05.
//
// Solidity: event TaskChallengedUnsuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerFilterer) FilterTaskChallengedUnsuccessfully(opts *bind.FilterOpts, taskIndex []uint32, challenger []common.Address) (*ContractConnect3DATaskManagerTaskChallengedUnsuccessfullyIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractConnect3DATaskManager.contract.FilterLogs(opts, "TaskChallengedUnsuccessfully", taskIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return &ContractConnect3DATaskManagerTaskChallengedUnsuccessfullyIterator{contract: _ContractConnect3DATaskManager.contract, event: "TaskChallengedUnsuccessfully", logs: logs, sub: sub}, nil
}

// WatchTaskChallengedUnsuccessfully is a free log subscription operation binding the contract event 0xfd3e26beeb5967fc5a57a0446914eabc45b4aa474c67a51b4b5160cac60ddb05.
//
// Solidity: event TaskChallengedUnsuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerFilterer) WatchTaskChallengedUnsuccessfully(opts *bind.WatchOpts, sink chan<- *ContractConnect3DATaskManagerTaskChallengedUnsuccessfully, taskIndex []uint32, challenger []common.Address) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractConnect3DATaskManager.contract.WatchLogs(opts, "TaskChallengedUnsuccessfully", taskIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractConnect3DATaskManagerTaskChallengedUnsuccessfully)
				if err := _ContractConnect3DATaskManager.contract.UnpackLog(event, "TaskChallengedUnsuccessfully", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTaskChallengedUnsuccessfully is a log parse operation binding the contract event 0xfd3e26beeb5967fc5a57a0446914eabc45b4aa474c67a51b4b5160cac60ddb05.
//
// Solidity: event TaskChallengedUnsuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerFilterer) ParseTaskChallengedUnsuccessfully(log types.Log) (*ContractConnect3DATaskManagerTaskChallengedUnsuccessfully, error) {
	event := new(ContractConnect3DATaskManagerTaskChallengedUnsuccessfully)
	if err := _ContractConnect3DATaskManager.contract.UnpackLog(event, "TaskChallengedUnsuccessfully", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractConnect3DATaskManagerTaskCompletedIterator is returned from FilterTaskCompleted and is used to iterate over the raw logs and unpacked data for TaskCompleted events raised by the ContractConnect3DATaskManager contract.
type ContractConnect3DATaskManagerTaskCompletedIterator struct {
	Event *ContractConnect3DATaskManagerTaskCompleted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractConnect3DATaskManagerTaskCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractConnect3DATaskManagerTaskCompleted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractConnect3DATaskManagerTaskCompleted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractConnect3DATaskManagerTaskCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractConnect3DATaskManagerTaskCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractConnect3DATaskManagerTaskCompleted represents a TaskCompleted event raised by the ContractConnect3DATaskManager contract.
type ContractConnect3DATaskManagerTaskCompleted struct {
	TaskIndex uint32
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTaskCompleted is a free log retrieval operation binding the contract event 0x9a144f228a931b9d0d1696fbcdaf310b24b5d2d21e799db623fc986a0f547430.
//
// Solidity: event TaskCompleted(uint32 indexed taskIndex)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerFilterer) FilterTaskCompleted(opts *bind.FilterOpts, taskIndex []uint32) (*ContractConnect3DATaskManagerTaskCompletedIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractConnect3DATaskManager.contract.FilterLogs(opts, "TaskCompleted", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return &ContractConnect3DATaskManagerTaskCompletedIterator{contract: _ContractConnect3DATaskManager.contract, event: "TaskCompleted", logs: logs, sub: sub}, nil
}

// WatchTaskCompleted is a free log subscription operation binding the contract event 0x9a144f228a931b9d0d1696fbcdaf310b24b5d2d21e799db623fc986a0f547430.
//
// Solidity: event TaskCompleted(uint32 indexed taskIndex)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerFilterer) WatchTaskCompleted(opts *bind.WatchOpts, sink chan<- *ContractConnect3DATaskManagerTaskCompleted, taskIndex []uint32) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractConnect3DATaskManager.contract.WatchLogs(opts, "TaskCompleted", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractConnect3DATaskManagerTaskCompleted)
				if err := _ContractConnect3DATaskManager.contract.UnpackLog(event, "TaskCompleted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTaskCompleted is a log parse operation binding the contract event 0x9a144f228a931b9d0d1696fbcdaf310b24b5d2d21e799db623fc986a0f547430.
//
// Solidity: event TaskCompleted(uint32 indexed taskIndex)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerFilterer) ParseTaskCompleted(log types.Log) (*ContractConnect3DATaskManagerTaskCompleted, error) {
	event := new(ContractConnect3DATaskManagerTaskCompleted)
	if err := _ContractConnect3DATaskManager.contract.UnpackLog(event, "TaskCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractConnect3DATaskManagerTaskRespondedIterator is returned from FilterTaskResponded and is used to iterate over the raw logs and unpacked data for TaskResponded events raised by the ContractConnect3DATaskManager contract.
type ContractConnect3DATaskManagerTaskRespondedIterator struct {
	Event *ContractConnect3DATaskManagerTaskResponded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractConnect3DATaskManagerTaskRespondedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractConnect3DATaskManagerTaskResponded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractConnect3DATaskManagerTaskResponded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractConnect3DATaskManagerTaskRespondedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractConnect3DATaskManagerTaskRespondedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractConnect3DATaskManagerTaskResponded represents a TaskResponded event raised by the ContractConnect3DATaskManager contract.
type ContractConnect3DATaskManagerTaskResponded struct {
	TaskResponse         IConnect3DATaskManagerTaskResponse
	TaskResponseMetadata IConnect3DATaskManagerTaskResponseMetadata
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterTaskResponded is a free log retrieval operation binding the contract event 0xb6a7079e92345de4742a12e32ffe70803bdaaac6f5417156288caefd356e4b02.
//
// Solidity: event TaskResponded((uint32,address,bytes32) taskResponse, (uint32,bytes32) taskResponseMetadata)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerFilterer) FilterTaskResponded(opts *bind.FilterOpts) (*ContractConnect3DATaskManagerTaskRespondedIterator, error) {

	logs, sub, err := _ContractConnect3DATaskManager.contract.FilterLogs(opts, "TaskResponded")
	if err != nil {
		return nil, err
	}
	return &ContractConnect3DATaskManagerTaskRespondedIterator{contract: _ContractConnect3DATaskManager.contract, event: "TaskResponded", logs: logs, sub: sub}, nil
}

// WatchTaskResponded is a free log subscription operation binding the contract event 0xb6a7079e92345de4742a12e32ffe70803bdaaac6f5417156288caefd356e4b02.
//
// Solidity: event TaskResponded((uint32,address,bytes32) taskResponse, (uint32,bytes32) taskResponseMetadata)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerFilterer) WatchTaskResponded(opts *bind.WatchOpts, sink chan<- *ContractConnect3DATaskManagerTaskResponded) (event.Subscription, error) {

	logs, sub, err := _ContractConnect3DATaskManager.contract.WatchLogs(opts, "TaskResponded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractConnect3DATaskManagerTaskResponded)
				if err := _ContractConnect3DATaskManager.contract.UnpackLog(event, "TaskResponded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTaskResponded is a log parse operation binding the contract event 0xb6a7079e92345de4742a12e32ffe70803bdaaac6f5417156288caefd356e4b02.
//
// Solidity: event TaskResponded((uint32,address,bytes32) taskResponse, (uint32,bytes32) taskResponseMetadata)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerFilterer) ParseTaskResponded(log types.Log) (*ContractConnect3DATaskManagerTaskResponded, error) {
	event := new(ContractConnect3DATaskManagerTaskResponded)
	if err := _ContractConnect3DATaskManager.contract.UnpackLog(event, "TaskResponded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractConnect3DATaskManagerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ContractConnect3DATaskManager contract.
type ContractConnect3DATaskManagerUnpausedIterator struct {
	Event *ContractConnect3DATaskManagerUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractConnect3DATaskManagerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractConnect3DATaskManagerUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractConnect3DATaskManagerUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractConnect3DATaskManagerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractConnect3DATaskManagerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractConnect3DATaskManagerUnpaused represents a Unpaused event raised by the ContractConnect3DATaskManager contract.
type ContractConnect3DATaskManagerUnpaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*ContractConnect3DATaskManagerUnpausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractConnect3DATaskManager.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractConnect3DATaskManagerUnpausedIterator{contract: _ContractConnect3DATaskManager.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ContractConnect3DATaskManagerUnpaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractConnect3DATaskManager.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractConnect3DATaskManagerUnpaused)
				if err := _ContractConnect3DATaskManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractConnect3DATaskManager *ContractConnect3DATaskManagerFilterer) ParseUnpaused(log types.Log) (*ContractConnect3DATaskManagerUnpaused, error) {
	event := new(ContractConnect3DATaskManagerUnpaused)
	if err := _ContractConnect3DATaskManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
