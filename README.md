## Akroma

Akroma is an EVM based application development platform (smart-contracts). Akroma will utilize a Masternode system, and build out an Oracle platform. Akroma is a fork of the Ethereum code with small changes being added to support the concept of Masternodes and a stable governance model via the Akroma Foundation. Akroma inherits all of the features of Ethereum and will merge in upstream changes from the Ethereum go-lang client.

- FAQ: https://medium.com/akroma/akroma-faq-and-ama-674fe45d7dfc
- Whitepaper/Roadmap: https://medium.com/akroma/hello-akroma-f413245a342b

## Build

[![Build Status](https://travis-ci.org/akroma-project/akroma.svg?branch=master)](https://travis-ci.org/akroma-project/akroma)

## Releases

Official releases are published at https://github.com/akroma-project/akroma/releases

## Attribution

This README has be adapted from the Go Ethereum official implementation, the original version is available here: https://github.com/akroma-project/akroma

## Building the source

Building geth requires both a Go (version 1.9 or later) and a C compiler.
You can install them using your favorite package manager.
Once the dependencies are installed, run

    make geth

or, to build the full suite of utilities:

    make all

Since Akroma is derived from the Ethereum codebase and has not diverged from it, you can reference the Ethereum build documentation for more details on building Akroma: https://github.com/akroma-project/akroma/wiki/Building-Ethereum

## Executables

The Akroma project comes with several wrappers/executables found in the `cmd` directory.

| Command    | Description |
|:----------:|-------------|
| **`geth`** | Our main CLI client. It is the entry point into the Akroma network, capable of running as a full node (default) archive node (retaining all historical state) or a light node (retrieving data live). It can be used by other processes as a gateway into the Akroma network via JSON RPC endpoints exposed on top of HTTP, WebSocket and/or IPC transports. `geth --help` for command line options. |

Since Akroma is derived from the Ethereum codebase and has not diverged from it, you can reference the Ethereum RPC documentation: https://github.com/akroma-project/akroma/wiki/Command-Line-Options

### Full node on the main Akroma network

By far the most common scenario is people wanting to simply interact with the Akroma network:
create accounts; transfer funds; deploy and interact with contracts. For this particular use-case
the user doesn't care about years-old historical data, so we can fast-sync quickly to the current
state of the network. To do so:

```
$ geth --fast --cache=512 console
```

This command will:

 * Start geth in fast sync mode (`--fast`), causing it to download more data in exchange for avoiding
   processing the entire history of the Akroma network, which is very CPU intensive.
 * Bump the memory allowance of the database to 512MB (`--cache=512`), which can help significantly in
   sync times especially for HDD users. This flag is optional and you can set it as high or as low as
   you'd like, though we'd recommend the 512MB - 2GB range.
 * Start up Geth's built-in interactive JavaScript console, (via the trailing `console` subcommand) through which you can invoke all official `web3` methods.
   This too is optional and if you leave it out you can always attach to an already running Geth instance
   with `geth attach`.

### Configuration

As an alternative to passing the numerous flags to the `geth` binary, you can also pass a configuration file via:

```
$ geth --config /path/to/your_config.toml
```

To get an idea how the file should look like you can use the `dumpconfig` subcommand to export your existing configuration:

```
$ geth --your-favourite-flags dumpconfig
```

*Note: This works only with geth v1.6.0 and above.*

#### Docker quick start

One of the quickest ways to get Akroma up and running on your machine is by using Docker:

```
docker run -d --name akroma-node -v /Users/alice/akroma:/root \
           -p 8545:8545 -p 40403:40403 \
           akroma/client-go --fast --cache=512
```

This will start geth in fast sync mode with a DB memory allowance of 512MB just as the above command does.  It will also create a persistent volume in your home directory for saving your blockchain as well as map the default ports. There is also an `alpine` tag available for a slim version of the image.

Do not forget `--rpcaddr 0.0.0.0`, if you want to access RPC from other containers and/or hosts. By default, `geth` binds to the local interface and RPC endpoints is not accessible from the outside.

### Programatically interfacing Geth nodes

As a developer, sooner rather than later you'll want to start interacting with Geth and the Akroma
network via your own programs and not manually through the console. To aid this, Geth has built in
support for a JSON-RPC based APIs. These can be exposed via HTTP, WebSockets and IPC (unix sockets on unix based platforms, and named pipes on Windows).

The IPC interface is enabled by default and exposes all the APIs supported by Geth, whereas the HTTP
and WS interfaces need to manually be enabled and only expose a subset of APIs due to security reasons.
These can be turned on/off and configured as you'd expect.

HTTP based JSON-RPC API options:

  * `--rpc` Enable the HTTP-RPC server
  * `--rpcaddr` HTTP-RPC server listening interface (default: "localhost")
  * `--rpcport` HTTP-RPC server listening port (default: 8545)
  * `--rpcapi` API's offered over the HTTP-RPC interface (default: "eth,net,web3")
  * `--rpccorsdomain` Comma separated list of domains from which to accept cross origin requests (browser enforced)
  * `--ws` Enable the WS-RPC server
  * `--wsaddr` WS-RPC server listening interface (default: "localhost")
  * `--wsport` WS-RPC server listening port (default: 8546)
  * `--wsapi` API's offered over the WS-RPC interface (default: "eth,net,web3")
  * `--wsorigins` Origins from which to accept websockets requests
  * `--ipcdisable` Disable the IPC-RPC server
  * `--ipcapi` API's offered over the IPC-RPC interface (default: "admin,debug,eth,miner,net,personal,shh,txpool,web3")
  * `--ipcpath` Filename for IPC socket/pipe within the datadir (explicit paths escape it)

You'll need to use your own programming environments' capabilities (libraries, tools, etc) to connect
via HTTP, WS or IPC to a Geth node configured with the above flags and you'll need to speak [JSON-RPC](https://www.jsonrpc.org/specification)
on all transports. You can reuse the same connection for multiple requests!

**Note: Please understand the security implications of opening up an HTTP/WS based transport before
doing so! Hackers on the internet are actively trying to subvert Akroma nodes with exposed APIs!
Further, all browser tabs can access locally running webservers, so malicious webpages could try to
subvert locally available APIs!**

## Contribution

Thank you for considering to help out with the source code! We welcome contributions from
anyone on the internet, and are grateful for even the smallest of fixes!

If you'd like to contribute to Akroma, please fork, fix, commit and send a pull request
for the maintainers to review and merge into the main code base. If you wish to submit more
complex changes though, please check up with the core devs first to ensure those changes are in line with the general philosophy of the project and/or get some
early feedback which can make both your efforts much lighter as well as our review and merge
procedures quick and simple.

Please make sure your contributions adhere to our coding guidelines:

 * Code must adhere to the official Go [formatting](https://golang.org/doc/effective_go.html#formatting) guidelines (i.e. uses [gofmt](https://golang.org/cmd/gofmt/)).
 * Code must be documented adhering to the official Go [commentary](https://golang.org/doc/effective_go.html#commentary) guidelines.
 * Pull requests need to be based on and opened against the `master` branch.
 * Commit messages should be prefixed with the package(s) they modify.
   * E.g. "eth, rpc: make trace configs optional"

Please see the Ethereum [Developers' Guide](https://github.com/akroma-project/akroma/wiki/Developers'-Guide)
for more details on configuring your environment, managing project dependencies and testing procedures.

## License

The Akroma library (i.e. all code outside of the `cmd` directory) is licensed under the
[GNU Lesser General Public License v3.0](https://www.gnu.org/licenses/lgpl-3.0.en.html), also
included in our repository in the `COPYING.LESSER` file.

The Akroma binaries (i.e. all code inside of the `cmd` directory) is licensed under the
[GNU General Public License v3.0](https://www.gnu.org/licenses/gpl-3.0.en.html), also included
in our repository in the `COPYING` file.
