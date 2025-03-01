package main

import (
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/log"
	"github.com/urfave/cli"

	batcher "github.com/wemixkanvas/kanvas/components/batcher"
	"github.com/wemixkanvas/kanvas/components/batcher/flags"
	klog "github.com/wemixkanvas/kanvas/utils/service/log"
)

var (
	Version   = ""
	GitCommit = ""
	GitDate   = ""
)

func main() {
	klog.SetupDefaults()

	app := cli.NewApp()
	app.Flags = flags.Flags
	app.Version = fmt.Sprintf("%s-%s-%s", Version, GitCommit, GitDate)
	app.Name = "kanvas-batcher"
	app.Usage = "Batch Submitter Service"
	app.Description = "Service for generating and submitting L2 tx batches " +
		"to L1"

	app.Action = curryMain(Version)
	err := app.Run(os.Args)
	if err != nil {
		log.Crit("Application failed", "message", err)
	}
}

// curryMain transforms the batcher.Main function into an app.Action
// This is done to capture the Version of the batcher.
func curryMain(version string) func(ctx *cli.Context) error {
	return func(ctx *cli.Context) error {
		return batcher.Main(version, ctx)
	}
}
