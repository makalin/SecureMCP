package main

import (
	"fmt"
	"os"

	"github.com/makalin/SecureMCP/internal/scanner"
	"github.com/makalin/SecureMCP/internal/server"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "securemcp",
	Short: "SecureMCP - Security auditing tool for MCP applications",
	Long:  `A comprehensive security auditing tool designed to detect vulnerabilities and misconfigurations in applications using the Model Context Protocol (MCP).`,
}

var scanCmd = &cobra.Command{
	Use:   "scan [target]",
	Short: "Scan an MCP server for vulnerabilities",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		target := args[0]
		fmt.Printf("[+] Scanning Target: %s\n", target)

		scanner := scanner.NewScanner()
		results, err := scanner.Scan(target)
		if err != nil {
			fmt.Printf("Error scanning target: %v\n", err)
			os.Exit(1)
		}

		// Print results
		for _, result := range results {
			fmt.Printf("[!] %s\n", result)
		}
	},
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the SecureMCP server",
	Run: func(cmd *cobra.Command, args []string) {
		srv := server.NewServer()
		if err := srv.Start(); err != nil {
			fmt.Printf("Error starting server: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)
	rootCmd.AddCommand(serverCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
