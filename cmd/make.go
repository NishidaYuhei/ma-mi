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
	"os/exec"
	"strings"

	"github.com/bellocha/ma-mi/lib"
	"github.com/spf13/cobra"
)

var (
	HomePath      = lib.GetHomePath()
	ConfigPath    = HomePath + "/.ma-mi/config"
	ConfigSetPath = lib.ReadFile(ConfigPath)
	MamiDirPath   = ConfigSetPath + "/" + lib.GetNowYearAndMonth()
	TempDirPath   = strings.TrimSuffix(ConfigPath, "config") + "temp"
)

// makeCmd represents the make command
var makeCmd = &cobra.Command{
	Use:   "make",
	Short: "make minutes command",
	Long:  `make minutes command`,
	Run: func(cmd *cobra.Command, args []string) {
		lib.MkdirAll(MamiDirPath)
		lib.MkdirAll(TempDirPath)
		lib.CreateNewFile(TempDirPath+"/default.md", "")
		editorCmd := exec.Command(lib.ReadFile(SetEditorConfPath), MamiDirPath+"/"+args[0]+".md")
		switch len(args) {
		case 1:
			lib.CreateNewFile(MamiDirPath+"/"+args[0]+".md", lib.ReadFile(TempDirPath+"/default.md"))
			fmt.Println(lib.ReadFile(SetEditorConfPath) + " " + MamiDirPath + "/" + args[0] + ".md")
			editorCmd.Run()
			break
		case 2:
			lib.CreateNewFile(MamiDirPath+"/"+args[0]+".md", lib.ReadFile(TempDirPath+"/"+args[1]+".md"))
			fmt.Println(lib.ReadFile(SetEditorConfPath), MamiDirPath+"/"+args[0]+".md")
			editorCmd.Run()
			break
		}
	},
}

func init() {
	rootCmd.AddCommand(makeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// makeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// makeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
