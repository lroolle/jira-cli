package versions

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/ankitpokhrel/jira-cli/api"
	"github.com/ankitpokhrel/jira-cli/internal/cmdutil"
	"github.com/ankitpokhrel/jira-cli/internal/view"
	"github.com/ankitpokhrel/jira-cli/pkg/jira"
)

// NewCmdVersions is a versions command.
func NewCmdVersions() *cobra.Command {
	return &cobra.Command{
		Use:     "versions",
		Short:   "List newly created versions in a project",
		Long:    "List newly created versions in a project.",
		Aliases: []string{"version", "ver"},
		Run:     Versions,
	}
}

// Versions displays a list view of newly created versions.
func Versions(cmd *cobra.Command, _ []string) {
	project := viper.GetString("project.key")

	debug, err := cmd.Flags().GetBool("debug")
	cmdutil.ExitIfError(err)

	versions, err := func() ([]*jira.Version, error) {
		s := cmdutil.Info(fmt.Sprintf("Fetching versions in project %s...", project))
		defer s.Stop()

		resp, err := api.DefaultClient(debug).GetProjectVersions(project)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}()
	cmdutil.ExitIfError(err)

	if len(versions) == 0 {
		fmt.Println()
		cmdutil.Failed("No versions found in project %q", project)
		return
	}

	v := view.NewVersion(versions)

	cmdutil.ExitIfError(v.Render())
}
