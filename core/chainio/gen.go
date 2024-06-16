package chainio

//go:generate mockgen -destination=./mocks/avs_subscriber.go -package=mocks github.com/Layr-Labs/connect3-da-avs/core/chainio AvsSubscriberer
//go:generate mockgen -destination=./mocks/avs_writer.go -package=mocks github.com/Layr-Labs/connect3-da-avs/core/chainio AvsWriterer
//go:generate mockgen -destination=./mocks/avs_reader.go -package=mocks github.com/Layr-Labs/connect3-da-avs/core/chainio AvsReaderer
