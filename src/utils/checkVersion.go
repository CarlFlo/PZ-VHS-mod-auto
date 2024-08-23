package utils

import (
	"io"
	"net/http"
	"regexp"

	"github.com/CarlFlo/malm"
)

func CheckVersion(currentVersion string) {
	// Handles checking if there is an update available for the bot
	upToDate, githubVersion, err := versonHandler(currentVersion)
	if err != nil {
		malm.Error("%s", err)
	}

	if upToDate {
		malm.Debug("Version %s", currentVersion)
	} else {
		malm.Info("New version available at '%s'! New version: '%s'. Your version: '%s'",
			"https://raw.githubusercontent.com/CarlFlo/PZ-VHS-mod-auto",
			githubVersion,
			currentVersion)
	}
}

// Return true or false if the version is up to date
// Return version on system
// Return version on github
// return error
func versonHandler(current string) (bool, string, error) {

	githubVersion, err := githubVersion()

	if err != nil {
		return false, "", err
	}

	upToDate := current == githubVersion

	return upToDate, githubVersion, nil
}

// Returns the online version or the error
func githubVersion() (string, error) {

	// get URL
	resp, err := http.Get("https://raw.githubusercontent.com/CarlFlo/PZ-VHS-mod-auto/main/main.go")
	if err != nil {
		return "", err
	}

	// read response
	defer resp.Body.Close()
	body, err := io.ReadAll(io.LimitReader(resp.Body, 1024*1024)) // Limited to 1 MB
	if err != nil {
		return "", err
	}

	// regex to find version
	pattern := regexp.MustCompile(`\d+-\d+-\d+`)
	version := pattern.FindString(string(body))

	return version, nil
}
