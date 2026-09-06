package cli

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/daandejongen/projectile-trajectory/app"
	"github.com/daandejongen/projectile-trajectory/trigonometry"
	"github.com/spf13/cobra"
)

const outputDir string = "output"
const plotDir string = "plots"
const logDir string = "logs"
const dataDir string = "data"

const defaultDragType = app.NoDrag
var defaultThrust = app.Thrust{Duration: 0, Force: 0, Angle: 45}

func NewSimulateCommand() *cobra.Command {
	var dragTypeInput string
	var thrustInput string
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
			dragType, dragErr := parseDragTypeInput(dragTypeInput)
			thrust, thrustErr := parseThrustInput(thrustInput)
			handleErrors(exitWithErrorStatus, angleParsingErr, speedParsingError, dragErr, thrustErr)

			// Simulate
			simulator := app.NewTrajectorySimulator().
				WithInitialAngle(trigonometry.RadiansFromDegrees(angleInDegrees)).
				WithInitialSpeed(speed).
				WithDragType(dragType).
				WithThrust(thrust)
			trajectory, trajectorySimulationErr := simulator.Simulate()
			handleErrors(printErrorToStOut, trajectorySimulationErr)

			// Output
			if trajectorySimulationErr == nil {
				printTrajectoryResults(os.Stdout, trajectory)
			}

			timeStamp := time.Now().Format("20060102-150405")

			if logSimulatorConditions {
				logFile, createLogFileErr := os.Create(fmt.Sprintf("./%s/%s/simulation-log-%s.txt", outputDir, logDir, timeStamp))
				handleErrors(printErrorToStOut, createLogFileErr)
				defer logFile.Close()
				simulator.Print(logFile)
				fmt.Fprint(logFile, "\nResult\n")
				if trajectorySimulationErr ==  nil {
					printTrajectoryResults(logFile, trajectory)
				} else {
					fmt.Fprintf(logFile, trajectorySimulationErr.Error())
				}
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
	command.Flags().StringVarP(&thrustInput, "thrust", "T", "", fmt.Sprintf("Add thrust of the form 'angle:45,duration:5,force:100' in degrees, seconds and Newton respectively. Default %s: You can omit properties to use the default for it.", defaultThrust.String()))
	command.Flags().BoolVarP(&makePlot, "plot", "P", false, fmt.Sprintf("Bool to indicate whether the trajectory should be plotted and saved to ./%s/%s/", outputDir, plotDir))
	command.Flags().BoolVarP(&saveData, "csv", "C", false, fmt.Sprintf("Bool to indicate whether the trajectory should be saved as csv in ./%s/%s/", outputDir, dataDir))
	command.Flags().BoolVarP(&logSimulatorConditions, "log", "L", false, fmt.Sprintf("Bool to indicate whether the simulator settings should be saved ./%s/%s/", outputDir, logDir))

	return command
}

func parseDragTypeInput(input string) (app.DragType, error) {
	switch input {
	case "l":
		return app.LinearDrag, nil
	case "q":
		return app.QuadraticDrag, nil
	case "":
		return app.NoDrag, nil
	default:
		return app.NoDrag, errors.New("dragtype not supported")
	}
}

func parseThrustInput(input string) (app.Thrust, error) {
	thrust := defaultThrust
	components := strings.Split(input, ",")
	for _, component := range components {
		propertyAndValue := strings.Split(component, ":")
		if len(propertyAndValue) != 2 {
			return thrust, errors.New("could not parse thrust input")
		}
		switch propertyAndValue[0] {
		case "angle":
			value, err := strconv.ParseFloat(propertyAndValue[1], 64)
			if err != nil {
				return thrust, err
			}
			thrust.Angle = trigonometry.RadiansFromDegrees(value)
		case "duration":
			value, err := strconv.ParseFloat(propertyAndValue[1], 64)
			if err != nil {
				return thrust, err
			}
			thrust.Duration = value
		case "force":
			value, err := strconv.ParseFloat(propertyAndValue[1], 64)
			if err != nil {
				return thrust, err
			}
			thrust.Force = value
		default:
			return thrust, errors.New(propertyAndValue[0] + " is not a thrust property")
		}
	}
	return thrust, nil
}

func printTrajectoryResults(writer io.Writer, trajectory app.Trajectory) {
	fmt.Fprintf(os.Stdout, "Landing point:   x = %f\n", trajectory.LandingPoint())
	fmt.Fprintf(os.Stdout, "Airtime:         %f seconds\n", trajectory.AirTime())
	fmt.Fprintf(os.Stdout, "Sim. Iterations: %d\n", trajectory.SimulationIterations())
}