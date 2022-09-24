package campaign

type Service interface {
	//GetAllCampaigns() ([]Campaign, error)
	GetCampaigns(userId int) ([]Campaign, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

// func (s *service) GetAllCampaigns() ([]Campaign, error) {
// 	campaigns, err := s.repository.FindAll()
// 	if err != nil {
// 		return campaigns, err
// 	}
// 	return campaigns, nil
// }

func (s *service) GetCampaigns(userId int) ([]Campaign, error) {

	if userId != 0 {
		campaigns, err := s.repository.FindByUserId(userId)
		if err != nil {
			return campaigns, err
		}
		return campaigns, nil
	}

	campaigns, err := s.repository.FindAll()
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil

}
