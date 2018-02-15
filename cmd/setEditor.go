// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
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
	"fmt"
	"log"
	"strings"

	"github.com/bellocha/ma-mi/lib"
	"github.com/spf13/cobra"
)

var (
	SetEditorConfPath = strings.TrimSuffix(ConfigPath, "config") + "/set_editor_conf"
)

// setEditorCmd represents the setEditor command
var setEditorCmd = &cobra.Command{
	Use:   "set-editor",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var editorPath string
		fmt.Println("Please default editor path.")
		fmt.Println("If you do not know the editor path...")
		fmt.Println("try: [which editor_command]")
		fmt.Println("example: which code")
		fmt.Scanf("%v", &editorPath)
		lib.RemoveFile(SetEditorConfPath)
		lib.CreateNewFile(SetEditorConfPath, editorPath)
		log.Printf("register editor %v", editorPath)
	},
}

func init() {
	rootCmd.AddCommand(setEditorCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setEditorCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setEditorCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
