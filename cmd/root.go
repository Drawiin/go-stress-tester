/*
Copyright © 2024 Vinicius S. Guimarães drawiinapps@gmail.com

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
	"fmt"
	"go-stress-tester/internal/usecase"
	"os"

	"github.com/spf13/cobra"
)

var url string
var requests, concurrency int
var runner usecase.StressTestRunner

var rootCmd = &cobra.Command{
	Use:   "run",
	Short: "Will run the stress test",
	Long:  `Will run the stress test with the given parameters.`,
	Run: func(cmd *cobra.Command, args []string) {
		result := runner.RunStressTest(usecase.ConfigDto{
			Url:         url,
			Requests:    requests,
			Concurrency: concurrency,
		})
		fmt.Println("Result:", result)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(stressTestRunner usecase.StressTestRunner) {
	runner = stressTestRunner
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&url, "url", "u", "", "Defines the url used for the stress test")
	rootCmd.Flags().IntVarP(&requests, "requests", "r", 0, "Defines the number of requests used for the stress test")
	rootCmd.Flags().IntVarP(&concurrency, "concurrency", "c", 0, "Defines the concurrency used for the stress test")
	rootCmd.MarkFlagRequired("url")
	rootCmd.MarkFlagRequired("requests")
	rootCmd.MarkFlagRequired("concurrency")
}
