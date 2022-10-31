package k8s_client

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-sdk/provider"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/test_helper"
	"github.com/spf13/viper"
	"testing"
)

type TestOptions struct {
	SkipEmptyJsonB bool
}

func MockTestHelper(t *testing.T, table *schema.Table, builder func(*testing.T, *gomock.Controller) K8sServices, _ TestOptions) {
	ctrl := gomock.NewController(t)
	testProvider := newTestProvider(t, ctrl, table, builder)
	config := "test : test"
	test_helper.RunProviderPullTables(testProvider, config, "./", "*")
}

func newTestProvider(t *testing.T, ctrl *gomock.Controller, table *schema.Table, builder func(*testing.T, *gomock.Controller) K8sServices) *provider.Provider {
	return &provider.Provider{
		Name:		"k8s",
		Version:	"v0.0.1",
		TableList:	[]*schema.Table{table},
		ClientMeta: schema.ClientMeta{
			InitClient: func(ctx context.Context, clientMeta *schema.ClientMeta, config *viper.Viper) ([]any, *schema.Diagnostics) {

				services := builder(t, ctrl)
				client := &Client{
					Context: "testContext",
				}
				client.SetServices(map[string]K8sServices{"testContext": services})
				return []any{client}, nil
			},
		},
		ConfigMeta: provider.ConfigMeta{
			GetDefaultConfigTemplate: func(ctx context.Context) string {
				return ``
			},
			Validation: func(ctx context.Context, config *viper.Viper) *schema.Diagnostics {
				return nil
			},
		},
		TransformerMeta: schema.TransformerMeta{
			DefaultColumnValueConvertorBlackList:	[]string{},
			DataSourcePullResultAutoExpand:		true,
		},
		ErrorsHandlerMeta: schema.ErrorsHandlerMeta{
			IgnoredErrors: []schema.IgnoredError{schema.IgnoredErrorOnSaveResult},
		},
	}
}
