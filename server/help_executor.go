package server

import (
	"github.com/hootsuite/atlantis/github"
	"github.com/spf13/viper"
)

type HelpExecutor struct {
	github *github.Client
}

var helpComment = "```cmake\n" +
	`atlantis - Terraform collaboration tool that enables you to collaborate on infrastructure
safely and securely. (v` + viper.GetString("version") + `)

Usage: atlantis <command> [environment] [--verbose]

Commands:
plan           Runs 'terraform plan' on the files changed in the pull request
apply          Runs 'terraform apply' using the plans generated by 'atlantis plan'
help           Get help

Examples:

# Generates a plan for staging environment
atlantis plan staging

# Generates a plan for a standalone terraform project
atlantis plan

# Applies a plan for staging environment
atlantis apply staging

# Applies a plan for a standalone terraform project
atlantis apply
`

func (h *HelpExecutor) execute(ctx *CommandContext) {
	ctx.Log.Info("generating help comment....")
	h.github.CreateComment(ctx.BaseRepo, ctx.Pull, helpComment)
	return
}
