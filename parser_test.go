package main

import (
	"bufio"
	"os"
	"testing"
)

// TestNewParser tests the NewParser function
func TestNewParser(t *testing.T) {
	p := NewParser()
	if p == nil {
		t.Error("Expected Parser instance, got nil")
	}
	if p.MyMap == nil {
		t.Error("Expected MyMap to be initialized, got nil")
	}
}

// TestFindKeyAndValue tests the findKeyAndValue method
func TestFindKeyAndValue(t *testing.T) {
	p := NewParser()
	key, value := p.findKeyAndValue("key = value")
	if key != "key" {
		t.Errorf("Expected key to be 'key', got '%s'", key)
	}
	if value != "value" {
		t.Errorf("Expected value to be 'value', got '%s'", value)
	}
}

// TestGetSectionsNames tests the getSectionsNames method
func TestGetSectionsNames(t *testing.T) {
	p := NewParser()
	p.SectionsNames = []string{"section1", "section2"}
	sections := p.getSectionsNames()
	if len(sections) != 2 {
		t.Errorf("Expected 2 sections, got %d", len(sections))
	}
	if sections[0] != "section1" || sections[1] != "section2" {
		t.Error("Section names do not match expected values")
	}
}

// TestGetSectionsMap tests the getSectionsMap method
func TestGetSectionsMap(t *testing.T) {
	p := NewParser()
	p.MyMap["section1"] = map[string]string{"key": "value"}
	sectionsMap := p.getSectionsMap()
	if len(sectionsMap) != 1 {
		t.Errorf("Expected 1 section, got %d", len(sectionsMap))
	}
	if sectionsMap["section1"]["key"] != "value" {
		t.Error("Section map does not match expected values")
	}
}

// TestGetValue tests the getValue method
func TestGetValue(t *testing.T) {
	p := NewParser()
	p.MyMap["section1"] = map[string]string{"key": "value"}
	value := p.getValue("section1", "key")
	if value != "value" {
		t.Errorf("Expected value 'value', got '%s'", value)
	}
}

// TestSetValue tests the setValue method
func TestSetValue(t *testing.T) {
	p := NewParser()
	p.setValue("section1", "key", "value")
	if p.MyMap["section1"]["key"] != "value" {
		t.Error("Value not set correctly")
	}
}

// TestParseLines tests the parseLines method
func TestParseLines(t *testing.T) {
	p := NewParser()
	lines := []string{
		"[section1]",
		"key = value",
	}
	p.parseLines(lines)
	if p.MyMap["section1"]["key"] != "value" {
		t.Error("Lines not parsed correctly")
	}
	if len(p.SectionsNames) != 1 || p.SectionsNames[0] != "section1" {
		t.Error("Section name not parsed correctly")
	}
}

// TestToString tests the toString method
func TestToString(t *testing.T) {
	p := NewParser()
	p.SectionsNames = []string{"section1"}
	p.MyMap["section1"] = map[string]string{"key": "value"}
	lines := p.toString()
	expectedLines := []string{
		"[section1]",
		"key = value",
	}
	for i, line := range lines {
		if line != expectedLines[i] {
			t.Errorf("Expected '%s', got '%s'", expectedLines[i], line)
		}
	}
}

// TestSaveToFile tests the saveToFile method
func TestSaveToFile(t *testing.T) {
	p := NewParser()
	p.SectionsNames = []string{"section1"}
	p.MyMap["section1"] = map[string]string{"key": "value"}
	p.saveToFile("test_output.txt")
	defer os.Remove("test_output.txt")

	file, err := os.Open("test_output.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	expectedLines := []string{
		"[section1]",
		"key = value",
	}
	for i, line := range lines {
		if line != expectedLines[i] {
			t.Errorf("Expected '%s', got '%s'", expectedLines[i], line)
		}
	}
}

// TestLoadFromString tests the loadFromString method
func TestLoadFromString(t *testing.T) {
	p := NewParser()
	content := "[section1]\nkey = value\n"
	p.loadFromString(content)
	if p.MyMap["section1"]["key"] != "value" {
		t.Error("Content not loaded correctly")
	}
}

// TestLoadFromFile tests the loadFromFile method
func TestLoadFromFile(t *testing.T) {
	content := "[section1]\nkey = value\n"
	err := os.WriteFile("test_input.txt", []byte(content), 0644)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove("test_input.txt")

	p := NewParser()
	p.loadFromFile("test_input.txt")
	if p.MyMap["section1"]["key"] != "value" {
		t.Error("File content not loaded correctly")
	}
}
