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

package certificates

import (
	"internal/apiclient"

	"github.com/GoogleCloudPlatform/application-integration-management-toolkit/client/certificates"

	"github.com/spf13/cobra"
)

// DelCmd to get integration flow
var DelCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a certificate from a region",
	Long:  "Delete a certificate from a region",
	Args: func(cmd *cobra.Command, args []string) (err error) {
		if err = apiclient.SetRegion(region); err != nil {
			return err
		}
		return apiclient.SetProjectID(project)
	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		_, err = certificates.Delete(name)
		return

	},
}

func init() {
	DelCmd.Flags().StringVarP(&name, "name", "n",
		"", "Integration flow name")

	_ = DelCmd.MarkFlagRequired("name")
}
