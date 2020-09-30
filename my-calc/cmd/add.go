/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Addition of arguments",
	Long: `
	This is for addition of arguments if you use integer no flags 
	are required but if you use float use the flag -f or --float.`,
	Run: func(cmd *cobra.Command, args []string) {
		fstatus, _ := cmd.Flags().GetBool("float")
		if fstatus {
			addFloats(args)
		} else {
			addInts(args)

		}

	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().BoolP("float", "f", false, "Add Floating numbers")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func addInts(args []string) {
	var sum int
	for _, ival := range args {
		itemp, err := strconv.Atoi(ival)

		if err != nil {
			fmt.Println(err)
		}
		sum = sum + itemp
	}
	fmt.Println(sum)
}

func addFloats(args []string) {
	var sum float64
	for _, ival := range args {
		itemp, err := strconv.ParseFloat(ival, 64)

		if err != nil {
			fmt.Println(err)
		}
		sum = sum + itemp
	}
	fmt.Println(sum)
}