/*
Copyright © 2025 Sean Kennedy <seanpk@outlook.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-for-rays",
	Short: "A ray tracer in Go",
	Long: `Following the book "Ray Tracer Challenge" by Jamis Buck, this project implements a ray tracer in Go.
It covers the fundamental concepts of ray tracing, including rays, intersections, transformations, lighting, and shadows.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
