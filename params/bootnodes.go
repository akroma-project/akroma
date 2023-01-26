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
	"enode://9a670c1cb783f0519c7ff5210017b6685cae4f3bd4f092b310a08f25c8fe4e052b2a999e7084be91902e61f4469bc398377446b31b4e5cbfa01762fe67deffd9@167.71.17.10:30303",    //akaboot2
	"enode://171c35270162bb4257e7392b5aa274bc6fe92fc79462a785d139d9bc8730e6b9d3e2d85794e39b23331463818883717bc6dd890af19fab52fb0401f165d770a0@147.182.158.117:30303", //akaboot3
	"enode://0c3e27e6271eaf4bee81912bf727e5232d8af0889ce9819a68b35cad8089c68d14ec7d98f49f827207377c374196a4115b36ba7b19934c0196246ea2485632fd@45.56.74.194:30303",    //akaboot4
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Akroma test network.
var TestnetBootnodes = []string{
	"enode://10f7ee9a4d6de968df7680a495d38154f4f313d168c382c89295ad7376f78af7af3b612be4ee9b476eeab3203e5072504209ae5c227bb42d746a18de2066e23e@161.35.118.238:30303",
}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{}
