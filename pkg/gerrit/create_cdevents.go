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
	sdk "github.com/cdevents/sdk-go/pkg/api"
)

func (projectCreated *ProjectCreated) RepositoryCreatedCDEvent() (string, error) {
	Log().Info("Creating CDEvent RepositoryCreatedEvent")
	cdEvent, err := sdk.NewRepositoryCreatedEvent()
	if err != nil {
		Log().Error("Error creating CDEvent RepositoryCreatedEvent %s\n", err)
		return "", err
	}
	cdEvent.SetSource(projectCreated.RepoURL)
	cdEvent.SetSubjectName(projectCreated.ProjectName)
	cdEvent.SetSubjectId(projectCreated.HeadName)
	cdEvent.SetSubjectUrl(projectCreated.RepoURL)
	cdEventStr, err := sdk.AsJsonString(cdEvent)
	if err != nil {
		Log().Error("Error creating RepositoryCreated CDEvent as Json string %s\n", err)
		return "", err
	}

	return cdEventStr, nil
}

func (projectHeadUpdated *ProjectHeadUpdated) RepositoryModifiedCDEvent() (string, error) {
	Log().Info("Creating CDEvent RepositoryModifiedEvent")
	cdEvent, err := sdk.NewRepositoryModifiedEvent()
	if err != nil {
		Log().Error("Error creating CDEvent RepositoryModified %s\n", err)
		return "", err
	}
	cdEvent.SetSource(projectHeadUpdated.RepoURL)
	cdEvent.SetSubjectName(projectHeadUpdated.ProjectName)
	cdEvent.SetSubjectId(projectHeadUpdated.NewHead)
	cdEvent.SetSubjectUrl(projectHeadUpdated.NewHead)
	cdEventStr, err := sdk.AsJsonString(cdEvent)
	if err != nil {
		Log().Error("Error creating RepositoryModified CDEvent as Json string %s\n", err)
		return "", err
	}

	return cdEventStr, nil
}

func (refUpdated *RefUpdated) BranchCreatedCDEvent() (string, error) {
	Log().Info("Creating CDEvent BranchCreatedEvent")
	cdEvent, err := sdk.NewBranchCreatedEvent()
	if err != nil {
		Log().Error("Error creating CDEvent BranchCreatedEvent %s\n", err)
		return "", err
	}
	cdEvent.SetSource(refUpdated.RepoURL)
	cdEvent.SetSubjectId(refUpdated.RefUpdate.NewRev)
	cdEvent.SetSubjectRepository(&sdk.Reference{Id: refUpdated.RefUpdate.RefName})
	cdEvent.SetSubjectSource(refUpdated.RefUpdate.Project)

	cdEventStr, err := sdk.AsJsonString(cdEvent)
	if err != nil {
		Log().Error("Error creating BranchCreated CDEvent as Json string %s\n", err)
		return "", err
	}
	return cdEventStr, nil
}

func (refUpdated *RefUpdated) BranchDeletedCDEvent() (string, error) {
	Log().Info("Creating CDEvent BranchDeletedEvent")
	cdEvent, err := sdk.NewBranchDeletedEvent()
	if err != nil {
		Log().Error("Error creating CDEvent BranchDeletedEvent %s\n", err)
		return "", err
	}
	cdEvent.SetSource(refUpdated.RepoURL)
	cdEvent.SetSubjectId(refUpdated.RefUpdate.OldRev)
	cdEvent.SetSubjectRepository(&sdk.Reference{Id: refUpdated.RefUpdate.RefName})
	cdEvent.SetSubjectSource(refUpdated.RefUpdate.Project)

	cdEventStr, err := sdk.AsJsonString(cdEvent)
	if err != nil {
		Log().Error("Error creating BranchDeleted CDEvent as Json string %s\n", err)
		return "", err
	}
	return cdEventStr, nil
}

func (patchsetCreated *PatchsetCreated) ChangeCreatedCDEvent() (string, error) {
	Log().Info("Creating CDEvent ChangeCreatedEvent")
	cdEvent, err := sdk.NewChangeCreatedEvent()
	if err != nil {
		Log().Error("Error creating CDEvent ChangeCreatedEvent %s\n", err)
		return "", err
	}
	cdEvent.SetSource(patchsetCreated.RepoURL)
	cdEvent.SetSubjectId(patchsetCreated.Change.Branch)
	cdEvent.SetSubjectSource(patchsetCreated.Change.Url)
	cdEvent.SetSubjectRepository(&sdk.Reference{Id: patchsetCreated.Project.Name})
	cdEventStr, err := sdk.AsJsonString(cdEvent)
	if err != nil {
		Log().Error("Error creating ChangeCreated CDEvent as Json string %s\n", err)
		return "", err
	}

	return cdEventStr, nil
}

func (patchsetCreated *PatchsetCreated) ChangeUpdatedCDEvent() (string, error) {
	Log().Info("Creating CDEvent ChangeUpdatedEvent")
	cdEvent, err := sdk.NewChangeUpdatedEvent()
	if err != nil {
		Log().Error("Error creating CDEvent ChangeUpdatedEvent %s\n", err)
		return "", err
	}
	cdEvent.SetSource(patchsetCreated.RepoURL)
	cdEvent.SetSubjectId(patchsetCreated.Change.Branch)
	cdEvent.SetSubjectSource(patchsetCreated.Change.Url)
	cdEvent.SetSubjectRepository(&sdk.Reference{Id: patchsetCreated.Project.Name})
	cdEventStr, err := sdk.AsJsonString(cdEvent)
	if err != nil {
		Log().Error("Error creating ChangeUpdated CDEvent as Json string %s\n", err)
		return "", err
	}

	return cdEventStr, nil
}

func (commentAdded *CommentAdded) ChangeReviewedCDEvent() (string, error) {
	Log().Info("Creating CDEvent ChangeReviewedEvent")
	cdEvent, err := sdk.NewChangeReviewedEvent()
	if err != nil {
		Log().Error("Error creating CDEvent ChangeReviewedEvent %s\n", err)
		return "", err
	}
	cdEvent.SetSource(commentAdded.RepoURL)
	cdEvent.SetSubjectId(commentAdded.Change.Branch)
	cdEvent.SetSubjectSource(commentAdded.Change.Url)
	cdEvent.SetSubjectRepository(&sdk.Reference{Id: commentAdded.Project.Name})
	cdEventStr, err := sdk.AsJsonString(cdEvent)
	if err != nil {
		Log().Error("Error creating ChangeReviewed CDEvent as Json string %s\n", err)
		return "", err
	}

	return cdEventStr, nil
}

func (changeMerged *ChangeMerged) ChangeMergedCDEvent() (string, error) {
	Log().Info("Creating CDEvent ChangeMergedEvent")
	cdEvent, err := sdk.NewChangeMergedEvent()
	if err != nil {
		Log().Error("Error creating CDEvent ChangeMergedEvent %s\n", err)
		return "", err
	}
	cdEvent.SetSource(changeMerged.RepoURL)
	cdEvent.SetSubjectId(changeMerged.Change.Branch)
	cdEvent.SetSubjectSource(changeMerged.Change.Url)
	cdEvent.SetSubjectRepository(&sdk.Reference{Id: changeMerged.Project.Name})
	cdEventStr, err := sdk.AsJsonString(cdEvent)
	if err != nil {
		Log().Error("Error creating ChangeMerged CDEvent as Json string %s\n", err)
		return "", err
	}

	return cdEventStr, nil
}

func (changeAbandoned *ChangeAbandoned) ChangeAbandonedCDEvent() (string, error) {
	Log().Info("Creating CDEvent ChangeAbandonedEvent")
	cdEvent, err := sdk.NewChangeAbandonedEvent()
	if err != nil {
		Log().Error("Error creating CDEvent ChangeAbandonedEvent %s\n", err)
		return "", err
	}
	cdEvent.SetSource(changeAbandoned.RepoURL)
	cdEvent.SetSubjectId(changeAbandoned.Change.Branch)
	cdEvent.SetSubjectSource(changeAbandoned.Change.Url)
	cdEvent.SetSubjectRepository(&sdk.Reference{Id: changeAbandoned.Project.Name})
	cdEventStr, err := sdk.AsJsonString(cdEvent)
	if err != nil {
		Log().Error("Error creating ChangeAbandoned CDEvent as Json string %s\n", err)
		return "", err
	}

	return cdEventStr, nil
}
