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

// [START metastore_v1alpha_generated_DataprocMetastore_DeleteService_sync]

package main

import (
	"context"

	metastore "cloud.google.com/go/metastore/apiv1alpha"
	metastorepb "google.golang.org/genproto/googleapis/cloud/metastore/v1alpha"
)

func main() {
	// import metastorepb "google.golang.org/genproto/googleapis/cloud/metastore/v1alpha"

	ctx := context.Background()
	c, err := metastore.NewDataprocMetastoreClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &metastorepb.DeleteServiceRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.DeleteService(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	err = op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
}

// [END metastore_v1alpha_generated_DataprocMetastore_DeleteService_sync]