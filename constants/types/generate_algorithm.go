package types

//go:generate libtools gen enum GenerateAlgorithm
// swagger:enum
type GenerateAlgorithm uint8

// 生成ID的算法
const (
	GENERATE_ALGORITHM_UNKNOWN    GenerateAlgorithm = iota
	GENERATE_ALGORITHM__SNOWFLAKE  // SnowFlake
)
