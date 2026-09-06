# Projectile trajectory simulator

This is a Command Line Interface (CLI) application to simulate the trajectory of a projectile in 2 dimensions, with three forces acting on the projectile: gravity, drag (air resistance) and thrust.

# Getting Started

1. [Install Go](https://go.dev/doc/install) on your local machine.
2. Fork this repo.
3. Clone the forked repo to your local machine.

You now have a local repository containing the go package `projectile-trajectory`. There is no need to [install the pacakge](https://go.dev/doc/tutorial/compile-install), you can run it directly from the command line. However, installing it will make the commands shorter.

# Usage

Examples are without having the package installed.

Simulate a trajectory with 45 degree shooting angle and inital speed 100 m/s.

    go run . sim 45 100

The following flags are available:

    --drag     > model drag linearly (l) or quadraticly (q)
    --thrust   > add thrust (for now only with default pars)
    --plot     > create a plot of the trajectory

Example:

    go run . sim 45 100 --drag q --thrust --plot

Using shorthands:

    go run . sim 45 100 -D q -T -P

# Output files

There are three output directories:

    ./output
        /logs    > to save simulation conditions
        /data    > to save the trajectory data
        /plots   > to save trajectory plots

Logs and data are filled automatically, plots only when the `--plot` flag is set.

# Mathematical background

The projectile's trajectory can be modeled with the following equations:

$$\dot{\mathbf{r}}_t = \mathbf{v}_t$$

where 
$\mathbf{r} = (x_t, y_t)$
is the position and 
$\mathbf{v} = (\dot{x}_t, \dot{y}_t)$
is the velocity of the projectile.

The velocity is changing according to

$$\dot{\mathbf{v}}_t 
= \mathbf{a}_t = \mathbf{g} + \frac{\mathbf{F}_{t,\mathrm{drag}}}{m} + \frac{\mathbf{F}_{t,\mathrm{thrust}}}{m},$$

where $\mathbf{g} = (0, -9.81).$

When drag is modeled quadratic, like

$$\mathbf{F}_{t,\mathrm{drag}} 
= -k |\mathbf{v}|^2 \hat{\mathbf{v}} 
= -k |\mathbf{v}|^2 \frac{\mathbf{v}}{|\mathbf{v}|} 
= -k |\mathbf{v}| \mathbf{v},$$

with

$$k := \frac{1}{2} \rho C_D A,$$

where $\rho$ is the air density, $C_D$ is the drag coefficient of the projectile and $A$ is its frontal area, the system of equations has no analytic solution, because the equations for the $y$ dimension will include components from $x$ (i.e., the vertical and horizontal movements are no longer independent of each other). Therefore, we need numerical methods to obtain the projectile's trajectory, like this simulator.