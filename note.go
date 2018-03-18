package main

import (
	"flag"
	"fmt"
	"os"
)

func usage() {
	fmt.Fprintf(os.Stderr,
		"Usage: %s [ options ]\noptions:\n", os.Args[0])
	flag.PrintDefaults()
}

type cmdLine interface {
	parseCommandLine()
}

type config struct {
	dataDir     string
	templateDir string
	host        string
	port        int
}

func (c *config) parseCommandLine() {
	flag.Usage = usage // override default usage printout
	pHost := flag.String("host", "0.0.0.0", "server listening to")
	pPort := flag.Int("port", 8080, "Server's port")
	pTemplateDir := flag.String("template_dir", "template", "template directory")
	pDataDir := flag.String("data_dir", "data", "Data directory")
	flag.Parse()

	c.dataDir = *pDataDir
	c.templateDir = *pTemplateDir
	c.host = *pHost
	c.port = *pPort
}

var appConfig config

func init() {
	appConfig = config{}
	appConfig.parseCommandLine()

	serverInit()
	persistInit()
}

func main() {
	serverRun(appConfig.host, appConfig.port)
}
