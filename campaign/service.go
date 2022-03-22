package campaign

type Service interface {
	GetCampaigns(userID int) ([]Campaign, error) 
	GetCampaignByID(input GetCampaignDetailInput) (Campaign, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func(s *service) GetCampaigns(userID int) ([]Campaign, error) {
	if userID != 0 {
		Campaigns, err := s.repository.FindByUserId(userID)
		if err != nil {
			return Campaigns, err
		}

		return Campaigns, nil
	}
	Campaigns, err := s.repository.FindAll()
	if err != nil {
		return Campaigns, err
	}

	return Campaigns, nil
}

func (s *service) GetCampaignByID(input GetCampaignDetailInput) (Campaign, error) {
	campaign, err := s.repository.FindById(input.ID)

	if err != nil {
		return campaign, err
	}

	return campaign, nil

}


