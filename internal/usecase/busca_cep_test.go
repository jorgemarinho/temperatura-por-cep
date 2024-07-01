package usecase

import (
	"testing"

	"github.com/jorgemarinho/temperatura-por-cep/internal/dto"
	"github.com/jorgemarinho/temperatura-por-cep/internal/entity"
)

func TestNewBuscaCepUseCase(t *testing.T) {
	type args struct {
		buscaCepInputDTO dto.BuscaCepInputDTO
	}
	tests := []struct {
		name string
		args args
		want *BuscaCepUseCase
	}{
		{
			name: "Teste de criação de um novo caso de uso",
			args: args{
				buscaCepInputDTO: dto.BuscaCepInputDTO{
					Cep: "72130360",
				},
			},
			want: &BuscaCepUseCase{
				BuscaCepInputDTO: dto.BuscaCepInputDTO{
					Cep: "74130011",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBuscaCepUseCase(tt.args.buscaCepInputDTO); got == nil {
				t.Errorf("NewBuscaCepUseCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuscaCepUseCase_Execute(t *testing.T) {
	type fields struct {
		BuscaCepInputDTO dto.BuscaCepInputDTO
	}
	tests := []struct {
		name    string
		fields  fields
		want    dto.BuscaCepOutputDTO
		wantErr bool
	}{
		{
			name: "Teste de execução do caso de uso",
			fields: fields{
				BuscaCepInputDTO: dto.BuscaCepInputDTO{
					Cep: "72130360",
				},
			},
			want: dto.BuscaCepOutputDTO{
				TempC: 0,
				TempF: 273,
				TempK: 32,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := BuscaCepUseCase{
				BuscaCepInputDTO: tt.fields.BuscaCepInputDTO,
			}
			got, err := b.Execute()
			if (err != nil) != tt.wantErr {
				t.Errorf("BuscaCepUseCase.Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.TempC != tt.want.TempC {
				t.Errorf("BuscaCepUseCase.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuscaCepUseCase_BuscaCep(t *testing.T) {
	type fields struct {
		BuscaCepInputDTO dto.BuscaCepInputDTO
	}
	type args struct {
		cep string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entity.Cep
		wantErr bool
	}{
		{
			name: "Teste de busca de CEP",
			fields: fields{
				BuscaCepInputDTO: dto.BuscaCepInputDTO{
					Cep: "72130360",
				},
			},
			args: args{
				cep: "72130360",
			},
			want: &entity.Cep{
				Cep:        "72130-360",
				Logradouro: "QNG 36",
				Bairro:     "Taguatinga Norte (Taguatinga)",
				Localidade: "Brasília",
				Uf:         "DF",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := BuscaCepUseCase{
				BuscaCepInputDTO: tt.fields.BuscaCepInputDTO,
			}
			got, err := b.BuscaCep(tt.args.cep)
			if (err != nil) != tt.wantErr {
				t.Errorf("BuscaCepUseCase.BuscaCep() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Cep != tt.want.Cep {
				t.Errorf("BuscaCepUseCase.BuscaCep() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuscaCepUseCase_BuscaTemperatura(t *testing.T) {
	type fields struct {
		BuscaCepInputDTO dto.BuscaCepInputDTO
	}
	type args struct {
		localidade string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entity.Temperatura
		wantErr bool
	}{
		{
			name: "Teste de busca de temperatura",
			fields: fields{
				BuscaCepInputDTO: dto.BuscaCepInputDTO{
					Cep: "72130360",
				},
			},
			args: args{
				localidade: "Brasília",
			},
			want: &entity.Temperatura{
				TempC: 0,
				TempF: 273,
				TempK: 32,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := BuscaCepUseCase{
				BuscaCepInputDTO: tt.fields.BuscaCepInputDTO,
			}
			got, err := b.BuscaTemperatura(tt.args.localidade)
			if (err != nil) != tt.wantErr {
				t.Errorf("BuscaCepUseCase.BuscaTemperatura() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.TempC != tt.want.TempC {
				t.Errorf("BuscaCepUseCase.BuscaTemperatura() = %v, want %v", got, tt.want)
			}
		})
	}
}
