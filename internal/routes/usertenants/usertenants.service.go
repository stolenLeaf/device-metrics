package usertenants

type UserTenantsService struct {
	Repo *UserTenantsRepo
}

func NewUserTenantsService(repo *UserTenantsRepo) *UserTenantsService {
	return &UserTenantsService{Repo: repo}
}

func (s *UserTenantsService) GetUsers() ([]UserTenants, error) {
	return s.Repo.GetAll()
}

func (s *UserTenantsService) CreateUserTenants(u *UserTenants) error {
	return s.Repo.Create(u)
}
