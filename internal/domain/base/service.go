package base

// Service Service
type Service struct {
}

// ProvideService ProvideService
func ProvideService() Service {
	return Service{}
}

// Release Release
func (s *Service) Release() Release {
	return Release{Version: "v1.0", Release: "2021.07.14"}
}
