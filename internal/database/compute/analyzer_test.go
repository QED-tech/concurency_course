package compute

import (
	"database/internal/database/commands"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAnalyzer_Analyze(t *testing.T) {
	t.Parallel()

	type args struct {
		tokens Tokens
	}

	tests := []struct {
		name    string
		args    args
		want    commands.Command
		wantErr error
	}{
		{
			name: "Should create a command",
			args: args{
				tokens: []string{"GET", "key"},
			},
			want: commands.Command{
				Operation: "GET",
				Arguments: []string{"key"},
			},
			wantErr: nil,
		},
		{
			name: "Should getting error - InvalidCommand",
			args: args{
				tokens: []string{"GE", "key", "value"},
			},
			want:    commands.Command{},
			wantErr: commands.ErrInvalidCommand,
		},
		{
			name: "Should getting error - InvalidArguments",
			args: args{
				tokens: []string{"DEL", "key", "value"},
			},
			want:    commands.Command{},
			wantErr: commands.ErrInvalidArgumentsCount,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			a := NewAnalyzer()

			got, err := a.Analyze(tt.args.tokens)

			is := require.New(t)

			if tt.wantErr != nil {
				is.EqualError(err, tt.wantErr.Error())
			}

			is.Equal(tt.want, got)
		})
	}
}
