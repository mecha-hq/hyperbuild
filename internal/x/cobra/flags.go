package cobra

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func InitEnvs(envPrefix string) *viper.Viper {
	runtimeViper := viper.New()

	runtimeViper.SetEnvPrefix(envPrefix)

	runtimeViper.AutomaticEnv()

	return runtimeViper
}

func BindFlags(cmd *cobra.Command, v *viper.Viper, logger func(v ...any), envPrefix string) {
	cmd.Flags().VisitAll(func(flag *pflag.Flag) {
		if strings.Contains(flag.Name, "-") {
			envSuffix := strings.ToUpper(strings.ReplaceAll(flag.Name, "-", "_"))

			env := envSuffix
			if envPrefix != "" {
				env = fmt.Sprintf("%s_%s", envPrefix, envSuffix)
			}

			if err := v.BindEnv(flag.Name, env); err != nil {
				logger(err)
			}
		}

		if !flag.Changed && v.IsSet(flag.Name) {
			val := v.Get(flag.Name)

			if err := cmd.Flags().Set(flag.Name, fmt.Sprintf("%v", val)); err != nil {
				logger(err)
			}
		}
	})
}
