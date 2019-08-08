# Go-Testimonium
This project contains a Go-library and command-line interface (CLI) to interact with the [Testimonium](todo add link) prototype.
 
The Testimonium prototype enables the cross-blockchain verification of transactions. 
That is, a "verifying" blockchain can verify that a certain transaction (receipt, state) is included 
in a different "target" blockchain without relying on third-party trust. 

> _Important: Testimonium is a research prototype. 
    It represents ongoing work conducted within the [TAST](https://dsg.tuwien.ac.at/projects/tast/) 
    research project. Use with care._
    
## Get Started
You need to have [Golang](https://golang.org/) installed.
To install follow the instructions [here](https://golang.org/doc/install).

Install the library and CLI with `$ go get github.com/pantos-io/go-testimonium`.  
Check that the CLI was installed correctly by running `$ go-testimonium --help`.

TODO How to create testimonium.yml file

TODO How to deploy the contracts

## Usage
The CLI can be started with `go-testimonium [command]` where `[command]` is one of the following:

`account`: Prints the address of the current account

`balance`: Prints the balance of the current account

`dispute [blockHash]`: Disputes the submitted block header with the specified hash

`get block [blockHash]`: Retrieves the block with the specified hash

`get transaction [txHash]`: Retrieves the transaction with the specified hash

`submit block [blocknumber]`: Submits the specified block header from the target chain to the verifying chain

`submit epoch [epoch]`: Sets the epoch data for the specified epoch on the verifying chain

`verify block [blockHash]`: Verifies a block from the target chain on the verifying chain

`verify transaction [txHash]`: Verifies a transaction from the target chain on the verifying chain

`verify receipt [txHash]`: Verifies a receipt from the target chain on the verifying chain

Use `go-testimonium [command] --help` for more information about a command.

## How to Contribute
Testimonium is a research prototype. We welcome anyone to contribute.
File a bug report or submit feature requests through the [issue tracker](https://github.com/pf92/go-testimonium/issues). 
If you want to contribute feel free to submit a pull request.

## Acknowledgement
The development of this prototype was funded by [Pantos](https://pantos.io/) within the [TAST](https://dsg.tuwien.ac.at/projects/tast/) research project.

## Licence
This project is licensed under the [MIT License](LICENSE).
