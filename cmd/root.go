/*
Copyright © 2026 Vardhan Battula <vardhanbattula7@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"os"

	"github.com/Vardhanb07/envbuild/internal/config"
	"github.com/Vardhanb07/envbuild/internal/dir"
	"github.com/Vardhanb07/envbuild/internal/env"
	"github.com/spf13/cobra"
)

var envFile string

var rootCmd = &cobra.Command{
	Use:   "envbuild",
	Short: "Build .env from a schema",
	Long: `Envbuild allows you to build your .env files from a schema.
Envbuild generates .env file using the env vars or by a key value store.`,
	Example: `
- envbuild <path>

This looks .envbuild.toml files in <path> and builds .env in it.

- envbuild --env-file <file-path>

This loads <file-path> and builds .env in the base path.`,
	Version: "0.1.0",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		path, err := dir.Resolve(envFile)
		if err != nil {
			return err
		}
		cfg, err := config.Read(path)
		if err != nil {
			return err
		}
		wd, err := os.Getwd()
		return env.Build(cfg, wd)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&envFile, "env-file", "e", ".envbuild.toml", "env file path")
}
