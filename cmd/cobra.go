package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"sme-jwk/cmd/api"
	"sme-jwk/cmd/version"
)

var rootCmd = &cobra.Command{
	Use:          "sme-jwk",
	Short:        "sme-jwk",
	SilenceUsage: true,
	Long:         `sme-jwk`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			tip()
			return errors.New("requires at least one arg")
		}
		return nil
	},
	PersistentPreRunE: func(*cobra.Command, []string) error {
		tip()
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		tip()
	},
}

func tip() {
	strUp := `
	+-+-+-+-+-+-+-+ +-+-+-+-+-+-+-+
	|w|e|l|c|o|m|e| |s|m|e|-|j|w|k|
	+-+-+-+-+-+-+-+ +-+-+-+-+-+-+-+`
	strDown := `	├── cobra
	├── viper
	├── wire
	└── clean
`
	color.Set(color.FgHiBlue, color.Bold)
	defer color.Unset()
	fmt.Println(strUp)
	color.Set(color.FgGreen, color.Bold)
	fmt.Println(strDown)
}

func init() {
	rootCmd.AddCommand(api.StartCmd)
	rootCmd.AddCommand(version.StartCmd)
}

// Execute Execute
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
