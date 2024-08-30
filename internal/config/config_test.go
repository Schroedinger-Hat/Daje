package config

import (
	"errors"
	"os"
	"path"
	"testing"

	"github.com/Schroedinger-Hat/Daje/constants"
)

const configNotExistsFileName = ".confignotexists"

func Test_getConfigFilePathShouldPassWithNoErrorsWhenFindConfigurationFile(t *testing.T) {
	workingDirectory, _ := os.Getwd()
	constants.ConfigBasepath = workingDirectory + "/../../testdata"

	result := getConfigFilePath()

	if result.Error != nil {
		t.Fatalf("getConfigFilePath should pass with no errors when find configuration file: expected=\"\" - get=\"%v\"", result.Error)
	}
}

func Test_getConfigFilePathShouldReturnTheExactPathWhenItFindTheConfigurationFile(t *testing.T) {
	workingDirectory, _ := os.Getwd()
	constants.ConfigBasepath = workingDirectory + "/../../testdata"
	expectedConfigPath := path.Join(constants.ConfigBasepath, "../testdata/.config/daje", constants.ConfigFileName)

	result := getConfigFilePath()

	if result.Value != expectedConfigPath {
		t.Fatalf("getConfigFilePath should return the exact path when find configuration file: expected=%v - get=%v", expectedConfigPath, result.Value)
	}
}

func Test_getConfigFilePathShouldFailWhenDontFindConfigurationFile(t *testing.T) {
	constants.ConfigFileName = configNotExistsFileName

	result := getConfigFilePath()
	errExpected := errors.New("getConfigFilePath: Configuration not found")

	if result.Error.Error() != errExpected.Error() {
		t.Fatalf("getConfigFilePath should fail when don't find configuration file: expected=%v - get=%v", errExpected, result.Error)
	}
}

func Test_getConfigFilePathShouldHaveEmptyPathWhenDontFindConfigurationFile(t *testing.T) {
	constants.ConfigFileName = configNotExistsFileName
	result := getConfigFilePath()

	if result.Value != "" {
		t.Fatalf("getConfigFilePath should have empty path when don't find configuration file: expected=\"\" - get=%v", result.Value)
	}
}
