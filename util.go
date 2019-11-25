package main

func contains(r []string, s string) bool {
	for _, s2 := range r {
		if s == s2 {
			return true
		}
	}
	return false
}