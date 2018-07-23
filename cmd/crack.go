package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// crackCmd represents the crack command

//Usage: corinda crack <trained> <target>
//
//example: corinda crack rockyou linkedin
var crackCmd = &cobra.Command{
        Use:   "crack",
        Short: "Use trained maps to crack target list",
        Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("crack called")
	},
}

func init() {
	rootCmd.AddCommand(crackCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// crackCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// crackCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
