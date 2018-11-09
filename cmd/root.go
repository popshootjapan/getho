package cmd

import (
	"fmt"
	"os"
	"path"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "getho",
	Short: "CLI tool for getho.io",
	Long: `this is a CLI tool for getho.io.
You can get nodes, contracts and accounts, also can upload your smart contract to your node.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.getho-cli.yaml)")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func exists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}

func initConfig() {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	condigDir := path.Join(home, ".getho")
	configName := "config"
	configPath := path.Join(condigDir, fmt.Sprintf("%s.%s", configName, "json"))

	if _, err := os.Stat(condigDir); os.IsNotExist(err) {
		if err = os.Mkdir(condigDir, 0755); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	if !exists(configPath) {
		file, err := os.OpenFile(configPath, os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer file.Close()
		fmt.Fprintln(file, "{}")
	}

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(condigDir)
		viper.SetConfigName(configName)
	}

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
