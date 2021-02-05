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
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// hmac256Cmd represents the hmac256 command
var hmac256Cmd = &cobra.Command{
	Use:   "hmac256",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		secret, _ := cmd.Flags().GetString("secret")
		plaintext, _ := cmd.Flags().GetString("text")
		textcase, _ := cmd.Flags().GetString("o")
		//fmt.Printf("secret: %s, text:%s\n", secret, plaintext)
		mac := hmac.New(sha256.New, []byte(secret))
		mac.Write([]byte(plaintext))
		hash := fmt.Sprintf("%x", mac.Sum(nil))
		if textcase == "u" {
			fmt.Println(strings.ToUpper(hash))
		} else {
			fmt.Println(hash)
		}
	},
}

func init() {
	rootCmd.AddCommand(hmac256Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// hmac256Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// hmac256Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	hmac256Cmd.Flags().String("secret", "", "Your secret key")
	hmac256Cmd.MarkFlagRequired("secret")
	hmac256Cmd.Flags().String("text", "", "Your text to hash (enclosed with quotes if contains spaces)")
	hmac256Cmd.Flags().String("o", "u", "output format l=lowercase, u=uppercase")
}
