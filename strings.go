package helpers

func filter(base []string, exclude []string) []string {
	filtered := []string{}
	for _, f := range base {
		if in(f, exclude) {
			continue
		}
		filtered = append(filtered, f)
	}
	return filtered
}

func wrapStrings(inputs []string, mark string) []string {
	outputs := []string{}
	for _, f := range inputs {
		outputs = append(outputs, mark+f+mark)
	}
	return outputs
}

func prependStrings(inputs []string, prefix string) []string {
	outputs := []string{}
	for _, f := range inputs {
		outputs = append(outputs, prefix+f)
	}
	return outputs
}