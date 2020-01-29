package accounting

import (
	"os"
	"testing"

	"github.com/markbates/goth"
	"github.com/medesk/xerogolang"
	"github.com/stretchr/testify/require"
)

func TestFindInvoices(t *testing.T) {
	accessToken := os.Getenv("TEST_XERO_ACCESS_TOKEN")
	tenantID := os.Getenv("TEST_XERO_TENANT_ID")

	type args struct {
		provider              *xerogolang.Provider
		session               goth.Session
		querystringParameters map[string]string
	}

	tests := []struct {
		name    string
		skip    bool
		args    args
		want    require.ValueAssertionFunc
		wantErr require.ErrorAssertionFunc
	}{
		{
			name: "should fail because access token is invalid",
			args: args{
				provider: xerogolang.NewOAuth2("invalid", tenantID),
			},
			want:    require.Nil,
			wantErr: require.Error,
		},
		{
			name: "should fail because tenant is invalid",
			args: args{
				provider: xerogolang.NewOAuth2(accessToken, "invalid"),
			},
			want:    require.Nil,
			wantErr: require.Error,
		},
		{
			name: "should get invoices",
			skip: accessToken == "" && tenantID == "",
			args: args{
				provider: xerogolang.NewOAuth2(accessToken, tenantID),
			},
			want:    require.NotNil,
			wantErr: require.NoError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.skip {
				t.Skip()
			}

			got, err := FindInvoices(tt.args.provider, tt.args.session, tt.args.querystringParameters)
			tt.wantErr(t, err)
			tt.want(t, got)
		})
	}
}
