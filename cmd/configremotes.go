package cmd

import (
	"errors"
	"strings"

	"github.com/cgxarrie-go/prq/internal/config"
	"github.com/spf13/cobra"
)

// CfgPATCmd config azure PAT command
var ConfigRemotes = &cobra.Command{
	Use:   "remotes",
	Short: "add or remove remotes to config",
	Long:  `add or remove remotes to config`,
	RunE: func(cmd *cobra.Command, args []string) error {
		add, _ := cmd.Flags().GetString("add")
		remove, _ := cmd.Flags().GetString("remove")

		errs := []error{}

		if add != "" {
			err := runAddRemote(add)
			if err != nil {
				errs = append(errs, err)
			}
		}

		if remove != "" {
			err := runRemoveRemote(remove)
			if err != nil {
				errs = append(errs, err)
			}			
		}
		
		if len(errs) > 0 {
			errorMessages := make([]string, 0, len(errs))
    		for _, err := range errs {
        		errorMessages = append(errorMessages, err.Error())
    		}
    		combinedError := errors.New(strings.Join(errorMessages, "\n"))
    		return combinedError
		}

		return nil
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// azurePATCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:

	ConfigRemotes.Flags().StringP("add", "a", "", "add")
	ConfigRemotes.Flags().StringP("remove", "r", "", "remove")

}

func runAddRemote(remote string) error {
	cfg := config.GetInstance()
	cfg.Load()
	cfg.Remotes = append(cfg.Remotes, remote)
	err := cfg.Save()
	if err != nil {
		return err
	}
	return nil
}

func runRemoveRemote(remote string) error {
	cfg := config.GetInstance()
	cfg.Load()

	for i, v := range cfg.Remotes {
    	if v == remote {
            cfg.Remotes = append(cfg.Remotes[:i], cfg.Remotes[i+1:]...)
        }
    }

	err := cfg.Save()
	if err != nil {
		return err
	}
	return nil
}
