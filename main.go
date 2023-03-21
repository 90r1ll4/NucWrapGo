package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"wrapper"

	"github.com/projectdiscovery/goflags"
)

func nucleiCheck() bool {
	// To check if nuclei configured or not
	command := "nuclei"
	_, err := exec.Command(command, "-silent", "--version").Output()
	// fmt.Println(outputs)
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			if exitError.ExitCode() == 127 {
				fmt.Printf("Command '%s' not found.", command)
				return false
			}
		} else {
			fmt.Println("Nuclei not found")
			return false
		}
	}

	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

	return true
}

type Options struct {
	url     string
	urlFile string
	output  string
	json    bool
	tables  bool
	// outputFile "output.txt"
}

func main() {
	opt := &Options{}
	flagSet := goflags.NewFlagSet()
	flagSet.SetDescription("Wrapper For Nuclei")
	flagSet.StringVarP(&opt.url, "url", "u", "", "url to scan")
	flagSet.StringVarP(&opt.urlFile, "url_file", "f", "", "list of urls")
	flagSet.StringVarP(&opt.output, "output", "o", "", "Output in text form[tables][json]")
	// flagSet.StringVarP(&opt.outputFile, "output File", "of", opt.outputFile, "Output file name")
	flagSet.BoolVarP(&opt.json, "json", "", false, "Output in json form")
	flagSet.BoolVarP(&opt.tables, "tables", "", false, "Output in table form")
	if err := flagSet.Parse(); err != nil {
		log.Fatalf("Could not parse flags: %s\n", err)
	}

	if opt.url == "" && opt.urlFile == "" {
		fmt.Println("Error: URL or URL File is required")
		return
	}

	wrapper.Nuclei(opt.url, opt.urlFile, opt.json, opt.tables)

}
