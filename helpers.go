package ga

func toStringSlice(is []interface{}) []string {
	ss := make([]string, 0, len(is))
	for _, i := range is {
		ss = append(ss, i.(string))
	}
	return ss
}
