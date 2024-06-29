package usecase

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

type BuscaCepInputDTO struct {
	Cep string `json:"cep"`
}

type BuscaCepOutputDTO struct {
	TempC float64 `json:"temp_C"`
	TempF float64 `json:"temp_F"`
	TempK float64 `json:"temp_K"`
}

type BuscaCepUseCase struct {
	BuscaCepInputDTO BuscaCepInputDTO
}

func NewBuscaCepUseCase(buscaCepInputDTO BuscaCepInputDTO) *BuscaCepUseCase {
	return &BuscaCepUseCase{
		BuscaCepInputDTO: buscaCepInputDTO,
	}
}

func (b BuscaCepUseCase) Execute() (BuscaCepOutputDTO, error) {

	if len(b.BuscaCepInputDTO.Cep) < 8 {
		err := fmt.Errorf("CEP must have 8 digits")
		return BuscaCepOutputDTO{}, err
	}

	resp, err := http.Get("https://viacep.com.br/ws/" + b.BuscaCepInputDTO.Cep + "/json/")

	if err != nil {
		return BuscaCepOutputDTO{}, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return BuscaCepOutputDTO{}, err
	}

	var c ViaCEP
	err = json.Unmarshal(body, &c)

	if err != nil {
		return BuscaCepOutputDTO{}, err
	}

	return BuscaCepOutputDTO{
		TempC: b.TempC,
		TempF: b.TempF,
		TempK: b.TempK,
	}, nil
}
