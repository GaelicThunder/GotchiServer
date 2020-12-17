package main

type MockDatabase struct{}

func (m *MockDatabase) GetGotchi(id string) ([]string, error) {
	return []string{"id1", "id2", "id3"}, nil
}

func (m *MockDatabase) SaveGotchi(myGotchiID, newGotchiID string) error {
	return nil
}

func NewMockDatabase() Database {
	return &MockDatabase{}
}
