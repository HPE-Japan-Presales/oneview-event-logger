package oneview

import (
	"testing"

	"github.com/HewlettPackard/oneview-golang/ov"
)

func createOvClient() *ov.OVClient {
	var ovClient *ov.OVClient

	c := ovClient.NewOVClient(
		"golang",
		"golangtest",
		"",
		"https://192.168.2.6",
		false, //ssl verificcation
		1200,
		"*")
	return c
}

func TestOvEvent(t *testing.T) {
	c := createOvClient()
	events, err := GetEventList(c)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", events)
}
