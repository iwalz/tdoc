package cmd

import (
	"io/ioutil"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/davecgh/go-spew/spew"
	"github.com/dnephin/cobra"
	"github.com/iwalz/tdoc/elements"
	"github.com/iwalz/tdoc/parser"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	SvgDir  string
	verbose bool
	debug   bool
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "tdoc file.tdoc",
	Short: "Short description",
	Long:  `Long description`,
	Args:  CheckFile,
	RunE: func(cmd *cobra.Command, args []string) error {
		spew.Dump(verbose)
		if verbose {
			log.SetLevel(log.InfoLevel)
		}
		if debug {
			log.SetLevel(log.DebugLevel)
		}

		content, err := ioutil.ReadFile(args[0])
		if err != nil {
			return err
		}
		p := &parser.TdocParserImpl{}
		l := parser.NewLexer(string(content), elements.NewComponentsList(SvgDir))
		p.Parse(l)

		spew.Dump(args)

		return nil
	},
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		log.Error(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags, which, if defined here,
	// will be global for your application.
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.tdoc.yaml)")
	RootCmd.PersistentFlags().StringVar(&SvgDir, "svgdir", "/home/ingo/svg", "Source directory for components. foo.svg will make component foo available")
	RootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Enables verbose mode")
	RootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "Enables debug mode")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName(".tdoc") // name of config file (without extension)
	viper.AddConfigPath("$HOME") // adding home directory as first search path
	viper.AutomaticEnv()         // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		log.Info("Using config file:", viper.ConfigFileUsed())
	}
}
