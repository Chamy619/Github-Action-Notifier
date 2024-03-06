package main

import "testing"

func TestNewConfig(t *testing.T) {
	filename := "config.yaml.sample"
	config, err := NewConfig(filename)

	if err != nil {
		t.Errorf("NewConfig fail: %s", err)
	}
	
	gotName := config.Notifies[0].Files[0].Name
	wantName := "apps/master-api/config.ini.sample"
	AssertEqual(t, wantName, gotName)

	gotType := config.Notifies[0].Files[0].Type
	wantType := "github-action"
	AssertEqual(t, wantType, gotType)

	gotAdded := config.Notifies[0].Files[0].Added
	wantAdded := false
	AssertEqual(t, wantAdded, gotAdded)
	
	gotModified := config.Notifies[0].Files[0].Modified
	wantModified := true
	AssertEqual(t, wantModified, gotModified)

	gotRemoved := config.Notifies[0].Files[0].Removed
	wantRemoved := false
	AssertEqual(t, wantRemoved, gotRemoved)

	gotMessage := config.Notifies[0].Files[0].Message
	wantMessage := "apps/master-api/config.ini.sample file is changed"
	AssertEqual(t, wantMessage, gotMessage)

	gotMessengerType := config.Notifies[0].Files[0].Messenger.Type
	wantMessengerType := "slack"
	AssertEqual(t, wantMessengerType, gotMessengerType)

	gotMessengerUrl := config.Notifies[0].Files[0].Messenger.Url
	wantMessengerUrl := "https://hooks.slack.com/services/key1/key2/key3"
	AssertEqual(t, wantMessengerUrl, gotMessengerUrl)
}

func TestNewConfigNotExistFile(t *testing.T) {
	filename := "notexists.yaml"
	_, err := NewConfig(filename)
	AssertEqual(t, "file is not readable", err.Error())
}

func AssertEqual(t *testing.T, want interface{}, got interface{}) {
	t.Helper()
	if want != got {
		t.Errorf("%v is not equal %v", want, got)
	}
}