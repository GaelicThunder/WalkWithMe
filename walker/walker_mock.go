package walker

import "time"

type MockedWalker struct{}

func (m *MockedWalker) PositionComputation(from, to string) (string, time.Duration, error) {
	return "1 Chome Jingumae, Shibuya City, Tokyo 150-0001, Japan", 1 * time.Hour, nil
}
