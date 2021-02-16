package primes

//IsPrime проверяет, делится ли переданное число a на каждое из уже найденных простых чисел из слайса primes
//в случае, если число a кратно хотя бы одному из уже найденных простых чисел, то число a простым не является
func IsPrime(a int, primes []int) bool {
	for i := 0; i < len(primes); i++ {
		if a%primes[i] == 0 {
			return false
		}
	}
	return true
}
