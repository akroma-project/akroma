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
	"enode://e9fb6b5bfa04bbfe134e2b34d761d669defe68a2ce311a59231d524ccd003b6f833703546a58dc9a6791939e63226fa04beaa02a6d4fc8aa603fff537bf9642d@137.184.120.92:30303",
	"enode://dc4fc0bb21ce46a048d51e0e4722a79918d2a671c24885d53d2bea815b32178cb28d3544170fe0404095d9fe8dc5770814300effba773854420aad563138eceb@192.99.244.230:30303",
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
