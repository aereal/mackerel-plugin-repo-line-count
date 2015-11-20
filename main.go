package main

import (
	"fmt"
	"os"
	"time"

	"github.com/codegangsta/cli"
)

type Param struct {
	repoName string
	repoPath string
	pattern  string
}

var Version string

var globalFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "path",
		Value: "",
		Usage: "repository path",
	},
	cli.StringFlag{
		Name:  "name",
		Value: "",
		Usage: "repository name",
	},
	cli.StringFlag{
		Name:  "pattern",
		Value: "",
		Usage: "search pattern",
	},
}

func main() {
	app := cli.NewApp()
	app.Name = "mackerel-plugin-repo-line-count"
	app.Version = Version
	app.Author = "aereal"
	app.Flags = globalFlags
	app.Action = invoke

	app.Run(os.Args)
}

func invoke(c *cli.Context) {
	param := &Param{}

	param.repoName = c.String("name")
	param.repoPath = c.String("path")
	param.pattern = c.String("pattern")

	run(param, c)
}

func run(param *Param, c *cli.Context) {
	result := Grep(param.repoPath, param.pattern)
	now := time.Now().Unix()
	for _, entry := range result {
		metricName := fmt.Sprintf("repo_lines.%s.%s", param.pattern, entry.SafeMetricNameComponent())
		fmt.Printf("%s\t%d\t%d\n", metricName, entry.Count, now)
	}
}
