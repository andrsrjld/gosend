package gosend

type GoSendService struct {
	Repository *GoSendRepository
}

type GoSendAgent interface {
}

func NewGoSendService() GoSendAgent {
	db := InitDB()

	repo := NewGoSendRepository(db)
	return &GoSendService{
		Repository: repo,
	}
}
