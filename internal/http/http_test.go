package http

import "testing"

func TestStatus_Get(t *testing.T) {
	type fields struct {
		code int
		text string
	}
	type args struct {
		page string
		auth string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "valid status and inputs",
			fields: fields{
				code: 200,
				text: "OK",
			},
			args: args{
				page: "https://www.google.com",
				auth: "valid_token",
			},
			wantErr: false,
		},
		{name: "invalid status Code",
			fields: fields{
				code: 400,
				text: "Bad Request",
			},
			args: args{
				page: "https://www.google.com'",
				auth: "invalid_token",
			},
			wantErr: true,
		},
		{name: "invalid page input",
			fields: fields{
				code: 200,
				text: "OK",
			},
			args: args{
				page: "invalid-url",
				auth: "valid_token",
			},
			wantErr: true,
		},
		{name: "Unauthorized access",
			fields: fields{
				code: 401,
				text: "Unauthorized",
			},
			args: args{
				page: "/dashboard",
				auth: "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Status{
				Code: tt.fields.code,
				Text: tt.fields.text,
			}
			if err := s.get(tt.args.page, tt.args.auth); err != nil != tt.wantErr {
				t.Errorf("get() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
