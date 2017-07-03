package editor

import "testing"

func TestEditor_vi(t *testing.T) {
	editor := NewEditor("vi")
	if editor == nil {
		t.Fatal("falied")
	}
	err := editor.Launch("test1.txt")
	if err != nil {
		t.Fatal(err.Error())
	}
}
func TestEditor_default(t *testing.T) {
	editor := NewEditor("abc")
	if editor == nil {
		t.Fatal("falied")
	}
	err := editor.Launch("test2.txt")
	if err != nil {
		t.Fatal(err.Error())
	}
}
