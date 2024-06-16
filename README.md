# Connect3 DA Demo - Fetching C3Data by Ethereum Address

[![Made by Connect3](https://img.shields.io/badge/made%20by-Connect3-purple.svg)](https://connect3.world/)
[![Chat on discord](https://img.shields.io/badge/join%20-discord-green.svg)](https://discord.com/invite/uBMruvxdNa)

Welcome to the AVS Project! This project is based on modifications of the "incredible-squaring-avs" project. It is designed to demonstrate how to fetch C3Data using an Ethereum address.
## Overview

The AVS Project integrates with the Connect3 ecosystem to fetch and manage data across various decentralized and on-chain storage systems. This project leverages smart contracts, off-chain components, and third-party services to provide a seamless data retrieval mechanism.

## Architecture

### AVS Architecture

![avs](https://github.com/connect3world/connect3-da-demo/assets/9106716/172258d0-6e1c-44f3-b54f-50280bdb063e)


The AVS architecture is designed to facilitate interaction between the core Ethereum components, AVS contracts, and off-chain applications. The key components include:

- **Consumer Client**: Interacts with the dAppUI.
- **dAppUI**: Provides the user interface for the on-chain dApp.
- **On-chain dApp**: Interfaces with the AVS contracts.
- **Ethereum**: The blockchain network where the core contracts and AVS contracts reside.
    - **EigenLayer Core Contracts**: Includes the Delegation Manager, Strategy Manager, and AVS Directory.
    - **AVS Contracts**: Contains the Service Manager and other service-related contracts.
- **Off-chain Components**: Includes C3 Nodes, third-party Ethereum RPC providers, and off-chain applications.
    - **C3 Nodes**: Executes the data retrieval logic and interacts with the RPC providers.

In this project demonstration, the functionality is to fetch C3Data using an Ethereum address and store the data hash on-chain through a smart contract. However, in our further design, a consumer will be able to create tasks to synchronize all post data over a period of time. Using a Merkle tree, the hash of this aggregated post data will be computed and stored on-chain. The actual data will be kept off-chain, where nodes will provide various aggregation services.

### Connect3 High-Level Architecture

![c3-high-level](https://github.com/connect3world/connect3-da-demo/assets/9106716/1fa6b2fa-2d6e-4279-bf2f-b00988fbb19a)


The Connect3 architecture is designed to showcase the comprehensive structure of our Connect3 project, highlighting how various components interact to ensure data availability and verification. Within this architecture, the AVS serves as the Data Availability (DA) layer, playing a crucial role in managing and retrieving data efficiently.

- **User**: Interacts with the system through the C3 SDK and C3 API.
- **C3 SDK and API**: Provide interfaces for interaction with the Connect3 ecosystem.
- **C3 Publisher**: Publishes data to the C3 Node.
- **C3 Node**: Core processing unit that interacts with decentralized storage, on-chain data, and data availability layers.
- **Decentralized Storage**: Includes IPFS and Arweave for off-chain data storage.
- **On-Chain Data**: Supports multiple blockchain networks like Ethereum, Solana, and Near.
- **C3 DA Layer**: Manages data availability storage.
- **Other DA Layers**: Integrates with platforms like Lens and Farcaster.
- **C3 Retriever**: Contains engines for query and recommendation.
- **Verification**: Ensures data integrity through a consensus mechanism and verifier.

## Features

- Fetches C3Data using a provided Ethereum address.
- Integrates with multiple decentralized storage solutions.
- Utilizes smart contracts for managing data access and retrieval.
- Includes off-chain components for enhanced functionality and performance.

## Getting Started

### Prerequisites

- Docker
- Docker Compose
- Go (Golang)

### Dependencies

You will need [foundry](https://book.getfoundry.sh/getting-started/installation) and [zap-pretty](https://github.com/maoueh/zap-pretty) and docker to run the examples below.
```
curl -L https://foundry.paradigm.xyz | bash
foundryup
go install github.com/maoueh/zap-pretty@latest
```
You will also need to [install docker](https://docs.docker.com/get-docker/), and build the contracts:
```
make build-contracts
```

### Running via make

This simple session illustrates the basic flow of the AVS. The makefile commands are hardcoded for a single operator, but it's however easy to create new operator config files, and start more operators manually (see the actual commands that the makefile calls).

Start anvil in a separate terminal:

```bash
make start-anvil-chain-with-el-and-avs-deployed
```

The above command starts a local anvil chain from a [saved state](./tests/anvil/avs-and-eigenlayer-deployed-anvil-state.json) with eigenlayer and connect3 DA contracts already deployed (but no operator registered).

Start the aggregator:

```bash
make start-aggregator
```

Register the operator with eigenlayer and connect3-da, and then start the process:

```bash
make start-operator
```

> By default, the `start-operator` command will also setup the operator (see `register_operator_on_startup` flag in `config-files/operator.anvil.yaml`). To disable this, set `register_operator_on_startup` to false, and run `make cli-setup-operator` before running `start-operator`.

### Configuration

Configuration files are located in the `config-files` directory. You can modify these files to suit your environment and requirements.

## Example C3Data

Below is an example of C3Data for the Ethereum address "0xd8da6bf26964af9d7eed9e03e53415d37aa96045":

```json
{
    "ucid": "p/cast:0x27d9dc6065a8153a8a1c96cd4881359b22b01a84",
    "cuid": "clxajpdwn288801rukhz2y0we",
    "title": "Post by @vitalik.eth",
    "body": "Welcome @optimism to the club of stage 1+ L2s! (meaning, L2s where the proof systems actually have teeth)\n\nI'm looking forward to seeing many more L2s join this club soon, especially some ZK ones.",
    "publishedOn": 1718117922,
    "url": "https://warpcast.com/vitalik.eth/0x27d9dc6065a8153a8a1c96cd4881359b22b01a84",
    "appId": "warpcast",
    "image_urls": [
        "https://media.firefly.land/farcaster/59ac61b8-af4d-46cf-899f-01dddd1baaba.png"
    ],
    "audio_urls": [],
    "video_urls": [],
    "thumbnail": {
        "url": "https://image.cdn.connect3.world/3379faab891fffcf6665ac9b24256efe.png",
    }
}
```

### Field Descriptions

- **ucid**: The unique content identifier for the post, indicating its origin and type.
- **cuid**: The unique identifier for the content within the Connect3 ecosystem.
- **title**: The title of the post.
- **body**: The main content of the post.
- **publishedOn**: The timestamp of when the post was published.
- **url**: The URL linking to the native DApp ui of the post.
- **appId**: The identifier for the application where the post was created.
- **image_urls**: An array of URLs linking to images associated with the post.
- **audio_urls**: An array of URLs linking to audio files associated with the post (empty in this example).
- **video_urls**: An array of URLs linking to video files associated with the post (empty in this example).
- **thumbnail**: An object containing information about the thumbnail image, including:
  - **id**: The identifier for the thumbnail.
  - **url**: The URL of the thumbnail image.
  - **transformed**: Any transformed versions of the thumbnail (null in this example).

This C3Data example demonstrates the structure and fields used to describe a piece of content fetched using an Ethereum address within the Connect3 ecosystem.


### Next Step: Aggregating Social Data with AVS
Connect3 currently provides capabilities for aggregating social data. In the next phase, AVS will integrate these capabilities through a series of planned enhancements. Tasks will be triggered to synchronize and fetch data from sources such as Farcaster Hub, Lens Momoka, and IPFS/Arweave. The Operator will fetch, aggregate, and restructure this social data into standard posts. Using a Merkle tree, the root hash of all aggregated data between tasks will be computed and stored on-chain. This root hash can be verified through challenges to ensure data integrity. The goal is to migrate this business logic to AVS, enhancing its functionality.

![next_step](https://github.com/connect3world/connect3-da-demo/assets/9106716/530565f3-0f5e-473c-a621-92e8459d2fb5)


## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contact

For any inquiries or issues, please open an issue on GitHub or leave a comment on Discord.

---

### Useful Links

- Medium: [Connect3](https://medium.com/@Connect3)
- Discord: [Connect3](https://discord.com/invite/uBMruvxdNa)
- GitHub: [connect3world](https://github.com/connect3world)
- Twitter: [connect3world](https://twitter.com/connect3world)

---
