package filemanager

import (
	"os"
	"path/filepath"
	"testing"
)

func TestIsEmptyDir(t *testing.T) {
	// Create a temp dir for testing
	tmpDir, err := os.MkdirTemp("", "test_isempty")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	fm := &defaultFileManager{}

	// Test 1: Empty directory should return true
	if !fm.IsEmptyDir(tmpDir) {
		t.Error("Expected empty directory to return true")
	}

	// Test 2: Directory with a file should return false
	testFile := filepath.Join(tmpDir, "test.txt")
	if err := os.WriteFile(testFile, []byte("test"), 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	if fm.IsEmptyDir(tmpDir) {
		t.Error("Expected directory with file to return false")
	}

	// Test 3: Remove file, add empty subdirectory - should return true
	os.Remove(testFile)
	subDir := filepath.Join(tmpDir, "subdir")
	if err := os.Mkdir(subDir, 0755); err != nil {
		t.Fatalf("Failed to create subdirectory: %v", err)
	}
	if !fm.IsEmptyDir(tmpDir) {
		t.Error("Expected directory with only empty subdir to return true")
	}

	// Test 4: Add file in subdirectory - should return false
	subFile := filepath.Join(subDir, "subfile.txt")
	if err := os.WriteFile(subFile, []byte("test"), 0644); err != nil {
		t.Fatalf("Failed to create subfile: %v", err)
	}
	if fm.IsEmptyDir(tmpDir) {
		t.Error("Expected directory with file in subdir to return false")
	}

	// Test 5: Non-existent directory should return false
	if fm.IsEmptyDir("/nonexistent/path") {
		t.Error("Expected non-existent directory to return false")
	}
}

func TestExpandTilde(t *testing.T) {
	fm := &defaultFileManager{}

	// Test 1: Path without tilde should return unchanged
	path := "/usr/local/bin"
	if result := fm.ExpandTilde(path); result != path {
		t.Errorf("Expected %s, got %s", path, result)
	}

	// Test 2: Path with tilde should expand to home
	pathWithTilde := "~/Documents"
	result := fm.ExpandTilde(pathWithTilde)
	home, err := os.UserHomeDir()
	if err != nil {
		t.Fatalf("Failed to get home dir: %v", err)
	}
	expected := filepath.Join(home, "Documents")
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}

	// Test 3: Path not starting with tilde should return unchanged
	pathPlain := "Documents"
	if result := fm.ExpandTilde(pathPlain); result != pathPlain {
		t.Errorf("Expected %s, got %s", pathPlain, result)
	}
}
