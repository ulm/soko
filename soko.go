package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"soko/pkg/app"
	"soko/pkg/config"
	"soko/pkg/logger"
	"soko/pkg/portage"
	"soko/pkg/portage/pkgcheck"
	"soko/pkg/portage/repology"
	"time"
)

func printHelp() {
	fmt.Println("Please specific one of the following options:")
	fmt.Println("  soko update                   -- incrementally update the database")
	fmt.Println("  soko fullupdate               -- update the database ")
	fmt.Println("  soko update-outdated-packages -- update the database containing all outdated gentoo packages")
	fmt.Println("  soko update-pgkcheck-results  -- update the database containing all pkgcheck results")
	fmt.Println("  soko serve                    -- serve the application")
}

func isCommand(command string) bool {
	return len(os.Args) > 1 && os.Args[1] == command
}

func main() {

	waitForPostgres()

	errorLogFile := logger.CreateLogFile(config.LogFile())
	defer errorLogFile.Close()
	initLoggers(os.Stdout, errorLogFile)

	if isCommand("serve") {
		app.Serve()
	} else if isCommand("update") {
		portage.Update()
	} else if isCommand("fullupdate") {
		portage.FullUpdate()
	} else if isCommand("update-outdated-packages") {
		repology.UpdateOutdated()
	} else if isCommand("update-pgkcheck-results") {
		pkgcheck.UpdatePkgCheckResults()
	} else {
		printHelp()
	}

}

// initialize the loggers depending on whether
// config.debug is set to true
func initLoggers(infoHandler io.Writer, errorHandler io.Writer) {
	if config.Debug() == "true" {
		logger.Init(os.Stdout, infoHandler, errorHandler)
	} else {
		logger.Init(ioutil.Discard, infoHandler, errorHandler)
	}
}

// TODO this has to be solved differently
// wait for postgres to come up
func waitForPostgres() {
	time.Sleep(5 * time.Second)
}
