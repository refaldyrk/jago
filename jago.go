package main

func Go(kata string) string {
	var jamet string

	for i := 0; i < len(kata); i++ {

		if kata[i] == 'a' {
			jamet += "4"
		} else if kata[i] == 'A' {
			jamet += "@"
		} else if kata[i] == 'i' {
			jamet += "1"
		} else if kata[i] == 'u' {
			jamet += "v"
		} else if kata[i] == 'e' {
			jamet += "3"
		} else if kata[i] == 'o' {
			jamet += "0"
		} else {
			jamet += string(kata[i])
		}
	}

	return jamet
}
