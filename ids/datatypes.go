package ids

import (
	"bufio"
	"embed"
	"fmt"
	"regexp"
	"strings"
	"sync"
)

//go:embed reference/official/v1.0.0/Documentation/DataTypes.md
var dataTypesDocs embed.FS

type dataTypeSpec struct {
	IFC2X3 bool
	IFC4   bool
	IFC43  bool
	Base   string
}

var (
	dataTypeSpecsOnce sync.Once
	dataTypeSpecs     map[string]dataTypeSpec
	dataTypeSpecsErr  error

	xmlBasePatterns = map[string]*regexp.Regexp{
		"xs:boolean":  regexp.MustCompile(`^(true|false|0|1)$`),
		"xs:date":     regexp.MustCompile(`^\d{4}-\d{2}-\d{2}(Z|([+-]\d{2}:\d{2}))?$`),
		"xs:dateTime": regexp.MustCompile(`^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|([+-]\d{2}:\d{2}))?$`),
		"xs:double":   regexp.MustCompile(`^([-+]?((\d+(\.\d*)?)|(\.\d+))([eE][-+]?\d+)?|NaN|\+INF|-INF)$`),
		"xs:duration": regexp.MustCompile(`^[-+]?P(\d+Y)?(\d+M)?(\d+D)?(T(\d+H)?(\d+M)?(\d+S)?)?$`),
		"xs:integer":  regexp.MustCompile(`^[+-]?(\d+)$`),
		"xs:string":   regexp.MustCompile(`(?s)^.*$`),
		"xs:time":     regexp.MustCompile(`^\d{2}:\d{2}:\d{2}(\.\d+)?(Z|([+-]\d{2}:\d{2}))?$`),
	}
)

func lookupDataTypeSpec(name string) (dataTypeSpec, bool, error) {
	dataTypeSpecsOnce.Do(func() {
		dataTypeSpecs, dataTypeSpecsErr = parseDataTypeSpecs()
	})
	if dataTypeSpecsErr != nil {
		return dataTypeSpec{}, false, dataTypeSpecsErr
	}
	spec, ok := dataTypeSpecs[name]
	return spec, ok, nil
}

func parseDataTypeSpecs() (map[string]dataTypeSpec, error) {
	data, err := dataTypesDocs.ReadFile("reference/official/v1.0.0/Documentation/DataTypes.md")
	if err != nil {
		return nil, fmt.Errorf("ids: read embedded DataTypes.md: %w", err)
	}

	specs := map[string]dataTypeSpec{}
	scanner := bufio.NewScanner(strings.NewReader(string(data)))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if !strings.HasPrefix(line, "| IFC") {
			continue
		}

		parts := strings.Split(line, "|")
		if len(parts) < 6 {
			continue
		}
		name := strings.TrimSpace(parts[1])
		specs[name] = dataTypeSpec{
			IFC2X3: strings.Contains(parts[2], "✔"),
			IFC4:   strings.Contains(parts[3], "✔"),
			IFC43:  strings.Contains(parts[4], "✔"),
			Base:   strings.TrimSpace(parts[5]),
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("ids: scan embedded DataTypes.md: %w", err)
	}
	return specs, nil
}

func (spec dataTypeSpec) supports(version IFCVersion) bool {
	switch version {
	case IFCVersionIFC2X3:
		return spec.IFC2X3
	case IFCVersionIFC4:
		return spec.IFC4
	case IFCVersionIFC4X3, IFCVersionIFC4X3ADD2:
		return spec.IFC43
	default:
		return false
	}
}
