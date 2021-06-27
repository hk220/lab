package main

import (
	"log"
	"os"

	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var (
	rootCmd = &cobra.Command{
		Use:  "hello-zpr",
		RunE: runRoot,
	}
	logger logr.Logger
)

type Server struct {
	Logger logr.Logger
}

func runRoot(cmd *cobra.Command, args []string) error {
	zapLogger, err := zap.NewProduction()
	if err != nil {
		return err
	}
	s := Server{Logger: zapr.NewLogger(zapLogger)}
	s.Logger.Info("Server")
	return nil
}

func main() {
	zapLogger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}
	logger = zapr.NewLogger(zapLogger)
	if err := rootCmd.Execute(); err != nil {
		logger.Error(err, "hello-zapr ran into an error")
		os.Exit(1)
	}
}
