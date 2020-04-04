package cmd

import (
	"github.com/spf13/cobra"
	"os"
	"time"
)

func NewImportCmd() *cobra.Command {
	type Options struct {
		dbHost     string
		dbUser     string
		dbPort     int
		dbPass     string
		dbSchema   string
		endpoint   string
		targetDate string
		globalIds  string
		verbose    bool
	}
	o := &Options{}
	cmd := &cobra.Command{
		Use:   "import",
		Short: "Import for presentation material such as financial results.",
		Run: func(cmd *cobra.Command, args []string) {
			con := config.NewConfig(o.dbHost, o.dbPort, o.dbUser, o.dbPass, o.dbSchema, o.verbose)
			t, err := time.Parse("2006-01-02", o.targetDate)
			if err != nil {
				cmd.PrintErr(err)
				os.Exit(1)
			}
			cmd.Printf("[%s] start import\n", time.Now().Format("2006-01-02 15:04:05"))
			if err := importReports(con, o.endpoint, t, o.globalIds); err != nil {
				cmd.PrintErr(err)
				os.Exit(1)
			} else {
				cmd.Printf("[%s] finish import\n", time.Now().Format("2006-01-02 15:04:05"))
			}
		},
	}
	o.dbPass = os.Getenv("COMMON_DB_PASSWORD")
	cmd.Flags().StringVarP(&o.dbHost, "dbhost", "d", "speeda-common-db", "common db host")
	cmd.Flags().StringVarP(&o.dbUser, "dbUser", "u", "sample", "common db dbPass")
	cmd.Flags().IntVarP(&o.dbPort, "dbPort", "P", 3306, "common db dbPass")
	cmd.Flags().BoolVarP(&o.verbose, "verbose", "v", false, "output debug logs")
	cmd.Flags().StringVarP(&o.endpoint, "endpoint", "e", "http://localhost:8080", "disclosure api endpoint")
	cmd.Flags().StringVarP(&o.globalIds, "globalids", "g", "", "target global id list")
	//TODO 時間で絞れるようにする
	cmd.Flags().StringVarP(&o.targetDate, "target-date", "t", "2019-02-01", "target date")
	return cmd
}
