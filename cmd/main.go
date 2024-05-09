/*
Copyright © 2024 Alexandre Pires

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
	"encoding/json"
	"fmt"
	"html"
	"io"
	"os"
	"regexp"
	"strings"
	"text/template"
	"time"

	"github.com/spf13/cobra"
)

// builtins is a map of functions that can be used in the template
var builtins = template.FuncMap{
	"env": os.Getenv,
	// Time functions
	"now": func() string {
		// return current time in RFC3339 format
		return time.Now().Format(time.RFC3339)
	},
	"date": func(format string, t string) string {
		tm, err := time.Parse(time.RFC3339, t)
		if err != nil {
			return ""
		}
		return tm.Format(format)
	},
	// Math functions
	"add": func(b, a float64) float64 {
		return a + b
	},
	"sub": func(b, a float64) float64 {
		return a - b
	},
	"mul": func(b, a float64) float64 {
		return a * b
	},
	"div": func(b, a float64) float64 {
		return a / b
	},
	// String functions
	"upper":       strings.ToUpper,
	"lower":       strings.ToLower,
	"trim":        strings.TrimSpace,
	"trimleft":    func(s string) string { return strings.TrimLeft(s, " ") },
	"trimright":   func(s string) string { return strings.TrimRight(s, " ") },
	"replace":     func(old, new, s string) string { return strings.Replace(s, old, new, -1) },
	"contains":    func(substr, s string) bool { return strings.Contains(s, substr) },
	"hasprefix":   func(prefix, s string) bool { return strings.HasPrefix(s, prefix) },
	"hassuffix":   func(suffix, s string) bool { return strings.HasSuffix(s, suffix) },
	"indexof":     func(substr, s string) int { return strings.Index(s, substr) },
	"lastindexof": func(substr, s string) int { return strings.LastIndex(s, substr) },
	"reverse": func(s string) string {
		runes := []rune(s)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		return string(runes)
	},
	"substr": func(start, count int, s string) string {
		return s[start : start+count]
	},
	"escapeString": func(s string) string {
		return html.EscapeString(s)
	},

	"len": func(a interface{}) int {
		switch a := a.(type) {
		case string:
			return len(a)
		case []string:
			return len(a)
		case map[string]string:
			return len(a)
		default:
			return 0
		}
	},
	"regexFind": func(pattern, s string) string {
		re := regexp.MustCompile(pattern)
		return re.FindString(s)
	},
	"empty": func(a interface{}) bool {
		switch a := a.(type) {
		case string:
			return a == ""
		case []string:
			return len(a) == 0
		case map[string]string:
			return len(a) == 0
		default:
			return true
		}
	},
}

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "gt",
	Args:  cobra.MinimumNArgs(1),
	Short: "A generic go template renderer",
	Long:  `A simple tool that can be used to render go templates from a JSON input.`,
	Run: func(cmd *cobra.Command, args []string) {

		// first argument is the template file
		templateFile := args[0]

		// if there is a second argument, it is the json file otherwise we use
		// input from stdin
		var inputData []byte
		var err error

		if len(args) > 1 {
			inputData, err = os.ReadFile(args[1])
			if err != nil {
				fmt.Println("Error reading json file")
				os.Exit(1)
			}
		} else {
			// Check if there's data available from stdin
			stat, _ := os.Stdin.Stat()
			if (stat.Mode() & os.ModeCharDevice) == 0 {
				inputData, err = io.ReadAll(os.Stdin)
				if err != nil {
					fmt.Println("Error reading input from stdin")
					os.Exit(1)
				}
			} else {
				cmd.Help()
				os.Exit(1)
			}
		}

		// Read the template file
		tplData, err := os.ReadFile(templateFile)
		if err != nil {
			fmt.Println("Error reading template file")
			os.Exit(1)
		}

		output, err := RenderTemplate(tplData, inputData)
		if err != nil {
			fmt.Println("Error rendering template")
			os.Exit(1)
		}

		// Write the output to stdout
		os.Stdout.WriteString(output)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the RootCmd.
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func RenderTemplate(tplData []byte, jsonData []byte) (string, error) {

	// Parse JSON data
	var data interface{}

	if jsonData != nil {

		err := json.Unmarshal(jsonData, &data)
		if err != nil {
			return "", err
		}
	} else {
		data = nil
	}

	// Define template
	tmpl, err := template.New("jsonTemplate").Funcs(builtins).Parse(string(tplData))

	if err != nil {
		return "", err
	}

	// Execute template and write to a string
	result := new(strings.Builder)
	err = tmpl.Execute(result, data)
	if err != nil {
		return "", err
	}

	return result.String(), nil
}

func init() {
	RootCmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		cmd.Println("Usage: gt <template file> [json file]")
	},
	)
}
