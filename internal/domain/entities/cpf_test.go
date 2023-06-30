package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCPF(t *testing.T) {
	type args struct {
		cpf string
	}
	validCPF := CPF("47397571093")
	tests := []struct {
		name    string
		args    args
		out     *CPF
		wantErr bool
		errOut  error
	}{
		{
			name: "should return an ErrCPFInvalidLength when cpf has less than 11 characters",
			args: args{
				cpf: "123.123.123-1",
			},
			out:     nil,
			wantErr: true,
			errOut:  ErrCPFInvalidLength,
		},
		{
			name: "should return an ErrCPFInvalidLength when cpf has more than 11 characters",
			args: args{
				cpf: "123.123.123-122",
			},
			out:     nil,
			wantErr: true,
			errOut:  ErrCPFInvalidLength,
		},
		{
			name: "should return an ErrCPFNonDigit if cpf have non digit characters",
			args: args{
				cpf: "123.1fds.123-122",
			},
			out:     nil,
			wantErr: true,
			errOut:  ErrCPFNonDigit,
		},
		{
			name: "should return an InvalidCPF error when cpf is invalid when d1 is invalid",
			args: args{
				cpf: "473.975.710-83",
			},
			out:     nil,
			wantErr: true,
			errOut:  ErrCPFInvalid,
		},
		{
			name: "should return an InvalidCPF error when cpf is invalid when d2 is invalid",
			args: args{
				cpf: "473.975.710-91",
			},
			out:     nil,
			wantErr: true,
			errOut:  ErrCPFInvalid,
		},
		{
			name: "should return a CPF on valid string",
			args: args{
				cpf: validCPF.String(),
			},
			out:     &validCPF,
			wantErr: false,
			errOut:  nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewCPF(tt.args.cpf)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tt.errOut, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.out, got)
			}
		})
	}
}
