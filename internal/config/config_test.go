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
	constants.DajeBasePath = workingDirectory + "/../../testdata"

	_, err := getConfigFilePath()

	if err != nil {
		t.Fatalf("getConfigFilePath should pass with no errors when find configuration file: expected=\"\" - get=\"%v\"", err)
	}
}

func Test_getConfigFilePathShouldReturnTheExactPathWhenItFindTheConfigurationFile(t *testing.T) {
	workingDirectory, _ := os.Getwd()
	constants.DajeBasePath = workingDirectory + "/../../testdata"
	expectedConfigPath := path.Join(constants.DajeBasePath, "../testdata/.config/daje", constants.DajeConfigFileName)

	path, _ := getConfigFilePath()

	if path != expectedConfigPath {
		t.Fatalf("getConfigFilePath should return the exact path when find configuration file: expected=%v - get=%v", expectedConfigPath, path)
	}
}

func Test_getConfigFilePathShouldFailWhenDontFindConfigurationFile(t *testing.T) {
	constants.DajeConfigFileName = configNotExistsFileName

	_, err := getConfigFilePath()
	errExpected := errors.New("getConfigFilePath: Configuration not found")

	if err.Error() != errExpected.Error() {
		t.Fatalf("getConfigFilePath should fail when don't find configuration file: expected=%v - get=%v", errExpected, err)
	}
}

func Test_getConfigFilePathShouldHaveEmptyPathWhenDontFindConfigurationFile(t *testing.T) {
	constants.DajeConfigFileName = configNotExistsFileName
	path, _ := getConfigFilePath()

	if path != "" {
		t.Fatalf("getConfigFilePath should have empty path when don't find configuration file: expected=\"\" - get=%v", path)
	}
}
