// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package attributes

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	conventions "go.opentelemetry.io/collector/semconv/v1.6.1"
)

func TestProcessExtractTags(t *testing.T) {
	pattrs := processAttributes{
		ExecutableName: "otelcol",
		ExecutablePath: "/usr/bin/cmd/otelcol",
		Command:        "cmd/otelcol",
		CommandLine:    "cmd/otelcol --config=\"/path/to/config.yaml\"",
		PID:            1,
		Owner:          "root",
	}

	assert.Equal(t, []string{
		fmt.Sprintf("%s:%s", conventions.AttributeProcessExecutableName, "otelcol"),
	}, pattrs.extractTags())

	pattrs = processAttributes{
		ExecutablePath: "/usr/bin/cmd/otelcol",
		Command:        "cmd/otelcol",
		CommandLine:    "cmd/otelcol --config=\"/path/to/config.yaml\"",
		PID:            1,
		Owner:          "root",
	}

	assert.Equal(t, []string{
		fmt.Sprintf("%s:%s", conventions.AttributeProcessExecutablePath, "/usr/bin/cmd/otelcol"),
	}, pattrs.extractTags())

	pattrs = processAttributes{
		Command:     "cmd/otelcol",
		CommandLine: "cmd/otelcol --config=\"/path/to/config.yaml\"",
		PID:         1,
		Owner:       "root",
	}

	assert.Equal(t, []string{
		fmt.Sprintf("%s:%s", conventions.AttributeProcessCommand, "cmd/otelcol"),
	}, pattrs.extractTags())

	pattrs = processAttributes{
		CommandLine: "cmd/otelcol --config=\"/path/to/config.yaml\"",
		PID:         1,
		Owner:       "root",
	}

	assert.Equal(t, []string{
		fmt.Sprintf("%s:%s", conventions.AttributeProcessCommandLine, "cmd/otelcol --config=\"/path/to/config.yaml\""),
	}, pattrs.extractTags())
}

func TestProcessExtractTagsEmpty(t *testing.T) {
	pattrs := processAttributes{}

	assert.Equal(t, []string{}, pattrs.extractTags())
}
