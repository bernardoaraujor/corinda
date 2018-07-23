package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/bernardoaraujor/corinda/train"
	"fmt"
)

// trainCmd represents the train command
var trainCmd = &cobra.Command{
	Use:   "train",
	Short: "Train models from csv list",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(os.Args) <= 2 {
			fmt.Println("error: which input list should I use?")
			os.Exit(1)
		}

		list := os.Args[2]

		fmt.Println("Starting the training process based on " + list + ".csv")
		train.Train(list)
	},
}

func init() {
	rootCmd.AddCommand(trainCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// trainCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// trainCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
