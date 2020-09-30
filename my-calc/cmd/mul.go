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

// mulCmd represents the mul command
var mulCmd = &cobra.Command{
	Use:   "mul",
	Short: "It's for the multiplication of arguments.",
	Long: `
	the command is useful in multiplication of arguments 
	it will calculate the multiplication for both the 
	Integer and floating point values`,
	Run: func(cmd *cobra.Command, args []string) {
		fval, _ := cmd.Flags().GetBool("float")
		if fval {
			mulFloats(args)
		} else {
			mulInts(args)
		}

	},
}

func init() {
	rootCmd.AddCommand(mulCmd)

	mulCmd.Flags().BoolP("float", "f", false, "Mulitplies float values")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mulCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// mulCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func mulInts(args []string) {
	result := 1
	for _, val := range args {
		ival, err := strconv.Atoi(val)

		if err != nil {
			fmt.Println(err)
		} else {
			result = result * ival
		}
	}

	fmt.Println(result)
}

func mulFloats(args []string) {
	var result float64 = 1
	for _, val := range args {
		ival, err := strconv.ParseFloat(val, 64)

		if err != nil {
			fmt.Println(err)
		} else {
			result = result * ival
		}
	}

	fmt.Println(result)
}
