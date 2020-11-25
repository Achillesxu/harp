// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package bundle

import (
	"context"
	"fmt"
	"regexp"

	"github.com/jmespath/go-jmespath"

	bundlev1 "github.com/elastic/harp/api/gen/go/harp/bundle/v1"
	"github.com/elastic/harp/pkg/bundle"
	"github.com/elastic/harp/pkg/bundle/selector"
	"github.com/elastic/harp/pkg/tasks"
)

// FilterTask implements secret container filtering task.
type FilterTask struct {
	ContainerReader tasks.ReaderProvider
	OutputWriter    tasks.WriterProvider
	KeepPaths       []string
	ExcludePaths    []string
	JMESPath        string
}

// Run the task.
//nolint:gocognit,gocyclo // to refactor
func (t *FilterTask) Run(ctx context.Context) error {
	// Create input reader
	reader, err := t.ContainerReader(ctx)
	if err != nil {
		return fmt.Errorf("unable to open input bundle: %w", err)
	}

	// Load bundle
	b, err := bundle.FromContainerReader(reader)
	if err != nil {
		return fmt.Errorf("unable to load bundle content: %w", err)
	}

	// Clean up bundle
	if len(t.KeepPaths) > 0 {
		pkgs := []*bundlev1.Package{}

		for _, includePath := range t.KeepPaths {
			includePathRegexp, errInclude := regexp.Compile(includePath)
			if errInclude != nil {
				return fmt.Errorf("unable to compile keep regexp '%s': %w", includePath, err)
			}

			for _, p := range b.Packages {
				if includePathRegexp.MatchString(p.Name) {
					pkgs = append(pkgs, p)
				}
			}
		}

		// Override packages
		b.Packages = pkgs
	}

	if len(t.ExcludePaths) > 0 {
		pkgs := []*bundlev1.Package{}

		for _, excludePath := range t.ExcludePaths {
			excludePathRegexp, errExclude := regexp.Compile(excludePath)
			if errExclude != nil {
				return fmt.Errorf("unable to compile exclusion regexp '%s': %w", excludePath, err)
			}

			for _, p := range b.Packages {
				if !excludePathRegexp.MatchString(p.Name) {
					pkgs = append(pkgs, p)
				}
			}
		}

		// Override packages
		b.Packages = pkgs
	}

	// Has JMESPath filter
	if t.JMESPath != "" {
		pkgs := []*bundlev1.Package{}

		// Compile expression first
		exp, errJMESPath := jmespath.Compile(t.JMESPath)
		if err != nil {
			return fmt.Errorf("unable to compile JMESPath filter '%s': %w", t.JMESPath, errJMESPath)
		}

		// Initialize selector
		s := selector.MatchJMESPath(exp)

		// Apply package filtering
		for _, p := range b.Packages {
			if s.IsSatisfiedBy(p) {
				pkgs = append(pkgs, p)
			}
		}

		// Override packages
		b.Packages = pkgs
	}

	// Create output writer
	writer, err := t.OutputWriter(ctx)
	if err != nil {
		return fmt.Errorf("unable to open output bundle: %w", err)
	}

	// Dump all content
	if err := bundle.ToContainerWriter(writer, b); err != nil {
		return fmt.Errorf("unable to dump bundle content: %w", err)
	}

	// No error
	return nil
}
