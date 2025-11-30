package config

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"

	"gopkg.in/yaml.v3"
)

func LoadConfig() Config {
	return LoadConfigFromPath("resources")
}

func LoadConfigFromPath(resourcePath string) Config {
	projectRoot, err := os.Getwd()
	if err != nil {
		panic(fmt.Sprintf("failed to get working directory: %v", err))
	}

	basePath := filepath.Join(projectRoot, resourcePath)

	masterPath := filepath.Join(basePath, "config.master.yml")
	masterConfig, err := loadYAMLFile[Config](masterPath)
	if err != nil {
		panic(fmt.Sprintf("failed to load master config: %v", err))
	}

	secretPath := filepath.Join(basePath, "config.secret.yml")
	secretConfig, err := loadYAMLFile[Config](secretPath)
	if err != nil {
		fmt.Printf("Warning: failed to load secret config: %v\n", err)
		secretConfig = Config{}
	}

	merged := deepMerge(masterConfig, secretConfig)
	merged = loadFileContents(merged, projectRoot)

	return merged
}

func loadFileContents(cfg Config, projectRoot string) Config {
	cfgVal := reflect.ValueOf(&cfg).Elem()
	loadFileContentsRecursive(cfgVal, projectRoot)
	return cfg
}

func loadFileContentsRecursive(val reflect.Value, projectRoot string) {
	if !val.IsValid() || !val.CanSet() {
		return
	}

	switch val.Kind() {
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			field := val.Field(i)
			fieldType := val.Type().Field(i)

			if !field.CanSet() {
				continue
			}

			if field.Kind() == reflect.String {
				filePath := field.String()
				if filePath == "" {
					continue
				}

				if isFilePath(filePath) {
					fullPath := filepath.Join(projectRoot, filePath)
					content, err := os.ReadFile(fullPath)
					if err != nil {
						panic(fmt.Sprintf("failed to load file content for field '%s' from %s: %v",
							fieldType.Name, fullPath, err))
					}
					field.SetString(string(content))
				}
			} else if field.Kind() == reflect.Struct {
				loadFileContentsRecursive(field, projectRoot)
			}
		}

	case reflect.Ptr:
		if !val.IsNil() {
			loadFileContentsRecursive(val.Elem(), projectRoot)
		}
	}
}

func isFilePath(s string) bool {
	if s == "" {
		return false
	}
	ext := filepath.Ext(s)
	validExts := []string{".txt", ".md", ".json", ".yaml", ".yml", ".xml", ".html", ".sql"}
	for _, validExt := range validExts {
		if ext == validExt {
			return true
		}
	}
	return false
}

func loadYAMLFile[T any](path string) (T, error) {
	var result T

	data, err := os.ReadFile(path)
	if err != nil {
		return result, fmt.Errorf("failed to read file %s: %w", path, err)
	}

	if err := yaml.Unmarshal(data, &result); err != nil {
		return result, fmt.Errorf("failed to parse YAML from %s: %w", path, err)
	}

	return result, nil
}

func deepMerge(base, override Config) Config {
	result := base

	baseVal := reflect.ValueOf(&result).Elem()
	overrideVal := reflect.ValueOf(override)

	mergeValues(baseVal, overrideVal)

	return result
}

func mergeValues(base, override reflect.Value) {
	if !override.IsValid() {
		return
	}

	switch override.Kind() {
	case reflect.Struct:
		for i := 0; i < override.NumField(); i++ {
			overrideField := override.Field(i)
			baseField := base.Field(i)

			if baseField.CanSet() {
				mergeValues(baseField, overrideField)
			}
		}

	case reflect.String:
		if override.String() != "" {
			base.SetString(override.String())
		}

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if override.Int() != 0 {
			base.SetInt(override.Int())
		}

	case reflect.Float32, reflect.Float64:
		if override.Float() != 0 {
			base.SetFloat(override.Float())
		}

	case reflect.Bool:
		base.SetBool(override.Bool())

	case reflect.Slice:
		if override.Len() > 0 {
			base.Set(override)
		}

	case reflect.Map:
		if override.Len() > 0 {
			base.Set(override)
		}

	case reflect.Ptr:
		if !override.IsNil() {
			if base.IsNil() {
				base.Set(reflect.New(base.Type().Elem()))
			}
			mergeValues(base.Elem(), override.Elem())
		}
	}
}
