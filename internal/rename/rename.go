package rename

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/nekovalue/unpckr/internal/config"
	"io"
	"os"
	"strings"
)

func RandomizeConflicts(config *config.ConfigurationType) {
	repeats := make(map[string]int)

	for _, value := range config.Storage.Destinations {
		repeats[value]++
	}

	for i, value := range config.Storage.Destinations {
		if repeats[value] > 1 {
			newFileNameParts := strings.Split(value, ".")
			newFileName := newFileNameParts[0] + "__" +
				generateRandomString(10, i) + "." + newFileNameParts[len(newFileNameParts)-1]
			config.Storage.Destinations[i] = newFileName
		}
	}
}

func HashingDestinations(config *config.ConfigurationType) error {
	for i, value := range config.Storage.Destinations {
		pathWithoutNameParts := strings.Split(value, string(os.PathSeparator))
		pathWithoutName := strings.Join(pathWithoutNameParts[:len(pathWithoutNameParts)-1], string(os.PathSeparator)) +
			string(os.PathSeparator)

		newFileNameParts := strings.Split(value, ".")
		hash, _ := hashFileMD5(config.Storage.Sources[i])
		newFileName := pathWithoutName + hash + "." + newFileNameParts[len(newFileNameParts)-1]

		config.Storage.Destinations[i] = newFileName
	}

	return nil
}

func hashFileMD5(filePath string) (string, error) {
	//Initialize variable returnMD5String now in case an error has to be returned
	var returnMD5String string

	//Open the passed argument and check for any error
	file, err := os.Open(filePath)
	if err != nil {
		return returnMD5String, err
	}

	//Open a new hash interface to write to
	hash := md5.New()

	//Copy the file in the hash interface and check for any error
	if _, err := io.Copy(hash, file); err != nil {
		return returnMD5String, err
	}

	//Get the 16 bytes hash
	hashInBytes := hash.Sum(nil)[:16]

	//Convert the bytes to a string
	returnMD5String = hex.EncodeToString(hashInBytes)

	err = file.Close()
	if err != nil {
		return returnMD5String, err
	}

	return returnMD5String, nil
}
