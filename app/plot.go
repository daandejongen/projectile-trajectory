package app

import (
	"log"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func Plot(trajectory Trajectory, path string) error {
    plot := plot.New()
	plot.Title.Text = "Projectile trajectory"
	plot.X.Label.Text = "x"
	plot.Y.Label.Text = "y"

	points := make(plotter.XYs, len(trajectory.Points))
	for i, _ := range points {
		points[i].X = trajectory.Points[i].Position.X
		points[i].Y = trajectory.Points[i].Position.Y
	}

	err := plotutil.AddLinePoints(plot, points)
	if err != nil {
		log.Fatalf("Failed to add line to plot: %v", err)
	}

	width := 6 * vg.Inch
	height := 4 * vg.Inch
	if err := plot.Save(width, height, path); err != nil {
		log.Fatalf("Failed to save plot: %v", err)
	}
	
	return nil
}