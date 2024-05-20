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

package main

import (
	"net/http"

	"github.com/cdevents/gerrit-translator/pkg/gerrit"
	"github.com/cdevents/webhook-adapter/pkg/cdevents"
	"github.com/hashicorp/go-plugin"
)

type EventTranslator struct{}

// TranslateEvent Invoked from external application to translate Gerrit event into CDEvent
func (EventTranslator) TranslateEvent(event string, headers http.Header) (string, error) {
	gerrit.Log().Info("Serving from gerrit-translator plugin")
	cdEvent, err := gerrit.HandleTranslateGerritEvent(event, headers)
	if err != nil {
		return "", err
	}
	return cdEvent, nil
}

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: cdevents.Handshake,
		Plugins: map[string]plugin.Plugin{
			"gerrit-translator-cdevents": &cdevents.TranslatorGRPCPlugin{Impl: &EventTranslator{}},
		},

		// A non-nil value here enables gRPC serving for this plugin...
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
