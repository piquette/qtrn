// Copyright © 2017 Michael Ackley <ackleymi@gmail.com>
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
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// cmdQtrn represents the base command when called without any subcommands
var cmdQtrn = &cobra.Command{
	Use:          "qtrn",
	SilenceUsage: true,
}

func init() {
	cobra.OnInitialize(initConfig)

	// Config.
	cmdQtrn.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.qtrn.yaml)")
	cmdQtrn.AddCommand(
		chartCmd,
		writeCmd,
		quoteCmd,
	)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName(".qtrn") // name of config file (without extension)
	viper.AddConfigPath("$HOME") // adding home directory as first search path
	viper.AutomaticEnv()         // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

// Main adds all child commands to the root command sets flags appropriately.
func Main() {
	if err := cmdQtrn.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
