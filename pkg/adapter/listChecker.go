// Copyright 2016 Google Inc.
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
// limitations under the License.

package adapter

type (
	// ListCheckerAspect checks the presence of a given symbol against a list.
	ListCheckerAspect interface {
		Aspect

		// CheckList verifies whether the given symbol is on the list.
		CheckList(symbol string) (bool, error)
	}

	// ListCheckerBuilder builds instances of the ListChecker aspect.
	ListCheckerBuilder interface {
		Builder

		// NewListChecker returns a new instance of the ListChecker aspect.
		NewListChecker(env Env, c AspectConfig) (ListCheckerAspect, error)
	}
)
