package dto

type BuscaCepInputDTO struct {
	Cep string `json:"cep"`
}

type BuscaCepOutputDTO struct {
	TempC float64 `json:"temp_C"`
	TempF float64 `json:"temp_F"`
	TempK float64 `json:"temp_K"`
}
