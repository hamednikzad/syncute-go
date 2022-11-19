package messages

import "testing"

func TestParseBadMessage(t *testing.T) {
	json := `{"Type":"bad_message"}`
	message, _ := parseBadMessage([]byte(json))

	if message.Type != BadMessageType {
		t.Errorf("parseBadMessage was wrong, got: %s, want: %s", message.Type, BadMessageType)
	}
}

func TestParseAllResourcesListMessage(t *testing.T) {
	json := `{"Type":"resources"}`
	message, _ := parseAllResourcesListMessage([]byte(json))

	if message.Type != AllResourcesListType {
		t.Errorf("parseAllResourcesListMessage was wrong, got: %s, want: %s", message.Type, AllResourcesListType)
	}
}

func TestParseNewResourceReceivedMessage(t *testing.T) {
	json := `{"Type":"new_resource"}`
	message, _ := parseNewResourceReceivedMessage([]byte(json))

	if message.Type != NewResourceReceivedType {
		t.Errorf("parseNewResourceReceivedMessage was wrong, got: %s, want: %s", message.Type, NewResourceReceivedType)
	}
}

func TestGetMessageType(t *testing.T) {
	tables := []struct {
		json    string
		msgType string
	}{
		{`{"Type":"bad_message"}`, BadMessageType},
		{`{"Type":"ready"}`, ReadyType},
		{`{"Type":"resources"}`, AllResourcesListType},
		{`{"Type":"new_resource"}`, NewResourceReceivedType},
	}
	for _, table := range tables {
		messageType, _ := getMessageType([]byte(table.json))
		if messageType != table.msgType {
			t.Errorf("getMessageType was wrong, got: %s, want: %s", messageType, table.msgType)
		}
	}
}
