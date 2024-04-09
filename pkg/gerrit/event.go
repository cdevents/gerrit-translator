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
	"fmt"
	"log"
	"net/http"
)

type GerritEvent struct {
	Event   string
	RepoURL string
}

func NewGerritEvent(event string, repoURL string) (pEvent *GerritEvent) {
	pEvent = &GerritEvent{event, repoURL}
	return
}

func HandleTranslateGerritEvent(event string, header http.Header) (string, error) {
	log.Printf("Handle translation into CDEvent from Gerrit event %s\n", event)
	repoURL := ""
	if header.Get("X-Origin-Url") != "" {
		repoURL = header.Get("X-Origin-Url")
	}
	gerritEvent := NewGerritEvent(event, repoURL)
	cdEvent, err := gerritEvent.TranslateIntoCDEvent()
	if err != nil {
		log.Printf("Error translating Gerrit event into CDEvent %s\n", err)
		return "", err
	}
	log.Printf("Gerrit Event translated into CDEvent %s\n", cdEvent)
	return cdEvent, nil
}

func (pEvent *GerritEvent) TranslateIntoCDEvent() (string, error) {
	eventMap := make(map[string]interface{})
	cdEvent := ""
	err := json.Unmarshal([]byte(pEvent.Event), &eventMap)
	if err != nil {
		log.Println("Error occurred while Unmarshal gerritEvent data into gerritEvent map", err)
		return "", err
	}
	eventType := eventMap["type"]
	log.Printf("handling translating to CDEvent from Gerrit Event type: %s\n", eventType)

	switch eventType {
	case "project-created":
		cdEvent, err = pEvent.HandleProjectCreatedEvent()
		if err != nil {
			return "", err
		}
	case "ref-updated":
		cdEvent, err = pEvent.HandleRefUpdatedEvent()
		if err != nil {
			return "", err
		}
	case "project-head-updated":
		cdEvent, err = pEvent.HandleProjectHeadUpdatedEvent()
		if err != nil {
			return "", err
		}
	case "patchset-created":
		cdEvent, err = pEvent.HandlePatchsetCreatedEvent()
		if err != nil {
			return "", err
		}
	case "comment-added":
		cdEvent, err = pEvent.HandleCommentAddedEvent()
		if err != nil {
			return "", err
		}
	case "change-merged":
		cdEvent, err = pEvent.HandleChangeMergedEvent()
		if err != nil {
			return "", err
		}
	case "change-abandoned":
		cdEvent, err = pEvent.HandleChangeAbandonedEvent()
		if err != nil {
			return "", err
		}
	default:
		log.Printf("Not handling CDEvent translation for Gerrit event type: %s\n", eventMap["type"])
		return "", fmt.Errorf("gerrit event type %s, not supported for translation", eventType)
	}
	return cdEvent, nil
}
