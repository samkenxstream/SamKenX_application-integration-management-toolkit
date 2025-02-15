// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package integrations

import (
	"internal/apiclient"

	"github.com/GoogleCloudPlatform/application-integration-management-toolkit/client/integrations"

	"github.com/spf13/cobra"
)

// ExportVerCmd to export integrations
var ExportVerCmd = &cobra.Command{
	Use:   "export",
	Short: "Export Integrations flow versions to a folder",
	Long:  "Export Integrations flow versions to a folder",
	Args: func(cmd *cobra.Command, args []string) (err error) {
		if err = apiclient.SetRegion(region); err != nil {
			return err
		}
		return apiclient.SetProjectID(project)
	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		if err = apiclient.FolderExists(folder); err != nil {
			return err
		}

		apiclient.SetExportToFile(folder)
		_, err = integrations.ListVersions(name, -1, "", "", "", true, true, false)
		return err
	},
}

var folder string
var allVersions bool
var numConnections int

func init() {
	ExportVerCmd.Flags().StringVarP(&folder, "folder", "f",
		"", "Folder to export Integration flows")
	ExportVerCmd.Flags().StringVarP(&name, "name", "n",
		"", "Integration flow name")

	_ = ExportVerCmd.MarkFlagRequired("folder")
	_ = ExportVerCmd.MarkFlagRequired("name")
}
