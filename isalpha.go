package piscine

func IsAlpha(s string) bool {
	flag := true
	for _, r := range s {
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') && (r < '0' || r > '9') {
			flag = false
		}
	}
	return flag
}

// func main() {
// 	fmt.Println(IsAlpha("Whatthis4"))
// }
