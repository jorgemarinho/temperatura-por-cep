package usecase

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"

	"github.com/jorgemarinho/temperatura-por-cep/internal/dto"
	"github.com/jorgemarinho/temperatura-por-cep/internal/entity"
	"github.com/jorgemarinho/temperatura-por-cep/internal/errors"
)

const (
	viaCepURL     = "https://viacep.com.br/ws/%s/json/"
	weatherAPI    = "http://api.weatherapi.com/v1/current.json?key=%s&q=%s"
	weatherAPIKey = "8887ae192b2343f9a32114928240104"
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
	if !isValidCep(b.BuscaCepInputDTO.Cep) {
		return dto.BuscaCepOutputDTO{}, &errors.HTTPError{Code: http.StatusUnprocessableEntity, Message: "CEP must have 8 digits and only contain numbers"}
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
		TempF: getTemperatureFahrenheit(temperatura.TempF),
		TempK: getTemperatureKelvin(temperatura.TempK),
	}, nil
}

func (b BuscaCepUseCase) BuscaCep(cep string) (*entity.Cep, error) {
	url := fmt.Sprintf(viaCepURL, cep)
	return b.makeHTTPRequestCep(url)
}

func (b BuscaCepUseCase) BuscaTemperatura(nomeCidade string) (*entity.Temperatura, error) {
	encodedNomeCidade := url.QueryEscape(nomeCidade)
	url := fmt.Sprintf(weatherAPI, weatherAPIKey, encodedNomeCidade)
	return b.makeHTTPRequestTemperatura(url)
}

func (b BuscaCepUseCase) makeHTTPRequestCep(url string) (*entity.Cep, error) {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}

	resp, err := client.Get(url)
	if err != nil {
		return nil, &errors.HTTPError{Code: http.StatusInternalServerError, Message: "error making HTTP request"}
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, &errors.HTTPError{Code: http.StatusInternalServerError, Message: "error reading response body"}
	}

	var tempResult map[string]interface{}
	if err := json.Unmarshal(body, &tempResult); err == nil {
		if errVal, ok := tempResult["erro"]; ok && errVal == "true" {
			return nil, &errors.HTTPError{Code: http.StatusNotFound, Message: "can not find zipcode"}
		}
	}

	var result entity.Cep
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, &errors.HTTPError{Code: http.StatusInternalServerError, Message: "error unmarshalling cep response"}
	}

	return &result, nil
}

func (b BuscaCepUseCase) makeHTTPRequestTemperatura(url string) (*entity.Temperatura, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, &errors.HTTPError{Code: http.StatusInternalServerError, Message: "error making HTTP request"}
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, &errors.HTTPError{Code: http.StatusInternalServerError, Message: "error reading response body"}
	}

	var response struct {
		Current entity.Temperatura `json:"current"`
	}

	if err := json.Unmarshal(body, &response); err != nil {
		return nil, &errors.HTTPError{Code: http.StatusInternalServerError, Message: "error unmarshalling temperatura response"}
	}

	return &response.Current, nil
}

func isValidCep(cep string) bool {
	return regexp.MustCompile(`^\d{8}$`).MatchString(cep)
}

func getTemperatureFahrenheit(celsius float64) float64 {
	return (celsius * 1.8) + 32
}

func getTemperatureKelvin(celsius float64) float64 {
	return celsius + 273.15
}
