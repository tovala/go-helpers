package helpers

func Filter(base []string, exclude []string) []string {
	filtered := []string{}
	for _, f := range base {
		if In(f, exclude) {
			continue
		}
		filtered = append(filtered, f)
	}
	return filtered
}

func WrapStrings(inputs []string, mark string) []string {
	outputs := []string{}
	for _, f := range inputs {
		outputs = append(outputs, mark+f+mark)
	}
	return outputs
}

func PrependStrings(inputs []string, prefix string) []string {
	outputs := []string{}
	for _, f := range inputs {
		outputs = append(outputs, prefix+f)
	}
	return outputs
}
