package usecase

import (
	"testing"

	"github.com/jorgemarinho/temperatura-por-cep/internal/dto"
	"github.com/jorgemarinho/temperatura-por-cep/internal/entity"
	"github.com/jorgemarinho/temperatura-por-cep/internal/usecase"
)

func TestNewBuscaCepUseCase(t *testing.T) {
	type args struct {
		buscaCepInputDTO dto.BuscaCepInputDTO
	}
	tests := []struct {
		name string
		args args
		want *usecase.BuscaCepUseCase
	}{
		{
			name: "Teste de criação de um novo caso de uso",
			args: args{
				buscaCepInputDTO: dto.BuscaCepInputDTO{
					Cep: "72130360",
				},
			},
			want: &usecase.BuscaCepUseCase{
				BuscaCepInputDTO: dto.BuscaCepInputDTO{
					Cep: "74130011",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := usecase.NewBuscaCepUseCase(tt.args.buscaCepInputDTO); got == nil {
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
		wantErr bool
	}{
		{
			name: "Teste de execução do caso de uso",
			fields: fields{
				BuscaCepInputDTO: dto.BuscaCepInputDTO{
					Cep: "72130360",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := usecase.BuscaCepUseCase{
				BuscaCepInputDTO: tt.fields.BuscaCepInputDTO,
			}
			got, err := b.Execute()
			if (err != nil) != tt.wantErr {
				t.Errorf("BuscaCepUseCase.Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got.TempC == 0 || got.TempF == 0 || got.TempK == 0 {
				t.Errorf("BuscaCepUseCase.Execute() returned invalid temperatures: %v", got)
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
			b := usecase.BuscaCepUseCase{
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
