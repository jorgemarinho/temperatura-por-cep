package usecase

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/jorgemarinho/temperatura-por-cep/internal/dto"
	"github.com/jorgemarinho/temperatura-por-cep/internal/entity"
)

type BuscaCepUseCase struct {
	BuscaCepInputDTO dto.BuscaCepInputDTO
}

func NewBuscaCepUseCase(buscaCepInputDTO dto.BuscaCepInputDTO) *BuscaCepUseCase {
	return &BuscaCepUseCase{
		BuscaCepInputDTO: buscaCepInputDTO,
	}
}

func (b BuscaCepUseCase) Execute() (dto.BuscaCepOutputDTO, error) {

	if len(b.BuscaCepInputDTO.Cep) < 8 {
		err := fmt.Errorf("CEP must have 8 digits")
		return dto.BuscaCepOutputDTO{}, err
	}

	cep, err := b.BuscaCep(b.BuscaCepInputDTO.Cep)

	if err != nil {
		return dto.BuscaCepOutputDTO{}, err
	}

	temperatura, err := b.BuscaTemperatura(cep.Localidade)

	if err != nil {
		return dto.BuscaCepOutputDTO{}, err
	}

	return dto.BuscaCepOutputDTO{
		TempC: temperatura.TempC,
		TempF: getTemperatureKelvin(temperatura.TempF),
		TempK: getTemperatureFahrenheit(temperatura.TempK),
	}, nil
}

func (b BuscaCepUseCase) BuscaCep(cep string) (*entity.Cep, error) {

	resp, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var c entity.Cep
	err = json.Unmarshal(body, &c)

	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (b BuscaCepUseCase) BuscaTemperatura(nomeCidade string) (*entity.Temperatura, error) {

	apiKey := "8887ae192b2343f9a32114928240104"
	returnNomeCidade := url.QueryEscape(nomeCidade)
	url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s", apiKey, returnNomeCidade)
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var t entity.Temperatura
	err = json.Unmarshal(body, &t)

	if err != nil {
		return nil, err
	}

	return &t, nil
}

func getTemperatureFahrenheit(celsius float64) float64 {
	return (celsius * 1.8) + 32
}

func getTemperatureKelvin(celsius float64) float64 {
	return celsius + 273
}
