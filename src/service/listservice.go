package service

//ListServices ...
//List services available
func ListServices() *[]string {
	var servs []string
	for key := range services {
		servs = append(servs, key)
	}

	return &servs
}
