package knode

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/log"
	"github.com/urfave/cli"

	"github.com/wemixkanvas/kanvas/components/node/chaincfg"
	"github.com/wemixkanvas/kanvas/components/node/flags"
	"github.com/wemixkanvas/kanvas/components/node/node"
	p2pcli "github.com/wemixkanvas/kanvas/components/node/p2p/cli"
	"github.com/wemixkanvas/kanvas/components/node/rollup"
	"github.com/wemixkanvas/kanvas/components/node/rollup/driver"
	"github.com/wemixkanvas/kanvas/components/node/sources"
	kpprof "github.com/wemixkanvas/kanvas/utils/service/pprof"
)

// NewConfig creates a Config from the provided flags or environment variables.
func NewConfig(ctx *cli.Context, log log.Logger) (*node.Config, error) {
	if err := flags.CheckRequired(ctx); err != nil {
		return nil, err
	}

	rollupConfig, err := NewRollupConfig(ctx)
	if err != nil {
		return nil, err
	}

	driverConfig := NewDriverConfig(ctx)

	p2pSignerSetup, err := p2pcli.LoadSignerSetup(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to load p2p signer: %w", err)
	}

	p2pConfig, err := p2pcli.NewConfig(ctx, rollupConfig.BlockTime)
	if err != nil {
		return nil, fmt.Errorf("failed to load p2p config: %w", err)
	}

	l1Endpoint := NewL1EndpointConfig(ctx)

	l2Endpoint, err := NewL2EndpointConfig(ctx, log)
	if err != nil {
		return nil, fmt.Errorf("failed to load l2 endpoints info: %w", err)
	}

	l2SyncEndpoint := NewL2SyncEndpointConfig(ctx)

	cfg := &node.Config{
		L1:     l1Endpoint,
		L2:     l2Endpoint,
		L2Sync: l2SyncEndpoint,
		Rollup: *rollupConfig,
		Driver: *driverConfig,
		RPC: node.RPCConfig{
			ListenAddr:  ctx.GlobalString(flags.RPCListenAddr.Name),
			ListenPort:  ctx.GlobalInt(flags.RPCListenPort.Name),
			EnableAdmin: ctx.GlobalBool(flags.RPCEnableAdmin.Name),
		},
		Metrics: node.MetricsConfig{
			Enabled:    ctx.GlobalBool(flags.MetricsEnabledFlag.Name),
			ListenAddr: ctx.GlobalString(flags.MetricsAddrFlag.Name),
			ListenPort: ctx.GlobalInt(flags.MetricsPortFlag.Name),
		},
		Pprof: kpprof.CLIConfig{
			Enabled:    ctx.GlobalBool(flags.PprofEnabledFlag.Name),
			ListenAddr: ctx.GlobalString(flags.PprofAddrFlag.Name),
			ListenPort: ctx.GlobalInt(flags.PprofPortFlag.Name),
		},
		P2P:                 p2pConfig,
		P2PSigner:           p2pSignerSetup,
		L1EpochPollInterval: ctx.GlobalDuration(flags.L1EpochPollIntervalFlag.Name),
		Heartbeat: node.HeartbeatConfig{
			Enabled: ctx.GlobalBool(flags.HeartbeatEnabledFlag.Name),
			Moniker: ctx.GlobalString(flags.HeartbeatMonikerFlag.Name),
			URL:     ctx.GlobalString(flags.HeartbeatURLFlag.Name),
		},
	}
	if err := cfg.Check(); err != nil {
		return nil, err
	}
	return cfg, nil
}

func NewL1EndpointConfig(ctx *cli.Context) *node.L1EndpointConfig {
	return &node.L1EndpointConfig{
		L1NodeAddr: ctx.GlobalString(flags.L1NodeAddr.Name),
		L1TrustRPC: ctx.GlobalBool(flags.L1TrustRPC.Name),
		L1RPCKind:  sources.RPCProviderKind(strings.ToLower(ctx.GlobalString(flags.L1RPCProviderKind.Name))),
	}
}

func NewL2EndpointConfig(ctx *cli.Context, log log.Logger) (*node.L2EndpointConfig, error) {
	l2Addr := ctx.GlobalString(flags.L2EngineAddr.Name)
	fileName := ctx.GlobalString(flags.L2EngineJWTSecret.Name)
	var secret [32]byte
	fileName = strings.TrimSpace(fileName)
	if fileName == "" {
		return nil, fmt.Errorf("file-name of jwt secret is empty")
	}
	if data, err := os.ReadFile(fileName); err == nil {
		jwtSecret := common.FromHex(strings.TrimSpace(string(data)))
		if len(jwtSecret) != 32 {
			return nil, fmt.Errorf("invalid jwt secret in path %s, not 32 hex-formatted bytes", fileName)
		}
		copy(secret[:], jwtSecret)
	} else {
		log.Warn("Failed to read JWT secret from file, generating a new one now. Configure L2 geth with --authrpc.jwt-secret=" + fmt.Sprintf("%q", fileName))
		if _, err := io.ReadFull(rand.Reader, secret[:]); err != nil {
			return nil, fmt.Errorf("failed to generate jwt secret: %w", err)
		}
		if err := os.WriteFile(fileName, []byte(hexutil.Encode(secret[:])), 0600); err != nil {
			return nil, err
		}
	}

	return &node.L2EndpointConfig{
		L2EngineAddr:      l2Addr,
		L2EngineJWTSecret: secret,
	}, nil
}

// NewL2SyncEndpointConfig returns a pointer to a L2SyncEndpointConfig if the
// flag is set, otherwise nil.
func NewL2SyncEndpointConfig(ctx *cli.Context) *node.L2SyncEndpointConfig {
	return &node.L2SyncEndpointConfig{
		L2NodeAddr: ctx.GlobalString(flags.BackupL2UnsafeSyncRPC.Name),
	}
}

func NewDriverConfig(ctx *cli.Context) *driver.Config {
	return &driver.Config{
		SyncerConfDepth:   ctx.GlobalUint64(flags.SyncerL1Confs.Name),
		ProposerConfDepth: ctx.GlobalUint64(flags.ProposerL1Confs.Name),
		ProposerEnabled:   ctx.GlobalBool(flags.ProposerEnabledFlag.Name),
		ProposerStopped:   ctx.GlobalBool(flags.ProposerStoppedFlag.Name),
	}
}

func NewRollupConfig(ctx *cli.Context) (*rollup.Config, error) {
	network := ctx.GlobalString(flags.Network.Name)
	if network != "" {
		config, err := chaincfg.GetRollupConfig(network)
		if err != nil {
			return nil, err
		}

		return &config, nil
	}

	rollupConfigPath := ctx.GlobalString(flags.RollupConfig.Name)
	file, err := os.Open(rollupConfigPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read rollup config: %w", err)
	}
	defer file.Close()

	var rollupConfig rollup.Config
	if err := json.NewDecoder(file).Decode(&rollupConfig); err != nil {
		return nil, fmt.Errorf("failed to decode rollup config: %w", err)
	}
	return &rollupConfig, nil
}

func NewSnapshotLogger(ctx *cli.Context) (log.Logger, error) {
	snapshotFile := ctx.GlobalString(flags.SnapshotLog.Name)
	handler := log.DiscardHandler()
	if snapshotFile != "" {
		var err error
		handler, err = log.FileHandler(snapshotFile, log.JSONFormat())
		if err != nil {
			return nil, err
		}
		handler = log.SyncHandler(handler)
	}
	logger := log.New()
	logger.SetHandler(handler)
	return logger, nil
}
