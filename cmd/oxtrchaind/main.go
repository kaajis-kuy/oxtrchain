package main

import (
    "github.com/oxtr/oxtrchain/app"
    "github.com/oxtr/oxtrchain/cmd/oxtrchaind/cmd"
)

func main() {
	rootCmd := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, clienthelpers.EnvPrefix, app.DefaultNodeHome); err != nil {
		fmt.Fprintln(rootCmd.OutOrStderr(), err)
		os.Exit(1)
	}
}
