package assets

type AssertModule struct {
	available float64 // 可用余额
	frozen    float64 // 冻结余额
}

func (m *AssertModule) GetFrozen() float64 {
	return m.frozen
}

func (m *AssertModule) GetAvailable() float64 {
	return m.available
}
func (m *AssertModule) GetTotal() float64 {
	return m.GetFrozen() + m.GetAvailable()
}
func NewAsset(available, frozen float64) *AssertModule {
	return &AssertModule{
		available: available,
		frozen:    frozen,
	}
}
