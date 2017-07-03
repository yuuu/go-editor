package editor

import (
	"os"
	"os/exec"
)

// Editor ...
type Editor struct {
	cmd string
}

// EditorError ...
type EditorError struct {
	msg    string
	cmderr error
}

// DEFAULT_EDITOR ...
const DEFAULT_EDITOR = "vi"

// NewEditor ...
func NewEditor(cmd string) *Editor {
	var editorCmd string
	var err error

	editorCmd = cmd
	err = nil
	_, err = exec.LookPath(editorCmd)
	if err != nil {
		osEditorCmd := os.Getenv("EDITOR")
		if osEditorCmd != "" {
			editorCmd = osEditorCmd
		} else {
			editorCmd = DEFAULT_EDITOR
		}
		err = nil
		_, err = exec.LookPath(editorCmd)
		if err != nil {
			return nil
		}
	}

	editor := Editor{editorCmd}
	return &editor
}

// Launch ...
func (editor *Editor) Launch(filepath string) error {
	var cmderr error

	cmd := exec.Command(editor.cmd, filepath)
	if cmd == nil {
		err := EditorError{"Invalid command", nil}
		return &err
	}

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmderr = cmd.Start()
	if cmderr != nil {
		err := EditorError{"Failed to Start", cmderr}
		return &err
	}

	cmderr = cmd.Wait()
	if cmderr != nil {
		err := EditorError{"Failed to Wait", cmderr}
		return &err
	}

	return nil
}

// Error ...
func (err *EditorError) Error() string {
	return err.msg + " [suberr:" + err.cmderr.Error() + "]"
}
