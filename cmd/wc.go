/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var char, bytes, line, max_line, help_ref bool

//help document for uses like no commands, user guide and default case
var help_document = `Usage: wc [OPTION]... [FILE]...
or:  wc [OPTION]... --files0-from=F
Print newline, word, and byte counts for each FILE, and a total line if
more than one FILE is specified.  A word is a non-zero-length sequence of
characters delimited by white space.
With no FILE, or when FILE is -, read standard input.
The options below may be used to select which counts are printed, always in
the following order: newline, word, character, byte, maximum line length.
-c, --bytes        	print the byte counts
-m, --chars        	print the character counts
-l, --lines        	print the newline counts
	--files0-from=F	read input from the files specified by
						 NUL-terminated names in file F;
						 If F is - then read names from standard input
-L, --max-line-length  print the maximum display width
-w, --words        	print the word counts
	--help 	display this help and exit
	--version  output version information and exit
GNU coreutils online help: <https://www.gnu.org/software/coreutils/>
Full documentation <https://www.gnu.org/software/coreutils/wc>
or available locally via: info '(coreutils) wc invocation'
`

// wcCmd represents the wc command
var wcCmd = &cobra.Command{
	Use:   "wc",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,

	// switch in accordance with the functions described below
	Run: func(cmd *cobra.Command, args []string) {
		switch {

		case char:
			fmt.Println(CharCount(args[0]))
		case bytes:
			fmt.Println(ByteCount(args[0]))
		case line:
			fmt.Println(LineCount(args[0]))
		case max_line:
			fmt.Println(MaxLine(args[0]))
		case help_ref:
			fmt.Println(help_document)
		default:
			fmt.Println(help_document)

		}
		//fmt.Println("wc called")
		//CharCount(args[0])
		//LineCount(args[0])
		//MaxLine(args[0])
		//ByteCount(args[0])
	},
}

//function to count number of characters in the input file
func CharCount(name string) (int, error) {
	file, err := os.Open(name)
	if err != nil {
		//fmt.Println("Error", err)
		return -1, fmt.Errorf("unable to process due to %v", err)
	}
	chars_count := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		chars_count += len(scanner.Text())
	}
	return chars_count, err
}

//function to count number of lines in the input file
func LineCount(name string) (int, error) {
	file, err := os.Open(name)
	if err != nil {
		//fmt.Println("Error", err)
		return -1, fmt.Errorf("unable to process due to %v", err)
	}
	lines_count := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines_count += len(scanner.Text())

	}

	return lines_count, err

}

//function to get the maximum line width
func MaxLine(name string) (int, error) {
	file, err := os.Open(name)
	if err != nil {
		//fmt.Println("Error", err)
		return -1, fmt.Errorf("unable to process due to %v", err)
	}
	//char_count := 0
	max_len := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		x := scanner.Text()
		if len(x) > max_len {
			max_len = len(scanner.Text())
		}

	}
	return max_len, err

}

//function to get the byte size of the input file
func ByteCount(name string) (int, error) {
	file, err := os.Stat(name)
	if err != nil {
		//fmt.Println("Error", err)
		return -1, fmt.Errorf("unable to process due to %v", err)

	}
	byte_size := int(file.Size())
	return byte_size, err
}

//flag definitions as per accordance to the help document

func init() {

	rootCmd.AddCommand(wcCmd)

	wcCmd.Flags().BoolVarP(&char, "chars_count", "m", false, "Display the number of characters")
	wcCmd.Flags().BoolVarP(&line, "lines_count", "l", false, "Display number of lines")
	wcCmd.Flags().BoolVarP(&max_line, "max_len", "L", false, "Display the length of the line having maximum length ")
	wcCmd.Flags().BoolVarP(&bytes, "byte_size", "c", false, "Display the size of the file")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// wcCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// wcCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
