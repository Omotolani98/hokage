package commands

import (
	"log"

	"github.com/Omotolani98/hokage/pkg/services"
	"github.com/Omotolani98/hokage/pkg/utils"

	"github.com/spf13/cobra"
)

func Apply() *cobra.Command {
	var path string

	cmd := &cobra.Command{
		Use:   "apply",
		Short: "applies arrangement and folder creation for provided path",
		Run: func(cmd *cobra.Command, args []string) {
			services.Manager(path)
		},
	}

	cmd.Flags().StringVarP(&path, "path", "p", "", "input path to sort out")
	folderExtensions := services.FolderFileExtensions
	for ffExt := range folderExtensions {
		ffExtSlice := folderExtensions[ffExt]
		key, ok := utils.Mapkey(folderExtensions, ffExtSlice)
		if !ok {
			log.Fatalf("could not extract key")
		}

		_ = cmd.MarkFlagFilename(path, services.FolderFileExtensions[key]...)
	}

	return cmd
}
