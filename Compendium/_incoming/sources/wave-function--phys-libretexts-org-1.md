---
id: 01KDKGVDWVZBRBJAMMM5K7T2DK
slug: "wave-function--phys-libretexts-org-1"
title: "Source: 7.2: Wave functions - Physics LibreTexts"
url: "https://phys.libretexts.org/Bookshelves/University_Physics/University_Physics_(OpenStax)/University_Physics_III_-_Optics_and_Modern_Physics_(OpenStax)/07%3A_Quantum_Mechanics/7.02%3A_Wavefunctions"
type: source
related_article: "wave-function"
created: 2025-12-28T22:22:16Z
tags: ["topic:quantum-mechanics", "topic:wavefunction", "topic:probability-density", "topic:normalization-condition", "topic:born-interpretation", "topic:copenhagen-interpretation", "topic:two-state-systems", "topic:qubits", "topic:schrdingers-cat", "topic:free-particle-wavefunction", "topic:complex-conjugate", "topic:expectation-value-of-position", "topic:expectation-value-of-momentum", "topic:kinetic-energy-operator", "topic:symmetry-of-wavefunctions", "topic:correspondence-principle", "topic:wave-function"]
people: ["person:niels-bohr", "person:erwin-schrdinger"]
summary: "Summarized source material for Wave Function"
researcher_version: "0.3.29"
---

# [Using the Wavefunction]
- The wavefunction Ψ(x,t) describes a particle's state and can be used to determine probabilities.
- The square of the wavefunction, |Ψ(x,t)|², represents the probability density of finding the particle at position x and time t.
- For an electron in a double-slit experiment, the probability density is proportional to the square of the total electric field, |E|².
- The Born interpretation states that the square of the wavefunction gives the probability density of finding a particle.

# [Normalization Condition]
- The normalization condition ensures that the total probability of finding a particle in all space is 1: ∫|Ψ(x,t)|² dx = 1.
- For a particle in a tube of length L, the wavefunction Ψ(x,0) = (2/L) cos(πx/L) satisfies the normalization condition.

# [Interpretation of the Wavefunction]
- The Copenhagen interpretation suggests that particles exist in superposition states and "collapse" to definite states upon measurement.
- Schrödinger's cat thought experiment illustrates the paradoxical implications of quantum superposition, where a cat could be simultaneously alive and dead until observed.

# [Two-State Systems and Quantum Computers]
- Two-state systems are fundamental in quantum mechanics, representing particles in mixed states.
- Quantum computers use qubits, which can exist in superpositions of states (e.g., 0 and 1) rather than binary digits.

# [Free Particle Wavefunction]
- The wavefunction for a free particle is given by Ψ(x,t) = A cos(kx - ωt) + iA sin(kx - ωt).
- Using Euler's formula, this can be rewritten as Ψ(x,t) = Ae^{i(kx - ωt)}.

# [Complex Conjugate and Probability]
- The probability of finding a particle in a narrow interval (x, x+dx) is given by P(x,x+dx) = |Ψ(x,t)|² dx.
- For a complex wavefunction Ψ(x,t), the product Ψ*(x,t)Ψ(x,t) is always real.

# [Example Calculation]
- If a = 3 + 4i, then a* a = (3 - 4i)(3 + 4i) = 9 + 16 = 25.

# Probability Density
- The probability density \( P(x, x+\Delta x) \approx \Psi^*(x,t)\Psi(x,t)\Delta x = |A|^2 \Delta x \).
- If \( A \) has real and complex parts (\( a + i b \)), then \( A^*A = a^2 + b^2 \), resulting in a real quantity.
- The interpretation of \( \Psi^*(x,t)\Psi(x,t) \) as a probability density ensures that the predictions of quantum mechanics can be checked in the “real world.”

# Wavefunction for a Particle Confined Between 0 and L
- A possible wavefunction is \( \psi(x,t) = A e^{-iEt/\hbar} \sin(\pi x/L) \) for \( 0 \leq x \leq L \).
- The normalization constant \( A \) is determined by ensuring the integral of \( |\psi(x,t)|^2 \) over all space equals 1.

# Expectation Value of Position
- The expectation value of position is given by:
  \[
  \langle x \rangle = \int_{-\infty}^\infty x P(x,t) dx = \int_{-\infty}^\infty x \Psi^*(x,t)\Psi(x,t) dx.
  \]
- For a symmetric wavefunction, the expectation value of position can be zero due to the symmetry of the probability density.

# Expectation Value of Momentum
- The expectation value of momentum is given by:
  \[
  \langle p \rangle = \int_{-\infty}^\infty \Psi^*(x,t) (-i\hbar \frac{d}{dx}) \Psi(x,t) dx.
  \]
- For the wavefunction \( \psi(x,t) = A e^{-iEt/\hbar} \sin(\pi x/L) \), the expectation value of momentum is zero because the integral involves an odd function.

# Expectation Value of Kinetic Energy
- The kinetic energy operator in position space is:
  \[
  K_{op} = -\frac{\hbar^2}{2m} \frac{d^2}{dx^2}.
  \]
- For the wavefunction \( \psi(x,t) = A e^{-iEt/\hbar} \sin(\pi x/L) \), the expectation value of kinetic energy is:
  \[
  \langle K \rangle = \int_0^L \Psi^*(x,t) \left( -\frac{\hbar^2}{2m} \frac{d^2}{dx^2} \right) \Psi(x,t) dx.
  \]
- The result is \( \langle K \rangle = \frac{\hbar^2}{8mL^2} \).

# Symmetry of Wavefunctions
- An even function satisfies \( \psi(x) = \psi(-x) \).
- An odd function satisfies \( \psi(x) = -\psi(-x) \).
- The product of two even functions or two odd functions results in an even function.
- The product of an even function and an odd function results in an odd function.

# Example: Normalized Wavefunction
- The normalized wavefunction is:
  \[
  \psi(x,t) = \frac{2}{L} e^{-iEt/\hbar} \sin(\pi x/L).
  \]
- The probability density is \( |\psi|^2 = \left( \frac{2}{L} \right)^2 \sin^2(\pi x/L) \).

# Probability of Locating a Particle
- For the particle in the state \( \psi(x,t) = A e^{-iEt/\hbar} \sin(\pi x/L) \), the probability of finding it between positions 0 and \( L/4 \) is:
  \[
  P(0 \leq x \leq L/4) = \int_0^{L/4} |\psi(x,t)|^2 dx.
  \]

# Correspondence Principle
- Niels Bohr asserted that quantum mechanics must agree with classical mechanics for macroscopic systems, leading to the correspondence principle.
- This principle suggests that classical mechanics is an approximation of quantum mechanics for systems with very large energies.
