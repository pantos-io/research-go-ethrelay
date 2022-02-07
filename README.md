# Go-ETH Relay

This project contains a Go library and a command-line interface (CLI) to interact with the [ETH Relay](https://github.com/pantos-io/ethrelay) prototype.

ETH Relay enables the cross-blockchain verification of transactions.
This makes it possible that a destination blockchain can verify if a certain transaction (or receipt, or state) is included
in a different source blockchain without relying or trusting a third-party.

Detailed information about how the prototype works can be found [here](https://dsg.tuwien.ac.at/projects/tast/pub/tast-white-paper-6.pdf).

> _Important: ETH Relay is a research prototype.
    It represents ongoing work conducted within the [TAST](https://dsg.tuwien.ac.at/projects/tast/)
    research project. Use with care._

## Prerequisites

You need to have [Golang](https://golang.org/doc/install) and [Ganache](https://www.trufflesuite.com/ganache) (>= 2.1.0) installed.

## Get Started

_The following setup will take you through the deployment of ETH Relay with a local Ethereum blockchain (Ganache)
as the destination chain and the main Ethereum chain as the source chain.
Information on how to connect to other blockchains can be found [here](#Configuration)._

1. Run `git clone https://github.com/pantos-io/go-ethrelay.git` to clone the repository on your local machine and then run `go get ./cmd/go-ethrelay` in the go-ethrelay folder to install the library and the CLI.
Check that the CLI was installed correctly by running `go-ethrelay`. In case `go-ethrelay` command is not found, have a look at the [Troubleshooting](#Troubleshooting) section.
If you want to install the library manually, you can simply clone this repository and run any command in the cloned directory with `go run ./cmd/go-ethrelay [command]`.

2. Run `go-ethrelay init` to initialize the client.
If you encounter any problems calling this command, get sure the rights are properly adjusted so Go can create the ethrelay.yml config file in the current folder.
It is also possible to generate the file by hand or change the example config file named ethrelay.example.yml contained in this repo.

3. Start Ganache (should start on the default port 7545, if not, change this in the config file).

4. Deploy the Ethash contract with `go-ethrelay deploy ethash`.
This deploys the contract responsible for verifying the Proof of Work (PoW) of a block.

5. Submit the correct epoch data to the Ethash contract with `go-ethrelay submit epoch <EPOCH_NO>`.
Depending on which block will be submitted as genesis block to the ETH Relay contract in the following step,
the correct epoch data can be calculated as `EPOCH_NO = BLOCK_NO / 30000` floored. This may take a while.
    > e.g., for genesis block 8084509, the correct epoch data is 269

6. Deploy the ETH Relay contract with `go-ethrelay deploy ethrelay --genesis <BLOCK_NO>`.
This deploys the contract responsible for the verification of transactions (or receipts, or state).
The `genesis` parameter specifies the first block of the source chain which will be submitted to
the ETH Relay contract. Verifications will be possible for all subsequent blocks.

ETH Relay is now setup. In order to submit blocks, a stake has to be submitted first. After you deposited some stake, you can submit block data from the source chain to the destination chain,
and request verifications of transactions, and dispute illegal blocks.

## Usage

The CLI can be started with `go-ethrelay [command]` where `[command]` is one of the commands below.

Use `go-ethrelay [command] --help` for more information about a command.

---

`init`: Initializes the client by creating an ethrelay.yml file in the current directory that acts as config file for all command calls.

`account`: Prints the address of the current account

`balance`: Prints the balance of the current account

`deploy ethash`: Deploys the Ethash smart contract on a destination chain

`deploy ethrelay`: Deploys the ETH Relay contract on a destination chain

`dispute [blockHash]`: Disputes the submitted block header with the specified hash

`generate [blockNumber]`: Generates and exports test data for the Ethrelay project

`get block [blockHash]`: Retrieves the block with the specified hash

`get transaction [txHash]`: Retrieves the transaction with the specified hash

`get longestchainendpoint`: Retrieves the most recent block hash of the longest chain in the ETH Relay contract on a destination chain

`stake get`: Retrieves the amount of stake deposited in the relay-contract on a destination chain

`stake deposit [amountInWei]`: Deposits amountInWei stake of the account balance in the contract

> e.g. `stake deposit 25000000000000000000` deposits 25 ETH

`stake withdraw [amountInWei]`: Withdraws the submitted stake back to the account balance. Remember that stake can be locked in the contract when a block was submitted and you have to wait until it is unlocked again.

`submit block [blockNumber or blockHash]`: Submits the specified block header from a source chain to a destination chain

`submit epoch [epoch]`: Sets the epoch data for the specified epoch on a destination chain

`verify block [blockHash]`: Verifies a block from the source chain on a destination chain

`verify transaction [txHash]`: Verifies a transaction from the source chain on a destination chain

`verify receipt [txHash]`: Verifies a receipt from the source chain on a destination chain

## Quick Setup

There is also a shell script in this repository named `setup-relay.sh`. This script helps researchers and developers to quickly setup
a working version of the relay with default-values and assuming a local Ganache instance. The call to this script is:

`sh setup-relay.sh [ganacheAccountPrivateKey] [genesisBlock] [stakeInETH]`

Example call:

```shell
sh setup-relay.sh 0x45b5ffd7266ec7131f31f94da843b99fd270b46d94bf01368ceeb936649dfc3b 11367417 25
```

## Configuration

The relay client uses a configuration file called `ethrelay.yml` file.

The default file looks like this:

```yml
chains:
  sources:
    mainnet:
      type: wss
      url: mainnet.infura.io/ws/v3/1e835672adba4b9b930a12a3ec58ebad
  destinations:
    local:
      port: 7545
      type: http
      url: localhost
privatekey: 0x0
```

Chain `mainnet` contains connection configuration for the main Ethereum chain (via Infura).
Chain `local` contains connection configuration for a local chain (e.g., run via Ganache).

You can configure the relay client for other Ethereum blockchains (there is no upper limit).
Just manually add or edit a chain entry under the `sources` or `destinations` key.
Key `type` refers to the connection type (e.g., http, https, ws, wss), `url` refers to the URL
and `port` refers to the port number under which the specific chain is reachable. If no type is specified, https is used.
If no port is defined, it is determined by the default port of the type.

If you have already deployed the Ethash and ETH Relay contracts, you might find further entries, i.e.
`ethashaddress` and `ethrelayaddress` under a specific chain config:

```yml
chains:
  destinations:
    local:
      ethashaddress: 0x123abc...
      ethrelayaddress: 0xabc123...
```

These are the addresses that the client uses to interact with the ETH Relay smart contracts.
If you deployed the contracts manually, just add the entries.

## Troubleshooting

### Dispute causes error: "VM Exception while processing transaction: revert"

If disputing a certain block causes a generic revert exception, make sure you are running Ganache version >= 2.1.0.

### Some Go libraries or dependencies could not be found or have wrong naming/versions

Maybe you have some old dependencies or versions installed and an error like the following occurs:

```text
go: github.com/pantos-io/go-ethrelay@v0.1.0: parsing go.mod:
    module declares its path as: github.com/pantos-io/go-testimonium
    but was required as: github.com/pantos-io/go-ethrelay
```

Try to clean the modules folder of Go with `go clean -modcache` and install the modules with `go get` again.

#### go-ethrelay command not found after installing from Github

Add the GOBIN and/or GOPATH to your PATH-variable. Find more information about GOBIN and GOPATH [here](https://golang.org/cmd/go/#hdr-GOPATH_environment_variable).

#### Client won't start because of Go problems

Get sure your Go-path variables like GOHOME, GOPATH und GOBIN are set properly.

#### Client won't work because issues to permissions or projectId (e.g. forbidden: project ID is required)

If you are using Infura, check that your url and protocol is correct and the permissions are properly set in the Infura admin panel.

## How to Contribute

ETH Relay is a research prototype. We welcome anyone to contribute.
File a bug report or submit feature requests through the [issue tracker](https://github.com/pantos-io/go-ethrelay/issues).
If you want to contribute feel free to submit a pull request.

## Useful resources

* [Etherscan](https://etherscan.io)
* [Ethereum Development with Go](https://goethereumbook.org)
* [Solidity documentation](https://solidity.readthedocs.io)

## Acknowledgement

The development of this prototype was funded by [Pantos](https://pantos.io/) within the [TAST](https://dsg.tuwien.ac.at/projects/tast/) research project.

## Licence

This project is licensed under the [MIT License](LICENSE).
