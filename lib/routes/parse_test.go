package routes

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {

	// Setup Environment

	r := MakeRouter()

	// ! All Handlers is nil for testing purpose
	r.RegisterRoute("/", nil)
	r.RegisterRoute("/faq/project", nil)
	r.RegisterRoute("/project/add", nil)
	r.RegisterRoute("/project/:id", nil)
	r.RegisterRoute("/user/:id/project/:name", nil)
	r.RegisterRoute("/user/:id/posts/:name/show", nil)

	// End of Setup

	type args struct {
		path       string
		wantRoute  *Route
		wantParams RouteParams
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "Parse Static Route [OK]",
			args: args{
				path: "/",
				wantRoute: &Route{
					Kind:    "common",
					Path:    "", // Index case
					Param:   "",
					Handler: nil,
					Routes:  MakeRouter(),
				},
				wantParams: RouteParams{},
			},
		},
		{
			name: "Parse Static Route With Two Parts [OK]",
			args: args{
				path: "/faq/project",
				wantRoute: &Route{
					Kind:    "common",
					Path:    "project",
					Param:   "",
					Handler: nil,
					Routes:  MakeRouter(),
				},
				wantParams: RouteParams{},
			},
		},
		{
			name: "Parse Static Route With Priority Over Dynamic Route [OK]",
			args: args{
				path: "/project/add",
				wantRoute: &Route{
					Kind:    "common",
					Path:    "add",
					Param:   "",
					Handler: nil,
					Routes:  MakeRouter(),
				},
				wantParams: RouteParams{},
			},
		},
		{
			name: "Parse Dynamic Route [OK]",
			args: args{
				path: "/project/1",
				wantRoute: &Route{
					Kind:    "special",
					Path:    "@",
					Param:   "id",
					Handler: nil,
					Routes:  MakeRouter(),
				},
				wantParams: RouteParams{"id": "1"},
			},
		},
		{
			name: "Parse Dynamic Route With Two Params [OK]",
			args: args{
				path: "/user/1/project/cool project",
				wantRoute: &Route{
					Kind:    "special",
					Path:    "@",
					Param:   "name",
					Handler: nil,
					Routes:  MakeRouter(),
				},
				wantParams: RouteParams{"id": "1", "name": "cool project"},
			},
		},
		{
			name: "Parse Dynamic Route With Two Params and Static Part in Final [OK]",
			args: args{
				path: "/user/1/posts/cool post/show",
				wantRoute: &Route{
					Kind:    "common",
					Path:    "show",
					Param:   "",
					Handler: nil,
					Routes:  MakeRouter(),
				},
				wantParams: RouteParams{"id": "1", "name": "cool post"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			route, params := parse(&r, tt.args.path)

			if !reflect.DeepEqual(route, tt.args.wantRoute) {
				t.Errorf("parse() got = %v, want %v", route, tt.args.wantRoute)
			}

			if !reflect.DeepEqual(params, tt.args.wantParams) {
				t.Errorf("parse() got1 = %v, want %v", params, tt.args.wantParams)
			}
		})
	}
}
