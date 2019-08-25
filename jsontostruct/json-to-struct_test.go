package jsontostruct

import (
	"testing"
)

func TestJsonObj(t *testing.T) {
	if err := generateStruct(jsonObj); err != nil {
		t.Error("generate error:", err)
	}
}

func TestJsonArr(t *testing.T) {
	if err := generateStruct(jsonArray); err != nil {
		t.Error("generate error:", err)
	}
}

	const jsonObj = `
	{
		"name": "sam",
		"gender": "m",
		"pet": null,
		"skills": [
		  "Eating",
		  "Sleeping",
		  "Crawing"
		],
		"xyz": {
			"aaa": 1,
			"bbb": 2,
			"ccc": 3
		},
		"hello.world": true
	}
	`
	const jsonStream = `
	{"Name": "Ed", "Text": "Knock knock."}
	{"Name": "Sam", "Text": "Who's there?"}
	{"Name": "Ed", "Text": "Go fmt."}
	{"Name": "Sam", "Text": "Go fmt who?"}
	{"Name": "Ed", "Text": "Go fmt yourself!"}
	`

	const jsonArray = `
	[
	{
		"name": "elgs",
		"gender": "m",
		"skills": [
		  "Golang",
		  "Java",
		  "C"
		]
	},
	{
		"name": "enny",
		"gender": "f",
		"skills": [
		  "IC",
		  "Electirc design",
		  "Verification"
		]
	},
	{
		"name": "sam",
		"gender": "m",
		"pet": null,
		"skills": [
		  "Eating",
		  "Sleeping",
		  "Crawing"
		]
	}
	]
	`
