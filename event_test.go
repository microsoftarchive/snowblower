package main

import (
	"encoding/json"
	"strings"
	"testing"
)

// sample article based on sample data collected from ObjC tracker
const eventArticle1 = `
{
  "aid":"Wunderlist/3.2.2",
  "res":"750x1334",
  "uid":"777",
  "p":"mob",
  "cx":"eyJzY2hlbWEiOiJpZ2x1OmNvbS5zbm93cGxvd2FuYWx5dGljcy5zbm93cGxvdy9jb250ZXh0cy9qc29uc2NoZW1hLzEtMC0wIiwiZGF0YSI6W119",
  "dtm":"1427898727829",
  "tv":"ios-0.3.2",
  "tna":"com.wunderlist",
  "ue_px":"eyJzY2hlbWEiOiJpZ2x1OmNvbS5zbm93cGxvd2FuYWx5dGljcy5zbm93cGxvdy91bnN0cnVjdF9ldmVudC9qc29uc2NoZW1hLzEtMC0wIiwiZGF0YSI6e319",
  "e":"ue",
  "lang":"en",
  "vp":"750x1334",
  "eid":"8dbe2937-beea-45ee-9e78-269343ba6561"
}
`

// boring test of serialization just to make sure we're catching the various
// things. Maybe not useful once we're done with primary development.
func TestEventArticle1Deserialization(t *testing.T) {
	a := Event{}
	if err := json.Unmarshal([]byte(eventArticle1), &a); err != nil {
		t.Errorf("Internal problem with article: %s", err)
		return
	}
	if a.AppID != "Wunderlist/3.2.2" {
		t.Error("AppID did not properly deserialize")
	}
	if a.Resolution != "750x1334" {
		t.Error("Resolution did not properly deserialize")
	}
	if a.UserID != "777" {
		t.Error("UserID did not properly deserialize")
	}
	if a.Platform != "mob" {
		t.Error("Platform did not properly deserialize")
	}
	if !strings.HasPrefix(a.ContextsEncoded, "eyJzY") {
		t.Error("Contexts did not properly deserialize")
	}
	if a.DeviceTimestamp != "1427898727829" {
		t.Error("Device Timestamp did not properly deserialize")
	}
	if a.TrackerVersion != "ios-0.3.2" {
		t.Error("Tracker Version did not properly deserialize")
	}
	if a.Namespace != "com.wunderlist" {
		t.Error("Namespace did not properly deserialize")
	}
	if !strings.HasPrefix(a.UnstructuredEventEncoded, "eyJzY2hlbWEiOiJpZ2x1O") {
		t.Error("UnstructuredEventEncoded did not properly deserialize")
	}
	if a.Event != "ue" {
		t.Error("Event did not properly deserialize")
	}
	if a.Language != "en" {
		t.Error("Language did not propery deserialize")
	}
	if a.Viewport != "750x1334" {
		t.Error("Viewport did not propery deserialize")
	}
	if a.EventID != "8dbe2937-beea-45ee-9e78-269343ba6561" {
		t.Error("Event ID did not propery deserialize")
	}
}
