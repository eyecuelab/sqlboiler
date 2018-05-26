package boilingcore

import "github.com/volatiletech/sqlboiler/bdb"

func nonNullVersion(c bdb.Column) string {
	if !c.Nullable {
		return c.Type
	}
	switch c.Type {
	case "null.Int64":
		return "int64"
	case "null.Int":
		return "int"
	case "null.Int16":
		return "int16"
	case "null.Float64":
		return "float64"
	case "null.Float32":
		return "float32"
	case "null.String":
		return "string"
	case "null.Byte":
		return "byte"
	case "null.JSON":
		return "*map[string]interface{}"  //have to use a pointer to map here because of jsonapi
	case "null.Bytes":
		return "[]byte"
	case "null.Bool":
		return "bool"
	case "null.Time":
		return "time.Time"
	case "types.Int64Array":
		return "[]int64"
	case "types.BytesArray":
		return "[][]byte"
	case "types.StringArray":
		return "[]string"
	case "types.BoolArray":
		return "[]bool"
	case "types.Float64Array":
		return "[]float64"
	}
	return ""
}

