package cmd

import (
	"fmt"

	"github.com/cgxarrie-go/prq/internal/config"
	"github.com/cgxarrie-go/prq/internal/ports"
	"github.com/cgxarrie-go/prq/internal/remoteclient"
	"github.com/cgxarrie-go/prq/internal/services"
	"github.com/cgxarrie-go/prq/internal/utils"
	"github.com/muesli/termenv"
	"github.com/spf13/cobra"
)

// ListCmd represents the list command
var createCmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"c"},
	Short:   "Create Pull Request",
	Long: `Create a Pull Request.
	
	Flags:
	-g, --target	: target branch. If blank, default is used
	-t, --title	: title. If blank, standard title is used
	-f, --draft	: draft. default is true
	-d, --desc	: description. default is emty
	-m, --templ	: template. default is none`,
	RunE: func(cmd *cobra.Command, args []string) error {
		dest, _ := cmd.Flags().GetString("target")
		ttl, _ := cmd.Flags().GetString("title")
		dft, _ := cmd.Flags().GetString("draft")
		desc, _ := cmd.Flags().GetString("desc")
		template, _ := cmd.Flags().GetString("templ")
		draft := !utils.IsFalse(dft)

		config.GetInstance().Load()

		r, err := utils.CurrentFolderRemote()
		if err != nil {
			return fmt.Errorf("getting remote: %w", err)
		}

		cl, err := remoteclient.NewRemoteClient(r)
		if err != nil {
			return fmt.Errorf("creating remote client: %w", err)
		}

		svc := services.NewCreatePRService(cl)
		svcReq := ports.CreatePRSvcRequest{
			Destination: dest,
			Title:       ttl,
			IsDraft:     draft,
			Description: desc,
			Template:    template,
		}

		pr, err := svc.Run(svcReq)
		if err != nil {
			return fmt.Errorf("creating PR: %w", err)
		}

		lnk := termenv.Hyperlink(pr.Link, "open PR")

		fmt.Printf("PR created with ID: %s (%s)\n", pr.ID, lnk)
		return nil
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// azCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// azCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	createCmd.Flags().StringP("target", "g", "", "target branch. If blank, default is used")
	createCmd.Flags().StringP("title", "t", "", "title. If blank, standard title is used")
	createCmd.Flags().BoolP("draft", "f", true, "draft. default is true")
	createCmd.Flags().StringP("desc", "d", "", "description. default is emty")
	createCmd.Flags().StringP("templ", "m", "", "template. default is none")

}
