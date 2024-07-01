package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Parser represents a configuration parser with sections and key-value pairs
type Parser struct {
	SectionsNames []string
	MyMap         map[string]map[string]string
}

// NewParser creates a new Parser instance
func NewParser() *Parser {
	return &Parser{
		MyMap: make(map[string]map[string]string),
	}
}

// findKeyAndValue splits a line into a key and value
func (p *Parser) findKeyAndValue(line string) (string, string) {
	res := strings.Split(line, " = ")
	return res[0], res[1]
}

// getSectionsNames returns the names of the sections
func (p *Parser) getSectionsNames() []string {
	return p.SectionsNames
}

// getSectionsMap returns the map of sections
func (p *Parser) getSectionsMap() map[string]map[string]string {
	return p.MyMap
}

// getValue returns the value for a given section and key
func (p *Parser) getValue(sectionName, key string) string {
	return p.MyMap[sectionName][key]
}

// setValue sets the value for a given section and key
func (p *Parser) setValue(sectionName, key, value string) {
	if _, exists := p.MyMap[sectionName]; !exists {
		p.MyMap[sectionName] = make(map[string]string)
	}
	p.MyMap[sectionName][key] = value
}

// parseLines parses lines to fill sections and key-value pairs
func (p *Parser) parseLines(lines []string) {
	var sectionName string
	for _, line := range lines {
		if len(line) == 0 || line[0] == ';' || line == "\n" {
			continue
		}
		if line[0] == '[' {
			sectionName = line[1 : len(line)-1]
			p.SectionsNames = append(p.SectionsNames, sectionName)
			p.MyMap[sectionName] = make(map[string]string)
		} else {
			key, value := p.findKeyAndValue(line)
			p.MyMap[sectionName][key] = value
		}
	}
}

// toString converts the parser content to a string slice
func (p *Parser) toString() []string {
	var lines []string
	for _, sectionName := range p.SectionsNames {
		lines = append(lines, "["+sectionName+"]")
		for key, value := range p.MyMap[sectionName] {
			lines = append(lines, key+" = "+value)
		}
	}
	return lines
}

// saveToFile saves the parser content to a file
func (p *Parser) saveToFile(filePath string) {
	lines := p.toString()
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
}

// loadFromString loads the parser content from a string
func (p *Parser) loadFromString(content string) {
	lines := strings.Split(content, "\n")
	p.parseLines(lines)
}

// loadFromFile loads the parser content from a file
func (p *Parser) loadFromFile(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Error while opening the file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 || line[0] == ';' || line == "\n" {
			continue
		}
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("Error while reading the file")
	}
	p.parseLines(lines)
}
func main() {
}
