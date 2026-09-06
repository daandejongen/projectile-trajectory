package cli

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/daandejongen/projectile-trajectory/app"
	"github.com/daandejongen/projectile-trajectory/trigonometry"
	"github.com/spf13/cobra"
)

const outputDir string = "output"
const plotDir string = "plots"
const logDir string = "logs"
const dataDir string = "data"

func NewSimulateCommand() *cobra.Command {
	var dragTypeInput string
	var addThrust bool
	var makePlot bool
	var saveData bool
	var logSimulatorConditions bool
	command := &cobra.Command{
		Use:   "sim [angle-in-degrees] [initial-speed-in-meters-per-second]",
		Short: "Simulate the trajectory of a projectile.",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			// Input
			angleInDegrees, angleParsingErr := strconv.ParseFloat(args[0], 64)
			speed, speedParsingError := strconv.ParseFloat(args[1], 64)
			drag, dragErr := translateDrag(dragTypeInput)
			handleErrors(exitWithErrorStatus, angleParsingErr, speedParsingError, dragErr)

			// Simulate
			simulator := app.NewTrajectorySimulator().
				WithInitialAngle(trigonometry.RadiansFromDegrees(angleInDegrees)).
				WithInitialSpeed(speed).
				WithDragType(drag)
			if addThrust {
				simulator.WithThrust(app.Thrust{Angle: trigonometry.RadiansFromDegrees(45), Force: 100, Duration: 5})
			}
			trajectory, trajectorySimulationErr := simulator.Simulate()
			handleErrors(printErrorToStOut, trajectorySimulationErr)

			// Output
			fmt.Fprintf(os.Stdout, "Landing point: x = %f\n", trajectory.LandingPoint())
			fmt.Fprintf(os.Stdout, "Airtime:       %f seconds\n", trajectory.AirTime())

			timeStamp := time.Now().Format("20060102-150405")

			if logSimulatorConditions {
				logFile, createLogFileErr := os.Create(fmt.Sprintf("./%s/%s/simulation-log-%s.txt", outputDir, logDir, timeStamp))
				handleErrors(printErrorToStOut, createLogFileErr)
				defer logFile.Close()
				simulator.Print(logFile)
				fmt.Fprint(logFile, "\nResults\n")
				fmt.Fprintf(logFile, "Landing point: x = %f\n", trajectory.LandingPoint())
				fmt.Fprintf(logFile, "Airtime: %f seconds\n", trajectory.AirTime())
			}

			if saveData {
				dataFile, createDataFileErr := os.Create(fmt.Sprintf("./%s/%s/simulation-data-%s.csv", outputDir, dataDir, timeStamp))
				handleErrors(printErrorToStOut, createDataFileErr)
				defer dataFile.Close()
				app.WriteCsv(dataFile, trajectory)
			}

			if makePlot {
				fmt.Println("generating plot...")
				plotErr := app.Plot(trajectory, fmt.Sprintf("./%s/%s/simulation-plot-%s", outputDir, plotDir, timeStamp))
				handleErrors(printErrorToStOut, plotErr)
			}
		},
	}

	command.Flags().StringVarP(&dragTypeInput, "drag", "D", "", "How drag is modeled proportional to the projectile's velocity: 'l' for linear or 'q' for quadratic. Leave empty for no drag.")
	command.Flags().BoolVarP(&addThrust, "thrust", "T", false, "Add a thrust of 100N in a constant 45 degree angle for the first 5 seconds.")
	command.Flags().BoolVarP(&makePlot, "plot", "P", false, fmt.Sprintf("Bool to indicate whether the trajectory should be plotted and saved to ./%s/%s/", outputDir, plotDir))
	command.Flags().BoolVarP(&saveData, "save", "S", false, fmt.Sprintf("Bool to indicate whether the trajectory should be saved as csv in ./%s/%s/", outputDir, dataDir))
	command.Flags().BoolVarP(&saveData, "log", "L", false, fmt.Sprintf("Bool to indicate whether the simulator settings should be saved ./%s/%s/", outputDir, logDir))
	
	return command
}

func translateDrag(input string) (app.DragType, error) {
	switch input {
	case "l":
		return app.Linear, nil
	case "q":
		return app.Quadratic, nil
	case "":
		return app.NoDrag, nil
	default:
		return app.NoDrag, errors.New("dragtype not supported")
	}
}
