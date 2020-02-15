package main

import (
	"github.com/Soontao/yaptranslator"
	"github.com/urfave/cli"
)

var options = []cli.Flag{
	cli.StringFlag{
		Name:  "file, f",
		Usage: "file or directory to be processed",
		Value: ".",
	},
	cli.StringFlag{
		Name:   "lang, l",
		Usage:  "Source Language",
		Value:  "en",
		EnvVar: "YAPT_SOURCE_LANG",
	},
	cli.StringFlag{
		Name:   "target, t",
		Usage:  "Target Language",
		Value:  "zh",
		EnvVar: "YAPT_TARGET_LANG",
	},
	cli.StringFlag{
		Name:   "service, s",
		EnvVar: "YAPT_SERVICE_PROVIDER",
		Usage:  "Service Provider for YAPT Tool",
		Value:  string(yaptranslator.ALICLOUD),
	},
	cli.StringFlag{
		Name:   "username, u",
		EnvVar: "YAPT_AUTH_USER",
		Usage:  "Auth User",
	},
	cli.StringFlag{
		Name:   "password, p",
		EnvVar: "YAPT_AUTH_PASSWORD",
		Usage:  "Auth User Password",
	},
	cli.StringFlag{
		Name:   "region, r",
		EnvVar: "YAPT_REGION",
		Usage:  "AUTH Region",
	},
	cli.StringFlag{
		Name:   "key, k",
		EnvVar: "YAPT_ACCESS_KEY",
		Usage:  "Auth Access Key",
	},
	cli.StringFlag{
		Name:   "secret",
		EnvVar: "YAPT_ACCESS_SECRET",
		Usage:  "Auth Access Secret",
	},
	cli.StringFlag{
		Name:   "token",
		EnvVar: "YAPT_AUTH_TOKEN",
		Usage:  "Auth Token",
	},
}
