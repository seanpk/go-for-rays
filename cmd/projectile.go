/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/seanpk/go-for-rays/internal/geometry"
	"github.com/spf13/cobra"
)

// projectileCmd represents the projectile command
var projectileCmd = &cobra.Command{
	Use:   "projectile",
	Short: "A projectile simulation",
	Long: `This command simulates the motion of a projectile under the influence of gravity and wind.
It provides options to configure the initial location and velocity, and the environmental factors affecting the projectile's trajectory.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		launchPointStr, _ := cmd.Flags().GetString("launch-point")
		velocityStr, _ := cmd.Flags().GetString("velocity")
		windStr, _ := cmd.Flags().GetString("wind")
		gravity, _ := cmd.Flags().GetFloat64("gravity")

		launchPoint, err := parseTuple(launchPointStr, parseTupleOptions{dimensions: 3, kind: "Point"})
		if err != nil {
			return fmt.Errorf("invalid launch point: %v", err)
		}

		velocity, err := parseTuple(velocityStr, parseTupleOptions{dimensions: 3, kind: "Vector"})
		if err != nil {
			return fmt.Errorf("invalid velocity: %v", err)
		}

		wind, err := parseTuple(windStr, parseTupleOptions{dimensions: 2, kind: "Vector"})
		if err != nil {
			return fmt.Errorf("invalid wind: %v", err)
		}

		gravityVector := geometry.NewVector(0, 0, -gravity)

		fmt.Printf("Projectile Simulation:\n")
		fmt.Printf("\tLaunch Point: %s\n", launchPoint.String())
		fmt.Printf("\tVelocity    : %s\n", velocity.String())
		fmt.Printf("\tWind        : %s\n", wind.String())
		fmt.Printf("\tGravity     : %s\n", gravityVector.String())

		return nil
	},
}

func init() {
	rootCmd.AddCommand(projectileCmd)

	projectileCmd.Flags().StringP("launch-point", "l", "", "Point from which the projectile is launched (x,y,z)")
	projectileCmd.Flags().StringP("velocity", "v", "", "Launch velocity vector of the projectile (x,y,z)")
	projectileCmd.Flags().StringP("wind", "w", "", "Wind vector (x,y)")
	projectileCmd.Flags().Float64P("gravity", "g", 9.81, "Gravity constant")
}
