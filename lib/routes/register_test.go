package routes

import (
	"testing"
)

func TestRegisterRoute(t *testing.T) {

	r := MakeRouter()

	type args struct {
		routes   *Routes
		path     string
		expected []string
		handler  RouteHandler
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "Register Static Route [OK]",
			args: args{
				routes:   &r,
				path:     "/",
				expected: []string{""}, // Index case
				handler:  nil,
			},
		},
		{
			name: "Register Static Route With Two Parts [OK]",
			args: args{
				routes:   &r,
				path:     "/faq/project",
				expected: []string{"faq", "project"},
				handler:  nil,
			},
		},
		{
			name: "Register Dynamic Route [OK]",
			args: args{
				routes:   &r,
				path:     "/project/:id",
				expected: []string{"project", "@"},
				handler:  nil,
			},
		},
		{
			name: "Register Dynamic Route With Two Params [OK]",
			args: args{
				routes:   &r,
				path:     "/user/:id/project/:name",
				expected: []string{"user", "@", "project", "@"},
				handler:  nil,
			},
		},
		{
			name: "Register Dynamic Route With Two Params and Static Part in Final [OK]",
			args: args{
				routes:   &r,
				path:     "/user/:id/project/:name/show",
				expected: []string{"user", "@", "project", "@", "show"},
				handler:  nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r.RegisterRoute(tt.args.path, tt.args.handler)

			node := tt.args.routes
			for _, part := range tt.args.expected {
				temp := (*node)[part]

				if temp == nil {
					t.Fail()
				} else {
					node = &temp.Routes
				}

			}
		})
	}
}
