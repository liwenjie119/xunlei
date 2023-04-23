package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
)

const (
	ENV_WEB_PORT    = "XL_WEB_PORT"
	ENV_WEB_ADDRESS = "XL_WEB_ADDRESS"
	ENV_DEBUG       = "XL_DEBUG"

	rootfs  = "/xunlei"
	dataDir = "/data"
	// downloadDir = "/downloads"
)

func init() {
	if os.Getenv(ENV_DEBUG) == "1" {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
	}
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	var mode string
	if len(os.Args) > 1 {
		mode = os.Args[1]
	}

	var err error
	switch mode {
	case "syno":
		err = syno(ctx)
	case "daemon":
		err = daemon(ctx)
	case "run":
		err = xlp(ctx)
	default:
		fmt.Fprintf(os.Stderr, "Usage: %s [syno|daemon|run]\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}
	if err != nil {
		log.Fatalf("%v", err)
	}
}
