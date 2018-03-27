// Copyright © 2018 Tommy Falgout <tommy@lastcoolnameleft.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"os"

	"github.com/lastcoolnameleft/yarm/yarm"
	"github.com/spf13/cobra"
)

var output string
var subscriptionID string
var resourceGroup string

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate the ARM template",
	Long:  `Requires an input of the values.yaml file to generate the ARM template`,
	Run: func(cmd *cobra.Command, args []string) {
		valuesFile := os.Args[2]
		yarm.CreateArm(valuesFile, output, subscriptionID, resourceGroup)
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	generateCmd.Flags().StringVarP(&output, "output", "o", "json", "Output format <yaml|json>")
	generateCmd.Flags().StringVarP(&subscriptionID, "subscriptionID", "s", "", "Subscription ID <uuid>")
	generateCmd.Flags().StringVarP(&resourceGroup, "resourceGroup", "g", "", "Resource Group <yarm>")
}
