package output

import (
	"encoding/json"

	"github.com/ArctisDev/frontgate/internal/types"
)

func JSONText(v interface{}) string {
	body, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return "{}"
	}
	return string(body)
}

func Success(v interface{}) types.ToolResult {
	return types.ToolResult{
		Text:           JSONText(v),
		StructuredData: v,
	}
}

func Error(err error) types.ToolResult {
	return types.ToolResult{
		Text:    err.Error(),
		IsError: true,
	}
}
