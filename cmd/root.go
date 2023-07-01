package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/tyler-m/palette/palette"
)

var RootCommand = &cobra.Command{
	Use:   "palette [flags] image_path [...image_path]",
	Short: "palette - Creates a palette based on the colors of a given image.",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		kFlag, err := cmd.PersistentFlags().GetInt("k")
		if err != nil {
			os.Exit(1)
		}

		seedFlag, err := cmd.PersistentFlags().GetInt64("seed")
		if err != nil {
			os.Exit(1)
		}

		fmt.Println(palette.Create(args, kFlag, seedFlag))
	},
}

func init() {
	RootCommand.PersistentFlags().IntP("k", "k", 5, "the size of the palette")
	RootCommand.PersistentFlags().Int64P("seed", "s", -1, "the seed used for initializing cluster means. -1 means no seed is used")
}
