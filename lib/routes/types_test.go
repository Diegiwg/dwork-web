package routes

import (
	"reflect"
	"testing"
)

func TestParseParamType(t *testing.T) {
	type args struct {
		value     string
		paramType ParamTypes
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name: "Parse::String [OK]",
			args: args{
				value:     "test",
				paramType: STRING,
			},
			want:    "test",
			wantErr: false,
		},
		{
			name: "Parse::Int [OK]",
			args: args{
				value:     "1",
				paramType: INT,
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "Parse::Float [OK]",
			args: args{
				value:     "1.1",
				paramType: FLOAT,
			},
			want:    1.1,
			wantErr: false,
		},
		{
			name: "Parse::Bool [OK]",
			args: args{
				value:     "true",
				paramType: BOOL,
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "Parse::UUID [OK]",
			args: args{
				value:     "00000000-0000-0000-0000-000000000000",
				paramType: UUID,
			},
			want:    "00000000-0000-0000-0000-000000000000",
			wantErr: false,
		},
		{
			name: "Parse::Invalid [Error]",
			args: args{
				value:     "",
				paramType: STRING,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Parse::String [Error]",
			args: args{
				value:     "",
				paramType: STRING,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Parse::Int [Error]",
			args: args{
				value:     "a",
				paramType: INT,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Parse::Float [Error]",
			args: args{
				value:     "b",
				paramType: FLOAT,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Parse::Bool [Error]",
			args: args{
				value:     "c",
				paramType: BOOL,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Parse::UUID [Error]",
			args: args{
				value:     "d",
				paramType: UUID,
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseParamType(tt.args.value, tt.args.paramType, "")
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseParamType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseParamType() = %v, want %v", got, tt.want)
			}
		})
	}
}
