// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License

package videostitcher

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/GoogleCloudPlatform/golang-samples/internal/testutil"
)

func TestListVodStitchDetails(t *testing.T) {
	tc := testutil.SystemTest(t)
	var buf bytes.Buffer

	sessionID := createTestVodSession(t)
	stitchDetailsNamePrefix := fmt.Sprintf("/locations/%s/vodSessions/%s/vodStitchDetails/", location, sessionID)

	testutil.Retry(t, 3, 2*time.Second, func(r *testutil.R) {
		if err := listVodStitchDetails(&buf, tc.ProjectID, sessionID); err != nil {
			r.Errorf("listVodStitchDetails got err: %v", err)
		}
		if got := buf.String(); !strings.Contains(got, stitchDetailsNamePrefix) {
			r.Errorf("listVodStitchDetails got: %v Want to contain: %v", got, stitchDetailsNamePrefix)
		}
	})
}
