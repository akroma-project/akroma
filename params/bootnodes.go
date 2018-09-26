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
	"enode://d8f2665393cc9cb24cd976189615c7f233a5e458fc8594dd081f0041b4fee45daf8d295bcd783d7a9ee2b3052c547c661a77f99597bed016850aefcbc99a2b33@188.166.175.111:30303", //london
	"enode://8b65cac4d74ada1233b7a7439e7f69d29c7a2be0a97c14e03be17d231014dc4dcaec1d014eb408a9b0442a89130a1f86e819552d48ef4a944457f85c6120febc@159.65.6.121:30303",    // singpore

	// Decentralized Bootnodes.
	"enode://30bc56125052b3dabb4f0916bbcf26c731e8d77a7f5cd908782ded9028d1b0d7a43e8db476989224b41abce122c1bf2cdf5fc4c3a84aa57063233e0e33e82efa@45.77.98.222:30303",   // 2313
	"enode://1184f5091d6eb5a91f468d89074d27eee699255bf7ae22db130d40208814234dfa15e614e179444c2b2d107509e9a4d0e7b1d744b0de442ff783f7133a16a8dd@207.148.15.103:30303", // 2315
	"enode://c81ef520130728fd6cc1e16de5f79483d4bfcbac2a20ce65fc3cbf57ae2cba4fa055d63274d629d80484eb59c6d5fe48a4a533327222ed93dcf2a53216dbd24b@140.82.34.133:30303",  // 2322
	"enode://3bfa6e4578328250d45c9098c936a886f9ec2948d671f7709e3ff601a37a42d8fca6bb3f7fc0314aea36fd1e7d389f8d5fdafd51e8719a46e29c940f72cd7e4d@192.99.244.169:30303", // 2338
	"enode://9d8ce3d53245cb6ba4e8c022101c8fdba6da46e21296f9abef93ef99826a8c0a11ee8b51e9297a93a2815c5ec82e0dec68d3f06f2fed626aed79e7c5b4502640@198.58.113.208:30303", // 2345
	"enode://d4d0370e7ee85a2c499ef67a3914e61565f6cfbc8b267a0551aeaf12d7dfd1daf9208f6cfd01f4b561a1d8f9db4f8c8d051cecc503a76dbd290ad909df71afb7@104.207.133.78:30303", // 2348
	"enode://9cf8cde650bff06c608e29b9813175774debcafb8427801a1dbcc73884295c64f78b5d5177c501f22b7e9454ad4f5f641f331c77f8c314c799318a68f14d9cb3@45.76.24.53:30303",    // 2349
	"enode://68c76b85b4c5790106b2f80e76bccc42bba5b430d83ad8fd6a5b5c283abfa04cc3b5a76ada72cf3413eb21a52c372a724f5d93b41b1cec69747020a0b28b59ae@108.61.204.221:30303", // 2350
	"enode://84cdefe23273ba841dc0be357e7c4a979357c3a2fdd8c214e277272a41bb4c8d777a3c308217d610ec28aedeaf79accebe1594fda51d65f34c6518ac793b33bb@51.38.132.205:30303",  // 2355
	"enode://8b65cac4d74ada1233b7a7439e7f69d29c7a2be0a97c14e03be17d231014dc4dcaec1d014eb408a9b0442a89130a1f86e819552d48ef4a944457f85c6120febc@159.65.6.121:30303",   // 2363
	"enode://fb3671c7539dac486379e3cf5f1a0f0ca2210471313f5489eb05c2c0cf21ab4b44a941a0ae474d4a07f7c994cf1da7ac7e4fe68879140e7e2f520a9b483a1a8b@168.235.85.209:30303", // 2577
	"enode://2931234cd3c561652e85b991e2f2458e42e7c9cb085093e9e422188ec226376e9af241d8b77a7fd035d7837312b951bc9ca697f13c886d2bcb00f33ef5329635@139.180.196.7:30303",  // 2684
	"enode://8f196d099309b93d76e769d88801198d5d0cda628c17e1dcc818ce43f9ea26ea8658aa00ec9270341b6bf6553d01edd9543c8d264cfedc98c2be642585da3e51@85.255.11.158:30303",  // 2691
	"enode://fb8b5b629f60d1edd821660f9904e9b8285addcd5dc32b9232bc4d81bbe4ad039855f010e2b36bc4217ae1a1223153ba3824a74906e7ea320a3349250dca3b9e@54.38.215.141:30303",  // 2718
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
