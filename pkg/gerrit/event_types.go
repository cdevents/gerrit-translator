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

// event fields and structures

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
type PatchSet struct {
	Number         int       `json:"number"`
	Revision       string    `json:"revision"`
	Parents        []string  `json:"parents"`
	Ref            string    `json:"ref"`
	Uploader       Submitter `json:"uploader"`
	CreatedOn      float64   `json:"createdOn"`
	Author         Submitter `json:"author"`
	Kind           string    `json:"kind"`
	SizeInsertions int       `json:"sizeInsertions"`
	SizeDeletions  int       `json:"sizeDeletions"`
}

type Change struct {
	Project       string    `json:"project"`
	Branch        string    `json:"branch"`
	Id            string    `json:"id"`
	Number        int       `json:"number"`
	Subject       string    `json:"subject"`
	Owner         Submitter `json:"owner"`
	Url           string    `json:"url"`
	CommitMessage string    `json:"commitMessage"`
	CreatedOn     float64   `json:"createdOn"`
	Status        string    `json:"status"`
}

type Project struct {
	Name string `json:"name"`
}
type ChangeKey struct {
	Key string `json:"key"`
}
type CommonChangeFields struct {
	PatchSet  PatchSet  `json:"patchSet"`
	Change    Change    `json:"change"`
	Project   Project   `json:"project"`
	RefName   string    `json:"refName"`
	ChangeKey ChangeKey `json:"changeKey"`
}
type Approval struct {
	Type        string `json:"type"`
	Description string `json:"description"`
	Value       string `json:"value"`
}

// Gerrit event types

// ProjectCreated project-created gerrit event structure
type ProjectCreated struct {
	ProjectName string `json:"projectName"`
	HeadName    string `json:"headName"`
	CommonFields
}

// ProjectHeadUpdated project-head-updated gerrit event structure
type ProjectHeadUpdated struct {
	ProjectName string `json:"projectName"`
	OldHead     string `json:"oldHead"`
	NewHead     string `json:"newHead"`
	CommonFields
}

// RefUpdated ref-updated gerrit event structure
type RefUpdated struct {
	Submitter Submitter `json:"submitter"`
	RefUpdate RefUpdate `json:"refUpdate"`
	CommonFields
}

// PatchsetCreated patchset-created gerrit event structure
type PatchsetCreated struct {
	Uploader Submitter `json:"submitter"`
	CommonChangeFields
	CommonFields
}

// CommentAdded comment-added gerrit event structure
type CommentAdded struct {
	Author    Submitter  `json:"author"`
	Approvals []Approval `json:"approvals"`
	Comment   string     `json:"comment"`
	CommonChangeFields
	CommonFields
}

// ChangeMerged change-merged gerrit event structure
type ChangeMerged struct {
	Submitter Submitter `json:"submitter"`
	NewRev    string    `json:"newRev"`
	CommonChangeFields
	CommonFields
}

// ChangeAbandoned change-abandoned gerrit event structure
type ChangeAbandoned struct {
	Abandoner Submitter `json:"abandoner"`
	Reason    string    `json:"reason"`
	CommonChangeFields
	CommonFields
}
