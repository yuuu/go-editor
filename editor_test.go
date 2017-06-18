package editor

import "testing"

func TestEditor_example(t *testing.T) {
	editor := NewEditor("vi")
	if editor == nil {
		t.Fatal("falied")
	}
	err := editor.Launch("test.txt")
	if err != nil {
		t.Fatal(err.Error())
	}
}
