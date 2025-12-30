---
id: 01KDPA1CMG30DS22YQ5J4DHX4Y
slug: "wave-function--labman-phys-utk-edu-919"
title: "Source: Wave functions - University of Tennessee"
url: "http://labman.phys.utk.edu/phys222core/modules/m10/wave_functions.html"
type: source
related_article: "wave-function"
created: 2025-12-30T00:20:47Z
tags: [""]
summary: "Summarized source material for Wave Function"
researcher_version: "0.3.29"
---

# Wave Functions
- In one dimension, wave functions are often denoted by the symbol ψ(x,t). They are functions of the coordinate x and the time t. 
- The wave function ψ(x,t) is a complex function, not a real one.
- The Schrödinger equation does not have real solutions but complex ones.
- The wave function contains all the information about the particle at a particular time.
- The wave function itself has no physical interpretation and is not measurable.
- The square of the absolute value of the wave function, |ψ(x,t)|², has a physical interpretation as a probability density (probability per unit length) of finding the particle at position x at time t.
- The probability of finding the particle in an interval ∆x about position x is proportional to |ψ(x,t)|²∆x.
- For the probability interpretation to make sense, the wave function must be single-valued and continuous.
- The wave function must be normalizable, meaning that the total probability of finding the particle anywhere is one (the area under the curve |ψ(x,t)|² must be 1).
- An example of an acceptable wave function is given as ψ(x,0) = (2/π)½cos(x) for x between -π/2 and +π/2.
- An example of an unacceptable wave function is given as ψ(x,0) which does not have a unique single value at x = 0 and is discontinuous.

# Probability Interpretation
- Each particle can be found in a small interval ∆x about any x between -π/2 and +π/2 with probability (2/π)cos²(x) ∆x.
- The average position of the particles is zero because they are equally likely to be found on either side of the origin.
- The most likely position to find a particle is also zero, where the square of the wave function has its maximum value.

# Energy Eigenfunctions and Stationary States
- If we know the eigenfunctions of the energy operator, we can determine the possible outcomes of an energy measurement.
- For a confined particle in a potential well, the energy eigenfunctions resemble standing waves (e.g., ψ(x,t) = ψ(x) * e^{-iEt/ħ}).
- The probability density for energy eigenstates is independent of time because the square of the complex time function is 1.
- Energy eigenstates are called stationary states because their probability density does not change with time.

# Incompatibility of Energy and Position
- Knowing the energy of a particle prevents us from knowing its exact position and tracking it over time.
- A stationary-state wave function provides only probabilities for finding the particle at different positions, not information about its motion or trajectory.

# Particle in a One-Dimensional Square Well
- An electron (mass m = 9.109×10⁻³¹ kg) confined in a one-dimensional square well with width L = 10 nm has quantized energy levels.
- The energy levels are given by En = n²π²ħ²/(2mL²), where n is an integer (n = 1, 2, 3, ...).
- The corresponding wave functions are ψn(x) = (2/L)½sin(nπx/L).
- The probability density functions are |ψn(x)|² = (2/L)sin²(nπx/L).

# Energy Transitions and Quantization
- A particle can transition between different energy levels when interacting with its environment.
- Absorbing or emitting a photon of energy hf = ΔE, where ΔE is the difference between two energy levels, allows transitions between states.
- The quantized energy levels lead to discrete emission and absorption spectra for atoms.

# Quantum Tunneling
- In quantum mechanics, particles can penetrate classically forbidden regions (regions where kinetic energy is negative).
- The probability of tunneling through a barrier decreases exponentially with the width and height of the barrier.
- The tunneling probability P is given by P = exp(-2αd), where α = [(2m/ħ²)(U - E)]½, d is the barrier width, U is the barrier height, and E is the particle's energy.

# Example Calculation
- For a potential barrier of height U = 6 eV and width d = 0.7 nm, the tunneling probability P = 0.001 corresponds to an electron energy E ≈ 5.07 eV.
- The calculation involves determining α from ln(P) = -2αd and then solving for E using E = U - ħ²α²/(2m).

# Applications of Tunneling
- **Tunneling Microscopy (STM):** 
  - A scanning tunneling microscope uses the phenomenon of quantum tunneling to image surfaces at the atomic level.
  - The tip is mounted on a piezoelectric tube, allowing tiny movements to maintain a constant tunneling current while scanning the surface.
  - This results in high-resolution images of surface topography, including individual atoms.

#Macroscopic Behavior and Quantum Mechanics
- If quantum mechanics governed macroscopic objects, uncertainty principles and tunneling effects would lead to unexpected behaviors (e.g., a parked car might spontaneously disappear through a wall).
- The likelihood of such events decreases with larger barriers and masses but is theoretically possible according to quantum mechanics.
