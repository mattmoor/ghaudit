// Copyright 2024 Chainguard, Inc.
// SPDX-License-Identifier: Apache-2.0

package org

import (
	"github.com/chainguard-dev/ghaudit/pkg/repo"
	"github.com/google/go-github/v60/github"
	"github.com/spf13/cobra"
)

func deployKeys(ghc *github.Client, org *string) *cobra.Command {
	rm := NewRepoMapper(ghc, org, repo.DeployKeys)

	return &cobra.Command{
		Use:           "deploy-keys",
		Short:         "Audit for usage of deploy keys.",
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return rm.Execute(cmd.Context())
		},
	}
}
