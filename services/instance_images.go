package services

func InstanceImages() []string {

	return []string{
		dindImage,
		"franela/dind:overlay2-dev",
		"franela/dind:overlay2-ucp",
	}

}
