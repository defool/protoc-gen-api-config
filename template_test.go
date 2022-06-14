package main

import (
	"html/template"
	"os"
	"testing"
)

func TestTemplate(t *testing.T) {
	val := Config{
		Methods: []MethodInfo{
			{
				Name: "User",
				Path: "/abc",
			},
			{
				Name: "User",
				Path: "/abc",
			},
		},
	}
	tp := template.Must(template.New("").Parse(fieldTemplate))
	err := tp.Execute(os.Stdout, val)
	checkErr(err)
}
