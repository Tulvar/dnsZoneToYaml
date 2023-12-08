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

func changeExtension(filename, newExt string) string {
	return strings.TrimSuffix(filename, filepath.Ext(filename)) + newExt
}

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
			fmt.Println("Zone name:", zoneName)

			if _, exists := config.Dns_records_map[zoneName]; !exists {
				config.Dns_records_map[zoneName] = RecordType{}
			}
		}

		if len(words) >= 2 {
			recordType := words[len(words)-2]
			recordValue := words[len(words)-1]
			recordName := words[0]

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
	inputFilename := "zone.bak"
	outputFilename := changeExtension(inputFilename, ".yaml")

	config, err := processFile(inputFilename)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	if err := writeConfigToFile(config, outputFilename); err != nil {
		log.Fatalf("error: %v", err)
	}

	fmt.Printf("YAML data written to %s\n", outputFilename)
}

func writeConfigToFile(config Config, filename string) error {
	data, err := yaml.Marshal(&config)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}
