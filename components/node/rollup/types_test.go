package rollup

import (
	"context"
	"encoding/json"
	"math/big"
	"math/rand"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"

	"github.com/wemixkanvas/kanvas/components/node/eth"
)

func randConfig() *Config {
	rng := rand.New(rand.NewSource(1234))
	randHash := func() (out [32]byte) {
		rng.Read(out[:])
		return
	}
	randAddr := func() (out common.Address) { // we need generics...
		rng.Read(out[:])
		return
	}
	return &Config{
		Genesis: Genesis{
			L1:     eth.BlockID{Hash: randHash(), Number: 424242},
			L2:     eth.BlockID{Hash: randHash(), Number: 1337},
			L2Time: uint64(time.Now().Unix()),
			SystemConfig: eth.SystemConfig{
				BatcherAddr: randAddr(),
				Overhead:    randHash(),
				Scalar:      randHash(),
				GasLimit:    1234567,
			},
		},
		BlockTime:              2,
		MaxProposerDrift:       100,
		ProposerWindowSize:     2,
		ChannelTimeout:         123,
		L1ChainID:              big.NewInt(900),
		L2ChainID:              big.NewInt(901),
		BatchInboxAddress:      randAddr(),
		DepositContractAddress: randAddr(),
		L1SystemConfigAddress:  randAddr(),
	}
}

func TestConfigJSON(t *testing.T) {
	config := randConfig()
	data, err := json.Marshal(config)
	assert.NoError(t, err)
	var roundTripped Config
	assert.NoError(t, json.Unmarshal(data, &roundTripped))
	assert.Equal(t, &roundTripped, config)
}

type mockL1Client struct {
	chainID *big.Int
	Hash    common.Hash
}

func (m *mockL1Client) ChainID(context.Context) (*big.Int, error) {
	return m.chainID, nil
}

func (m *mockL1Client) L1BlockRefByNumber(ctx context.Context, number uint64) (eth.L1BlockRef, error) {
	return eth.L1BlockRef{
		Hash:   m.Hash,
		Number: 100,
	}, nil
}

func TestValidateL1Config(t *testing.T) {
	config := randConfig()
	config.L1ChainID = big.NewInt(100)
	config.Genesis.L1.Number = 100
	config.Genesis.L1.Hash = [32]byte{0x01}
	mockClient := mockL1Client{chainID: big.NewInt(100), Hash: common.Hash{0x01}}
	err := config.ValidateL1Config(context.TODO(), &mockClient)
	assert.NoError(t, err)
}

func TestValidateL1ConfigInvalidChainIdFails(t *testing.T) {
	config := randConfig()
	config.L1ChainID = big.NewInt(101)
	config.Genesis.L1.Number = 100
	config.Genesis.L1.Hash = [32]byte{0x01}
	mockClient := mockL1Client{chainID: big.NewInt(100), Hash: common.Hash{0x01}}
	err := config.ValidateL1Config(context.TODO(), &mockClient)
	assert.Error(t, err)
	config.L1ChainID = big.NewInt(99)
	err = config.ValidateL1Config(context.TODO(), &mockClient)
	assert.Error(t, err)
}

func TestValidateL1ConfigInvalidGenesisHashFails(t *testing.T) {
	config := randConfig()
	config.L1ChainID = big.NewInt(100)
	config.Genesis.L1.Number = 100
	config.Genesis.L1.Hash = [32]byte{0x00}
	mockClient := mockL1Client{chainID: big.NewInt(100), Hash: common.Hash{0x01}}
	err := config.ValidateL1Config(context.TODO(), &mockClient)
	assert.Error(t, err)
	config.Genesis.L1.Hash = [32]byte{0x02}
	err = config.ValidateL1Config(context.TODO(), &mockClient)
	assert.Error(t, err)
}

func TestCheckL1ChainID(t *testing.T) {
	config := randConfig()
	config.L1ChainID = big.NewInt(100)
	err := config.CheckL1ChainID(context.TODO(), &mockL1Client{chainID: big.NewInt(100)})
	assert.NoError(t, err)
	err = config.CheckL1ChainID(context.TODO(), &mockL1Client{chainID: big.NewInt(101)})
	assert.Error(t, err)
	err = config.CheckL1ChainID(context.TODO(), &mockL1Client{chainID: big.NewInt(99)})
	assert.Error(t, err)
}

func TestCheckL1BlockRefByNumber(t *testing.T) {
	config := randConfig()
	config.Genesis.L1.Number = 100
	config.Genesis.L1.Hash = [32]byte{0x01}
	mockClient := mockL1Client{chainID: big.NewInt(100), Hash: common.Hash{0x01}}
	err := config.CheckL1GenesisBlockHash(context.TODO(), &mockClient)
	assert.NoError(t, err)
	mockClient.Hash = common.Hash{0x02}
	err = config.CheckL1GenesisBlockHash(context.TODO(), &mockClient)
	assert.Error(t, err)
	mockClient.Hash = common.Hash{0x00}
	err = config.CheckL1GenesisBlockHash(context.TODO(), &mockClient)
	assert.Error(t, err)
}

type mockL2Client struct {
	chainID *big.Int
	Hash    common.Hash
}

func (m *mockL2Client) ChainID(context.Context) (*big.Int, error) {
	return m.chainID, nil
}

func (m *mockL2Client) L2BlockRefByNumber(ctx context.Context, number uint64) (eth.L2BlockRef, error) {
	return eth.L2BlockRef{
		Hash:   m.Hash,
		Number: 100,
	}, nil
}

func TestValidateL2Config(t *testing.T) {
	config := randConfig()
	config.L2ChainID = big.NewInt(100)
	config.Genesis.L2.Number = 100
	config.Genesis.L2.Hash = [32]byte{0x01}
	mockClient := mockL2Client{chainID: big.NewInt(100), Hash: common.Hash{0x01}}
	err := config.ValidateL2Config(context.TODO(), &mockClient)
	assert.NoError(t, err)
}

func TestValidateL2ConfigInvalidChainIdFails(t *testing.T) {
	config := randConfig()
	config.L2ChainID = big.NewInt(101)
	config.Genesis.L2.Number = 100
	config.Genesis.L2.Hash = [32]byte{0x01}
	mockClient := mockL2Client{chainID: big.NewInt(100), Hash: common.Hash{0x01}}
	err := config.ValidateL2Config(context.TODO(), &mockClient)
	assert.Error(t, err)
	config.L2ChainID = big.NewInt(99)
	err = config.ValidateL2Config(context.TODO(), &mockClient)
	assert.Error(t, err)
}

func TestValidateL2ConfigInvalidGenesisHashFails(t *testing.T) {
	config := randConfig()
	config.L2ChainID = big.NewInt(100)
	config.Genesis.L2.Number = 100
	config.Genesis.L2.Hash = [32]byte{0x00}
	mockClient := mockL2Client{chainID: big.NewInt(100), Hash: common.Hash{0x01}}
	err := config.ValidateL2Config(context.TODO(), &mockClient)
	assert.Error(t, err)
	config.Genesis.L2.Hash = [32]byte{0x02}
	err = config.ValidateL2Config(context.TODO(), &mockClient)
	assert.Error(t, err)
}

func TestCheckL2ChainID(t *testing.T) {
	config := randConfig()
	config.L2ChainID = big.NewInt(100)
	err := config.CheckL2ChainID(context.TODO(), &mockL2Client{chainID: big.NewInt(100)})
	assert.NoError(t, err)
	err = config.CheckL2ChainID(context.TODO(), &mockL2Client{chainID: big.NewInt(101)})
	assert.Error(t, err)
	err = config.CheckL2ChainID(context.TODO(), &mockL2Client{chainID: big.NewInt(99)})
	assert.Error(t, err)
}

func TestCheckL2BlockRefByNumber(t *testing.T) {
	config := randConfig()
	config.Genesis.L2.Number = 100
	config.Genesis.L2.Hash = [32]byte{0x01}
	mockClient := mockL2Client{chainID: big.NewInt(100), Hash: common.Hash{0x01}}
	err := config.CheckL2GenesisBlockHash(context.TODO(), &mockClient)
	assert.NoError(t, err)
	mockClient.Hash = common.Hash{0x02}
	err = config.CheckL2GenesisBlockHash(context.TODO(), &mockClient)
	assert.Error(t, err)
	mockClient.Hash = common.Hash{0x00}
	err = config.CheckL2GenesisBlockHash(context.TODO(), &mockClient)
	assert.Error(t, err)
}

func TestConfig_Check(t *testing.T) {
	tests := []struct {
		name        string
		modifier    func(cfg *Config)
		expectedErr error
	}{
		{
			name:        "BlockTimeZero",
			modifier:    func(cfg *Config) { cfg.BlockTime = 0 },
			expectedErr: ErrBlockTimeZero,
		},
		{
			name:        "ChannelTimeoutZero",
			modifier:    func(cfg *Config) { cfg.ChannelTimeout = 0 },
			expectedErr: ErrMissingChannelTimeout,
		},
		{
			name:        "ProposerWindowSizeZero",
			modifier:    func(cfg *Config) { cfg.ProposerWindowSize = 0 },
			expectedErr: ErrInvalidProposerWindowSize,
		},
		{
			name:        "ProposerWindowSizeOne",
			modifier:    func(cfg *Config) { cfg.ProposerWindowSize = 1 },
			expectedErr: ErrInvalidProposerWindowSize,
		},
		{
			name:        "NoL1Genesis",
			modifier:    func(cfg *Config) { cfg.Genesis.L1.Hash = common.Hash{} },
			expectedErr: ErrMissingGenesisL1Hash,
		},
		{
			name:        "NoL2Genesis",
			modifier:    func(cfg *Config) { cfg.Genesis.L2.Hash = common.Hash{} },
			expectedErr: ErrMissingGenesisL2Hash,
		},
		{
			name:        "GenesisHashesEqual",
			modifier:    func(cfg *Config) { cfg.Genesis.L2.Hash = cfg.Genesis.L1.Hash },
			expectedErr: ErrGenesisHashesSame,
		},
		{
			name:        "GenesisL2TimeZero",
			modifier:    func(cfg *Config) { cfg.Genesis.L2Time = 0 },
			expectedErr: ErrMissingGenesisL2Time,
		},
		{
			name:        "NoBatcherAddr",
			modifier:    func(cfg *Config) { cfg.Genesis.SystemConfig.BatcherAddr = common.Address{} },
			expectedErr: ErrMissingBatcherAddr,
		},
		{
			name:        "NoOverhead",
			modifier:    func(cfg *Config) { cfg.Genesis.SystemConfig.Overhead = eth.Bytes32{} },
			expectedErr: ErrMissingOverhead,
		},
		{
			name:        "NoScalar",
			modifier:    func(cfg *Config) { cfg.Genesis.SystemConfig.Scalar = eth.Bytes32{} },
			expectedErr: ErrMissingScalar,
		},
		{
			name:        "NoGasLimit",
			modifier:    func(cfg *Config) { cfg.Genesis.SystemConfig.GasLimit = 0 },
			expectedErr: ErrMissingGasLimit,
		},
		{
			name:        "NoBatchInboxAddress",
			modifier:    func(cfg *Config) { cfg.BatchInboxAddress = common.Address{} },
			expectedErr: ErrMissingBatchInboxAddress,
		},
		{
			name:        "NoDepositContractAddress",
			modifier:    func(cfg *Config) { cfg.DepositContractAddress = common.Address{} },
			expectedErr: ErrMissingDepositContractAddress,
		},
		{
			name:        "NoL1ChainId",
			modifier:    func(cfg *Config) { cfg.L1ChainID = nil },
			expectedErr: ErrMissingL1ChainID,
		},
		{
			name:        "NoL2ChainId",
			modifier:    func(cfg *Config) { cfg.L2ChainID = nil },
			expectedErr: ErrMissingL2ChainID,
		},
		{
			name:        "ChainIDsEqual",
			modifier:    func(cfg *Config) { cfg.L2ChainID = cfg.L1ChainID },
			expectedErr: ErrChainIDsSame,
		},
		{
			name:        "L1ChainIdNegative",
			modifier:    func(cfg *Config) { cfg.L1ChainID = big.NewInt(-1) },
			expectedErr: ErrL1ChainIDNotPositive,
		},
		{
			name:        "L1ChainIdZero",
			modifier:    func(cfg *Config) { cfg.L1ChainID = big.NewInt(0) },
			expectedErr: ErrL1ChainIDNotPositive,
		},
		{
			name:        "L2ChainIdNegative",
			modifier:    func(cfg *Config) { cfg.L2ChainID = big.NewInt(-1) },
			expectedErr: ErrL2ChainIDNotPositive,
		},
		{
			name:        "L2ChainIdZero",
			modifier:    func(cfg *Config) { cfg.L2ChainID = big.NewInt(0) },
			expectedErr: ErrL2ChainIDNotPositive,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cfg := randConfig()
			test.modifier(cfg)
			err := cfg.Check()
			assert.Same(t, err, test.expectedErr)
		})
	}
}
