package main

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/urfave/cli"
)

// Version for command line
var Version string

func main() {
	app := cli.NewApp()
	app.Name = "facebook plugin"
	app.Usage = "facebook plugin"
	app.Action = run
	app.Version = Version
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "page.token",
			Usage:  "facebook page token",
			EnvVar: "PLUGIN_FB_PAGE_TOKEN,FB_PAGE_TOKEN",
		},
		cli.StringFlag{
			Name:   "verify.token",
			Usage:  "facebook verify token",
			EnvVar: "PLUGIN_FB_VERIFY_TOKEN,FB_VERIFY_TOKEN",
		},
		cli.BoolFlag{
			Name:   "verify",
			Usage:  "verifying webhooks on the Facebook Developer Portal",
			EnvVar: "PLUGIN_VERIFY",
		},
		cli.StringSliceFlag{
			Name:   "to",
			Usage:  "send message to user",
			EnvVar: "PLUGIN_TO",
		},
		cli.StringSliceFlag{
			Name:   "message",
			Usage:  "facebook message",
			EnvVar: "PLUGIN_MESSAGE",
		},
		cli.StringFlag{
			Name:   "repo.owner",
			Usage:  "repository owner",
			EnvVar: "DRONE_REPO_OWNER",
		},
		cli.StringFlag{
			Name:   "repo.name",
			Usage:  "repository name",
			EnvVar: "DRONE_REPO_NAME",
		},
		cli.StringFlag{
			Name:   "commit.sha",
			Usage:  "git commit sha",
			EnvVar: "DRONE_COMMIT_SHA",
		},
		cli.StringFlag{
			Name:   "commit.branch",
			Value:  "master",
			Usage:  "git commit branch",
			EnvVar: "DRONE_COMMIT_BRANCH",
		},
		cli.StringFlag{
			Name:   "commit.author",
			Usage:  "git author name",
			EnvVar: "DRONE_COMMIT_AUTHOR",
		},
		cli.StringFlag{
			Name:   "commit.message",
			Usage:  "commit message",
			EnvVar: "DRONE_COMMIT_MESSAGE",
		},
		cli.StringFlag{
			Name:   "build.event",
			Value:  "push",
			Usage:  "build event",
			EnvVar: "DRONE_BUILD_EVENT",
		},
		cli.IntFlag{
			Name:   "build.number",
			Usage:  "build number",
			EnvVar: "DRONE_BUILD_NUMBER",
		},
		cli.StringFlag{
			Name:   "build.status",
			Usage:  "build status",
			Value:  "success",
			EnvVar: "DRONE_BUILD_STATUS",
		},
		cli.StringFlag{
			Name:   "build.link",
			Usage:  "build link",
			EnvVar: "DRONE_BUILD_LINK",
		},
	}
	app.Run(os.Args)
}

func run(c *cli.Context) error {
	plugin := Plugin{
		Repo: Repo{
			Owner: c.String("repo.owner"),
			Name:  c.String("repo.name"),
		},
		Build: Build{
			Number:  c.Int("build.number"),
			Event:   c.String("build.event"),
			Status:  c.String("build.status"),
			Commit:  c.String("commit.sha"),
			Branch:  c.String("commit.branch"),
			Author:  c.String("commit.author"),
			Message: c.String("commit.message"),
			Link:    c.String("build.link"),
		},
		Config: Config{
			PageToken:   c.String("page.token"),
			VerifyToken: c.String("verify.token"),
			Verify:      c.Bool("verify"),
			To:          c.StringSlice("to"),
			Message:     c.StringSlice("message"),
		},
	}

	return plugin.Exec()
}