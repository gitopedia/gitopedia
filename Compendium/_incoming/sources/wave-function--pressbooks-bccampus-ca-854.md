---
id: 01KDPA6NGDXR3856TE68592CYP
slug: "wave-function--pressbooks-bccampus-ca-854"
title: "Source: 3.1 Wave Functions - BCIT Phys8400: Modern Physics"
url: "https://pressbooks.bccampus.ca/bcitphys8400/chapter/3-1-wave-functions/"
type: source
related_article: "wave-function"
created: 2025-12-30T00:23:36Z
tags: [""]
summary: "Summarized source material for Wave Function"
researcher_version: "0.3.29"
---

# Chapter 3: Quantum Mechanics
## By the end of this section, you will be able to:
- Describe the statistical interpretation of the wave function
- Use the wave function to determine probabilities
- Calculate expectation values of position, momentum, and kinetic energy

---

# A Clue to the Physical Meaning of the Wave Function
- The wave function Ψ(x,t) is analogous to the electric field strength E(x,t) in light waves.
- The square of the electric field strength |E|² corresponds to the energy density of the light wave.
- For matter particles like electrons, the square of the wave function |Ψ|² gives the probability density of finding the particle at a specific position and time.

---

# The Square of the Wave Function
- The probability (P) that a particle is found in a narrow interval (x, x + dx) at time t is proportional to |Ψ(x,t)|².
- For a constant wave function Ψ(x,t) = C over an interval (0 < x < L), the normalization condition ensures that the total probability is 1:
  \[
  \int_0^L |C|^2 dx = 1
  \]
- The probability of finding the particle in the first half of the tube (0 < x < L/2) is calculated as:
  \[
  P(x=0, L/2) = \int_0^{L/2} \left|\frac{1}{L}\right|^2 dx = 0.50
  \]

---

# Where Is the Ball? (Part I)
- A ball in a tube of length L has a wave function Ψ(x,t) = C for 0 < x < L.
- The normalization constant C is determined by ensuring the total probability over the entire interval is 1:
  \[
  \int_0^L |C|^2 dx = 1
  \]
- The probability of finding the ball in the first half of the tube is 50%.

---

# Where Is the Ball? (Part II)
- A ball with a wave function Ψ(x,t) = Asin(kx - ωt) has a definite wavelength λ = 2L.
- If the tube length L = 1 m, the momentum p of the ball is:
  \[
  p = \frac{h}{\lambda} = \frac{h}{2L} \approx 10^{-36} \, \text{m/s}
  \]
- This momentum is too small to be measured by human instruments.

---

# An Interpretation of the Wave Function
- The wave function Ψ(x,t) describes the probability distribution of finding a particle at position x and time t.
- According to the Copenhagen interpretation:
  - When not being observed, a particle exists in a superposition of states (e.g., multiple possible positions).
  - Upon measurement, the wave function collapses to a specific state with a probability determined by |Ψ|².

---

# Thought Experiment: Schrödinger's Cat
- A cat in a box is described as being simultaneously alive and dead until observed.
- This illustrates the concept of superposition in quantum mechanics.

---

# Two-State Systems and Quantum Computers
- Two-state systems are fundamental in quantum mechanics, with applications in electron spin and quantum computing.
- A qubit exists in a mixed state, with probabilities p (for 0) and q = 1 - p (for 1).

---

# Complex Wave Functions
- The probability density is given by Ψ*(x,t)Ψ(x,t), where Ψ* is the complex conjugate of Ψ.
- This ensures that all predictions are real numbers.

---

# Expectation Values
- The expectation value of position is:
  \[
  \langle x \rangle = \int_{-\infty}^\infty x |\Psi(x,t)|^2 dx
  \]
- The expectation value of momentum is:
  \[
  \langle p \rangle = \int_{-\infty}^\infty p |\Psi(p,t)|^2 dp
  \]

---

# Summary
- The wave function Ψ(x,t) provides a mathematical description of a particle's state.
- |Ψ|² gives the probability density of finding the particle at a specific position and time.
- Quantum mechanics introduces concepts like superposition and wave function collapse, which differ from classical physics.

# Momentum Operators
- The momentum operator in the x-direction is denoted as \(-i\hbar \frac{d}{dx}\).
- Momentum operators for the y- and z-directions are defined similarly to the x-direction.
- The kinetic energy operator is given by \(-\frac{\hbar^2}{2m} \frac{d^2}{dx^2}\).

# Symmetry of Wave Functions
- A symmetric wave function can be even or odd.
- An even function satisfies \( \psi(x) = \psi(-x) \).
- An odd function satisfies \( \psi(x) = -\psi(-x) \).
- Examples of even functions include \( x^2 e^{-x^2} \), and examples of odd functions include \( x e^{-x^2} \).
- The integral over all space of an odd function is zero.

# Expectation Value (Part I)
- The normalized wave function is \( \psi(x) = e^{-|x|/x_0} \).
- The expectation value of position is calculated as:
  \[
  \langle x \rangle = \int_{-\infty}^{\infty} x |\psi(x)|^2 dx = 0
  \]
- This result is due to the symmetry of the wave function about \( x = 0 \).

# Expectation Value (Part II)
- The time-dependent wave function is \( \psi(x,t) = A e^{-i\omega t} \sin(\pi x / L) \).
- The normalization constant \( A \) is calculated as:
  \[
  A = \frac{2}{L}
  \]
- The expectation value of position is:
  \[
  \langle x \rangle = \int_0^L x |\psi(x)|^2 dx = \frac{L}{2}
  \]
- The expectation value of momentum is calculated as:
  \[
  \langle p \rangle = \int_0^L \psi^*(x) (-i\hbar \frac{d}{dx}) \psi(x) dx = 0
  \]
- The expectation value of kinetic energy is:
  \[
  \langle K \rangle = \int_0^L \psi^*(x) \left(-\frac{\hbar^2}{2m} \frac{d^2}{dx^2}\right) \psi(x) dx = \frac{\hbar^2}{2mL^2}
  \]
- The probability density is:
  \[
  |\psi(x)|^2 = \left(\frac{2}{L}\right) \sin^2\left(\frac{\pi x}{L}\right)
  \]
- The probability density is largest at \( x = L/2 \) and zero at \( x = 0 \) and \( x = L \).

# Correspondence Principle
- Niels Bohr asserted that quantum mechanics must agree with classical mechanics for macroscopic systems.
- This principle suggests that classical mechanics is an approximation of quantum mechanics for systems with very large energies.
