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

import (
	"encoding/json"
	"errors"
	"strings"
)

func (pEvent *GerritEvent) HandleProjectCreatedEvent() (string, error) {
	var projectCreated ProjectCreated
	err := json.Unmarshal([]byte(pEvent.Event), &projectCreated)
	if err != nil {
		Log().Error("Error occurred while Unmarshal GerritEvent into ProjectCreated struct", err)
		return "", err
	}
	Log().Info("ProjectCreated GerritEvent received : ", projectCreated.ProjectName, projectCreated.HeadName, projectCreated.CommonFields.Type)
	projectCreated.RepoURL = pEvent.RepoURL
	cdEvent, err := projectCreated.RepositoryCreatedCDEvent()
	if err != nil {
		return "", err
	}
	Log().Info("Translated project-created gerrit event into dev.cdevents.repository.created CDEvent: ", cdEvent)
	return cdEvent, nil
}

func (pEvent *GerritEvent) HandleProjectHeadUpdatedEvent() (string, error) {
	var projectHeadUpdated ProjectHeadUpdated
	err := json.Unmarshal([]byte(pEvent.Event), &projectHeadUpdated)
	if err != nil {
		Log().Error("Error occurred while Unmarshal GerritEvent into ProjectHeadUpdated struct", err)
		return "", err
	}
	Log().Info("ProjectHeadUpdated GerritEvent received for project : ", projectHeadUpdated.ProjectName)
	projectHeadUpdated.RepoURL = pEvent.RepoURL
	cdEvent, err := projectHeadUpdated.RepositoryModifiedCDEvent()
	if err != nil {
		return "", err
	}
	Log().Info("Translated project-head-updated gerrit event into dev.cdevents.repository.modified CDEvent: ", cdEvent)
	return cdEvent, nil
}

func (pEvent *GerritEvent) HandleRefUpdatedEvent() (string, error) {
	cdEvent := ""
	var refUpdated RefUpdated
	err := json.Unmarshal([]byte(pEvent.Event), &refUpdated)
	if err != nil {
		Log().Error("Error occurred while Unmarshal GerritEvent into RefUpdated struct", err)
		return "", err
	}
	Log().Info("RefUpdated GerritEvent received : ", refUpdated.RefUpdate.RefName, refUpdated.Submitter.Name, refUpdated.CommonFields.Type)
	refUpdated.RepoURL = pEvent.RepoURL
	if strings.Contains(refUpdated.RefUpdate.RefName, "refs/changes") {
		Log().Info("Ignoring handling ref-updated gerrit event as this is followed by patchset/change events: ", refUpdated)
		return "", errors.New("ignoring translating ref-updated gerrit event")
	} else if refUpdated.RefUpdate.OldRev == "0000000000000000000000000000000000000000" {
		cdEvent, err = refUpdated.BranchCreatedCDEvent()
		if err != nil {
			return "", err
		}
		Log().Info("Translated ref-updated gerrit event into dev.cdevents.branch.created CDEvent: ", cdEvent)
	} else if refUpdated.RefUpdate.NewRev == "0000000000000000000000000000000000000000" {
		cdEvent, err = refUpdated.BranchDeletedCDEvent()
		if err != nil {
			return "", err
		}
		Log().Info("Translated ref-updated gerrit event into dev.cdevents.branch.deleted CDEvent: ", cdEvent)
	} else {
		Log().Info("Ignoring handling ref-updated gerrit event for refName : ", refUpdated.RefUpdate.RefName)
		return "", errors.New("ignoring translating ref-updated gerrit event")
	}

	return cdEvent, nil
}

func (pEvent *GerritEvent) HandlePatchsetCreatedEvent() (string, error) {
	cdEvent := ""
	var patchsetCreated PatchsetCreated
	err := json.Unmarshal([]byte(pEvent.Event), &patchsetCreated)
	if err != nil {
		Log().Error("Error occurred while Unmarshal GerritEvent into PatchsetCreated struct", err)
		return "", err
	}
	Log().Info("PatchsetCreated GerritEvent received for project : ", patchsetCreated.Project.Name)
	patchsetCreated.RepoURL = pEvent.RepoURL
	if patchsetCreated.PatchSet.Number == 1 {
		cdEvent, err = patchsetCreated.ChangeCreatedCDEvent()
		if err != nil {
			return "", err
		}
		Log().Info("Translated patchset-created gerrit event into dev.cdevents.change.created CDEvent: ", cdEvent)
	} else {
		cdEvent, err = patchsetCreated.ChangeUpdatedCDEvent()
		if err != nil {
			return "", err
		}
		Log().Info("Translated patchset-created gerrit event into dev.cdevents.change.updated CDEvent: ", cdEvent)
	}

	return cdEvent, nil
}

func (pEvent *GerritEvent) HandleCommentAddedEvent() (string, error) {
	var commentAdded CommentAdded
	err := json.Unmarshal([]byte(pEvent.Event), &commentAdded)
	if err != nil {
		Log().Error("Error occurred while Unmarshal GerritEvent into CommentAdded struct", err)
		return "", err
	}
	Log().Info("CommentAdded GerritEvent received for project : ", commentAdded.Project.Name)
	commentAdded.RepoURL = pEvent.RepoURL
	cdEvent, err := commentAdded.ChangeReviewedCDEvent()
	if err != nil {
		return "", err
	}
	Log().Info("Translated comment-added gerrit event into dev.cdevents.change.reviewed CDEvent: ", cdEvent)
	return cdEvent, nil
}
func (pEvent *GerritEvent) HandleChangeMergedEvent() (string, error) {
	var changeMerged ChangeMerged
	err := json.Unmarshal([]byte(pEvent.Event), &changeMerged)
	if err != nil {
		Log().Error("Error occurred while Unmarshal GerritEvent into ChangeMerged struct", err)
		return "", err
	}
	Log().Info("ChangeMerged GerritEvent received for project : ", changeMerged.Project.Name)
	changeMerged.RepoURL = pEvent.RepoURL
	cdEvent, err := changeMerged.ChangeMergedCDEvent()
	if err != nil {
		return "", err
	}
	Log().Info("Translated change-merged gerrit event into dev.cdevents.change.merged CDEvent: ", cdEvent)
	return cdEvent, nil
}

func (pEvent *GerritEvent) HandleChangeAbandonedEvent() (string, error) {
	var changeAbandoned ChangeAbandoned
	err := json.Unmarshal([]byte(pEvent.Event), &changeAbandoned)
	if err != nil {
		Log().Error("Error occurred while Unmarshal GerritEvent into ChangeAbandoned struct", err)
		return "", err
	}
	Log().Info("ChangeAbandoned GerritEvent received for project : ", changeAbandoned.Project.Name)
	changeAbandoned.RepoURL = pEvent.RepoURL
	cdEvent, err := changeAbandoned.ChangeAbandonedCDEvent()
	if err != nil {
		return "", err
	}
	Log().Info("Translated change-abandoned gerrit event into dev.cdevents.change.abandoned CDEvent: ", cdEvent)
	return cdEvent, nil
}
