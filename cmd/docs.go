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

	"github.com/Vardhanb07/envbuild/internal/dir"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

var docsCmd = &cobra.Command{
	Use:   "docs",
	Short: "Generate docs",
	Example: `
- envbuild docs

Generates docs in ./docs.

- envbuild docs --directory <dir>

Generates docs in <dir>.

- envbuild docs -d <dir>

Generates docs in <dir>.`,
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		d := cmd.Flag("directory")
		path, err := dir.Resolve(d.Value.String())
		if err != nil {
			return err
		}
		if err := os.MkdirAll(path, 0777); err != nil {
			return err
		}
		return doc.GenMarkdownTree(cmd.Root(), path)
	},
}

func init() {
	rootCmd.AddCommand(docsCmd)

	docsCmd.Flags().StringP("directory", "d", "./docs", "generated docs are placed in this directory")
}
