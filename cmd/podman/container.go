package main

import (
	"strings"

	"github.com/containers/libpod/cmd/podman/cliconfig"
	"github.com/spf13/cobra"
)

var (
	containerDescription = "Manage containers"
	containerCommand     = cliconfig.PodmanCommand{
		Command: &cobra.Command{
			Use:              "container",
			Short:            "Manage Containers",
			Long:             containerDescription,
			TraverseChildren: true,
			RunE:             commandRunE(),
		},
	}

	contInspectSubCommand  cliconfig.InspectValues
	_contInspectSubCommand = &cobra.Command{
		Use:   strings.Replace(_inspectCommand.Use, "| IMAGE", "", 1),
		Short: "Display the configuration of a container",
		Long:  `Displays the low-level information on a container identified by name or ID.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			contInspectSubCommand.InputArgs = args
			contInspectSubCommand.GlobalFlags = MainGlobalOpts
			return inspectCmd(&contInspectSubCommand)
		},
		Example: `podman container inspect myCtr
  podman container inspect -l --format '{{.Id}} {{.Config.Labels}}'`,
	}

	listSubCommand  cliconfig.PsValues
	_listSubCommand = &cobra.Command{
		Use:     strings.Replace(_psCommand.Use, "ps", "list", 1),
		Args:    noSubArgs,
		Short:   _psCommand.Short,
		Long:    _psCommand.Long,
		Aliases: []string{"ls"},
		RunE: func(cmd *cobra.Command, args []string) error {
			listSubCommand.InputArgs = args
			listSubCommand.GlobalFlags = MainGlobalOpts
			return psCmd(&listSubCommand)
		},
		Example: strings.Replace(_psCommand.Example, "podman ps", "podman container list", -1),
	}

	// Commands that are universally implemented.
	containerCommands = []*cobra.Command{
		_attachCommand,
		_containerExistsCommand,
		_contInspectSubCommand,
		_diffCommand,
		_exportCommand,
		_createCommand,
		_killCommand,
		_listSubCommand,
		_logsCommand,
		_runCommand,
		_rmCommand,
		_waitCommand,
	}
)

func init() {
	contInspectSubCommand.Command = _contInspectSubCommand
	inspectInit(&contInspectSubCommand)

	listSubCommand.Command = _listSubCommand
	psInit(&listSubCommand)

	containerCommand.AddCommand(containerCommands...)
	containerCommand.AddCommand(getContainerSubCommands()...)
	containerCommand.SetUsageTemplate(UsageTemplate())

	rootCmd.AddCommand(containerCommand.Command)
}
