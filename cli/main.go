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

package cli

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"

	finance "github.com/piquette/finance-go"
	"github.com/piquette/qtrn/version"
	"github.com/spf13/cobra"
)

var (
	// cmdQtrn is the root command.
	cmdQtrn = &cobra.Command{
		Use: "qtrn",
		Run: func(cmd *cobra.Command, args []string) {
			if flagPrintVersion {
				version.PrintVersion()
				return
			}
			cmd.Usage()
		},
	}
	//flagPrintVersion set flag to show current qtrn version.
	flagPrintVersion bool
)

func init() {
	//
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	finance.SetHTTPClient(client)

	//	cmdQtrn.AddCommand(chartCmd)
	//	cmdQtrn.AddCommand(writeCmd)
	cmdQtrn.AddCommand(equityCmd)
	cmdQtrn.Flags().BoolVarP(&flagPrintVersion, "version", "v", false, "show the version and exit")
}

// MainFunc adds all child commands to the root command sets flags appropriately.
func MainFunc() {
	if err := cmdQtrn.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
