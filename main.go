package main

import "github.com/spf13/cobra"

func main() {

	var collectorCmd = &cobra.Command{
		Use:   "collect",
		Short: "Run the collector",
		Run: func(cmd *cobra.Command, args []string) {
			startCollector()
		},
	}

	var etlCmd = &cobra.Command{
		Use:   "etl",
		Short: "Run the ETL processor",
		Run: func(cd *cobra.Command, args []string) {
			startETL()
		},
	}

	var rootCmd = &cobra.Command{Use: "snowblower"}
	rootCmd.AddCommand(collectorCmd)
	rootCmd.AddCommand(etlCmd)
	rootCmd.Execute()

}
