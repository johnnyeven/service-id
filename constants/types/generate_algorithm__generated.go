package types

import (
	"bytes"
	"encoding"
	"errors"

	profzone_libtools_courier_enumeration "github.com/profzone/libtools/courier/enumeration"
)

var InvalidGenerateAlgorithm = errors.New("invalid GenerateAlgorithm")

func init() {
	profzone_libtools_courier_enumeration.RegisterEnums("GenerateAlgorithm", map[string]string{
		"SNOWFLAKE": "SnowFlake",
	})
}

func ParseGenerateAlgorithmFromString(s string) (GenerateAlgorithm, error) {
	switch s {
	case "":
		return GENERATE_ALGORITHM_UNKNOWN, nil
	case "SNOWFLAKE":
		return GENERATE_ALGORITHM__SNOWFLAKE, nil
	}
	return GENERATE_ALGORITHM_UNKNOWN, InvalidGenerateAlgorithm
}

func ParseGenerateAlgorithmFromLabelString(s string) (GenerateAlgorithm, error) {
	switch s {
	case "":
		return GENERATE_ALGORITHM_UNKNOWN, nil
	case "SnowFlake":
		return GENERATE_ALGORITHM__SNOWFLAKE, nil
	}
	return GENERATE_ALGORITHM_UNKNOWN, InvalidGenerateAlgorithm
}

func (GenerateAlgorithm) EnumType() string {
	return "GenerateAlgorithm"
}

func (GenerateAlgorithm) Enums() map[int][]string {
	return map[int][]string{
		int(GENERATE_ALGORITHM__SNOWFLAKE): {"SNOWFLAKE", "SnowFlake"},
	}
}
func (v GenerateAlgorithm) String() string {
	switch v {
	case GENERATE_ALGORITHM_UNKNOWN:
		return ""
	case GENERATE_ALGORITHM__SNOWFLAKE:
		return "SNOWFLAKE"
	}
	return "UNKNOWN"
}

func (v GenerateAlgorithm) Label() string {
	switch v {
	case GENERATE_ALGORITHM_UNKNOWN:
		return ""
	case GENERATE_ALGORITHM__SNOWFLAKE:
		return "SnowFlake"
	}
	return "UNKNOWN"
}

var _ interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
} = (*GenerateAlgorithm)(nil)

func (v GenerateAlgorithm) MarshalText() ([]byte, error) {
	str := v.String()
	if str == "UNKNOWN" {
		return nil, InvalidGenerateAlgorithm
	}
	return []byte(str), nil
}

func (v *GenerateAlgorithm) UnmarshalText(data []byte) (err error) {
	*v, err = ParseGenerateAlgorithmFromString(string(bytes.ToUpper(data)))
	return
}
