package objective

import (
	"testing"
	"pcne/assistant"
)

func TestListObjectives(t *testing.T) {
	got := ListObjectives()
	want := ""

	assistant.AssertCorrect(t, got, want)
}