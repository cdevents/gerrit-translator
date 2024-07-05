package gerrit

import (
	"net/http"
	"os"
	"testing"
)

func TestHandleTranslateProjectCreatedEvent(t *testing.T) {
	event, err := os.ReadFile("testdata/project-created.json")
	if err != nil {
		t.Fatalf("Failed to read project-created.json file: %v", err)
	}
	headers := http.Header{}
	headers.Set("X-Origin-Url", "http://gerrit.est.tech")

	cdEvent, err := HandleTranslateGerritEvent(string(event), headers)
	if err != nil {
		t.Errorf("Expected RepositoryCreated CDEvent to be successful.")
		return
	}
	Log().Info("Handle project-created gerrit event into dev.cdevents.repository.created is successful ", cdEvent)
}

func TestHandleTranslateProjectHeadUpdatedEvent(t *testing.T) {

	event, err := os.ReadFile("testdata/project-head-updated.json")
	if err != nil {
		t.Fatalf("Failed to read project-head-updated.json file: %v", err)
	}
	headers := http.Header{}
	headers.Set("X-Origin-Url", "http://gerrit.est.tech")

	cdEvent, err := HandleTranslateGerritEvent(string(event), headers)
	if err != nil {
		t.Errorf("Expected RepositoryModified CDEvent to be successful.")
		return
	}
	Log().Info("Handle project-head-updated gerrit event into dev.cdevents.repository.modified is successful ", cdEvent)
}
