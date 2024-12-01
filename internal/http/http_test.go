package http

import "testing"

func TestStatus_get(t *testing.T) {
	type fields struct {
		Code []int
		Text []string
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
			args: args{
				page: "https://www.google.com",
				auth: "valid_token",
			},
			wantErr: false,
		},
		{name: "invalid auth",
			args: args{
				page: "https://www.google.com'",
				auth: "invalid_token",
			},
			wantErr: true,
		},
		{name: "invalid urls input",
			args: args{
				page: "invalid-url",
				auth: "valid_token",
			},
			wantErr: true,
		},
		{name: "Unauthorized access",
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
				Code: tt.fields.Code,
				Text: tt.fields.Text,
			}
			if err := s.get(tt.args.page, tt.args.auth); (err != nil) != tt.wantErr {
				t.Errorf("get() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
