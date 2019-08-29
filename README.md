# Go-Testimonium
This project contains a Go-library and command-line interface (CLI) to interact with the [Testimonium](https://github.com/pantos-io/testimonium) prototype.
 
The Testimonium prototype enables the cross-blockchain verification of transactions. 
That is, a "verifying" blockchain can verify that a certain transaction (receipt, state) is included 
in a different "target" blockchain without relying on third-party trust.

Detailed information about how the prototype works can be found [here](https://dsg.tuwien.ac.at/projects/tast/pub/tast-white-paper-6.pdf).

> _Important: Testimonium is a research prototype. 
    It represents ongoing work conducted within the [TAST](https://dsg.tuwien.ac.at/projects/tast/) 
    research project. Use with care._
## Prerequisites
You need to have [Golang](https://golang.org/doc/install) and [Ganache](https://www.trufflesuite.com/ganache) installed. 
## Get Started
_The following setup will take you through the deployment of Testimonium with a local Ethereum blockchain (Ganache)
as verifying chain and the main Ethereum chain as target chain.
Information on how to connect other blockchains can be found [here](#Configuration)._


1. Install the library and CLI with `$ go get github.com/pantos-io/go-testimonium`.  
Check that the CLI was installed correctly by running `$ go-testimonium --help`.

2. Run `go-testimonium init` to initialize the client.

3. Start Ganache (port 8545)

4. Deploy the Ethash contract with `go-testimonium deploy ethash`. 
This deploys the contract responsible for verifying the Proof of Work (PoW) of a block.

5. Submit the correct epoch data to the Ethash contract with `go-testimonium submit epoch <EPOCH_NO>`.
Depending on which block will be submitted as genesis block to the Testimonium contract, 
the correct epoch data can be calculated as `EPOCH_NO = BLOCK_NO / 30000`. 
    > e.g., for genesis block 8084509 the correct epoch data is 269

6. Deploy the Testimonium contract with `go-testimonium deploy testimonium --genesis <BLOCK_NO>`.
This deploys the contract responsible for the verification of transactions (receipts, state). 
The `genesis` parameter specifies the first block of the target chain which will be submitted to 
the Testimonium contract. Verifications will be possible for all subsequent blocks.

###
The Testimonium prototype is now setup. 
You can now submit block data from the target chain to the verifying chain, 
and request verifications of transactions, and dispute illegal blocks. 

## Usage
The CLI can be started with `go-testimonium [command]` where `[command]` is one of the commands below.

Use `go-testimonium [command] --help` for more information about a command.

---

`init`: Initializes the client by creating a testimonium.yml file in the current directory.

`account`: Prints the address of the current account

`balance`: Prints the balance of the current account

`deploy ethash`: Deploys the Ethash smart contract on the verifying chain

`deploy testimonium`: Deploys the Testimonium contract on the verifying chain

`dispute [blockHash]`: Disputes the submitted block header with the specified hash

`get block [blockHash]`: Retrieves the block with the specified hash

`get transaction [txHash]`: Retrieves the transaction with the specified hash

`submit block [blocknumber]`: Submits the specified block header from the target chain to the verifying chain

`submit epoch [epoch]`: Sets the epoch data for the specified epoch on the verifying chain

`verify block [blockHash]`: Verifies a block from the target chain on the verifying chain

`verify transaction [txHash]`: Verifies a transaction from the target chain on the verifying chain

`verify receipt [txHash]`: Verifies a receipt from the target chain on the verifying chain


## Configuration
The Testimonium client uses a configuration file called `testimonium.yml` file.

The default file looks like this:

    privateKey: <YOUR PRIVATE KEY>
    chains:
        0:
            url: mainnet.infura.io
        1:
            type: http
            url: localhost
            port: 8545

Chain ID 0 contains connection configuration for the main Ethereum chain (via Infura).
Chain ID 1 contains connection configuration for a local chain (e.g., run via Ganache).

You can configure the Testimonium client for other Ethereum blockchains (there is no upper limit).
Just manually add or edit a chain entry under the `chains` key.
Key `type` refers to the connections type (e.g., http, https, ws, wss), 
`url` refers to the URL, 
and `port` refers to the port number under which the specific chain is reachable. If no type is specified, https is used.
If no port is defined, it is determined by the default port of the type.


If you have already deployed the Ethash and Testimonium contracts, you might find further entries
`ethashAddress` and `testimoniumAddress` under a specific chain config:

    ...
    chains:
        ...
        1:
            type: http
            url: localhost
            port: 8545
            testimoniumAddress: 0xabc123...
            ethashAddress: 0x123abc...

These are the addresses that the client uses to interact with the Testimonium smart contracts.
If you deployed the contracts manually, just add the entries. 

## How to Contribute
Testimonium is a research prototype. We welcome anyone to contribute.
File a bug report or submit feature requests through the [issue tracker](https://github.com/pantos-io/go-testimonium/issues). 
If you want to contribute feel free to submit a pull request.

## Acknowledgement
The development of this prototype was funded by [Pantos](https://pantos.io/) within the [TAST](https://dsg.tuwien.ac.at/projects/tast/) research project.

## Licence
This project is licensed under the [MIT License](LICENSE).
