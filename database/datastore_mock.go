package database

type MockDatabase struct {
	stored map[string]*WalkStatus
}

func (m *MockDatabase) GetWalk() (*WalkStatus, error) {
	return &WalkStatus{From: "Tokyo", To: "Kyoto", ActualPosition: "4 Chome-2-8 Shibakoen, Minato City, Tokyo 105-0011, Japan", Status: "actual", ID: "mock"}, nil
}

func (m *MockDatabase) SaveWalk(walk *WalkStatus) error {
	m.stored[walk.ID] = walk
	return nil
}

func NewMockDatabase() Database {
	return &MockDatabase{
		stored: make(map[string]*WalkStatus),
	}
}
