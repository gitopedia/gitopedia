---
id: 01KDKEEY295QNGW48GFZ471PFY
slug: "wave-function--phys-libretexts-org-1"
title: "Source: 7.2: Wave functions - Physics LibreTexts"
url: "https://phys.libretexts.org/Bookshelves/University_Physics/University_Physics_(OpenStax)/University_Physics_III_-_Optics_and_Modern_Physics_(OpenStax)/07%3A_Quantum_Mechanics/7.02%3A_Wavefunctions"
type: source
related_article: "wave-function"
created: 2025-12-28T21:41:05Z
tags: ["topic:wavefunction", "topic:probability-density", "topic:born-interpretation", "topic:superposition", "topic:copenhagen-interpretation", "topic:schrdingers-thought-experiment", "topic:two-state-systems", "topic:complex-wavefunctions", "topic:momentum-operator", "topic:kinetic-energy-operator", "topic:symmetry-of-wavefunctions", "topic:expectation-value", "topic:probability-interpretation", "topic:correspondence-principle", "topic:wave-function"]
people: ["person:niels-bohr", "person:erwin-schrdinger", "person:max-born"]
summary: "Summarized source material for Wave Function"
researcher_version: "0.3.29"
---

# Using the Wavefunction
- The wavefunction Ψ(x,t) describes how a particle behaves like a wave.
- The square of the wavefunction's amplitude, |Ψ(x,t)|², represents the probability density of finding the particle at position x and time t.
- For light waves, the energy density is proportional to the square of the electric field strength, |E|². Similarly, for matter waves, the probability density is given by |Ψ(x,t)|².

# Probability Interpretation
- The Born interpretation states that the square of the wavefunction gives the probability density of finding a particle at a specific position.
- The probability of finding a particle in a narrow interval (x, x+dx) is P(x,x+dx) = |Ψ(x,t)|² dx.
- For macroscopic objects, the probability of finding the particle in a specific location is almost 1 due to the wavefunction's large amplitude.

# Wavefunction and State
- A particle can exist in multiple states simultaneously until measured. This is known as superposition.
- The Copenhagen interpretation suggests that before measurement, the particle exists in all possible states (e.g., both alive and dead in Schrödinger's thought experiment).
- Upon measurement, the wavefunction collapses to a single state.

# Two-State Systems
- Quantum systems can exist in two or more states at once. For example, an electron can have spin-up and spin-down simultaneously.
- These systems are fundamental in quantum computing, where qubits represent mixed states of 0 and 1.

# Complex Wavefunctions
- The wavefunction is typically a complex function containing imaginary numbers (e.g., i = √-1).
- To calculate probabilities, the complex conjugate of the wavefunction must be used. This ensures all predictions are real numbers.
- For a complex number a = 3 + 4i, its product with its conjugate is a* a = (3 + 4i)(3 - 4i) = 25.

# Free Particles
- A free particle's wavefunction can be expressed as Ψ(x,t) = A cos(kx - ωt) + iA sin(kx - ωt).
- Using Euler’s formula, this can be rewritten as Ψ(x,t) = A e^{i(kx - ωt)}.
- The real and imaginary parts of the wavefunction describe the particle's motion and probability distribution.

# Applications
- Quantum mechanics is used to describe particles in various states, from free particles to those bound by forces.
- The principles of superposition and wavefunction collapse are central to understanding quantum phenomena.

# Probability Density
- The probability density \( P(x, x+\Delta x) \) is given by \( \Psi^*(x,t)\Psi(x,t)\Delta x = |A|^2 \Delta x \), where \( A \) is a complex constant.
- If \( A \) has real and complex parts (\( a + i b \)), then \( A^*A = a^2 + b^2 \), which ensures that the probability density is a real quantity.

# Wavefunction for a Particle Confined in Region [0, L]
- The wavefunction is \( \psi(x,t) = A e^{-iEt/\hbar} \sin(\pi x / L) \) for \( 0 \leq x \leq L \).
- The normalization constant \( A \) is determined by ensuring the integral of \( |\psi(x,t)|^2 \) over the region [0, L] equals 1.

# Expectation Value of Position
- The expectation value of position is given by \( \langle x \rangle = \int_{-\infty}^\infty x P(x,t) dx \).
- For a symmetric wavefunction, such as \( \psi(x) = e^{-|x|/x_0}/x_0 \), the expectation value of position is zero due to the symmetry of the probability density.

# Momentum Operator
- The momentum operator in the x-direction is \( -i\hbar \frac{d}{dx} \).
- The expectation value of momentum is calculated using \( \langle p \rangle = \int_{-\infty}^\infty \Psi^*(x,t) (-i\hbar \frac{d}{dx}) \Psi(x,t) dx \).

# Kinetic Energy Operator
- The kinetic energy operator is derived from the momentum operator and is given by \( -\frac{\hbar^2}{2m} \frac{d^2}{dx^2} \).
- The expectation value of kinetic energy involves applying this operator to the wavefunction and integrating over all space.

# Symmetry of Wavefunctions
- Even functions satisfy \( \psi(x) = \psi(-x) \), while odd functions satisfy \( \psi(x) = -\psi(-x) \).
- The product of two even functions or two odd functions results in an even function, while the product of an even and an odd function results in an odd function.

# Example Calculation for Expectation Values
- For the wavefunction \( \psi(x,t) = A e^{-iEt/\hbar} \sin(\pi x / L) \), the normalization constant is found to be \( A = 2/L \).
- The expectation value of position is calculated as \( \langle x \rangle = \frac{L}{2} \).
- The expectation value of momentum is zero due to the symmetry of the wavefunction.
- The expectation value of kinetic energy is non-zero and given by \( \langle K \rangle = \frac{\hbar^2}{8mL^2} \).

# Probability of Locating a Particle
- The probability of finding a particle between positions 0 and \( L/4 \) involves integrating the probability density \( |\psi(x,t)|^2 \) over this interval.

# Correspondence Principle
- Niels Bohr's correspondence principle states that quantum mechanics must agree with classical mechanics for macroscopic systems.
- This principle suggests that classical mechanics is an approximation of quantum mechanics for systems with very large energies.
