package cmd

import (
	"os"
	"path/filepath"

	"charm.land/lipgloss/v2"
	"charm.land/lipgloss/v2/table"

	"github.com/charmbracelet/shedu/internal/config"
	"github.com/charmbracelet/x/term"
	"github.com/spf13/cobra"
)

var dirsCmd = &cobra.Command{
	Use:     "dirs",
	Short:   "printa diretórios utilizados pelo shedu",
	Long:    `printa os diretórios aonde o shedu armazena seus arquivos de configurações e dados.
isso inclui o diretório da configuração global e diretório de dados`,
	Example: `
# printa todos os diretórios
shedu dirs

# printa apenas o diretório de configuração
shedu dirs config

# printa apenas o diretório de dados
shedu dirs data
`,
	Run: func(cmd *cobra.Command, args []string) {
		if term.IsTerminal(os.Stdout.Fd()) {
			// estamos em um tty: deixar elegante
			t := table.New().
				Border(lipgloss.RoundedBorder()).
				StyleFunc(func(row, col int) lipgloss.Style {
					return lipgloss.NewStyle().Padding(0, 2)
				}).
				Row("Config", filepath.Dir(config.GlobalConfig())).
				Row("Data", filepath.Dir(config.GlobalConfigData()))

			lipgloss.Println(t)

			return
		}

		// não é um tty
		cmd.Println(filepath.Dir(config.GlobalConfig()))
		cmd.Println(filepath.Dir(config.GlobalConfigData()))
	}
}

var configDirCmd = &cobra.Command{
	Use:   "config",
	Short: "printa o diretório de configuração utilizado pelo shedu",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println(filepath.Dir(config.GlobalConfig()))
	}
}

var dataDirCmd = &cobra.Command{
	Use:   "data",
	Short: "printa o diretório de dados utilizado pelo shedu",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println(filepath.Dir(config.GlobalConfigData()))
	}
}

func init() {
	dirsCmd.AddCommand(configDirCmd, dataDirCmd)
}