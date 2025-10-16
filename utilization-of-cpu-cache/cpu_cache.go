package utilization_of_cpu_cache

var users []User

//type Image []byte

type Image [128 * 128]byte

type User struct {
	Login   string
	Active  bool
	Icon    Image
	Country string
}

func init() {
	const size = 10_000

	countries := []string{"US", "CA", "JP", "KR"}

	users := make([]User, size)

	for i := range users {
		users[i].Active = i%5 > 0
		users[i].Country = countries[i%len(countries)]
		//users[i].Icon = make([]byte, 128*128)
	}
}

func CountryCount(users []User) map[string]int {

	counts := make(map[string]int) // country -> count

	for _, u := range users {
		if !u.Active {
			continue
		}
		counts[u.Country]++
	}

	return counts
}
