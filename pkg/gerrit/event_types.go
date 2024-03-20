/*
Copyright (C) 2024 Nordix Foundation.
For a full list of individual contributors, please see the commit history.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
SPDX-License-Identifier: Apache-2.0
*/

package gerrit

type CommonFields struct {
	Type           string  `json:"type"`
	EventCreatedOn float64 `json:"eventCreatedOn"`
	RepoURL        string  `json:"repoURL,omitempty"`
}

type Submitter struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
}
type RefUpdate struct {
	OldRev  string `json:"oldRev"`
	NewRev  string `json:"newRev"`
	RefName string `json:"refName"`
	Project string `json:"project"`
}

// Gerrit event types

type ProjectCreated struct {
	ProjectName string `json:"projectName"`
	HeadName    string `json:"headName"`
	CommonFields
}
type ProjectHeadUpdated struct {
	ProjectName string `json:"projectName"`
	OldHead     string `json:"oldHead"`
	NewHead     string `json:"newHead"`
	CommonFields
}
type RefUpdated struct {
	Submitter Submitter `json:"submitter"`
	RefUpdate RefUpdate `json:"refUpdate"`
	CommonFields
}
