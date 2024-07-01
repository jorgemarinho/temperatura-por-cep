package usecase

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "net/url"
    "os"
    "regexp"

    "github.com/jorgemarinho/temperatura-por-cep/internal/dto"
    "github.com/jorgemarinho/temperatura-por-cep/internal/entity"
)

const (
    viaCepURL  = "https://viacep.com.br/ws/%s/json/"
    weatherAPI = "http://api.weatherapi.com/v1/current.json?key=%s&q=%s"
)

var weatherAPIKey = os.Getenv("WEATHER_API_KEY")

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
        return dto.BuscaCepOutputDTO{}, fmt.Errorf("CEP must have 8 digits and only contain numbers")
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
        TempF: getTemperatureFahrenheit(temperatura.TempC),
        TempK: getTemperatureKelvin(temperatura.TempC),
    }, nil
}

func (b BuscaCepUseCase) BuscaCep(cep string) (*entity.Cep, error) {
    url := fmt.Sprintf(viaCepURL, cep)
    return b.makeHTTPRequest[entity.Cep](url)
}

func (b BuscaCepUseCase) BuscaTemperatura(nomeCidade string) (*entity.Temperatura, error) {
    encodedNomeCidade := url.QueryEscape(nomeCidade)
    url := fmt.Sprintf(weatherAPI, weatherAPIKey, encodedNomeCidade)
    return b.makeHTTPRequest[entity.Temperatura](url)
}

func (b BuscaCepUseCase) makeHTTPRequest[T any](url string) (*T, error) {
    resp, err := http.Get(url)
    if err != nil {
        return nil, fmt.Errorf("error making HTTP request: %w", err)
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, fmt.Errorf("error reading response body: %w", err)
    }

    var result T
    if err := json.Unmarshal(body, &result); err != nil {
        return nil, fmt.Errorf("error unmarshalling response: %w", err)
    }

    return &result, nil
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