package provider

import (
	"context"
	"github.com/selefra/selefra-provider-k8s/k8s_client"
	"strings"

	"github.com/selefra/selefra-provider-sdk/provider"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/spf13/viper"
)

const Version = "v0.0.1"

func GetProvider() *provider.Provider {
	return &provider.Provider{
		Name:      "k8s",
		Version:   Version,
		TableList: GenTables(),
		ClientMeta: schema.ClientMeta{
			InitClient: func(ctx context.Context, clientMeta *schema.ClientMeta, config *viper.Viper) ([]any, *schema.Diagnostics) {

				diagnostics := schema.NewDiagnostics()

				//k8sConfigPath := config.GetString("providers.0.config-path")
				k8sConfigPath := config.GetString("config-path")
				if k8sConfigPath != "" {
					if !strings.HasPrefix(k8sConfigPath, "/") && !strings.HasPrefix(k8sConfigPath, "./") {
						return nil, diagnostics.AddErrorMsg("config-path must start with / or ./, %s is not ok", k8sConfigPath)
					}
				}
				client, err := k8s_client.NewClient(ctx, k8sConfigPath, nil)
				if err != nil {
					return nil, schema.NewDiagnostics().AddError(err)
				}
				return []any{client}, nil
			},
		},
		ConfigMeta: provider.ConfigMeta{
			GetDefaultConfigTemplate: func(ctx context.Context) string {
				return `#providers:
#    - name: k8s
#      config-path: `
			},
			Validation: func(ctx context.Context, config *viper.Viper) *schema.Diagnostics {
				return nil
			},
		},
		TransformerMeta: schema.TransformerMeta{
			DefaultColumnValueConvertorBlackList: []string{
				"",
			},
			DataSourcePullResultAutoExpand: true,
		},
		ErrorsHandlerMeta: schema.ErrorsHandlerMeta{
			IgnoredErrors: []schema.IgnoredError{schema.IgnoredErrorOnSaveResult},
		},
	}
}
