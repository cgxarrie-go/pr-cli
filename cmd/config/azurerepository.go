package config

import (
	"fmt"

	"github.com/spf13/cobra"

	appCfg "github.com/cgxarrie-go/prq/config"
)

// azureRepositoryCmd represents the azureRepository command
var azureRepositoryCmd = &cobra.Command{
	Use:   "repo",
	Short: "set azure Repository",
	Long:  `Set the Azure Repository in the configuration file`,
	RunE: func(cmd *cobra.Command, args []string) error {
		action, project, repo, err := getAzureRepositoryAction(cmd)
		if err != nil {
			return err
		}

		err = action(project, repo)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	ConfigCmd.AddCommand(azureRepositoryCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// azureRepositoryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	azureRepositoryCmd.Flags().StringP("add", "a", "", "Add a repo")
	azureRepositoryCmd.Flags().StringP("del", "d", "", "Remove a repo")
	azureRepositoryCmd.Flags().StringP("project", "p", "", "The project")
}

func getAzureRepositoryAction(cmd *cobra.Command) (
	action func(project, repo string) error, project, repo string, err error) {

	pFlag, _ := cmd.Flags().GetString("project")
	if pFlag == "" {
		return action, project, repo,
			fmt.Errorf("missing mandatory -p <project-name>")
	}

	addFlag, _ := cmd.Flags().GetString("add")
	if addFlag != "" {
		return runAddAzureRepositoryCmd, pFlag, addFlag, nil
	}

	delFlag, _ := cmd.Flags().GetString("del")
	if delFlag != "" {
		return runDeleteAzureRepositoryCmd, pFlag, delFlag, nil
	}

	return action, project, repo,
		fmt.Errorf("missig action flag for azure repo command")

}

func runAddAzureRepositoryCmd(project, repo string) error {
	errAlreadyInConfig := "repo %s already in azure config"
	cfg := appCfg.GetInstance()
	cfg.Load()

	projectIdx, err := getAzureRepositoryProjectIndex(cfg.Azure, project)
	if err != nil {
		return err
	}

	prj := cfg.Azure.Projects[projectIdx]
	if len(prj.RepositoryIDs) == 1 && prj.RepositoryIDs[0] == repo {
		return fmt.Errorf(errAlreadyInConfig, repo)
	}

	if len(prj.RepositoryIDs) > 1 {
		for _, v := range prj.RepositoryIDs {
			if v == repo {
				return fmt.Errorf(errAlreadyInConfig, repo)
			}
		}
	}

	cfg.Azure.Projects[projectIdx].RepositoryIDs =
		append(prj.RepositoryIDs, repo)
	err = cfg.Save()
	if err != nil {
		return err
	}
	return nil
}

func runDeleteAzureRepositoryCmd(project, repo string) error {

	errNotFound := "project %s not found in azure config"

	cfg := appCfg.GetInstance()
	cfg.Load()

	projectIdx, err := getAzureRepositoryProjectIndex(cfg.Azure, project)
	if err != nil {
		return err
	}

	prj := cfg.Azure.Projects[projectIdx]
	if len(prj.RepositoryIDs) == 0 {
		return fmt.Errorf(errNotFound, repo)
	}

	if len(prj.RepositoryIDs) == 1 && prj.RepositoryIDs[0] == repo {
		prj.RepositoryIDs = make([]string, 0)
		err := cfg.Save()
		if err != nil {
			return err
		}
		return nil
	}

	for i, v := range prj.RepositoryIDs {
		if v == repo {
			cfg.Azure.Projects[projectIdx].RepositoryIDs =
				append(prj.RepositoryIDs[:i], prj.RepositoryIDs[i+1:]...)
			err := cfg.Save()
			if err != nil {
				return err
			}
			return nil
		}
	}

	return fmt.Errorf(errNotFound, repo)
}

func getAzureRepositoryProjectIndex(cfg appCfg.AzureConfig, name string) (
	int, error) {

	errNotFound := "project %s not found in azure config"
	if len(cfg.Projects) == 0 {
		return -1, fmt.Errorf(errNotFound, name)
	}

	if len(cfg.Projects) == 1 && cfg.Projects[0].ID == name {
		return 0, nil
	}

	for i, v := range cfg.Projects {
		if v.ID == name {
			return i, nil
		}
	}

	return -1, fmt.Errorf(errNotFound, name)
}
