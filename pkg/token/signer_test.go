package token

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewSigner(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    Signer
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "test",
			args:    args{"name"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewSigner(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewSigner() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSigner() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSignerVerifier(t *testing.T) {
	signer, err := NewSigner("mysecret")
	if err != nil {
		panic(err)
	}
	verifier, err := NewVerifier("mysecret")
	if err != nil {
		panic(err)
	}

	content := "details"
	j, err := signer.Sign(&content)
	if err != nil {
		panic(err)
	}

	repContent, err := verifier.Verify(j)
	if err != nil {
		panic(err)
	}

	fmt.Printf("reproduced %v", *repContent)

}
