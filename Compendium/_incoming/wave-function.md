---
id: 01KDKGWT7VPAK4BVQ038GTAJCV
title: "Wave Function"
slug: "wave-function"
created: 2025-12-28T22:22:36Z
tags: ["topic:quantum-mechanics", "topic:wave-function"]
people: ["person:niels-bohr"]
researcher_version: "1"
model: "deepseek-r1:14b"
iterations: 0
summary: "Initial overview based on 7.2: Wave functions - Physics LibreTexts"
---

# Wave Function

## Overview
The wave function, denoted as Ψ(x,t), is a fundamental concept in quantum mechanics that describes the state of a particle. It provides a mathematical description of the probability distribution of a particle's position and momentum.

## Key Concepts

### Probability Density
The square of the wavefunction, |Ψ(x,t)|², represents the probability density of finding a particle at position x and time t. This concept is derived from the Born interpretation, which states that the square of the wavefunction gives the probability density of finding a particle.

### Wavefunction Equation
For a free particle, the wavefunction can be expressed as:
\[ \Psi(x,t) = A e^{i(kx - \omega t)} \]
This equation describes the particle's state in terms of its wave number (k) and angular frequency (ω).

## History and Interpretations

### Copenhagen Interpretation
The Copenhagen interpretation suggests that particles exist in superposition states, meaning they can be in multiple states simultaneously until measured. This is famously illustrated by Schrödinger's cat thought experiment, where a cat could theoretically be both alive and dead until observed.

## Applications

### Two-State Systems and Quantum Computing
Two-state systems are fundamental in quantum mechanics, representing particles in mixed states. In quantum computing, qubits can exist in superpositions of states (e.g., 0 and 1), unlike classical binary digits, enabling more complex computations.

## Examples of Wavefunctions

### Particle Confined Between 0 and L
A possible wavefunction for a particle confined between 0 and L is:
\[ \psi(x,t) = A e^{-iEt/\hbar} \sin(\pi x/L) \]
The normalization constant A ensures that the integral of the probability density over all space equals 1.

## Expectation Values

### Expectation Value of Position
The expectation value of position is given by:
\[ \langle x \rangle = \int_{-\infty}^\infty x |\Psi(x,t)|^2 dx \]
For symmetric wavefunctions, this value can be zero due to the symmetry of the probability density.

## Symmetry in Wavefunctions

### Even and Odd Functions
An even function satisfies \( \psi(x) = \psi(-x) \), while an odd function satisfies \( \psi(x) = -\psi(-x) \). The product of two even or two odd functions results in an even function, whereas the product of an even and an odd function results in an odd function.

## Theoretical Principles

### Correspondence Principle
Niels Bohr's correspondence principle asserts that quantum mechanics must agree with classical mechanics for macroscopic systems. This principle underpins the idea that classical mechanics is a limiting case of quantum mechanics at large scales.

## References

[^1]: [7.2: Wave functions - Physics LibreTexts](https://phys.libretexts.org/Bookshelves/University_Physics/University_Physics_(OpenStax)/University_Physics_III_-_Optics_and_Modern_Physics_(OpenStax)/07%3A_Quantum_Mechanics/7.02%3A_Wavefunctions)