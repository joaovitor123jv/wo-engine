package woutils

import (
	"encoding/xml"
	"io"
	"os"
	"strings"
)

func GetFileName(path string) string {
	return path[strings.LastIndex(path, "/")+1:]
}

func GetDirFromPath(path string) string {
	return path[:len(path)-len(GetFileName(path))]
}

func AppendOnPath(path, toAppend string) string {
	trimmedPath := strings.TrimSuffix(path, "/")
	trimmedToAppend := strings.TrimPrefix(toAppend, "/")

	return trimmedPath + "/" + trimmedToAppend
}

// ReadXml reads an XML file and unmarshals it into the xmlStruct interface
func ReadXml(filePath string, xmlStruct interface{}) error {
	xmlFile, err := os.Open(filePath)

	if err != nil {
		return err
	}
	defer xmlFile.Close()

	bytes, err := io.ReadAll(xmlFile)
	if err != nil {
		return err
	}

	err = xml.Unmarshal(bytes, &xmlStruct)
	if err != nil {
		return err
	}

	return nil
}
