# Example 1

The projectile is a perfect sphere with a radius of $3 \mathrm{cm}$ and a density of $7.8 \mathrm{g}/\mathrm{cm}^3$.
The projectile is launched at an $45$ degree angle, from $x=y=0$, with a velocity of $200 \mathrm{m/s}$. 
The drag is modeled quadraticly.
The question is: At what time and distance will the projectile hit the ground?

### Command

    go run . sim 45 200 --drag q --log

### Result

    Landing point:   x = 4061.853913
    Airtime:         28.802500 seconds
    Sim. Iterations: 288025

### Log

    Simulator conditions
    projectile:        Shape=sphere, radius=3.000000cm, density=7.800000g/cm^3
    initial position:  {0.000000 0.000000}
    initial angle:     0.785398
    initial speed:     200.000000
    thrust:            angle=45.000000rad, force=0.000000N, duration=0.000000s
    timeStepInterval:  0.000100
    drag type:         quadratic
    integrationMethod: Forward Euler
    maxTimeSteps:      1000000
    errTolerAtLanding: 0.000010

### Command using simplectic Euler integration

We can also use simplectic Euler integration, which uses the velocity at time $t+1$ rather than the velocity at time $t$ to update the projectile's position between $t$ and $t+1$.

    go run . sim 45 200 --drag q --method s --log

### Result
    Result
    Landing point:   x = 4061.825708
    Airtime:         28.802300 seconds
    Sim. Iterations: 288023

### Log

    Simulator conditions
    projectile:        Shape=sphere, radius=3.000000cm, density=7.800000g/cm^3
    initial position:  {0.000000 0.000000}
    initial angle:     0.785398
    initial speed:     200.000000
    thrust:            angle=45.000000rad, force=0.000000N, duration=0.000000s
    timeStepInterval:  0.000100
    drag type:         quadratic
    integrationMethod: Simplectic Euler
    maxTimeSteps:      1000000
    errTolerAtLanding: 0.000010

# Example 2

To the situation of example 1, we add a thrust of $100\mathrm{N}$, for the first $5$ seconds, that stays in a $45$ degree angle. What is the airtime and distance now?

### Command

    go run . sim 45 200 --drag q --thrust angle:45,force:100,duration:5 --log

### Result

    Landing point:   x = 4089.856897
    Airtime:         28.901800 seconds
    Sim. Iterations: 289018

### Log

    Simulator conditions
    projectile:        Shape=sphere, radius=3.000000cm, density=7.800000g/cm^3
    initial position:  {0.000000 0.000000}
    initial angle:     0.785398
    initial speed:     200.000000
    thrust:            angle=0.785398rad, force=100.000000N, duration=5.000000s
    timeStepInterval:  0.000100
    drag type:         quadratic
    integrationMethod: Forward Euler
    maxTimeSteps:      1000000
    errTolerAtLanding: 0.000010