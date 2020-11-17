package cast

// String ...
func String(v string) *string {
	return &v
}

// StringValue ...
func StringValue(v *string) string {
	if v != nil {
		return *v
	}

	return ""
}
