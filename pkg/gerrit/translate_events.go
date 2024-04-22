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
	if refUpdated.RefUpdate.OldRev == "0000000000000000000000000000000000000000" {
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
	}

	return cdEvent, nil
}
