package circuit

type API interface {
	ValidateCurrent(old, new float64) float64
}
