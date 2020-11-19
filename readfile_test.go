package readfile_test

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/caioreix/readfile"
)

var path = "readfilepackagetest.test"
var content = "Lorem ipsum dolor sit amet, consectetur adipiscing elit.\nPhasellus pharetra finibus turpis, non mattis augue finibus vel.\nCras arcu mi."

func createTestFile() {

	err := ioutil.WriteFile(path, []byte(content), 0755)
	if err != nil {
		log.Fatalf("Unable to write file: %v", err)
	}
}

func deleteTestFile() {
	err := os.Remove(path)
	if err != nil {
		log.Fatalf("Unable to delete file: %v", err)
	}
}

func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestFile(t *testing.T) {
	createTestFile()
	defer deleteTestFile()
	ctt, err := readfile.File(path)
	if err != nil {
		t.Error(err)
	}

	if ctt != content {
		t.Errorf("have:\n%s\nwant:\n%s", ctt, content)
	}
}

func TestByLine(t *testing.T) {
	createTestFile()
	defer deleteTestFile()
	lines, err := readfile.ByLine(path)
	if err != nil {
		t.Error(err)
	}

	contentLines := strings.SplitAfter(content, "\n")

	if equal(lines, contentLines) {
		t.Errorf("have:\n%#v\nwant:\n%#v", lines, contentLines)
	}
}

func TestByWord(t *testing.T) {
	createTestFile()
	defer deleteTestFile()
	words, err := readfile.ByWord(path)
	if err != nil {
		t.Error(err)
	}

	contentWords := strings.Split(content, " ")

	if equal(words, contentWords) {
		t.Errorf("have:\n%#v\nwant:\n%#v", words, contentWords)
	}
}
