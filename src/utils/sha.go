package utils

import (
	"crypto/sha1"
	"fmt"
	"io"
	"os"

	"github.com/CarlFlo/malm"
)

func CompareSha1(path, sha1 string) bool {

	newSha1, err := MakeSha1(path)
	if err != nil {
		malm.Fatal("'%v'", err)
	}

	return sha1 == newSha1
}

func MakeSha1(filePath string) (string, error) {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	// Create a new SHA-1 hash
	hasher := sha1.New()

	// Copy the file's content into the hasher
	if _, err := io.Copy(hasher, file); err != nil {
		return "", fmt.Errorf("error hashing file: %w", err)
	}

	// Calculate the SHA-1 checksum
	hashInBytes := hasher.Sum(nil)

	// Convert the checksum to a hex string
	hashString := fmt.Sprintf("%x", hashInBytes)

	return hashString, nil
}
