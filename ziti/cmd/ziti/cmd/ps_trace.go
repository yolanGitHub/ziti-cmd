/*
	Copyright NetFoundry, Inc.

	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

	https://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/

package cmd

import (
	cmdutil "github.com/netfoundry/ziti-cmd/ziti/cmd/ziti/cmd/factory"
	cmdhelper "github.com/netfoundry/ziti-cmd/ziti/cmd/ziti/cmd/helpers"
	"errors"
	"github.com/spf13/cobra"
	"io"
	"strconv"
)

// PsTraceOptions the options for the create spring command
type PsTraceOptions struct {
	PsOptions
	CtrlListener string
}

// NewCmdPsTrace creates a command object for the "create" command
func NewCmdPsTrace(f cmdutil.Factory, out io.Writer, errOut io.Writer) *cobra.Command {
	options := &PsTraceOptions{
		PsOptions: PsOptions{
			CommonOptions: CommonOptions{
				Factory: f,
				Out:     out,
				Err:     errOut,
			},
		},
	}

	cmd := &cobra.Command{
		Use: "trace",
		Run: func(cmd *cobra.Command, args []string) {
			options.Cmd = cmd
			options.Args = args
			err := options.Run()
			cmdhelper.CheckErr(err)
		},
	}

	options.addCommonFlags(cmd)
	cmd.Flags().StringVarP(&options.Flags.Pid, "pid", "", "", "pid (target of sub-cmd)")

	return cmd
}

// Run implements the command
func (o *PsTraceOptions) Run() error {
	pidStr := o.Args[0]
	_, err := strconv.Atoi(o.Args[0])
	if err != nil {
		return err
	}
	addr, err := targetToAddr(pidStr)
	if err != nil {
		return errors.New("Could not resolve pid -- may not have 'cliagent' enabled")
	}
	fn, ok := pscmds["trace"]
	if !ok {
		return errors.New("internal error")
	}
	var params []string
	params = append(params, pidStr)
	if err := fn(*addr, params); err != nil {
		return err
	}
	return nil
}
