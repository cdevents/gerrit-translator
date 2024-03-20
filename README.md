# Gerrit Translator CDEvents
A translator plugin for translating [Gerrit events](https://gerrit-review.googlesource.com/Documentation/cmd-stream-events.html#events) into [Source Code Control CDEvents](https://cdevents.dev/docs/source-code-version-control/).
This plugin is served using Hashicorp's [go-plugin](https://github.com/hashicorp/go-plugin/). 

The binary of this plugin is published with a release URL and is used by external applications like [cdevents/webhook-adapter](https://github.com/cdevents/webhook-adapter)

The published plugin's binary can be downloaded and loaded by creating a new plugin client using HashiCorp's go-plugin, which manages the lifecycle of this plugin and establishes the RPC connection.

### How to run locally
Run the below command from the project root directory, which creates a plugin's binary with the name `gerrit-translator-cdevents`
````go
go build -o gerrit-translator-cdevents ./pkg/
````

### Gerrit-CDEvents type mapping for translation
Below are the Gerrit events that currently have mappings with CDEvents and are supported by this translator.

| CDEvent Type  | Gerrit Event Type  |
| :------------ |:-------------------|
| dev.cdevents.repository.created| project-created |
|  dev.cdevents.repository.modified   | project-head-updated    |
| dev.cdevents.branch.created   |  ref-updated     |
| dev.cdevents.branch.deleted |   ref-updated    |
| dev.cdevents.change.created |    patchset-created     |
| dev.cdevents.change.reviewed |    comment-added     |
| dev.cdevents.change.merged |      change-merged   |
| dev.cdevents.change.abandoned |   change-abandoned      | 
| dev.cdevents.change.updated |   patchset-created    |




