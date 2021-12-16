package cmd

import (
	"fmt"
	"github.com/NICEXAI/go2struct-tool/internal/conver"
	"github.com/NICEXAI/go2struct-tool/internal/errorx"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"os"
	"runtime"
)

const version = "0.1.0"

var (
	inputFilePath  string
	outputFilePath string
	moduleName     string
	fieldTagName   string
	watchMode      bool
)

var rootCmd = &cobra.Command{
	Use:   "go2struct",
	Short: "lazy-go is an Easy-to-use format conversion tool.",
	Long:  "Use the command to convert arbitrary formats to Go Struct (including json, toml, yaml, etc.)",
	Run: func(cmd *cobra.Command, args []string) {
		if inputFilePath == "" || outputFilePath == "" {
			if err := cmd.Help(); err != nil {
				color.Red(err.Error())
			}
			return
		}

		if err := conver.Convert(inputFilePath, outputFilePath, moduleName, fieldTagName); err != nil {
			color.Red(err.Error())
			return
		}

		color.Green("file convert success")

		if watchMode {
			fsTask, err := conver.Watch(inputFilePath, outputFilePath, moduleName, fieldTagName)
			if err != nil {
				color.Red("%v: %s", errorx.ErrCovertFailed, err.Error())
				return
			}

			color.Blue("file listening...")
			fsTask.Wait()
		}
	},
	Version: fmt.Sprintf("%s %s/%s", version, runtime.GOOS, runtime.GOARCH),
}

func init() {
	rootCmd.Flags().StringVarP(&inputFilePath, "input", "i", "", "source file path")
	rootCmd.Flags().StringVarP(&outputFilePath, "output", "o", "", "target file path")
	rootCmd.Flags().StringVarP(&moduleName, "module", "m", "", "module name of the target file, default: target file name")
	rootCmd.Flags().StringVarP(&fieldTagName, "tag", "t", "json", "set filed tag name, default: json")
	rootCmd.Flags().BoolVarP(&watchMode, "watch", "w", false, "listening file changes and auto convert content to Go Struct")
}

// Execute entrypoint
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
