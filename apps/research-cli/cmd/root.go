package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

// rootCmd represents the base command when called without any subcommands
func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "disclosure-importer",
		Short: "A financial results importer",
		Long: `Query financial results report paths from speeda.ir_rep on common database. After that,
get file data and post metadata and reoprt pdf to sales api.`,
	}
	cobra.OnInitialize(initConfig)
	cmd.AddCommand(NewImportCmd())
	return cmd
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cmd := NewRootCmd()
	cmd.SetOut(os.Stdout)
	if err := cmd.Execute(); err != nil {
		cmd.SetOut(os.Stderr)
		cmd.Println(err)
		os.Exit(1)
	}
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	//if err := viper.BindEnv("common_db_password", "COMMON_DB_PASSWORD"); err != nil {
	//	fmt.Print(err)
	//}
	viper.AutomaticEnv() // read in environment variables that match
}

