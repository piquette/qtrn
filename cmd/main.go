// Copyright © 2018 Piquette Capital, LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"crypto/tls"
	"net/http"

	finance "github.com/piquette/finance-go"
	quote "github.com/piquette/qtrn/cmd/quote"

	"github.com/piquette/qtrn/version"
	"github.com/spf13/cobra"
)

var (
	// versionF flag prints the version and exits.
	versionF bool
	// insecureF flag skips tls verification when executing requests.
	insecureF bool
)

// Execute adds all child commands to the root command sets flags appropriately.
func Execute() error {

	c := &cobra.Command{
		Use: "qtrn",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if insecureF {
				tr := &http.Transport{
					TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
				}
				client := &http.Client{Transport: tr}
				finance.SetHTTPClient(client)
			}
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if versionF {
				version.PrintVersion()
				return nil
			}
			return cmd.Usage()
		},
	}

	//	cmdQtrn.AddCommand(chartCmd)
	//	cmdQtrn.AddCommand(writeCmd)
	c.AddCommand(quote.Cmd)
	c.PersistentFlags().BoolVarP(&insecureF, "insecure", "x", false, "set `--insecure` or `-x` to skip tls verification during requests")
	c.Flags().BoolVarP(&versionF, "version", "v", false, "show the version and exit")
	return c.Execute()
}
