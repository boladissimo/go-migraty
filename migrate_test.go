package migraty

import (
	"testing"
)

const testScriptsPath string = "test/scripts/"

func TestNormalizePath_givenPathWithSlash_givenPath(t *testing.T) {

	path := "foo/bar/"
	actual := normalizeMigrationPath(path)

	if actual != path {
		t.Errorf("Expected %s. Got %s", path, actual)
	}

}

func TestNormalizePath_givenPathWithoutSlash_givenPathWithFinalSlash(t *testing.T) {

	path := "foo/bar"
	expected := path + "/"
	actual := normalizeMigrationPath(path)

	if actual != expected {
		t.Errorf("Expected %s. Got %s", expected, actual)
	}
}

func TestGetMigrationScript_givenTableNameAndScriptPath_scriptContent(t *testing.T) {
	expected := "bar"
	actual := getMigrationScript("foo", testScriptsPath)

	if actual != expected {
		t.Errorf("Expected %s. Got %s", expected, actual)
	}
}

func TestGetTableNames_givenScriptPath_fooAndBar(t *testing.T) {
	expected := stringList{"bar", "foo"}
	actual := getTableNames(testScriptsPath)

	if !expected.equals(actual) {
		t.Errorf("Expected %s. Got %s", expected, actual)
	}
}

type stringList []string

func (a stringList) equals(b stringList) bool {
	logInfo(a, b)
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
