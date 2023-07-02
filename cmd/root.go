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
		k, err := cmd.PersistentFlags().GetInt("k")
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error retrieving value for flag --k.")
			os.Exit(1)
		}

		seed, err := cmd.PersistentFlags().GetInt64("seed")
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error retrieving value for flag --seed.")
			os.Exit(1)
		}

		dsFactor, err := cmd.PersistentFlags().GetFloat64("downsample")
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error retrieving value for flag --downsample.")
			os.Exit(1)
		}
		if dsFactor <= 0 || dsFactor > 1 {
			fmt.Fprintln(os.Stderr, "Downsample factor must be greater than 0 and less than or equal to 1.")
			os.Exit(1)
		}

		fmt.Println(palette.Create(args, k, seed, dsFactor))
	},
}

func init() {
	RootCommand.PersistentFlags().IntP("k", "k", 5, "the size of the palette")
	RootCommand.PersistentFlags().Int64P("seed", "s", -1, "the seed used for initializing cluster means. -1 means no seed is used")
	RootCommand.PersistentFlags().Float64P("downsample", "d", 1, "the factor by which to downsample the image before creating a palette")
}
