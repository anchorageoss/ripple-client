# Updated version of [rubblelabs/riple](https://github.com/rubblelabs/ripple) client

Go packages to interact with the Ripple protocol.

[![Go Reference](https://pkg.go.dev/badge/github.com/anchorageoss/ripple-client.svg)](https://pkg.go.dev/github.com/anchorageoss/ripple-client)

The data, crypto, and websockets packages are very functional and quite well tested. Most websockets commands are implemented but not all.

The peers and ledger packages are the least polished packages currently, and they are very much unfinished (and the tests might be non-existent or non-functional), but better to get the code out in the open.

We've included command-line tools to show how to apply the library:

* listener: connects to rippled servers with the peering protocol and displays the traffic
* subscribe: tracks ledgers and transactions via websockets and explains each transaction's metadata
* tx: creates transactions, signs them, and submits them via websockets
* vanity: generates new ripple wallets in search of vanity addresses