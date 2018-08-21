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
	// Akroma Foundation Go Bootnodes
	"enode://b68cec80d402504826e3a0c55c9974633b34fdbc1cc79284a0bea333420541d62e4b95b74b23a1cf2dd3dbe4e9aca41addb53aa2871a1a168faf75536dc752fd@40.79.84.146:30303",    //seed1
	"enode://dab9ed40ef32b279452612cb345b0c0e68feca8f9dd76534efe21abe48142f7418b98aa58f078b7cffde2fcc141275e87120d9becc39ff8d01a28b1e4ef98921@52.191.171.167:30303",  //seed2 (correct)
	"enode://d8f2665393cc9cb24cd976189615c7f233a5e458fc8594dd081f0041b4fee45daf8d295bcd783d7a9ee2b3052c547c661a77f99597bed016850aefcbc99a2b33@188.166.175.111:30303", //london
	"enode://8b65cac4d74ada1233b7a7439e7f69d29c7a2be0a97c14e03be17d231014dc4dcaec1d014eb408a9b0442a89130a1f86e819552d48ef4a944457f85c6120febc@159.65.6.121:30303",    // singpore

	// Decentralized Bootnodes.
	"enode://30bc56125052b3dabb4f0916bbcf26c731e8d77a7f5cd908782ded9028d1b0d7a43e8db476989224b41abce122c1bf2cdf5fc4c3a84aa57063233e0e33e82efa@45.77.98.222:30303", // 2313

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
