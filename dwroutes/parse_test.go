package dwroutes

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {

	// Setup Environment

	r := MakeRouter()

	// All Handlers is nil for testing purpose
	r.RegisterRoute(GET, "/", nil)
	r.RegisterRoute(GET, "/faq/project", nil)
	r.RegisterRoute(GET, "/project/add", nil)
	r.RegisterRoute(GET, "/project/<int:id>", nil)
	r.RegisterRoute(GET, "/user/<uuid:id>/project/<string:name>", nil)
	r.RegisterRoute(GET, "/user/<uuid:id>/posts/<string:name>/show", nil)

	r.Dump()

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
					Kind:      "common",
					Path:      "", // Index case
					Param:     "",
					ParamType: NULL,
					Handler:   nil,
					Routes:    MakeRouter(),
				},
				wantParams: RouteParams{},
			},
		},
		{
			name: "Parse Static Route With Two Parts [OK]",
			args: args{
				path: "/faq/project",
				wantRoute: &Route{
					Kind:      "common",
					Path:      "project",
					Param:     "",
					ParamType: NULL,
					Handler:   nil,
					Routes:    MakeRouter(),
				},
				wantParams: RouteParams{},
			},
		},
		{
			name: "Parse Static Route With Priority Over Dynamic Route [OK]",
			args: args{
				path: "/project/add",
				wantRoute: &Route{
					Kind:      "common",
					Path:      "add",
					Param:     "",
					ParamType: NULL,
					Handler:   nil,
					Routes:    MakeRouter(),
				},
				wantParams: RouteParams{},
			},
		},
		{
			name: "Parse Dynamic Route [OK]",
			args: args{
				path: "/project/1",
				wantRoute: &Route{
					Kind:      "special",
					Path:      "@",
					Param:     "id",
					ParamType: INT,
					Handler:   nil,
					Routes:    MakeRouter(),
				},
				wantParams: RouteParams{"id": 1},
			},
		},
		{
			name: "Parse Dynamic Route With Two Params [OK]",
			args: args{
				path: "/user/550e8400-e29b-41d4-a716-446655440000/project/cool project",
				wantRoute: &Route{
					Kind:      "special",
					Path:      "@",
					Param:     "name",
					ParamType: STRING,
					Handler:   nil,
					Routes:    MakeRouter(),
				},
				wantParams: RouteParams{"id": "550e8400-e29b-41d4-a716-446655440000", "name": "cool project"},
			},
		},
		{
			name: "Parse Dynamic Route With Two Params and Static Part in Final [OK]",
			args: args{
				path: "/user/550e8400-e29b-41d4-a716-446655440000/posts/cool post/show",
				wantRoute: &Route{
					Kind:      "common",
					Path:      "show",
					Param:     "",
					ParamType: NULL,
					Handler:   nil,
					Routes:    MakeRouter(),
				},
				wantParams: RouteParams{"id": "550e8400-e29b-41d4-a716-446655440000", "name": "cool post"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			route, params := parse(&r, tt.args.path, "GET")

			if !reflect.DeepEqual(route, tt.args.wantRoute) {
				t.Errorf("parse() got = %v, want %v", route, tt.args.wantRoute)
			}

			if !reflect.DeepEqual(params, tt.args.wantParams) {
				t.Errorf("parse() got1 = %v, want %v", params, tt.args.wantParams)
			}
		})
	}
}
