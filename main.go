package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

type Record struct {
	Name   string   `yaml:"name"`
	Values []string `yaml:"values"`
}

type RecordType struct {
	A     []Record `yaml:"A,omitempty"`
	MX    []Record `yaml:"MX,omitempty"`
	CNAME []Record `yaml:"CNAME,omitempty"`
	NS    []Record `yaml:"NS,omitempty"`
	SRV   []Record `yaml:"SRV,omitempty"`
}

type Config struct {
	Dns_records_map map[string]RecordType `yaml:"dns_records_map"`
}

// функция возвращает имена исходных файлов в директории с программой.
func FindAllTXTFiles() ([]string, error) {
	var filesList []string

	files, err := os.ReadDir(".")
	if err != nil {
		log.Fatalf("error: %v", err)
		return nil, err
	}

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".txt") {
			filesList = append(filesList, file.Name())
		}
	}
	return filesList, nil
}

// функция дублирует имя выходного файда на такое что было у исходного и добавляет расширение yaml
func changeExtension(filename, newExt string) string {
	return strings.TrimSuffix(filename, filepath.Ext(filename)) + newExt
}

// функция которая подготавливает карту с dns зоной
func processFile(filename string) (Config, error) {
	config := Config{
		Dns_records_map: make(map[string]RecordType),
	}

	file, err := os.Open(filename)
	if err != nil {
		return Config{}, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNumber := 0
	var zoneName string

	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		lineNumber++

		if lineNumber == 2 {
			zoneName = words[len(words)-1]
			fmt.Printf("Zone processing begins: %s\n", zoneName)

			if _, exists := config.Dns_records_map[zoneName]; !exists {
				config.Dns_records_map[zoneName] = RecordType{}
			}
		}

		if len(words) >= 2 {
			recordType := words[len(words)-2]
			recordValue := words[len(words)-1]
			recordName := fmt.Sprintf("%s.%s", words[0], zoneName)

			recordTypeData := config.Dns_records_map[zoneName]

			switch recordType {
			case "A":
				recordTypeData.A = append(recordTypeData.A, Record{Name: recordName, Values: []string{recordValue}})
			case "MX":
				recordTypeData.MX = append(recordTypeData.MX, Record{Name: recordName, Values: []string{recordValue}})
			case "CNAME":
				recordTypeData.CNAME = append(recordTypeData.CNAME, Record{Name: recordName, Values: []string{recordValue}})
			case "NS":
				recordTypeData.NS = append(recordTypeData.NS, Record{Name: recordName, Values: []string{recordValue}})
			case "SRV":
				recordTypeData.SRV = append(recordTypeData.SRV, Record{Name: recordName, Values: []string{recordValue}})
			}
			config.Dns_records_map[zoneName] = recordTypeData
		}
	}

	if err := scanner.Err(); err != nil {
		return Config{}, err
	}

	return config, nil
}

func main() {
	txtFiles, err := FindAllTXTFiles()
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	for _, txtFile := range txtFiles {
		fmt.Printf("Files with zones found: %s\n", txtFile)

		outputFilename := changeExtension(txtFile, ".yaml")

		config, err := processFile(txtFile)
		if err != nil {
			log.Fatalf("error: %v", err)
		}

		if err := writeConfigToFile(config, outputFilename); err != nil {
			log.Fatalf("error: %v", err)
		}

		fmt.Printf("YAML data written to file %s\n", outputFilename)
	}
}

// функция непосредственно занимается конвертация карты в yaml формат
func writeConfigToFile(config Config, filename string) error {
	data, err := yaml.Marshal(&config)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}
