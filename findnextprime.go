package piscine

func FindNextPrime(nb int) int {
	in := nb
	for IsPrimeT(in) == false {
		in++
	}
	return in
}

func IsPrimeT(nb int) bool {
	if nb < 2 {
		return false
	}
	for i := 2; i*i <= nb; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}
