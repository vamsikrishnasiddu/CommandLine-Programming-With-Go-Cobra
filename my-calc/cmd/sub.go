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

// subCmd represents the sub command
var subCmd = &cobra.Command{
	Use:   "sub",
	Short: "It's for subtraction of arguments",
	Long: `
	  This will calculate the subtraction of more than two arguments .
	  This is useful for both the Integer and floatingPoint Values.
	  Use -f or --float after the command to calculate floatingpoint
	  numbers subtraction.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		floatval, err := cmd.Flags().GetBool("float")
		if err != nil {
			fmt.Println(err)
		} else {
			if floatval {
				subFloats(args)
			} else {
				subInts(args)
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(subCmd)

	subCmd.Flags().BoolP("float", "f", false, "sub float numbers")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// subCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// subCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func subInts(args []string) {
	var diff int
	diff, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println(err)
	}

	for i := 1; i < len(args); i++ {
		ival, err := strconv.Atoi(args[i])

		if err != nil {
			fmt.Println(err)
		} else {
			diff = diff - ival
		}

	}
	fmt.Println(diff)

}

func subFloats(args []string) {
	var diff float64
	diff, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		fmt.Println(err)
	}

	for i := 1; i < len(args); i++ {
		ival, err := strconv.ParseFloat(args[i], 64)

		if err != nil {
			fmt.Println(err)
		} else {
			diff = diff - ival
		}

	}
	fmt.Println(diff)

}
