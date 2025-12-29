---
id: 01KDP8PXJ4N24C627NFAVYN7TY
slug: "wave-function--phys-libretexts-org-1"
title: "Source: 7.2: Wave functions - Physics LibreTexts"
url: "https://phys.libretexts.org/Bookshelves/University_Physics/University_Physics_(OpenStax)/University_Physics_III_-_Optics_and_Modern_Physics_(OpenStax)/07%3A_Quantum_Mechanics/7.02%3A_Wavefunctions"
type: source
related_article: "wave-function"
created: 2025-12-29T23:57:29Z
tags: ["topic:quantum-mechanics", "topic:wavefunction", "topic:two-state-systems", "topic:superposition-principle", "topic:schrdingers-thought-experiment", "topic:wave-function"]
people: ["person:erwin-schrdinger", "person:niels-bohr"]
summary: "Summarized source material for Wave Function"
researcher_version: "0.3.29"
---

# 7.1: Prelude to Quantum Mechanics
- Particles can act like waves and particles depending on the situation.
- The wavefunction Ψ(x,t) describes how a particle behaves as a wave.
- The purpose of this chapter is to explain the meaning of the wavefunction, its use in making predictions, and the concept of wave-particle duality.

# 7.2: Wavefunctions
- A clue to the physical meaning of the wavefunction Ψ(x,t) is provided by two-slit interference of monochromatic light.
- The square of the matter wave |Ψ|² in one dimension gives the probability density that a particle will be found at a particular position and time per unit length.
- The probability (P) a particle is found in a narrow interval (x, x+dx) at time t is given by P(x, x+dx) = |Ψ(x,t)|² dx.
- The wavefunction can be used to describe particles that are free or bound by forces to other particles. These functions are usually complex functions containing imaginary numbers.

# 7.3: Two-state systems
- Two-state systems (left and right, atom decays and does not decay) are often used to illustrate the principles of quantum mechanics.
- A qubit in a quantum computer is not in a state of zero or one but rather in a mixed state of zero and one.

# 7.4: Free particles
- A free particle experiences no forces and moves with a constant velocity.
- The wavefunction for a free particle is given by Ψ(x,t) = A cos(kx - ωt) + iA sin(kx - ωt), where A is the amplitude, k is the wave number, and ω is the angular frequency.

# 7.5: Complex conjugate
- The complex conjugate of a function is obtained by replacing every occurrence of i with -i.
- The product Ψ*(x,t)Ψ(x,t) is always a real number, ensuring that experimental measurements produce real numbers only.

# 7.6: Schrödinger's thought experiment
- Erwin Schrödinger’s thought experiment involves a cat in a steel box with a radioactive substance, illustrating the superposition principle.
- The Copenhagen interpretation states that until observed, the cat is simultaneously alive and dead, but this interpretation remains the most commonly taught view of quantum mechanics.

# 7.7: Quantum computers
- Quantum computers use qubits instead of binary digits (zeroes and ones).
- A qubit can be in a mixed state, producing zero with probability p and one with probability q = 1 - p upon measurement.
- Some scientists believe that quantum computers have the potential to revolutionize the computer industry.

# 7.8: Wavefunction collapse
- The Copenhagen interpretation includes the concept of wavefunction collapse, where the wavefunction "jumps into" a particular position state when observed.
- This process is called state reduction or wavefunction collapse and results in a definite measurement outcome.

# 7.9: Superposition principle
- The superposition principle states that a particle can exist in multiple states simultaneously until measured.
- This principle applies to various measurable quantities, such as momentum and energy, and finds applications in nature and quantum computing.

# 7.10: Formal quantum mechanical treatment
- A formal quantum mechanical treatment of a free particle indicates that its wavefunction has real and complex parts.
- The wavefunction can be rewritten using Euler’s formula, which combines cosine and sine terms into a single exponential form.

# 7.11: Probability calculation
- If the wavefunction varies slowly over an interval Δx, the probability of finding the particle in that interval is approximately P(x, x+dx) = |Ψ(x,t)|² dx.
- This procedure ensures that experimental measurements produce real numbers only by taking the product of the complex conjugate and the wavefunction.

# Wavefunctions
- The probability density is given by \( P(x, x+\Delta x) \approx \Psi^*(x,t)\Psi(x,t)\Delta x = |A|^2 \Delta x \), where \( A \) is a complex constant.
- If \( A \) has real and complex parts (\( a + i b \)), then \( A^*A = a^2 + b^2 \), ensuring the probability density is real.
- The wavefunction for a particle confined between 0 and \( L \) is \( \psi(x,t) = A e^{-iEt/\hbar} \sin(\pi x/L) \) for \( 0 \leq x \leq L \) and zero otherwise.
- The normalization constant \( A \) is determined by ensuring the integral of \( |\psi(x,t)|^2 \) over all space equals 1.

# Expectation Values
- The expectation value of position is given by \( \langle x \rangle = \int_{-\infty}^\infty x P(x,t) dx = \int_{-\infty}^\infty x \Psi^*(x,t) \Psi(x,t) dx \).
- The expectation value of momentum in the \( x \)-direction is calculated using the operator \( -i\hbar d/dx \), resulting in \( \langle p \rangle = \int_{-\infty}^\infty \Psi^*(x,t) (-i\hbar d/dx) \Psi(x,t) dx \).
- The expectation value of kinetic energy is derived from the second derivative of the wavefunction, leading to \( \langle K \rangle = -\frac{\hbar^2}{2m} \int_{-\infty}^\infty \Psi^*(x,t) d^2/dx^2 \Psi(x,t) dx \).

# Symmetry and Odd Functions
- An even function satisfies \( \psi(x) = \psi(-x) \), while an odd function satisfies \( \psi(x) = -\psi(-x) \).
- The product of two even functions or two odd functions results in an even function, while the product of an even and an odd function results in an odd function.
- The integral of an odd function over all space is zero due to symmetry.

# Example Calculations
- For the wavefunction \( \psi(x) = e^{-|x|/x_0}/x_0 \), the expectation value of position is zero because the probability density is symmetric about \( x=0 \).
- The normalized wavefunction for a particle in a region between 0 and \( L \) is \( \psi(x,t) = A e^{-i\omega t} \sin(\pi x/L) \), with normalization constant \( A = 2/L \).
- The expectation value of position for this wavefunction is \( \langle x \rangle = L/2 \), while the expectation value of momentum is zero due to symmetry.

# Probability Density
- The probability density for the particle in a region between 0 and \( L \) is \( |\psi(x,t)|^2 = (2/L) \sin^2(\pi x/L) \).
- This probability density is largest at \( x = L/2 \) and zero at \( x=0 \) and \( x=L \).

# Correspondence Principle
- Niels Bohr asserted that quantum mechanics must agree with classical mechanics for macroscopic systems, suggesting classical mechanics is an approximation of quantum mechanics for large energies.

# Probability of Locating a Particle
- The probability of finding the particle between positions 0 and \( L/4 \) can be calculated by integrating the probability density over this interval.
