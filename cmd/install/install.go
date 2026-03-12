/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/
package install

import (

	"github.com/spf13/cobra"
)

var (
	packageName string
)
// InstallCmd represents the install command
var InstallCmd = &cobra.Command{
	Use:   "install",
	Short: "Runs `npm install` after giving sonme warnings about the package",
	Long: "Runs `npm install` after printing pre-install warnings. For any package <P>, the warnings are:\n\n" +
	"1. <P> requires <NUM_DEPS> additional dependencies\n" +
	"   a. Direct: <NUM_DIRECT_DEPS>\n" +
	"   b. Transitive: <NUM_TRANSITIVE_DEPS>\n\n" +
	"2. <P> adds <NUM_MB> MB to project size\n\n" +
	"3. <P> would introduce <NUM_VULNERABILITIES> vulnerabilities\n\n" +
	"4. <P> would introduce <NUM_CONFLICTS> dependency conflicts\n\n" +
	"Using `-v` prints a detailed breakdown of each warning in the following format:\n\n" +
	"1. <P> requires <NUM_DEPS> additional dependencies\n" +
	"   a. Direct: <NUM_DIRECT_DEPS>\n" +
	"      i. {\"<DEP_NAME_1>\": \"<VERSION_1>\", \"<DEP_NAME_2>\": \"<VERSION_2>\", ...}\n" +
	"   b. Transitive: <NUM_TRANSITIVE_DEPS>\n" +
	"      i. {\"<DEP_NAME_1>\": \"<VERSION_1>\", \"<DEP_NAME_2>\": \"<VERSION_2>\", ...}\n\n" +
	"2. <P> adds <NUM_MB> MB to project size\n" +
	"   a. {\"<DEP_NAME_1>\": {\"version\": \"<VERSION_1>\", \"sizeMB\": <SIZE_MB_1>}, " +
	"\"<DEP_NAME_2>\": {\"version\": \"<VERSION_2>\", \"sizeMB\": <SIZE_MB_2>}, ...}\n\n" +
	"3. <P> would introduce <NUM_VULNERABILITIES> vulnerabilities\n" +
	"   a. {\"critical\": <NUM_CRITICAL>, \"high\": <NUM_HIGH>, \"moderate\": <NUM_MODERATE>, \"low\": <NUM_LOW>}\n" +
	"   b. {\"<VULN_ID_1>\": {\"package\": \"<DEP_NAME>\", \"severity\": \"<SEVERITY>\", \"title\": \"<TITLE>\"}, ...}\n\n" +
	"4. <P> would introduce <NUM_CONFLICTS> dependency conflicts\n" +
	"   a. [{\"package\": \"<DEP_NAME_1>\", \"version\": \"<VERSION_1>\", " +
	"\"conflictsWith\": {\"package\": \"<DEP_NAME_2>\", \"version\": \"<VERSION_2>\"}, " +
	"\"reason\": \"<PEER_DEP_OR_ENGINE_REASON>\"}, ...]\n\n" +
	"Using `--json` outputs the same analysis as structured JSON",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	InstallCmd.Flags().StringVarP(&packageName, "package", "p", "", "The package to install")
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// installCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// installCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


