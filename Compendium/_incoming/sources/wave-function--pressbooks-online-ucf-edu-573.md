---
id: 01KDP8WN194JXGNYSB54DNH6ZW
slug: "wave-function--pressbooks-online-ucf-edu-573"
title: "Source: Wave Functions - University Physics Volume 3"
url: "https://pressbooks.online.ucf.edu/osuniversityphysics3/chapter/wave-functions/"
type: source
related_article: "wave-function"
created: 2025-12-30T00:00:43Z
tags: [""]
summary: "Summarized source material for Wave Function"
researcher_version: "0.3.29"
---

# Using the Wave Function
- A clue to the physical meaning of the wave function is provided by the two-slit interference of monochromatic light.
- The wave function of a light wave is given by \( E(x,t) \), and its energy density is proportional to \( E^2 \).
- The probability (per unit area) that a single photon will strike a particular spot on the screen is proportional to the square of the total electric field at that point.
- For matter particles, such as electrons, the square of the matter wave in one dimension has a similar interpretation as the square of the electric field. It gives the probability density that a particle will be found at a particular position and time per unit length.
- The probability \( P \) a particle is found in a narrow interval \( (x, x + dx) \) at time \( t \) is given by \( |\psi(x,t)|^2 dx \).
- This probabilistic interpretation of the wave function is called the Born interpretation.

# Wave Function and Probability
- The square of the wave function ensures that the probability is positive.
- For a particle in two dimensions, the integration is over an area; for three dimensions, it is over a volume. For simplicity, one-dimensional cases are considered initially.
- The normalization condition for the wave function is \( \int_{-\infty}^{\infty} |\psi(x,t)|^2 dx = 1 \).

# Where Is the Ball? (Part I)
- A ball is constrained to move along a line inside a tube of length \( L \). It is equally likely to be found anywhere in the tube.
- The wave function can be written as \( \psi(x) = C \), where \( C \) is a constant, and \( 0 \leq x \leq L \).
- Applying the normalization condition: \( \int_{-\infty}^{\infty} |\psi(x)|^2 dx = 1 \) simplifies to \( C^2 L = 1 \), giving \( C = \frac{1}{\sqrt{L}} \).
- The probability of finding the ball in the left half of the tube is 50%, as expected.

# Where Is the Ball? (Part II)
- A ball is constrained to move along a line inside a tube of length \( L \). It is found preferentially in the middle of the tube.
- The wave function can be written as \( \psi(x) = A \cos(kx + \phi) \), where \( A \) is the amplitude and \( k \) is the wave number.
- Boundary conditions give \( \psi(0) = 0 \) and \( \psi(L) = 0 \), leading to specific values of \( k \).
- Applying the normalization condition: \( \int_{-\infty}^{\infty} |\psi(x)|^2 dx = 1 \) gives \( A = \frac{1}{\sqrt{L}} \).
- The probability of finding the ball in the last one-quarter of the tube is approximately 9.1%.

# An Interpretation of the Wave Function
- The wave function can be used to determine where a particle is likely to be when a measurement is made.
- According to the Copenhagen interpretation, until observed, a particle exists in a superposition of states.
- Schrödinger's thought experiment with a cat illustrates the absurdity of this interpretation, but it remains widely accepted.

# Two-state Systems and Quantum Computers
- Two-state systems are used to illustrate principles of quantum mechanics, such as electron spin and mixed states.
- In quantum computing, qubits can be in a superposition of states (e.g., 0 and 1).
- If a large number of qubits are placed in the same state, measuring an individual qubit would yield 0 with probability \( p \) and 1 with probability \( 1-p \).

# Free Particles and Wave Functions
- A free particle's wave function is given by \( \psi(x,t) = A e^{i(kx - \omega t)} \), where \( A \) is the amplitude, \( k \) is the wave number, and \( \omega \) is the angular frequency.
- Using Euler’s formula, this can be written as \( \psi(x,t) = A (\cos(kx - \omega t) + i \sin(kx - \omega t)) \).
- The probability of finding the particle in a narrow interval \( (x, x + dx) \) is given by \( |\psi(x,t)|^2 dx \).

# Expectation Values
- The expectation value of position is given by \( \langle x \rangle = \int_{-\infty}^{\infty} \psi^*(x,t) x \psi(x,t) dx \).
- The expectation value of momentum is given by \( \langle p \rangle = \int_{-\infty}^{\infty} \psi^*(x,t) (-i\hbar \frac{d}{dx}) \psi(x,t) dx \).

# Normalization Constant
- For a particle with energy \( E \) confined between 0 and \( L \), the wave function is \( \psi(x) = \sqrt{\frac{2}{L}} \sin\left(\frac{n\pi x}{L}\right) \), where \( n \) is an integer.
- The normalization constant ensures that \( \int_{0}^{L} |\psi(x)|^2 dx = 1 \).

# Conclusion
- Quantum mechanics provides a probabilistic description of particle positions and other measurable quantities.
- The wave function encapsulates the probabilities of finding a particle in different states, and its interpretation remains a fundamental aspect of quantum theory.

# Wave Functions
- The momentum operator in the x-direction is denoted as \( \hat{p}_x = -i\hbar \frac{\partial}{\partial x} \), where \( \hbar \) is the reduced Planck's constant.
- The kinetic energy operator is given by \( \hat{T} = -\frac{\hbar^2}{2m} \frac{\partial^2}{\partial x^2} \), where \( m \) is the mass of the particle.
- A symmetric wave function can be either even or odd. An even function satisfies \( f(x) = f(-x) \), while an odd function satisfies \( f(x) = -f(-x) \).
- The expectation value of position for a normalized wave function \( \psi(x) \) is calculated as:
  \[
  \langle x \rangle = \int_{-\infty}^{\infty} x |\psi(x)|^2 dx
  \]
- For the wave function \( \psi(x) = A \sin\left(\frac{n\pi x}{L}\right) \), the normalization constant \( A \) is determined by:
  \[
  A = \sqrt{\frac{2}{L}}
  \]
- The expectation value of momentum for a particle in a one-dimensional box is zero due to the symmetry of the wave function.
- The expectation value of kinetic energy for a particle in a one-dimensional box is given by:
  \[
  \langle T \rangle = \frac{n^2\pi^2\hbar^2}{2mL^2}
  \]
- The probability density function for a particle in a one-dimensional box is largest at the midpoint \( x = L/2 \) and zero at the boundaries \( x = 0 \) and \( x = L \).
- Niels Bohr's correspondence principle states that quantum mechanics must agree with classical mechanics for macroscopic systems.
- The square of the wave function represents the probability density of finding a particle at a specific location in space (Born’s interpretation).
- A wave function must be normalized before making predictions about physical quantities.
- The physical unit of a wave function is \( \text{m}^{-1/2} \), and the square of the wave function has units of \( \text{m}^{-1} \).
- The magnitude of a wave function cannot be negative because it represents a probability amplitude, which must have a non-negative magnitude.
- A wave function represents the quantum state of a physical system and is used to predict measurable quantities such as position, momentum, and energy.
- The physical meaning of a wave function is that it encodes all the information about the state of a particle in quantum mechanics.
- The expectation value of a quantity is the average value that would be obtained for that quantity over many measurements on a large number of particles in the same quantum state.
- For the function \( \psi(x) = e^{ikx} \), where \( k \) is a real constant, the wave function represents a free particle with momentum \( \hbar k \).
- Among the given functions, only (a) and (e) can be normalized to represent valid wave functions.
- For a particle with mass \( m \) moving along the x-axis and wave function \( \psi(x) = A e^{-x^2/(2\sigma^2)} \):
  - The normalization constant \( A \) is \( \frac{1}{\sqrt{\sqrt{\pi} \sigma}} \).
  - The probability of finding the particle in the interval \( [0, L/4] \) can be calculated using:
    \[
    P = \int_{0}^{L/4} |\psi(x)|^2 dx
    \]
  - The expectation value of position is \( \langle x \rangle = \sigma^2 \).
  - The expectation value of momentum is zero due to the symmetry of the wave function.
  - The expectation value of kinetic energy is:
    \[
    \langle T \rangle = \frac{\hbar^2}{4m\sigma^2}
    \]
- A particle with mass \( m \) and wave function \( \psi(x) = e^{ikx} \):
  - The normalization constant is already satisfied as the wave function is normalized.
  - The probability of finding the particle in any interval can be calculated using:
    \[
    P = \int_{a}^{b} |\psi(x)|^2 dx
    \]
  - The average position is \( \langle x \rangle = 0 \) due to the symmetry of the wave function.
  - The expectation value of momentum is \( \hbar k \).
  - The expectation value of kinetic energy is:
    \[
    \langle T \rangle = \frac{\hbar^2 k^2}{2m}
    \]
