package testutils

type MockFileEvaluator struct {
	ShouldReturnPulumiStateFile bool
}

func (m *MockFileEvaluator) IsPulumiStateFile(path string) (bool, error) {
	return m.ShouldReturnPulumiStateFile, nil
}
