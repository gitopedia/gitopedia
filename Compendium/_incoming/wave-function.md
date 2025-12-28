---
id: 01KDKEHY29W3V52M4K8HWTC3TZ
title: "Wave Function"
slug: "wave-function"
created: 2025-12-28T21:41:42Z
tags: ["topic:wave-function", "topic:quantum-mechanics", "topic:copenhagen-interpretation", "topic:wave-function"]
people: ["person:niels-bohr"]
researcher_version: "1"
model: "deepseek-r1:14b"
iterations: 0
summary: "Initial overview based on 7.2: Wave functions - Physics LibreTexts"
---

# Wave Function

## Overview
The wave function is a fundamental concept in quantum mechanics that describes the behavior of particles at both microscopic and macroscopic levels. It encapsulates the probabilistic nature of particles and their possible states, providing a mathematical framework to predict particle behaviors.

## Key Concepts

### Probability Interpretation
- The Born interpretation posits that the square of the wave function's amplitude, |Ψ(x,t)|², gives the probability density of finding a particle at position x and time t.
- The probability of locating a particle in an infinitesimal interval (x, x+dx) is given by P(x,x+dx) = |Ψ(x,t)|² dx.

### Wavefunction and State
- A particle can exist in multiple states simultaneously until measured, a phenomenon known as superposition.
- According to the Copenhagen interpretation, before measurement, a particle exists in all possible states. Upon measurement, the wave function collapses to a single state.

### Two-State Systems
- Quantum systems can occupy two or more states concurrently. For instance, an electron's spin can be both up and down until measured. This principle is foundational in quantum computing, where qubits represent mixed states of 0 and 1.

## Mathematical Foundations

### Complex Wavefunctions
- The wave function is typically complex, involving imaginary numbers (e.g., i = √-1). To compute probabilities, the complex conjugate must be used.
- Example: For a complex number a = 3 + 4i, its product with its conjugate is 25.

### Free Particles
- A free particle's wave function can be expressed as Ψ(x,t) = A cos(kx - ωt) + iA sin(kx - ωt), which simplifies using Euler’s formula to Ψ(x,t) = A e^{i(kx - ωt)}.

### Operators and Expectation Values
- **Momentum Operator**: The x-direction momentum operator is \( -i\hbar \frac{d}{dx} \).
- **Kinetic Energy Operator**: Derived from the momentum operator, it is \( -\frac{\hbar^2}{2m} \frac{d^2}{dx^2} \).
- **Expectation Values**: Calculated by integrating the product of the wave function's conjugate and the operator applied to the wave function.

### Symmetry
- Wave functions can be even (ψ(x) = ψ(-x)) or odd (ψ(x) = -ψ(-x)). The product of two even or two odd functions results in an even function, while mixing even and odd yields an odd function.

## Applications

### Confined Particles
- For a particle confined to [0, L], the wave function is \( \psi(x,t) = A e^{-iEt/\hbar} \sin(\pi x / L) \), with normalization constant A = 2/L.

### Expectation Values
- **Position**: Symmetric wave functions yield ⟨x⟩ = 0.
- **Momentum**: Zero due to symmetry.
- **Kinetic Energy**: Non-zero, calculated as \( \langle K \rangle = \frac{\hbar^2}{8mL^2} \).

### Probability of Detection
- The probability of finding a particle between 0 and L/4 involves integrating the probability density over this interval.

## Correspondence Principle
Niels Bohr's principle asserts that quantum mechanics aligns with classical mechanics for macroscopic systems, underscoring that classical physics is an approximation of quantum principles at large scales.

## References

[^1]: [7.2: Wave functions - Physics LibreTexts](https://phys.libretexts.org/Bookshelves/University_Physics/University_Physics_(OpenStax)/University_Physics_III_-_Optics_and_Modern_Physics_(OpenStax)/07%3A_Quantum_Mechanics/7.02%3A_Wavefunctions)