// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package workspacemode

import (
	"fmt"

	"github.com/daytonaio/daytona/internal/util"
	"github.com/daytonaio/daytona/internal/util/apiclient"
	workspace_cmd "github.com/daytonaio/daytona/pkg/cmd/workspace"
	"github.com/daytonaio/daytona/pkg/views"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:     "stop",
	Short:   "Stop the project",
	Args:    cobra.NoArgs,
	GroupID: util.WORKSPACE_GROUP,
	Run: func(cmd *cobra.Command, args []string) {
		apiClient, err := apiclient.GetApiClient(nil)
		if err != nil {
			log.Fatal(err)
		}

		err = workspace_cmd.StopWorkspace(apiClient, workspaceId, projectName)
		if err != nil {
			log.Fatal(err)
		}

		if projectName != "" {
			views.RenderInfoMessage(fmt.Sprintf("Project '%s' from workspace '%s' successfully stopped", projectName, workspaceId))
		} else {
			views.RenderInfoMessage(fmt.Sprintf("Workspace '%s' successfully stopped", workspaceId))
		}
	},
}
