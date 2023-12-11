# [ENG]
# DNS Records Processing Tool
This Go package provides a tool for processing DNS records from text files and converting them to YAML format. It is particularly useful for managing DNS configurations in a structured and automated way.

## Features
* File Discovery: Automatically finds all .txt files in the current directory.
* DNS Record Parsing: Parses various DNS record types (A, MX, CNAME, NS, SRV) from the text files.
* YAML Conversion: Converts the parsed records into a YAML file, maintaining the structure and integrity of the original data.
## Requirements
* Go (Golang) environment setup is required to run this tool.
* gopkg.in/yaml.v3 package for YAML marshaling.
## Installation
1. Clone the repository or download the source code.
2. Ensure you have Go installed on your system.
3. Install the YAML package: go get gopkg.in/yaml.v3
## Usage
1. Place your .txt DNS record files in the same directory as the tool.
2. Run the program using go run main.go.
3 .The tool will find all .txt files, process them, and output corresponding .yaml files in the same directory.
## Functions Overview
* FindAllTXTFiles(): Scans the current directory and returns a list of .txt files.
* changeExtension(filename, newExt): Changes the file extension to .yaml.
* processFile(filename): Parses the DNS records from a given text file and returns a Config struct containing the DNS records map.
* writeConfigToFile(config, filename): Converts the Config struct to YAML format and writes it to a file.
## Structure
* Record: Represents a DNS record with Name and Values.
* RecordType: Contains slices of Record for different DNS record types.
* Config: Main struct holding the map of RecordType keyed by zone names.
## Example
Comig soon 

----------------------------------------------------------------------------------------------------------------------------------------------------------------

# [RUS]
# Инструмент Обработки DNS Записей
Этот пакет на языке программирования Go предлагает инструмент для обработки DNS записей из текстовых файлов и их преобразования в формат YAML. Это особенно полезно для управления конфигурациями DNS структурированным и автоматизированным способом.

## Особенности
* Поиск Файлов: Автоматический поиск всех файлов .txt в текущем каталоге.
* Анализ DNS Записей: Анализ различных типов DNS записей (A, MX, CNAME, NS, SRV) из текстовых файлов.
* Преобразование в YAML: Конвертация проанализированных записей в файл YAML, сохраняя структуру и целостность исходных данных.
## Требования
* Необходима установка среды Go (Golang) для использования этого инструмента.
* Пакет gopkg.in/yaml.v3 для маршалинга YAML.
## Установка
1. Склонируйте репозиторий или загрузите исходный код.
2. Убедитесь, что у вас установлен Go на вашей системе.
3. Установите пакет YAML: go get gopkg.in/yaml.v3
## Использование
1. Разместите ваши .txt файлы с DNS записями в той же директории, что и инструмент.
2. Запустите программу с помощью go run main.go.
3. Инструмент найдет все файлы .txt, обработает их и выведет соответствующие файлы .yaml в той же директории.
## Обзор Функций
* FindAllTXTFiles(): Сканирует текущий каталог и возвращает список файлов .txt.
* changeExtension(filename, newExt): Изменяет расширение файла на .yaml.
* processFile(filename): Анализирует DNS записи из данного текстового файла и возвращает структуру Config, содержащую карту записей DNS.
* writeConfigToFile(config, filename): Преобразует структуру Config в формат YAML и записывает ее в файл.
## Структура
Record: Представляет собой DNS запись с Name и Values.
RecordType: Содержит массивы Record для различных типов DNS записей.
Config: Основная структура, содержащая карту RecordType, ключеванных именами зон.
## Пример
Будет добавлен позже
