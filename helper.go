package jasanya_response_backend

import (
	"fmt"
)

func ErrorLoginMessage(field1 string, field2 string) map[string][]string {
	return map[string][]string{
		field1: {
			fmt.Sprintf("invalid %s atau %s", field1, field2),
		},
		field2: {
			fmt.Sprintf("invalid %s atau %s", field2, field1),
		},
	}
}

func RegisterErrMapOfSlices(field string, msg string) map[string][]string {
	return map[string][]string{
		field: {
			msg,
		},
	}
}

func MergeMapOfSlices(maps ...map[string][]string) map[string][]string {
	merged := make(map[string][]string)
	for _, m := range maps {
		for key, value := range m {
			merged[key] = append(merged[key], value...)
		}
	}
	return merged
}
