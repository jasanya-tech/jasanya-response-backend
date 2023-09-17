package response

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
