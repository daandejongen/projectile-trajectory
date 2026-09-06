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

const plotDir string = "output/plots"
const logDir string = "output/logs"
const dataDir string = "output/data"

const defaultDragType = app.NoDrag
var defaultThrust = app.Thrust{Duration: 0, Force: 0, Angle: 45}

func NewSimulateCommand() *cobra.Command {
	var dragTypeInput string
	var thrustInput string
	var makePlot bool
	var saveData bool
	var logSimulatorConditions bool
	var integrationMethodInput string
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
			integrationMethod, integrationMethodErr := parseIntegrationMethodInput(integrationMethodInput)
			handleErrors(exitWithErrorStatus, angleParsingErr, speedParsingError, dragErr, thrustErr, integrationMethodErr)

			// Simulate
			simulator := app.NewTrajectorySimulator().
				WithInitialAngle(trigonometry.RadiansFromDegrees(angleInDegrees)).
				WithInitialSpeed(speed).
				WithDragType(dragType).
				WithThrust(thrust).
				WithIntegrationMethod(integrationMethod)
			trajectory, trajectorySimulationErr := simulator.Simulate()
			handleErrors(printErrorToStOut, trajectorySimulationErr)

			// Output
			if trajectorySimulationErr == nil {
				printTrajectoryResults(os.Stdout, trajectory)
			}

			timeStamp := time.Now().Format("20060102-150405")

			if logSimulatorConditions {
				os.MkdirAll(logDir, 0755)
				logFile, createLogFileErr := os.Create(fmt.Sprintf("%s/simulation-log-%s.txt", logDir, timeStamp))
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
				os.MkdirAll(dataDir, 0755)
				fmt.Print("writing csv...")
				dataFile, createDataFileErr := os.Create(fmt.Sprintf("%s/simulation-data-%s.csv", dataDir, timeStamp))
				handleErrors(printErrorToStOut, createDataFileErr)
				defer dataFile.Close()
				app.WriteCsv(dataFile, trajectory)
				fmt.Println(" succeeded")
			}

			if makePlot {
				os.MkdirAll(plotDir, 0755)
				fmt.Print("generating plot...")
				plotErr := app.Plot(trajectory, fmt.Sprintf("%s/simulation-plot-%s", plotDir, timeStamp))
				handleErrors(printErrorToStOut, plotErr)
				fmt.Println(" succeeded")
			}
		},
	}

	command.Flags().StringVarP(&dragTypeInput, "drag", "D", "", "How drag is modeled proportional to the projectile's velocity: 'l' for linear or 'q' for quadratic. Leave empty for no drag.")
	command.Flags().StringVarP(&thrustInput, "thrust", "T", "", fmt.Sprintf("Add thrust of the form 'angle:45,duration:5,force:100' in degrees, seconds and Newton respectively. Default %s: You can omit properties to use the default for it.", defaultThrust.String()))
	command.Flags().BoolVarP(&makePlot, "plot", "P", false, fmt.Sprintf("Bool to indicate whether the trajectory should be plotted and saved to ./%s/", plotDir))
	command.Flags().BoolVarP(&saveData, "csv", "C", false, fmt.Sprintf("Bool to indicate whether the trajectory should be saved as csv in ./%s/", dataDir))
	command.Flags().BoolVarP(&logSimulatorConditions, "log", "L", false, fmt.Sprintf("Bool to indicate whether the simulator settings should be saved ./%s/", logDir))
	command.Flags().StringVarP(&integrationMethodInput, "method", "M", "f", fmt.Sprint("Integration method: f = Forward Euler, s = Simplectic Euler"))

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
	if input == "" {
		return thrust, nil
	}
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

func parseIntegrationMethodInput(input string) (app.IntegrationMethod, error) {
	switch input {
	case "f":
		return app.ForwardEulerIntegration, nil
	case "s":
		return app.SimplecticEulerIntegration, nil
	case "":
		return app.ForwardEulerIntegration, nil
	default:
		return app.NoIntegrationMethod, errors.New("integration method not supported")
	}
}

func printTrajectoryResults(writer io.Writer, trajectory app.Trajectory) {
	fmt.Fprintf(writer, "Landing point:   x = %f\n", trajectory.LandingPoint())
	fmt.Fprintf(writer, "Airtime:         %f seconds\n", trajectory.AirTime())
	fmt.Fprintf(writer, "Sim. Iterations: %d\n", trajectory.SimulationIterations())
}