package main

import (
  "testing"
  . "github.com/diasporism/file_editor"
)

func TestPage(t *testing.T) {
  page := Page{ Title: "title", Body: []byte("body") }

  if page.Title != "title" {
    t.Error(
      "For", page.Title,
      "expected", "title",
      "got", page.Title,
    )
  }

  if string(page.Body) != "body" {
    t.Error(
      "For", page.Body,
      "expected", "body",
      "got", page.Body,
    )
  }
}
