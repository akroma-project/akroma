// Copyright 2016 The go-ethereum Authors
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

import (
	"fmt"
	"math/big"

	"github.com/akroma-project/akroma/common"
)

// Genesis hashes to enforce below configs on.
var (
	MainnetGenesisHash = common.HexToHash("0x679ee3d5213ddab6aad2c53c2b5a7a1021d113f868b93929893a42c93ae61efd")
	TestnetGenesisHash = common.HexToHash("0xc0aa9949ba05d4e30bff37bcdc4e5d9523a55a0fb34f7ad6abab1a4981ab5971")
	RinkebyGenesisHash = common.HexToHash("0x6341fd3daf94b748c72ced5a5b26028f2474f5f00d824504e4fa37a75767e177")
)

var (
	// MainnetChainConfig is the chain parameters to run a node on the main network.
	MainnetChainConfig = &ChainConfig{
		ChainID:                 big.NewInt(200625),
		HomesteadBlock:          big.NewInt(0),
		DAOForkBlock:            nil,
		DAOForkSupport:          false,
		EIP150Block:             big.NewInt(0),
		EIP150Hash:              common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
		EIP155Block:             big.NewInt(1),
		EIP158Block:             big.NewInt(1),
		ByzantiumBlock:          big.NewInt(0),
		ConstantinopleBlock:     nil,
		PetersburgBlock:         nil,
		AkromaBlock:             big.NewInt(300000),
		BaneslayerBlock:         big.NewInt(1200000),
		CopperLeafBlock:         big.NewInt(2200000),
		DawnbreakBlock:          big.NewInt(3200000),
		ExaltedAngelBlock:       big.NewInt(6500000),
		FlamebladeBlock:         big.NewInt(7500000),
		GabrielAngelfireBlock:   big.NewInt(8500000),
		HailstormValkyrieBlock:  big.NewInt(9500000),
		IonaBlock:               big.NewInt(10500000),
		JenaraBlock:             big.NewInt(11500000),
		KarmicGuideBlock:        big.NewInt(12500000),
		LinvalaBlock:            big.NewInt(13500000),
		MaelstromArchangelBlock: big.NewInt(14500000),
		PlatinumAngelBlock:      big.NewInt(15500000),
		RestorationAngelBlock:   big.NewInt(16500000),
		SerraAngelBlock:         big.NewInt(17500000),
		TwilightShepherdBlock:   big.NewInt(18500000),
		ValkyrieHarbingerBlock:  big.NewInt(19500000),
		WarriorAngelBlock:       big.NewInt(20500000),
		XathridDemonBlock:       big.NewInt(21500000),
		YouthfulValkyrieBlock:   big.NewInt(22500000),
		ZuranOrbBlock:           big.NewInt(23500000),
		Ethash:                  new(EthashConfig),
	}

	// MainnetTrustedCheckpoint contains the light client trusted checkpoint for the main network.
	MainnetTrustedCheckpoint = &TrustedCheckpoint{
		Name:         "mainnet",
		SectionIndex: 216,
		SectionHead:  common.HexToHash("0xae3e551c8d60d06fd411a8e6008e90625d3bb0cbbf664b65d5ed90b318553541"),
		CHTRoot:      common.HexToHash("0xeea7d2ab3545a37deecc66fc43c9556ae337c3ea1c6893e401428207bdb8e434"),
		BloomRoot:    common.HexToHash("0xb0d4176d160d67b99a9f963281e52bce0583a566b74b4497fe3ed24ae04004ff"),
	}

	// TestnetChainConfig contains the chain parameters to run a node on the Ropsten test network.
	TestnetChainConfig = &ChainConfig{
		ChainID:                 big.NewInt(200624),
		HomesteadBlock:          big.NewInt(0),
		DAOForkBlock:            nil,
		DAOForkSupport:          false,
		EIP150Block:             big.NewInt(0),
		EIP150Hash:              common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
		EIP155Block:             big.NewInt(1),
		EIP158Block:             big.NewInt(1),
		ByzantiumBlock:          big.NewInt(0),
		ConstantinopleBlock:     big.NewInt(5),
		PetersburgBlock:         big.NewInt(10),
		AkromaBlock:             big.NewInt(15),
		BaneslayerBlock:         big.NewInt(20),
		CopperLeafBlock:         big.NewInt(25),
		DawnbreakBlock:          big.NewInt(30),
		ExaltedAngelBlock:       big.NewInt(35),
		FlamebladeBlock:         big.NewInt(40),
		GabrielAngelfireBlock:   big.NewInt(45),
		HailstormValkyrieBlock:  big.NewInt(50),
		IonaBlock:               big.NewInt(55),
		JenaraBlock:             big.NewInt(60),
		KarmicGuideBlock:        big.NewInt(65),
		LinvalaBlock:            big.NewInt(70),
		MaelstromArchangelBlock: big.NewInt(75),
		PlatinumAngelBlock:      big.NewInt(80),
		RestorationAngelBlock:   big.NewInt(85),
		SerraAngelBlock:         big.NewInt(90),
		TwilightShepherdBlock:   big.NewInt(95),
		ValkyrieHarbingerBlock:  big.NewInt(100),
		WarriorAngelBlock:       big.NewInt(105),
		XathridDemonBlock:       big.NewInt(110),
		YouthfulValkyrieBlock:   big.NewInt(115),
		ZuranOrbBlock:           big.NewInt(120),
		Ethash:                  new(EthashConfig),
	}

	// TestnetTrustedCheckpoint contains the light client trusted checkpoint for the Ropsten test network.
	TestnetTrustedCheckpoint = &TrustedCheckpoint{
		Name:         "testnet",
		SectionIndex: 148,
		SectionHead:  common.HexToHash("0x4d3181bedb6aa96a6f3efa866c71f7802400d0fb4a6906946c453630d850efc0"),
		CHTRoot:      common.HexToHash("0x25df2f9d63a5f84b2852988f0f0f7af5a7877da061c11b85c812780b5a27a5ec"),
		BloomRoot:    common.HexToHash("0x0584834e5222471a06c669d210e302ca602780eaaddd04634fd65471c2a91419"),
	}

	// RinkebyChainConfig contains the chain parameters to run a node on the Rinkeby test network.
	RinkebyChainConfig = &ChainConfig{
		ChainID:             big.NewInt(4),
		HomesteadBlock:      big.NewInt(1),
		DAOForkBlock:        nil,
		DAOForkSupport:      true,
		EIP150Block:         big.NewInt(2),
		EIP150Hash:          common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
		EIP155Block:         big.NewInt(3),
		EIP158Block:         big.NewInt(3),
		ByzantiumBlock:      big.NewInt(1035301),
		ConstantinopleBlock: nil,
		PetersburgBlock:     nil, //TODO! Insert Rinkeby block number
		AkromaBlock:         nil,
		BaneslayerBlock:     nil,
		DawnbreakBlock:      nil,
		Clique: &CliqueConfig{
			Period: 15,
			Epoch:  30000,
		},
	}

	// RinkebyTrustedCheckpoint contains the light client trusted checkpoint for the Rinkeby test network.
	RinkebyTrustedCheckpoint = &TrustedCheckpoint{
		Name:         "rinkeby",
		SectionIndex: 113,
		SectionHead:  common.HexToHash("0xb812f3095af3af1cb2de7d7c2086ee807736a7315992c461b0986699185daf77"),
		CHTRoot:      common.HexToHash("0x5416d0924925eb835987ad3d1f059ecc66778c51959c8246a7a35b22ec5f3109"),
		BloomRoot:    common.HexToHash("0xcf74ca2c14e843b366561dab4fc64237bf6bb335119cbc97d723f3b501863470"),
	}

	// AllEthashProtocolChanges contains every protocol change (EIPs) introduced
	// and accepted by the Ethereum core developers into the Ethash consensus.
	//
	// This configuration is intentionally not using keyed fields to force anyone
	// adding flags to the config to also have to set these fields.
	AllEthashProtocolChanges = &ChainConfig{big.NewInt(1337), big.NewInt(0), nil, false, big.NewInt(0), common.Hash{}, big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), new(EthashConfig), nil}

	// AllCliqueProtocolChanges contains every protocol change (EIPs) introduced
	// and accepted by the Ethereum core developers into the Clique consensus.
	//
	// This configuration is intentionally not using keyed fields to force anyone
	// adding flags to the config to also have to set these fields.
	AllCliqueProtocolChanges = &ChainConfig{big.NewInt(1337), big.NewInt(0), nil, false, big.NewInt(0), common.Hash{}, big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), nil, &CliqueConfig{Period: 0, Epoch: 30000}}

	TestChainConfig = &ChainConfig{big.NewInt(1), big.NewInt(0), nil, false, big.NewInt(0), common.Hash{}, big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), new(EthashConfig), nil}
	TestRules       = TestChainConfig.Rules(new(big.Int))
)

// TrustedCheckpoint represents a set of post-processed trie roots (CHT and
// BloomTrie) associated with the appropriate section index and head hash. It is
// used to start light syncing from this checkpoint and avoid downloading the
// entire header chain while still being able to securely access old headers/logs.
type TrustedCheckpoint struct {
	Name         string      `json:"-"`
	SectionIndex uint64      `json:"sectionIndex"`
	SectionHead  common.Hash `json:"sectionHead"`
	CHTRoot      common.Hash `json:"chtRoot"`
	BloomRoot    common.Hash `json:"bloomRoot"`
}

// ChainConfig is the core config which determines the blockchain settings.
//
// ChainConfig is stored in the database on a per block basis. This means
// that any network, identified by its genesis block, can have its own
// set of configuration options.
type ChainConfig struct {
	ChainID *big.Int `json:"chainId"` // chainId identifies the current chain and is used for replay protection

	HomesteadBlock *big.Int `json:"homesteadBlock,omitempty"` // Homestead switch block (nil = no fork, 0 = already homestead)

	DAOForkBlock   *big.Int `json:"daoForkBlock,omitempty"`   // TheDAO hard-fork switch block (nil = no fork)
	DAOForkSupport bool     `json:"daoForkSupport,omitempty"` // Whether the nodes supports or opposes the DAO hard-fork

	// EIP150 implements the Gas price changes (https://github.com/ethereum/EIPs/issues/150)
	EIP150Block *big.Int    `json:"eip150Block,omitempty"` // EIP150 HF block (nil = no fork)
	EIP150Hash  common.Hash `json:"eip150Hash,omitempty"`  // EIP150 HF hash (needed for header only clients as only gas pricing changed)

	EIP155Block *big.Int `json:"eip155Block,omitempty"` // EIP155 HF block
	EIP158Block *big.Int `json:"eip158Block,omitempty"` // EIP158 HF block

	ByzantiumBlock          *big.Int `json:"byzantiumBlock,omitempty"`          // Byzantium switch block (nil = no fork, 0 = already on byzantium)
	ConstantinopleBlock     *big.Int `json:"constantinopleBlock,omitempty"`     // Constantinople switch block (nil = no fork, 0 = already activated)
	PetersburgBlock         *big.Int `json:"petersburgBlock,omitempty"`         // Petersburg switch block (nil = same as Constantinople)
	EWASMBlock              *big.Int `json:"ewasmBlock,omitempty"`              // EWASM switch block (nil = no fork, 0 = already activated)
	AkromaBlock             *big.Int `json:"akromaBlock,omitempty"`             // Akroma switch block (nil = no fork, 0 = already on akroma)
	BaneslayerBlock         *big.Int `json:"baneslayerBlock,omitempty"`         // second major akroma release
	CopperLeafBlock         *big.Int `json:"copperLeafBlock,omitempty"`         // third major akroma release
	DawnbreakBlock          *big.Int `json:"dawnbreakBlock,omitempty"`          // fourth major akroma release
	ExaltedAngelBlock       *big.Int `json:"exaltedAngelBlock,omitempty"`       // x major akroma release
	FlamebladeBlock         *big.Int `json:"flamebladeBlock,omitempty"`         // x major akroma release
	GabrielAngelfireBlock   *big.Int `json:"gabrielAngelfireBlock,omitempty"`   // x major akroma release
	HailstormValkyrieBlock  *big.Int `json:"hailstormValkyrieBlock,omitempty"`  // x major akroma release
	IonaBlock               *big.Int `json:"ionaBlock,omitempty"`               // x major akroma release
	JenaraBlock             *big.Int `json:"jenaraBlock,omitempty"`             // x major akroma release
	KarmicGuideBlock        *big.Int `json:"karmicGuideBlock,omitempty"`        // x major akroma release
	LinvalaBlock            *big.Int `json:"linvalaBlock,omitempty"`            // x major akroma release
	MaelstromArchangelBlock *big.Int `json:"maelstromArchangelBlock,omitempty"` // x major akroma release
	PlatinumAngelBlock      *big.Int `json:"platinumAngelBlock,omitempty"`      // x major akroma release
	RestorationAngelBlock   *big.Int `json:"restorationAngelBlock,omitempty"`   // x major akroma release
	SerraAngelBlock         *big.Int `json:"serraAngelBlock,omitempty"`         // x major akroma release
	TwilightShepherdBlock   *big.Int `json:"twilightShepherdBlock,omitempty"`   // x major akroma release
	ValkyrieHarbingerBlock  *big.Int `json:"valkyrieHarbingerBlock,omitempty"`  // x major akroma release
	WarriorAngelBlock       *big.Int `json:"warriorAngelBlock,omitempty"`       // x major akroma release
	XathridDemonBlock       *big.Int `json:"xathridDemonBlock,omitempty"`       // x major akroma release
	YouthfulValkyrieBlock   *big.Int `json:"youthfulValkyrieBlock,omitempty"`   // x major akroma release
	ZuranOrbBlock           *big.Int `json:"zuranOrbBlock,omitempty"`           // x major akroma release

	// Various consensus engines
	Ethash *EthashConfig `json:"ethash,omitempty"`
	Clique *CliqueConfig `json:"clique,omitempty"`
}

// EthashConfig is the consensus engine configs for proof-of-work based sealing.
type EthashConfig struct{}

// String implements the stringer interface, returning the consensus engine details.
func (c *EthashConfig) String() string {
	return "ethash"
}

// CliqueConfig is the consensus engine configs for proof-of-authority based sealing.
type CliqueConfig struct {
	Period uint64 `json:"period"` // Number of seconds between blocks to enforce
	Epoch  uint64 `json:"epoch"`  // Epoch length to reset votes and checkpoint
}

// String implements the stringer interface, returning the consensus engine details.
func (c *CliqueConfig) String() string {
	return "clique"
}

// String implements the fmt.Stringer interface.
func (c *ChainConfig) String() string {
	var engine interface{}
	switch {
	case c.Ethash != nil:
		engine = c.Ethash
	case c.Clique != nil:
		engine = c.Clique
	default:
		engine = "unknown"
	}
	return fmt.Sprintf("{ChainID: %v Homestead: %v DAO: %v DAOSupport: %v EIP150: %v EIP155: %v EIP158: %v Byzantium: %v Constantinople: %v  ConstantinopleFix: %v Akroma: %v Baneslayer: %v CopperLeaf: %v Dawnbreak: %v  ExaltedAngel: %v Flameblade: %v GabrielAngelfire: %v HailstormValkyrie: %v Iona: %v Jenara: %v KarmicGuide: %v Linvala: %v MaelstromArchangel: %v  PlatinumAngel: %v  RestorationAngel: %v SerraAngel: %v TwilightShepherd: %v WarriorAngel: %v ValkyrieHarbinger: %v XathridDemon: %v YouthfulValkyrie: %v ZuranOrb: %v Engine: %v}",
		c.ChainID,
		c.HomesteadBlock,
		c.DAOForkBlock,
		c.DAOForkSupport,
		c.EIP150Block,
		c.EIP155Block,
		c.EIP158Block,
		c.ByzantiumBlock,
		c.ConstantinopleBlock,
		c.PetersburgBlock,
		c.AkromaBlock,
		c.BaneslayerBlock,
		c.CopperLeafBlock,
		c.DawnbreakBlock,
		c.ExaltedAngelBlock,
		c.FlamebladeBlock,
		c.GabrielAngelfireBlock,
		c.HailstormValkyrieBlock,
		c.IonaBlock,
		c.JenaraBlock,
		c.KarmicGuideBlock,
		c.LinvalaBlock,
		c.MaelstromArchangelBlock,
		c.PlatinumAngelBlock,
		c.RestorationAngelBlock,
		c.SerraAngelBlock,
		c.TwilightShepherdBlock,
		c.ValkyrieHarbingerBlock,
		c.WarriorAngelBlock,
		c.XathridDemonBlock,
		c.YouthfulValkyrieBlock,
		c.ZuranOrbBlock,
		engine,
	)
}

// IsHomestead returns whether num is either equal to the homestead block or greater.
func (c *ChainConfig) IsHomestead(num *big.Int) bool {
	return isForked(c.HomesteadBlock, num)
}

// IsDAOFork returns whether num is either equal to the DAO fork block or greater.
func (c *ChainConfig) IsDAOFork(num *big.Int) bool {
	return isForked(c.DAOForkBlock, num)
}

// IsEIP150 returns whether num is either equal to the EIP150 fork block or greater.
func (c *ChainConfig) IsEIP150(num *big.Int) bool {
	return isForked(c.EIP150Block, num)
}

// IsEIP155 returns whether num is either equal to the EIP155 fork block or greater.
func (c *ChainConfig) IsEIP155(num *big.Int) bool {
	return isForked(c.EIP155Block, num)
}

// IsEIP158 returns whether num is either equal to the EIP158 fork block or greater.
func (c *ChainConfig) IsEIP158(num *big.Int) bool {
	return isForked(c.EIP158Block, num)
}

// IsByzantium returns whether num is either equal to the Byzantium fork block or greater.
func (c *ChainConfig) IsByzantium(num *big.Int) bool {
	return isForked(c.ByzantiumBlock, num)
}

// IsConstantinople returns whether num is either equal to the Constantinople fork block or greater.
func (c *ChainConfig) IsConstantinople(num *big.Int) bool {
	return isForked(c.ConstantinopleBlock, num)
}

// IsPetersburg returns whether num is either
// - equal to or greater than the PetersburgBlock fork block,
// - OR is nil, and Constantinople is active
func (c *ChainConfig) IsPetersburg(num *big.Int) bool {
	return isForked(c.PetersburgBlock, num) || c.PetersburgBlock == nil && isForked(c.ConstantinopleBlock, num)
}

// IsEWASM returns whether num represents a block number after the EWASM fork
func (c *ChainConfig) IsEWASM(num *big.Int) bool {
	return isForked(c.EWASMBlock, num)
}

// IsAkroma returns whether num is either equal to the Akroma fork block or greater.
func (c *ChainConfig) IsAkroma(num *big.Int) bool {
	return isForked(c.AkromaBlock, num)
}

// IsBaneslayer returns whether num is either equal to the Baneslayer fork block or greater.
func (c *ChainConfig) IsBaneslayer(num *big.Int) bool {
	return isForked(c.BaneslayerBlock, num)
}

// IsCopperLeaf returns whether num is either equal to the CopperLeaf fork block or greater.
func (c *ChainConfig) IsCopperLeaf(num *big.Int) bool {
	return isForked(c.CopperLeafBlock, num)
}

// IsDawnbreak returns whether num is either equal to the Dawnbreak fork block or greater.
func (c *ChainConfig) IsDawnbreak(num *big.Int) bool {
	return isForked(c.DawnbreakBlock, num)
}

// IsExaltedAngel returns whether num is either equal to the ExaltedAngel fork block or greater.
func (c *ChainConfig) IsExaltedAngel(num *big.Int) bool {
	return isForked(c.DawnbreakBlock, num)
}

// IsFlameblade returns whether num is either equal to the Flameblade fork block or greater.
func (c *ChainConfig) IsFlameblade(num *big.Int) bool {
	return isForked(c.DawnbreakBlock, num)
}

// IsGabrielAngelfire returns whether num is either equal to the GabrielAngelfire fork block or greater.
func (c *ChainConfig) IsGabrielAngelfire(num *big.Int) bool {
	return isForked(c.DawnbreakBlock, num)
}

// IsHailstormValkyrie returns whether num is either equal to the HailstormValkyrie fork block or greater.
func (c *ChainConfig) IsHailstormValkyrie(num *big.Int) bool {
	return isForked(c.DawnbreakBlock, num)
}

// IsIona returns whether num is either equal to the Iona fork block or greater.
func (c *ChainConfig) IsIona(num *big.Int) bool {
	return isForked(c.DawnbreakBlock, num)
}

// IsJenara returns whether num is either equal to the Jenara fork block or greater.
func (c *ChainConfig) IsJenara(num *big.Int) bool {
	return isForked(c.DawnbreakBlock, num)
}

// IsKarmicGuide returns whether num is either equal to the KarmicGuide fork block or greater.
func (c *ChainConfig) IsKarmicGuide(num *big.Int) bool {
	return isForked(c.DawnbreakBlock, num)
}

// IsLinvala returns whether num is either equal to the Linvala fork block or greater.
func (c *ChainConfig) IsLinvala(num *big.Int) bool {
	return isForked(c.DawnbreakBlock, num)
}

// IsMaelstromArchangel returns whether num is either equal to the MaelstromArchangel fork block or greater.
func (c *ChainConfig) IsMaelstromArchangel(num *big.Int) bool {
	return isForked(c.DawnbreakBlock, num)
}

// IsPlatinumAngel returns whether num is either equal to the PlatinumAngel fork block or greater.
func (c *ChainConfig) IsPlatinumAngel(num *big.Int) bool {
	return isForked(c.DawnbreakBlock, num)
}

// IsRestorationAngel returns whether num is either equal to the RestorationAngel fork block or greater.
func (c *ChainConfig) IsRestorationAngel(num *big.Int) bool {
	return isForked(c.DawnbreakBlock, num)
}

// IsSerraAngel returns whether num is either equal to the SerraAngel fork block or greater.
func (c *ChainConfig) IsSerraAngel(num *big.Int) bool {
	return isForked(c.DawnbreakBlock, num)
}

// IsTwilightShepherd returns whether num is either equal to the TwilightShepherd fork block or greater.
func (c *ChainConfig) IsTwilightShepherd(num *big.Int) bool {
	return isForked(c.DawnbreakBlock, num)
}

// IsValkyrieHarbinger returns whether num is either equal to the ValkyrieHarbinger fork block or greater.
func (c *ChainConfig) IsValkyrieHarbinger(num *big.Int) bool {
	return isForked(c.DawnbreakBlock, num)
}

// IsWarriorAngel returns whether num is either equal to the WarriorAngel fork block or greater.
func (c *ChainConfig) IsWarriorAngel(num *big.Int) bool {
	return isForked(c.DawnbreakBlock, num)
}

// IsXathridDemon returns whether num is either equal to the XathridDemon fork block or greater.
func (c *ChainConfig) IsXathridDemon(num *big.Int) bool {
	return isForked(c.DawnbreakBlock, num)
}

// IsYouthfulValkyrie returns whether num is either equal to the YouthfulValkyrie fork block or greater.
func (c *ChainConfig) IsYouthfulValkyrie(num *big.Int) bool {
	return isForked(c.DawnbreakBlock, num)
}

// IsZuranOrb returns whether num is either equal to the ZuranOrb fork block or greater.
func (c *ChainConfig) IsZuranOrb(num *big.Int) bool {
	return isForked(c.DawnbreakBlock, num)
}

// GasTable returns the gas table corresponding to the current phase (homestead or homestead reprice).
//
// The returned GasTable's fields shouldn't, under any circumstances, be changed.
func (c *ChainConfig) GasTable(num *big.Int) GasTable {
	if num == nil {
		return GasTableHomestead
	}
	switch {
	case c.IsConstantinople(num):
		return GasTableConstantinople
	case c.IsEIP158(num):
		return GasTableEIP158
	case c.IsEIP150(num):
		return GasTableEIP150
	default:
		return GasTableHomestead
	}
}

// CheckCompatible checks whether scheduled fork transitions have been imported
// with a mismatching chain configuration.
func (c *ChainConfig) CheckCompatible(newcfg *ChainConfig, height uint64) *ConfigCompatError {
	bhead := new(big.Int).SetUint64(height)

	// Iterate checkCompatible to find the lowest conflict.
	var lasterr *ConfigCompatError
	for {
		err := c.checkCompatible(newcfg, bhead)
		if err == nil || (lasterr != nil && err.RewindTo == lasterr.RewindTo) {
			break
		}
		lasterr = err
		bhead.SetUint64(err.RewindTo)
	}
	return lasterr
}

func (c *ChainConfig) checkCompatible(newcfg *ChainConfig, head *big.Int) *ConfigCompatError {
	if isForkIncompatible(c.HomesteadBlock, newcfg.HomesteadBlock, head) {
		return newCompatError("Homestead fork block", c.HomesteadBlock, newcfg.HomesteadBlock)
	}
	if isForkIncompatible(c.DAOForkBlock, newcfg.DAOForkBlock, head) {
		return newCompatError("DAO fork block", c.DAOForkBlock, newcfg.DAOForkBlock)
	}
	if c.IsDAOFork(head) && c.DAOForkSupport != newcfg.DAOForkSupport {
		return newCompatError("DAO fork support flag", c.DAOForkBlock, newcfg.DAOForkBlock)
	}
	if isForkIncompatible(c.EIP150Block, newcfg.EIP150Block, head) {
		return newCompatError("EIP150 fork block", c.EIP150Block, newcfg.EIP150Block)
	}
	if isForkIncompatible(c.EIP155Block, newcfg.EIP155Block, head) {
		return newCompatError("EIP155 fork block", c.EIP155Block, newcfg.EIP155Block)
	}
	if isForkIncompatible(c.EIP158Block, newcfg.EIP158Block, head) {
		return newCompatError("EIP158 fork block", c.EIP158Block, newcfg.EIP158Block)
	}
	if c.IsEIP158(head) && !configNumEqual(c.ChainID, newcfg.ChainID) {
		return newCompatError("EIP158 chain ID", c.EIP158Block, newcfg.EIP158Block)
	}
	if isForkIncompatible(c.ByzantiumBlock, newcfg.ByzantiumBlock, head) {
		return newCompatError("Byzantium fork block", c.ByzantiumBlock, newcfg.ByzantiumBlock)
	}
	if isForkIncompatible(c.ConstantinopleBlock, newcfg.ConstantinopleBlock, head) {
		return newCompatError("Constantinople fork block", c.ConstantinopleBlock, newcfg.ConstantinopleBlock)
	}
	if isForkIncompatible(c.PetersburgBlock, newcfg.PetersburgBlock, head) {
		return newCompatError("ConstantinopleFix fork block", c.PetersburgBlock, newcfg.PetersburgBlock)
	}
	if isForkIncompatible(c.EWASMBlock, newcfg.EWASMBlock, head) {
		return newCompatError("ewasm fork block", c.EWASMBlock, newcfg.EWASMBlock)
	}
	if isForkIncompatible(c.AkromaBlock, newcfg.AkromaBlock, head) {
		return newCompatError("Akroma fork block", c.AkromaBlock, newcfg.AkromaBlock)
	}
	if isForkIncompatible(c.BaneslayerBlock, newcfg.BaneslayerBlock, head) {
		return newCompatError("Baneslayer fork block", c.BaneslayerBlock, newcfg.BaneslayerBlock)
	}
	if isForkIncompatible(c.CopperLeafBlock, newcfg.CopperLeafBlock, head) {
		return newCompatError("CopperLeaf fork block", c.CopperLeafBlock, newcfg.CopperLeafBlock)
	}
	if isForkIncompatible(c.DawnbreakBlock, newcfg.DawnbreakBlock, head) {
		return newCompatError("Dawnbreak fork block", c.DawnbreakBlock, newcfg.DawnbreakBlock)
	}
	if isForkIncompatible(c.ExaltedAngelBlock, newcfg.ExaltedAngelBlock, head) {
		return newCompatError("ExaltedAngel fork block", c.ExaltedAngelBlock, newcfg.ExaltedAngelBlock)
	}
	if isForkIncompatible(c.FlamebladeBlock, newcfg.FlamebladeBlock, head) {
		return newCompatError("Flameblade fork block", c.FlamebladeBlock, newcfg.FlamebladeBlock)
	}
	if isForkIncompatible(c.GabrielAngelfireBlock, newcfg.GabrielAngelfireBlock, head) {
		return newCompatError("GabrielAngelfire fork block", c.GabrielAngelfireBlock, newcfg.GabrielAngelfireBlock)
	}
	if isForkIncompatible(c.HailstormValkyrieBlock, newcfg.HailstormValkyrieBlock, head) {
		return newCompatError("HailstormValkyrie fork block", c.HailstormValkyrieBlock, newcfg.HailstormValkyrieBlock)
	}
	if isForkIncompatible(c.IonaBlock, newcfg.IonaBlock, head) {
		return newCompatError("Iona fork block", c.IonaBlock, newcfg.IonaBlock)
	}
	if isForkIncompatible(c.JenaraBlock, newcfg.JenaraBlock, head) {
		return newCompatError("Jenara fork block", c.JenaraBlock, newcfg.JenaraBlock)
	}
	if isForkIncompatible(c.KarmicGuideBlock, newcfg.KarmicGuideBlock, head) {
		return newCompatError("KarmicGuide fork block", c.KarmicGuideBlock, newcfg.KarmicGuideBlock)
	}
	if isForkIncompatible(c.LinvalaBlock, newcfg.LinvalaBlock, head) {
		return newCompatError("Dawnbreak fork block", c.LinvalaBlock, newcfg.LinvalaBlock)
	}
	if isForkIncompatible(c.MaelstromArchangelBlock, newcfg.MaelstromArchangelBlock, head) {
		return newCompatError("MaelstromArchangel fork block", c.MaelstromArchangelBlock, newcfg.MaelstromArchangelBlock)
	}
	if isForkIncompatible(c.PlatinumAngelBlock, newcfg.PlatinumAngelBlock, head) {
		return newCompatError("PlatinumAngel fork block", c.PlatinumAngelBlock, newcfg.PlatinumAngelBlock)
	}
	if isForkIncompatible(c.RestorationAngelBlock, newcfg.RestorationAngelBlock, head) {
		return newCompatError("RestorationAngel fork block", c.RestorationAngelBlock, newcfg.RestorationAngelBlock)
	}
	if isForkIncompatible(c.SerraAngelBlock, newcfg.SerraAngelBlock, head) {
		return newCompatError("SerraAngel fork block", c.SerraAngelBlock, newcfg.SerraAngelBlock)
	}
	if isForkIncompatible(c.TwilightShepherdBlock, newcfg.TwilightShepherdBlock, head) {
		return newCompatError("TwilightShepherd fork block", c.TwilightShepherdBlock, newcfg.TwilightShepherdBlock)
	}
	if isForkIncompatible(c.ValkyrieHarbingerBlock, newcfg.ValkyrieHarbingerBlock, head) {
		return newCompatError("ValkyrieHarbinger fork block", c.ValkyrieHarbingerBlock, newcfg.ValkyrieHarbingerBlock)
	}
	if isForkIncompatible(c.WarriorAngelBlock, newcfg.WarriorAngelBlock, head) {
		return newCompatError("WarriorAngel fork block", c.WarriorAngelBlock, newcfg.WarriorAngelBlock)
	}
	if isForkIncompatible(c.XathridDemonBlock, newcfg.XathridDemonBlock, head) {
		return newCompatError("XathridDemon fork block", c.XathridDemonBlock, newcfg.XathridDemonBlock)
	}
	if isForkIncompatible(c.YouthfulValkyrieBlock, newcfg.YouthfulValkyrieBlock, head) {
		return newCompatError("YouthfulValkyrie fork block", c.YouthfulValkyrieBlock, newcfg.YouthfulValkyrieBlock)
	}
	if isForkIncompatible(c.ZuranOrbBlock, newcfg.ZuranOrbBlock, head) {
		return newCompatError("ZuranOrb fork block", c.ZuranOrbBlock, newcfg.ZuranOrbBlock)
	}
	return nil
}

// isForkIncompatible returns true if a fork scheduled at s1 cannot be rescheduled to
// block s2 because head is already past the fork.
func isForkIncompatible(s1, s2, head *big.Int) bool {
	return (isForked(s1, head) || isForked(s2, head)) && !configNumEqual(s1, s2)
}

// isForked returns whether a fork scheduled at block s is active at the given head block.
func isForked(s, head *big.Int) bool {
	if s == nil || head == nil {
		return false
	}
	return s.Cmp(head) <= 0
}

func configNumEqual(x, y *big.Int) bool {
	if x == nil {
		return y == nil
	}
	if y == nil {
		return x == nil
	}
	return x.Cmp(y) == 0
}

// ConfigCompatError is raised if the locally-stored blockchain is initialised with a
// ChainConfig that would alter the past.
type ConfigCompatError struct {
	What string
	// block numbers of the stored and new configurations
	StoredConfig, NewConfig *big.Int
	// the block number to which the local chain must be rewound to correct the error
	RewindTo uint64
}

func newCompatError(what string, storedblock, newblock *big.Int) *ConfigCompatError {
	var rew *big.Int
	switch {
	case storedblock == nil:
		rew = newblock
	case newblock == nil || storedblock.Cmp(newblock) < 0:
		rew = storedblock
	default:
		rew = newblock
	}
	err := &ConfigCompatError{what, storedblock, newblock, 0}
	if rew != nil && rew.Sign() > 0 {
		err.RewindTo = rew.Uint64() - 1
	}
	return err
}

func (err *ConfigCompatError) Error() string {
	return fmt.Sprintf("mismatching %s in database (have %d, want %d, rewindto %d)", err.What, err.StoredConfig, err.NewConfig, err.RewindTo)
}

// Rules wraps ChainConfig and is merely syntactic sugar or can be used for functions
// that do not have or require information about the block.
//
// Rules is a one time interface meaning that it shouldn't be used in between transition
// phases.
type Rules struct {
	ChainID                                     *big.Int
	IsHomestead, IsEIP150, IsEIP155, IsEIP158   bool
	IsByzantium, IsConstantinople, IsPetersburg bool
	IsAkroma                                    bool
	IsBaneslayer                                bool
	IsCopperLeaf                                bool
	IsDawnbreak                                 bool
	IsExaltedAngel                              bool
	IsFlameblade                                bool
	IsGabrielAngelfire                          bool
	IsHailstormValkyrie                         bool
	IsIona                                      bool
	IsJenara                                    bool
	IsKarmicGuide                               bool
	IsLinvala                                   bool
	IsMaelstromArchangel                        bool
	IsPlatinumAngel                             bool
	IsRestorationAngel                          bool
	IsSerraAngel                                bool
	IsTwilightShepherd                          bool
	IsValkyrieHarbinger                         bool
	IsWarriorAngel                              bool
	IsXathridDemon                              bool
	IsYouthfulValkyrie                          bool
	IsZuranOrb                                  bool
}

// Rules ensures c's ChainID is not nil.
func (c *ChainConfig) Rules(num *big.Int) Rules {
	chainID := c.ChainID
	if chainID == nil {
		chainID = new(big.Int)
	}
	return Rules{
		ChainID:              new(big.Int).Set(chainID),
		IsHomestead:          c.IsHomestead(num),
		IsEIP150:             c.IsEIP150(num),
		IsEIP155:             c.IsEIP155(num),
		IsEIP158:             c.IsEIP158(num),
		IsByzantium:          c.IsByzantium(num),
		IsConstantinople:     c.IsConstantinople(num),
		IsPetersburg:         c.IsPetersburg(num),
		IsAkroma:             c.IsAkroma(num),
		IsBaneslayer:         c.IsBaneslayer(num),
		IsCopperLeaf:         c.IsCopperLeaf(num),
		IsDawnbreak:          c.IsDawnbreak(num),
		IsExaltedAngel:       c.IsExaltedAngel(num),
		IsFlameblade:         c.IsFlameblade(num),
		IsGabrielAngelfire:   c.IsGabrielAngelfire(num),
		IsHailstormValkyrie:  c.IsHailstormValkyrie(num),
		IsIona:               c.IsIona(num),
		IsJenara:             c.IsJenara(num),
		IsKarmicGuide:        c.IsKarmicGuide(num),
		IsLinvala:            c.IsLinvala(num),
		IsMaelstromArchangel: c.IsMaelstromArchangel(num),
		IsPlatinumAngel:      c.IsPlatinumAngel(num),
		IsRestorationAngel:   c.IsRestorationAngel(num),
		IsSerraAngel:         c.IsSerraAngel(num),
		IsTwilightShepherd:   c.IsTwilightShepherd(num),
		IsValkyrieHarbinger:  c.IsValkyrieHarbinger(num),
		IsWarriorAngel:       c.IsWarriorAngel(num),
		IsXathridDemon:       c.IsXathridDemon(num),
		IsYouthfulValkyrie:   c.IsYouthfulValkyrie(num),
		IsZuranOrb:           c.IsZuranOrb(num),
	}
}
