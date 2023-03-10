package config

import (
	"fmt"

	"github.com/spf13/cobra"

	appCfg "github.com/cgxarrie-go/prq/config"
	"github.com/cgxarrie-go/prq/domain/errors"
)

// azureProjectCmd represents the azureProject command
var azureProjectCmd = &cobra.Command{
	Use:   "project",
	Short: "set azure Project",
	Long:  `Set the Azure Project in the configuration file`,
	Run: func(cmd *cobra.Command, args []string) {
		nextFlag := true
		var err error
		addFlag, _ := cmd.Flags().GetString("add")
		if addFlag != "" {
			nextFlag = false
			err = runAddAzureProjectCmd(addFlag)
		}

		if nextFlag {
			delFlag, _ := cmd.Flags().GetString("del")
			if delFlag != "" {
				nextFlag = false
				err = runDeleteAzureProjectCmd(delFlag)
			}
		}

		if nextFlag {
			err = fmt.Errorf("missig action flag for azure project command")
		}

		if err != nil {
			errors.Print(err)
			return
		}

	},
}

func init() {
	azureCmd.AddCommand(azureProjectCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// azureProjectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	azureProjectCmd.Flags().StringP("add", "a", "", "Add a project")
	azureProjectCmd.Flags().StringP("del", "d", "", "Remove a project")
}

func runAddAzureProjectCmd(name string) error {
	errAlreadyInConfig := "project %s already in azure config"
	cfg := appCfg.GetInstance()
	cfg.Load()

	if len(cfg.Azure.Projects) == 1 && cfg.Azure.Projects[0].ID == name {
		return fmt.Errorf(errAlreadyInConfig, name)
	}

	if len(cfg.Azure.Projects) > 1 {
		for _, v := range cfg.Azure.Projects {
			if v.ID == name {
				return fmt.Errorf(errAlreadyInConfig, name)
			}
		}
	}

	prj := appCfg.AzureProjectConfig{
		ID:            name,
		RepositoryIDs: []string{},
	}

	cfg.Azure.Projects = append(cfg.Azure.Projects, prj)
	err := cfg.Save()
	if err != nil {
		return err
	}
	return nil
}

func runDeleteAzureProjectCmd(name string) error {

	errNotFound := "project %s not found in azure config"

	cfg := appCfg.GetInstance()
	cfg.Load()

	if len(cfg.Azure.Projects) == 0 {
		return fmt.Errorf(errNotFound, name)
	}

	if len(cfg.Azure.Projects) == 1 && cfg.Azure.Projects[0].ID == name {
		cfg.Azure.Projects = make([]appCfg.AzureProjectConfig, 0)
		err := cfg.Save()
		if err != nil {
			return err
		}
		return nil
	}

	for i, v := range cfg.Azure.Projects {
		if v.ID == name {
			cfg.Azure.Projects = append(cfg.Azure.Projects[:i],
				cfg.Azure.Projects[i+1:]...)
			err := cfg.Save()
			if err != nil {
				return err
			}
			return nil
		}
	}

	return fmt.Errorf(errNotFound, name)
}
