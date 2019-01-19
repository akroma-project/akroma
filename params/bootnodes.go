// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Akroma network.
var MainnetBootnodes = []string{
	"enode://2b2895b069aac74f1289ddbf6c083df7f549817fb488065261d9b5b8c4c175bc9103670ed32b70f3f7110a3658dfe9fbe39ff8828d9fc48df0a777971a4fc2b3@51.38.98.51:30303",
	"enode://bfe245cebd3f028880ea6844687baa8915475826ab74d9daf14bdbb6e931d35a30ebf6f78a5c66aff3ac34e3501e5d6ddf1e182ee3ce65907e84bac921133013@51.38.133.221:30303",
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var TestnetBootnodes = []string{
	"enode://7fdacaac41d84d56c0713797f75392c74419c56478e39e9831ecc3685e752dd53373df515b5caf1cb9b62af6682c13c2417be8dc059e3dae8ade9660409bd3f3@159.89.41.132:30303", //west-us
}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{}
