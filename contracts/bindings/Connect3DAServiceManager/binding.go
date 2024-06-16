// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractConnect3DAServiceManager

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

// IPaymentCoordinatorRangePayment is an auto generated low-level Go binding around an user-defined struct.
type IPaymentCoordinatorRangePayment struct {
	StrategiesAndMultipliers []IPaymentCoordinatorStrategyAndMultiplier
	Token                    common.Address
	Amount                   *big.Int
	StartTimestamp           uint32
	Duration                 uint32
}

// IPaymentCoordinatorStrategyAndMultiplier is an auto generated low-level Go binding around an user-defined struct.
type IPaymentCoordinatorStrategyAndMultiplier struct {
	Strategy   common.Address
	Multiplier *big.Int
}

// ISignatureUtilsSignatureWithSaltAndExpiry is an auto generated low-level Go binding around an user-defined struct.
type ISignatureUtilsSignatureWithSaltAndExpiry struct {
	Signature []byte
	Salt      [32]byte
	Expiry    *big.Int
}

// ContractConnect3DAServiceManagerMetaData contains all meta data concerning the ContractConnect3DAServiceManager contract.
var ContractConnect3DAServiceManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_avsDirectory\",\"type\":\"address\",\"internalType\":\"contractIAVSDirectory\"},{\"name\":\"_registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"_stakeRegistry\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"},{\"name\":\"_connect3DATaskManager\",\"type\":\"address\",\"internalType\":\"contractIConnect3DATaskManager\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"avsDirectory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"connect3DATaskManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIConnect3DATaskManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deregisterOperatorFromAVS\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"freezeOperator\",\"inputs\":[{\"name\":\"operatorAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getOperatorRestakedStrategies\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRestakeableStrategies\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"payForRange\",\"inputs\":[{\"name\":\"rangePayments\",\"type\":\"tuple[]\",\"internalType\":\"structIPaymentCoordinator.RangePayment[]\",\"components\":[{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[]\",\"internalType\":\"structIPaymentCoordinator.StrategyAndMultiplier[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"duration\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerOperatorToAVS\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateAVSMetadataURI\",\"inputs\":[{\"name\":\"_metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false}]",
	Bin: "0x6101206040523480156200001257600080fd5b5060405162001b5738038062001b5783398101604081905262000035916200015a565b6001600160a01b03808516608052600060a081905281851660c05290831660e05284908484620000646200007f565b505050506001600160a01b03166101005250620001c2915050565b600054610100900460ff1615620000ec5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff90811610156200013f576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b03811681146200015757600080fd5b50565b600080600080608085870312156200017157600080fd5b84516200017e8162000141565b6020860151909450620001918162000141565b6040860151909350620001a48162000141565b6060860151909250620001b78162000141565b939692955090935050565b60805160a05160c05160e051610100516118e0620002776000396000818161019b015261093b01526000818161065a015281816107b60152818161084d01528181610c6e01528181610df20152610e91015260008181610485015281816105140152818161059401528181610a0001528181610a9601528181610bac0152610d4d01526000818161031501526103f301526000818161010c01528181610a5401528181610af20152610b7101526118e06000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c80639926ee7d116100715780639926ee7d1461015d578063a364f4da14610170578063a98fb35514610183578063b548d68214610196578063e481af9d146101bd578063f2fde38b146101c557600080fd5b80631b445516146100b957806333cfb7b7146100ce57806338c8ee64146100f75780636b3aa72e1461010a578063715018a6146101445780638da5cb5b1461014c575b600080fd5b6100cc6100c7366004611186565b6101d8565b005b6100e16100dc366004611210565b610460565b6040516100ee9190611234565b60405180910390f35b6100cc610105366004611210565b610930565b7f00000000000000000000000000000000000000000000000000000000000000005b6040516001600160a01b0390911681526020016100ee565b6100cc6109e1565b6033546001600160a01b031661012c565b6100cc61016b366004611336565b6109f5565b6100cc61017e366004611210565b610a8b565b6100cc6101913660046113e1565b610b52565b61012c7f000000000000000000000000000000000000000000000000000000000000000081565b6100e1610ba6565b6100cc6101d3366004611210565b610f70565b6101e0610fe6565b60005b818110156103db578282828181106101fd576101fd611432565b905060200281019061020f9190611448565b610220906040810190602001611210565b6001600160a01b03166323b872dd333086868681811061024257610242611432565b90506020028101906102549190611448565b604080516001600160e01b031960e087901b1681526001600160a01b039485166004820152939092166024840152013560448201526064016020604051808303816000875af11580156102ab573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102cf9190611478565b508282828181106102e2576102e2611432565b90506020028101906102f49190611448565b610305906040810190602001611210565b6001600160a01b031663095ea7b37f000000000000000000000000000000000000000000000000000000000000000085858581811061034657610346611432565b90506020028101906103589190611448565b604080516001600160e01b031960e086901b1681526001600160a01b039093166004840152013560248201526044016020604051808303816000875af11580156103a6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103ca9190611478565b506103d4816114b0565b90506101e3565b50604051630da22a8b60e11b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690631b4455169061042a9085908590600401611564565b600060405180830381600087803b15801561044457600080fd5b505af1158015610458573d6000803e3d6000fd5b505050505050565b6040516309aa152760e11b81526001600160a01b0382811660048301526060916000917f000000000000000000000000000000000000000000000000000000000000000016906313542a4e90602401602060405180830381865afa1580156104cc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104f09190611672565b60405163871ef04960e01b8152600481018290529091506000906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063871ef04990602401602060405180830381865afa15801561055b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061057f919061168b565b90506001600160c01b038116158061061957507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156105f0573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061061491906116b4565b60ff16155b1561063557505060408051600081526020810190915292915050565b6000610649826001600160c01b0316611040565b90506000805b825181101561071f577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316633ca5a5f584838151811061069957610699611432565b01602001516040516001600160e01b031960e084901b16815260f89190911c6004820152602401602060405180830381865afa1580156106dd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107019190611672565b61070b90836116d7565b915080610717816114b0565b91505061064f565b5060008167ffffffffffffffff81111561073b5761073b611281565b604051908082528060200260200182016040528015610764578160200160208202803683370190505b5090506000805b845181101561092357600085828151811061078857610788611432565b0160200151604051633ca5a5f560e01b815260f89190911c6004820181905291506000906001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690633ca5a5f590602401602060405180830381865afa1580156107fd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108219190611672565b905060005b8181101561090d576040516356e4026d60e11b815260ff84166004820152602481018290527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063adc804da906044016040805180830381865afa15801561089b573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108bf91906116ef565b600001518686815181106108d5576108d5611432565b6001600160a01b0390921660209283029190910190910152846108f7816114b0565b9550508080610905906114b0565b915050610826565b505050808061091b906114b0565b91505061076b565b5090979650505050505050565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146109de5760405162461bcd60e51b815260206004820152604260248201527f6f6e6c79436f6e6e6563743344415461736b4d616e616765723a206e6f74206660448201527f726f6d206372656469626c65207371756172696e67207461736b206d616e616760648201526132b960f11b608482015260a4015b60405180910390fd5b50565b6109e9610fe6565b6109f36000611103565b565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610a3d5760405162461bcd60e51b81526004016109d59061174e565b604051639926ee7d60e01b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690639926ee7d9061042a9085908590600401611813565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610ad35760405162461bcd60e51b81526004016109d59061174e565b6040516351b27a6d60e11b81526001600160a01b0382811660048301527f0000000000000000000000000000000000000000000000000000000000000000169063a364f4da906024015b600060405180830381600087803b158015610b3757600080fd5b505af1158015610b4b573d6000803e3d6000fd5b5050505050565b610b5a610fe6565b60405163a98fb35560e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063a98fb35590610b1d90849060040161185e565b606060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610c08573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c2c91906116b4565b60ff16905080610c4a57505060408051600081526020810190915290565b6000805b82811015610cff57604051633ca5a5f560e01b815260ff821660048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690633ca5a5f590602401602060405180830381865afa158015610cbd573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ce19190611672565b610ceb90836116d7565b915080610cf7816114b0565b915050610c4e565b5060008167ffffffffffffffff811115610d1b57610d1b611281565b604051908082528060200260200182016040528015610d44578160200160208202803683370190505b5090506000805b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610da9573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610dcd91906116b4565b60ff16811015610f6657604051633ca5a5f560e01b815260ff821660048201526000907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690633ca5a5f590602401602060405180830381865afa158015610e41573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e659190611672565b905060005b81811015610f51576040516356e4026d60e11b815260ff84166004820152602481018290527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063adc804da906044016040805180830381865afa158015610edf573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f0391906116ef565b60000151858581518110610f1957610f19611432565b6001600160a01b039092166020928302919091019091015283610f3b816114b0565b9450508080610f49906114b0565b915050610e6a565b50508080610f5e906114b0565b915050610d4b565b5090949350505050565b610f78610fe6565b6001600160a01b038116610fdd5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016109d5565b6109de81611103565b6033546001600160a01b031633146109f35760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016109d5565b606060008061104e84611155565b61ffff1667ffffffffffffffff81111561106a5761106a611281565b6040519080825280601f01601f191660200182016040528015611094576020820181803683370190505b5090506000805b8251821080156110ac575061010081105b15610f66576001811b9350858416156110f3578060f81b8383815181106110d5576110d5611432565b60200101906001600160f81b031916908160001a9053508160010191505b6110fc816114b0565b905061109b565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6000805b82156111805761116a600184611871565b909216918061117881611888565b915050611159565b92915050565b6000806020838503121561119957600080fd5b823567ffffffffffffffff808211156111b157600080fd5b818501915085601f8301126111c557600080fd5b8135818111156111d457600080fd5b8660208260051b85010111156111e957600080fd5b60209290920196919550909350505050565b6001600160a01b03811681146109de57600080fd5b60006020828403121561122257600080fd5b813561122d816111fb565b9392505050565b6020808252825182820181905260009190848201906040850190845b818110156112755783516001600160a01b031683529284019291840191600101611250565b50909695505050505050565b634e487b7160e01b600052604160045260246000fd5b6040516060810167ffffffffffffffff811182821017156112ba576112ba611281565b60405290565b600067ffffffffffffffff808411156112db576112db611281565b604051601f8501601f19908116603f0116810190828211818310171561130357611303611281565b8160405280935085815286868601111561131c57600080fd5b858560208301376000602087830101525050509392505050565b6000806040838503121561134957600080fd5b8235611354816111fb565b9150602083013567ffffffffffffffff8082111561137157600080fd5b908401906060828703121561138557600080fd5b61138d611297565b82358281111561139c57600080fd5b83019150601f820187136113af57600080fd5b6113be878335602085016112c0565b815260208301356020820152604083013560408201528093505050509250929050565b6000602082840312156113f357600080fd5b813567ffffffffffffffff81111561140a57600080fd5b8201601f8101841361141b57600080fd5b61142a848235602084016112c0565b949350505050565b634e487b7160e01b600052603260045260246000fd5b60008235609e1983360301811261145e57600080fd5b9190910192915050565b8035611473816111fb565b919050565b60006020828403121561148a57600080fd5b8151801515811461122d57600080fd5b634e487b7160e01b600052601160045260246000fd5b60006000198214156114c4576114c461149a565b5060010190565b6bffffffffffffffffffffffff811681146109de57600080fd5b8183526000602080850194508260005b85811015611545578135611508816111fb565b6001600160a01b0316875281830135611520816114cb565b6bffffffffffffffffffffffff168784015260409687019691909101906001016114f5565b509495945050505050565b803563ffffffff8116811461147357600080fd5b60208082528181018390526000906040808401600586901b8501820187855b8881101561166457878303603f190184528135368b9003609e190181126115a957600080fd5b8a0160a0813536839003601e190181126115c257600080fd5b8201803567ffffffffffffffff8111156115db57600080fd5b8060061b36038413156115ed57600080fd5b8287526115ff838801828c85016114e5565b9250505061160e888301611468565b6001600160a01b03168886015281870135878601526060611630818401611550565b63ffffffff16908601526080611647838201611550565b63ffffffff16950194909452509285019290850190600101611583565b509098975050505050505050565b60006020828403121561168457600080fd5b5051919050565b60006020828403121561169d57600080fd5b81516001600160c01b038116811461122d57600080fd5b6000602082840312156116c657600080fd5b815160ff8116811461122d57600080fd5b600082198211156116ea576116ea61149a565b500190565b60006040828403121561170157600080fd5b6040516040810181811067ffffffffffffffff8211171561172457611724611281565b6040528251611732816111fb565b81526020830151611742816114cb565b60208201529392505050565b60208082526052908201527f536572766963654d616e61676572426173652e6f6e6c7952656769737472794360408201527f6f6f7264696e61746f723a2063616c6c6572206973206e6f742074686520726560608201527133b4b9ba393c9031b7b7b93234b730ba37b960711b608082015260a00190565b6000815180845260005b818110156117ec576020818501810151868301820152016117d0565b818111156117fe576000602083870101525b50601f01601f19169290920160200192915050565b60018060a01b038316815260406020820152600082516060604084015261183d60a08401826117c6565b90506020840151606084015260408401516080840152809150509392505050565b60208152600061122d60208301846117c6565b6000828210156118835761188361149a565b500390565b600061ffff808316818114156118a0576118a061149a565b600101939250505056fea2646970667358221220a1ff9596944c65f2b9a1da25cd25a4e0cdecf084e562caf98201c9e5f726e5f464736f6c634300080c0033",
}

// ContractConnect3DAServiceManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractConnect3DAServiceManagerMetaData.ABI instead.
var ContractConnect3DAServiceManagerABI = ContractConnect3DAServiceManagerMetaData.ABI

// ContractConnect3DAServiceManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractConnect3DAServiceManagerMetaData.Bin instead.
var ContractConnect3DAServiceManagerBin = ContractConnect3DAServiceManagerMetaData.Bin

// DeployContractConnect3DAServiceManager deploys a new Ethereum contract, binding an instance of ContractConnect3DAServiceManager to it.
func DeployContractConnect3DAServiceManager(auth *bind.TransactOpts, backend bind.ContractBackend, _avsDirectory common.Address, _registryCoordinator common.Address, _stakeRegistry common.Address, _connect3DATaskManager common.Address) (common.Address, *types.Transaction, *ContractConnect3DAServiceManager, error) {
	parsed, err := ContractConnect3DAServiceManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractConnect3DAServiceManagerBin), backend, _avsDirectory, _registryCoordinator, _stakeRegistry, _connect3DATaskManager)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractConnect3DAServiceManager{ContractConnect3DAServiceManagerCaller: ContractConnect3DAServiceManagerCaller{contract: contract}, ContractConnect3DAServiceManagerTransactor: ContractConnect3DAServiceManagerTransactor{contract: contract}, ContractConnect3DAServiceManagerFilterer: ContractConnect3DAServiceManagerFilterer{contract: contract}}, nil
}

// ContractConnect3DAServiceManager is an auto generated Go binding around an Ethereum contract.
type ContractConnect3DAServiceManager struct {
	ContractConnect3DAServiceManagerCaller     // Read-only binding to the contract
	ContractConnect3DAServiceManagerTransactor // Write-only binding to the contract
	ContractConnect3DAServiceManagerFilterer   // Log filterer for contract events
}

// ContractConnect3DAServiceManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractConnect3DAServiceManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractConnect3DAServiceManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractConnect3DAServiceManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractConnect3DAServiceManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractConnect3DAServiceManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractConnect3DAServiceManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractConnect3DAServiceManagerSession struct {
	Contract     *ContractConnect3DAServiceManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                     // Call options to use throughout this session
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// ContractConnect3DAServiceManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractConnect3DAServiceManagerCallerSession struct {
	Contract *ContractConnect3DAServiceManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                           // Call options to use throughout this session
}

// ContractConnect3DAServiceManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractConnect3DAServiceManagerTransactorSession struct {
	Contract     *ContractConnect3DAServiceManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                           // Transaction auth options to use throughout this session
}

// ContractConnect3DAServiceManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractConnect3DAServiceManagerRaw struct {
	Contract *ContractConnect3DAServiceManager // Generic contract binding to access the raw methods on
}

// ContractConnect3DAServiceManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractConnect3DAServiceManagerCallerRaw struct {
	Contract *ContractConnect3DAServiceManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ContractConnect3DAServiceManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractConnect3DAServiceManagerTransactorRaw struct {
	Contract *ContractConnect3DAServiceManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractConnect3DAServiceManager creates a new instance of ContractConnect3DAServiceManager, bound to a specific deployed contract.
func NewContractConnect3DAServiceManager(address common.Address, backend bind.ContractBackend) (*ContractConnect3DAServiceManager, error) {
	contract, err := bindContractConnect3DAServiceManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractConnect3DAServiceManager{ContractConnect3DAServiceManagerCaller: ContractConnect3DAServiceManagerCaller{contract: contract}, ContractConnect3DAServiceManagerTransactor: ContractConnect3DAServiceManagerTransactor{contract: contract}, ContractConnect3DAServiceManagerFilterer: ContractConnect3DAServiceManagerFilterer{contract: contract}}, nil
}

// NewContractConnect3DAServiceManagerCaller creates a new read-only instance of ContractConnect3DAServiceManager, bound to a specific deployed contract.
func NewContractConnect3DAServiceManagerCaller(address common.Address, caller bind.ContractCaller) (*ContractConnect3DAServiceManagerCaller, error) {
	contract, err := bindContractConnect3DAServiceManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractConnect3DAServiceManagerCaller{contract: contract}, nil
}

// NewContractConnect3DAServiceManagerTransactor creates a new write-only instance of ContractConnect3DAServiceManager, bound to a specific deployed contract.
func NewContractConnect3DAServiceManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractConnect3DAServiceManagerTransactor, error) {
	contract, err := bindContractConnect3DAServiceManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractConnect3DAServiceManagerTransactor{contract: contract}, nil
}

// NewContractConnect3DAServiceManagerFilterer creates a new log filterer instance of ContractConnect3DAServiceManager, bound to a specific deployed contract.
func NewContractConnect3DAServiceManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractConnect3DAServiceManagerFilterer, error) {
	contract, err := bindContractConnect3DAServiceManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractConnect3DAServiceManagerFilterer{contract: contract}, nil
}

// bindContractConnect3DAServiceManager binds a generic wrapper to an already deployed contract.
func bindContractConnect3DAServiceManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractConnect3DAServiceManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractConnect3DAServiceManager *ContractConnect3DAServiceManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractConnect3DAServiceManager.Contract.ContractConnect3DAServiceManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractConnect3DAServiceManager *ContractConnect3DAServiceManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractConnect3DAServiceManager.Contract.ContractConnect3DAServiceManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractConnect3DAServiceManager *ContractConnect3DAServiceManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractConnect3DAServiceManager.Contract.ContractConnect3DAServiceManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractConnect3DAServiceManager *ContractConnect3DAServiceManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractConnect3DAServiceManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractConnect3DAServiceManager *ContractConnect3DAServiceManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractConnect3DAServiceManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractConnect3DAServiceManager *ContractConnect3DAServiceManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractConnect3DAServiceManager.Contract.contract.Transact(opts, method, params...)
}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_ContractConnect3DAServiceManager *ContractConnect3DAServiceManagerCaller) AvsDirectory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractConnect3DAServiceManager.contract.Call(opts, &out, "avsDirectory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_ContractConnect3DAServiceManager *ContractConnect3DAServiceManagerSession) AvsDirectory() (common.Address, error) {
	return _ContractConnect3DAServiceManager.Contract.AvsDirectory(&_ContractConnect3DAServiceManager.CallOpts)
}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_ContractConnect3DAServiceManager *ContractConnect3DAServiceManagerCallerSession) AvsDirectory() (common.Address, error) {
	return _ContractConnect3DAServiceManager.Contract.AvsDirectory(&_ContractConnect3DAServiceManager.CallOpts)
}

// Connect3DATaskManager is a free data retrieval call binding the contract method 0xb548d682.
//
// Solidity: function connect3DATaskManager() view returns(address)
func (_ContractConnect3DAServiceManager *ContractConnect3DAServiceManagerCaller) Connect3DATaskManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractConnect3DAServiceManager.contract.Call(opts, &out, "connect3DATaskManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Connect3DATaskManager is a free data retrieval call binding the contract method 0xb548d682.
//
// Solidity: function connect3DATaskManager() view returns(address)
func (_ContractConnect3DAServiceManager *ContractConnect3DAServiceManagerSession) Connect3DATaskManager() (common.Address, error) {
	return _ContractConnect3DAServiceManager.Contract.Connect3DATaskManager(&_ContractConnect3DAServiceManager.CallOpts)
}

// Connect3DATaskManager is a free data retrieval call binding the contract method 0xb548d682.
//
// Solidity: function connect3DATaskManager() view returns(address)
func (_ContractConnect3DAServiceManager *ContractConnect3DAServiceManagerCallerSession) Connect3DATaskManager() (common.Address, error) {
	return _ContractConnect3DAServiceManager.Contract.Connect3DATaskManager(&_ContractConnect3DAServiceManager.CallOpts)
}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_ContractConnect3DAServiceManager *ContractConnect3DAServiceManagerCaller) GetOperatorRestakedStrategies(opts *bind.CallOpts, operator common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _ContractConnect3DAServiceManager.contract.Call(opts, &out, "getOperatorRestakedStrategies", operator)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_ContractConnect3DAServiceManager *ContractConnect3DAServiceManagerSession) GetOperatorRestakedStrategies(operator common.Address) ([]common.Address, error) {
	return _ContractConnect3DAServiceManager.Contract.GetOperatorRestakedStrategies(&_ContractConnect3DAServiceManager.CallOpts, operator)
}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_ContractConnect3DAServiceManager *ContractConnect3DAServiceManagerCallerSession) GetOperatorRestakedStrategies(operator common.Address) ([]common.Address, error) {
	return _ContractConnect3DAServiceManager.Contract.GetOperatorRestakedStrategies(&_ContractConnect3DAServiceManager.CallOpts, operator)
}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_ContractConnect3DAServiceManager *ContractConnect3DAServiceManagerCaller) GetRestakeableStrategies(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ContractConnect3DAServiceManager.contract.Call(opts, &out, "getRestakeableStrategies")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_ContractConnect3DAServiceManager *ContractConnect3DAServiceManagerSession) GetRestakeableStrategies() ([]common.Address, error) {
	return _ContractConnect3DAServiceManager.Contract.GetRestakeableStrategies(&_ContractConnect3DAServiceManager.CallOpts)
}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_ContractConnect3DAServiceManager *ContractConnect3DAServiceManagerCallerSession) GetRestakeableStrategies() ([]common.Address, error) {
	return _ContractConnect3DAServiceManager.Contract.GetRestakeableStrategies(&_ContractConnect3DAServiceManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractConnect3DAServiceManager *ContractConnect3DAServiceManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractConnect3DAServiceManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractConnect3DAServiceManager *ContractConnect3DAServiceManagerSession) Owner() (common.Address, error) {
	return _ContractConnect3DAServiceManager.Contract.Owner(&_ContractConnect3DAServiceManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractConnect3DAServiceManager *ContractConnect3DAServiceManagerCallerSession) Owner() (common.Address, error) {
	return _ContractConnect3DAServiceManager.Contract.Owner(&_ContractConnect3DAServiceManager.CallOpts)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_ContractConnect3DAServiceManager *ContractConnect3DAServiceManagerTransactor) DeregisterOperatorFromAVS(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _ContractConnect3DAServiceManager.contract.Transact(opts, "deregisterOperatorFromAVS", operator)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_ContractConnect3DAServiceManager *ContractConnect3DAServiceManagerSession) DeregisterOperatorFromAVS(operator common.Address) (*types.Transaction, error) {
	return _ContractConnect3DAServiceManager.Contract.DeregisterOperatorFromAVS(&_ContractConnect3DAServiceManager.TransactOpts, operator)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_ContractConnect3DAServiceManager *ContractConnect3DAServiceManagerTransactorSession) DeregisterOperatorFromAVS(operator common.Address) (*types.Transaction, error) {
	return _ContractConnect3DAServiceManager.Contract.DeregisterOperatorFromAVS(&_ContractConnect3DAServiceManager.TransactOpts, operator)
}

// FreezeOperator is a paid mutator transaction binding the contract method 0x38c8ee64.
//
// Solidity: function freezeOperator(address operatorAddr) returns()
func (_ContractConnect3DAServiceManager *ContractConnect3DAServiceManagerTransactor) FreezeOperator(opts *bind.TransactOpts, operatorAddr common.Address) (*types.Transaction, error) {
	return _ContractConnect3DAServiceManager.contract.Transact(opts, "freezeOperator", operatorAddr)
}

// FreezeOperator is a paid mutator transaction binding the contract method 0x38c8ee64.
//
// Solidity: function freezeOperator(address operatorAddr) returns()
func (_ContractConnect3DAServiceManager *ContractConnect3DAServiceManagerSession) FreezeOperator(operatorAddr common.Address) (*types.Transaction, error) {
	return _ContractConnect3DAServiceManager.Contract.FreezeOperator(&_ContractConnect3DAServiceManager.TransactOpts, operatorAddr)
}

// FreezeOperator is a paid mutator transaction binding the contract method 0x38c8ee64.
//
// Solidity: function freezeOperator(address operatorAddr) returns()
func (_ContractConnect3DAServiceManager *ContractConnect3DAServiceManagerTransactorSession) FreezeOperator(operatorAddr common.Address) (*types.Transaction, error) {
	return _ContractConnect3DAServiceManager.Contract.FreezeOperator(&_ContractConnect3DAServiceManager.TransactOpts, operatorAddr)
}

// PayForRange is a paid mutator transaction binding the contract method 0x1b445516.
//
// Solidity: function payForRange(((address,uint96)[],address,uint256,uint32,uint32)[] rangePayments) returns()
func (_ContractConnect3DAServiceManager *ContractConnect3DAServiceManagerTransactor) PayForRange(opts *bind.TransactOpts, rangePayments []IPaymentCoordinatorRangePayment) (*types.Transaction, error) {
	return _ContractConnect3DAServiceManager.contract.Transact(opts, "payForRange", rangePayments)
}

// PayForRange is a paid mutator transaction binding the contract method 0x1b445516.
//
// Solidity: function payForRange(((address,uint96)[],address,uint256,uint32,uint32)[] rangePayments) returns()
func (_ContractConnect3DAServiceManager *ContractConnect3DAServiceManagerSession) PayForRange(rangePayments []IPaymentCoordinatorRangePayment) (*types.Transaction, error) {
	return _ContractConnect3DAServiceManager.Contract.PayForRange(&_ContractConnect3DAServiceManager.TransactOpts, rangePayments)
}

// PayForRange is a paid mutator transaction binding the contract method 0x1b445516.
//
// Solidity: function payForRange(((address,uint96)[],address,uint256,uint32,uint32)[] rangePayments) returns()
func (_ContractConnect3DAServiceManager *ContractConnect3DAServiceManagerTransactorSession) PayForRange(rangePayments []IPaymentCoordinatorRangePayment) (*types.Transaction, error) {
	return _ContractConnect3DAServiceManager.Contract.PayForRange(&_ContractConnect3DAServiceManager.TransactOpts, rangePayments)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractConnect3DAServiceManager *ContractConnect3DAServiceManagerTransactor) RegisterOperatorToAVS(opts *bind.TransactOpts, operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractConnect3DAServiceManager.contract.Transact(opts, "registerOperatorToAVS", operator, operatorSignature)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractConnect3DAServiceManager *ContractConnect3DAServiceManagerSession) RegisterOperatorToAVS(operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractConnect3DAServiceManager.Contract.RegisterOperatorToAVS(&_ContractConnect3DAServiceManager.TransactOpts, operator, operatorSignature)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractConnect3DAServiceManager *ContractConnect3DAServiceManagerTransactorSession) RegisterOperatorToAVS(operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractConnect3DAServiceManager.Contract.RegisterOperatorToAVS(&_ContractConnect3DAServiceManager.TransactOpts, operator, operatorSignature)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractConnect3DAServiceManager *ContractConnect3DAServiceManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractConnect3DAServiceManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractConnect3DAServiceManager *ContractConnect3DAServiceManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractConnect3DAServiceManager.Contract.RenounceOwnership(&_ContractConnect3DAServiceManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractConnect3DAServiceManager *ContractConnect3DAServiceManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractConnect3DAServiceManager.Contract.RenounceOwnership(&_ContractConnect3DAServiceManager.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractConnect3DAServiceManager *ContractConnect3DAServiceManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ContractConnect3DAServiceManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractConnect3DAServiceManager *ContractConnect3DAServiceManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractConnect3DAServiceManager.Contract.TransferOwnership(&_ContractConnect3DAServiceManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractConnect3DAServiceManager *ContractConnect3DAServiceManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractConnect3DAServiceManager.Contract.TransferOwnership(&_ContractConnect3DAServiceManager.TransactOpts, newOwner)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_ContractConnect3DAServiceManager *ContractConnect3DAServiceManagerTransactor) UpdateAVSMetadataURI(opts *bind.TransactOpts, _metadataURI string) (*types.Transaction, error) {
	return _ContractConnect3DAServiceManager.contract.Transact(opts, "updateAVSMetadataURI", _metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_ContractConnect3DAServiceManager *ContractConnect3DAServiceManagerSession) UpdateAVSMetadataURI(_metadataURI string) (*types.Transaction, error) {
	return _ContractConnect3DAServiceManager.Contract.UpdateAVSMetadataURI(&_ContractConnect3DAServiceManager.TransactOpts, _metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_ContractConnect3DAServiceManager *ContractConnect3DAServiceManagerTransactorSession) UpdateAVSMetadataURI(_metadataURI string) (*types.Transaction, error) {
	return _ContractConnect3DAServiceManager.Contract.UpdateAVSMetadataURI(&_ContractConnect3DAServiceManager.TransactOpts, _metadataURI)
}

// ContractConnect3DAServiceManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContractConnect3DAServiceManager contract.
type ContractConnect3DAServiceManagerInitializedIterator struct {
	Event *ContractConnect3DAServiceManagerInitialized // Event containing the contract specifics and raw log

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
func (it *ContractConnect3DAServiceManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractConnect3DAServiceManagerInitialized)
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
		it.Event = new(ContractConnect3DAServiceManagerInitialized)
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
func (it *ContractConnect3DAServiceManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractConnect3DAServiceManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractConnect3DAServiceManagerInitialized represents a Initialized event raised by the ContractConnect3DAServiceManager contract.
type ContractConnect3DAServiceManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractConnect3DAServiceManager *ContractConnect3DAServiceManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractConnect3DAServiceManagerInitializedIterator, error) {

	logs, sub, err := _ContractConnect3DAServiceManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractConnect3DAServiceManagerInitializedIterator{contract: _ContractConnect3DAServiceManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractConnect3DAServiceManager *ContractConnect3DAServiceManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractConnect3DAServiceManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _ContractConnect3DAServiceManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractConnect3DAServiceManagerInitialized)
				if err := _ContractConnect3DAServiceManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ContractConnect3DAServiceManager *ContractConnect3DAServiceManagerFilterer) ParseInitialized(log types.Log) (*ContractConnect3DAServiceManagerInitialized, error) {
	event := new(ContractConnect3DAServiceManagerInitialized)
	if err := _ContractConnect3DAServiceManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractConnect3DAServiceManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ContractConnect3DAServiceManager contract.
type ContractConnect3DAServiceManagerOwnershipTransferredIterator struct {
	Event *ContractConnect3DAServiceManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractConnect3DAServiceManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractConnect3DAServiceManagerOwnershipTransferred)
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
		it.Event = new(ContractConnect3DAServiceManagerOwnershipTransferred)
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
func (it *ContractConnect3DAServiceManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractConnect3DAServiceManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractConnect3DAServiceManagerOwnershipTransferred represents a OwnershipTransferred event raised by the ContractConnect3DAServiceManager contract.
type ContractConnect3DAServiceManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractConnect3DAServiceManager *ContractConnect3DAServiceManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractConnect3DAServiceManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractConnect3DAServiceManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractConnect3DAServiceManagerOwnershipTransferredIterator{contract: _ContractConnect3DAServiceManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractConnect3DAServiceManager *ContractConnect3DAServiceManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractConnect3DAServiceManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractConnect3DAServiceManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractConnect3DAServiceManagerOwnershipTransferred)
				if err := _ContractConnect3DAServiceManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ContractConnect3DAServiceManager *ContractConnect3DAServiceManagerFilterer) ParseOwnershipTransferred(log types.Log) (*ContractConnect3DAServiceManagerOwnershipTransferred, error) {
	event := new(ContractConnect3DAServiceManagerOwnershipTransferred)
	if err := _ContractConnect3DAServiceManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
