---
id: 01KDP9QMF655ZG89YAYHSVNNN9
slug: "wave-function--courses-lumenlearning-com-1025"
title: "Source: 16.2 Mathematics of Waves | University Physics Volume 1"
url: "https://courses.lumenlearning.com/suny-osuniversityphysics/chapter/16-2-mathematics-of-waves/"
type: source
related_article: "wave-function"
created: 2025-12-30T00:15:24Z
tags: [""]
summary: "Summarized source material for Wave Function"
researcher_version: "0.3.29"
---

# 16.2 Mathematics of Waves
- By the end of this section, you will be able to model a wave moving with a constant wave velocity using a mathematical expression.
- You will also be able to calculate the velocity and acceleration of the medium and show how the velocity of the medium differs from the wave velocity (propagation velocity).
- A pulse can be described as a wave consisting of a single disturbance that moves through the medium with a constant amplitude. The pulse moves as a pattern that maintains its shape as it propagates with a constant wave speed.
- The distance the pulse moves in a time Δt is equal to Δx = vΔt.
- A sinusoidal wave can be modeled using a wave function, which can be used to find the position, velocity, and acceleration of particles in the medium at any time.
- For a string under tension, the wave function is given by y(x,t) = Asin(kx - ωt + φ), where A is the amplitude, k is the wave number, ω is the angular frequency, and φ is the initial phase shift.
- The wave number k is defined as 2π/λ, where λ is the wavelength. The angular frequency ω is defined as 2π/T, where T is the period.
- The velocity of the medium can be found by taking the partial derivative of the position function with respect to time: vy = ∂y/∂t.
- The maximum velocity of the medium is |vymax| = Aω.
- The acceleration of the medium can be found by taking the partial derivative of the velocity function with respect to time: ay = ∂vy/∂t.
- The maximum acceleration of the medium is |aymax| = Aω².
- The linear wave equation is given by ∇²y = (1/v²) ∂²y/∂t², where v is the wave speed. This equation describes how the curvature of the wave relates to its acceleration.
- If two wave functions are solutions to the linear wave equation, their sum is also a solution. This property is known as superposition and applies to any linear wave functions.

# Key Equations
- Wave function: y(x,t) = Asin(kx - ωt + φ)
- Wave number: k = 2π/λ
- Angular frequency: ω = 2π/T
- Velocity of the medium: vy = ∂y/∂t
- Acceleration of the medium: ay = ∂vy/∂t
- Linear wave equation: ∇²y = (1/v²) ∂²y/∂t²

# Example Calculation
Given the wave function y(x,t) = 0.2msin(6.28m⁻¹x - 1.57s⁻¹t):
- Amplitude A = 0.2m
- Wave number k = 6.28m⁻¹
- Angular frequency ω = 1.57s⁻¹
- Wavelength λ = 2π/k = 2π/6.28 ≈ 1.0m
- Period T = 2π/ω = 2π/1.57 ≈ 4.0s
- Wave speed v = ω/k = 1.57/6.28 ≈ 0.25m/s

# Graphical Representation
- Figure 16.11 shows the height of the wave y as a function of position x for two times, t=0.00s and t=0.80s. The wave moves to the right with constant speed.
- Figure 16.12 shows the height of the wave y as a function of time t at position x=0.6m. The medium oscillates between y=+0.20m and y=-0.20m every period of 4.0s.

# Check Your Understanding
- Yes, a cosine function can be used instead of a sine function in the wave equation. The choice depends on the initial conditions of the wave.

# Principle of Superposition
- The principle of superposition states that if wave functions \( y_1(x,t) = f(x \mp vt) \) and \( y_2(x,t) = g(x \mp vt) \) are solutions to the linear wave equation, then any linear combination \( Ay_1(x,t) + By_2(x,t) \), where \( A \) and \( B \) are constants, is also a solution.
- This property applies to all types of waves, including sound waves and electromagnetic waves.

# Interference of Waves on a String
- Consider two waves on a string: 
  - Wave 1: \( y_1(x,t) = A\sin(kx - \omega t) \)
  - Wave 2: \( y_2(x,t) = A\sin(2kx + 2\omega t) \)
- The resulting wave function is the sum of the two individual wave functions: 
  - \( y_R(x,t) = y_1(x,t) + y_2(x,t) = A\sin(kx - \omega t) + A\sin(2kx + 2\omega t) \)
- The velocity of the resulting wave is equal to the speed of the original waves: \( v = \omega k \).

# Wave Equation Analysis
- The linear wave equation is given by:
  - \( \frac{\partial^2 y(x,t)}{\partial x^2} = \frac{1}{v^2} \frac{\partial^2 y(x,t)}{\partial t^2} \)
- For a wave function of the form \( y(x,t) = f(x \mp vt) \), this equation holds true.
- The wave speed \( v \) is related to the angular frequency \( \omega \) and wave number \( k \) by \( v = \frac{\omega}{k} \).

# Characteristics of a Sinusoidal Wave
- A sinusoidal wave can be described by the wave function:
  - \( y(x,t) = A\sin(kx - \omega t + \phi) \)
- Key parameters include:
  - Amplitude (\( A \)): The maximum displacement from equilibrium.
  - Wavelength (\( \lambda \)): The distance between two consecutive points in phase on the wave.
  - Wave speed (\( v \)): The speed at which the wave propagates.
  - Period (\( T \)): The time it takes for one complete cycle of the wave.

# Example Calculations
- **Example 1**: A wave is modeled by \( y(x,t) = (0.25\,m)\cos(0.30\,m^{-1}x - 0.90\,s^{-1}t + \frac{\pi}{3}) \).
  - Amplitude (\( A \)): \( 0.25\,m \)
  - Wave number (\( k \)): \( 0.30\,m^{-1} \)
  - Angular frequency (\( \omega \)): \( 0.90\,s^{-1} \)
  - Wave speed (\( v \)): \( 3.0\,m/s \)
  - Phase shift (\( \phi \)): \( \frac{\pi}{3}\,rad \)
  - Wavelength (\( \lambda \)): \( 20.93\,m \)
  - Period (\( T \)): \( 6.98\,s \)

- **Example 2**: A surface ocean wave has an amplitude of \( 0.60\,m \) and a wavelength of \( 8.00\,m \). It moves at \( 1.50\,m/s \) in the positive x-direction.
  - Wave function: \( y(x,t) = (0.60\,m)\sin(6.28\,m^{-1}x - 1.50\,s^{-1}t) \)

# Particle Motion
- Each particle on the string moves a distance of \( 4A \) each period.
- The time required for a particle to move through a distance of \( 5.00\,km \) is calculated as:
  - Time (\( t \)): \( 10.42\,s \)

# Wave Comparison
- **Wave 1**: \( y_1(x,t) = 0.50\,msin(2\pi3.00\,mx + 2\pi4.00\,st) \)
- **Wave 2**: \( y_2(x,t) = 0.50\,msin(2\pi6.00\,mx - 2\pi4.00\,st) \)
  - Similarities: Same angular frequency, frequency, and period.
  - Differences: Traveling in opposite directions; \( y_2(x,t) \) has twice the wavelength and half the wave speed.

# Ocean Waves
- A swimmer observes ocean waves with a vertical distance between crest and trough of \( 0.45\,m \) and a distance between crests of \( 1.8\,m \). The frequency is \( 12\,waves/minute \).
  - Wave function: \( y(x,t) = 0.23\,msin(3.49\,m^{-1}x - 0.63\,s^{-1}t) \)

# Particle Displacement
- For a wave described by \( y(x,t) = 0.3\,msin(2.00\,m^{-1}x - 628.00\,s^{-1}t) \):
  - Number of crests passing an observer in \( 2.00\,minutes \): 
    - Crests per second: \( 628.00/2\pi = 100\,Hz \)
    - Total crests: \( 100 \times 120 = 12000 \)
  - Distance traveled by the wave in \( 2.00\,minutes \): 
    - Wave speed: \( v = f\lambda = (628.00/2\pi) \times 0.5 = 314\,m/s \)
    - Distance: \( 314 \times 120 = 37680\,m \)

# Conclusion
- The principles of superposition and interference are fundamental to understanding wave behavior.
- The linear wave equation governs the propagation of waves, and key parameters such as amplitude, wavelength, and wave speed can be derived from the wave function.
