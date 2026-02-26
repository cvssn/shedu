package main

import (
	"log/slog"
	"net/http"
	_ "net/http/pprof"
	"os"

	"github.com/charmbracelet/shedu/internal/cmd"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	if os.Getenv("SHEDU_PROFILE") != "" {
		go func() {
			slog.Info("servindo pprof em localhost:6060")

			if httpErr := http.ListenAndServe("localhost:6060", nil); httpErr != nil {
				slog.Error("falha em pprof listen", "error", httpErr)
			}
		}()
	}

	cmd.Execute()
}