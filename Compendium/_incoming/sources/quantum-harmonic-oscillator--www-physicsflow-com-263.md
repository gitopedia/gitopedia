---
id: 01KFWRQK407NJP3BKS60N1PA8N
slug: "quantum-harmonic-oscillator--www-physicsflow-com-263"
title: "Source: Quantum harmonic oscillator and coherent states"
url: "https://www.physicsflow.com/g/4.1.3"
type: source
related_article: "quantum-harmonic-oscillator"
created: 2026-01-26T09:04:08Z
summary: "Summarized source material for Quantum Harmonic Oscillator"
researcher_version: "0.3.29"
---

# Quantum Harmonic Oscillator
- The quantum harmonic oscillator is one of the most important model systems in quantum mechanics.
- It describes a particle subject to a restoring force proportional to its displacement from an equilibrium position, similar to a weight on a spring or a pendulum.
- In quantum mechanics, it serves as the basis for understanding more complex systems including fields and particles.

# Basics of Quantum Harmonic Oscillator
- The classical harmonic oscillator can be represented by the second-order differential equation: 
  \[ m\frac{d^2x}{dt^2} + kx = 0 \]
  where \( m \) is the mass, \( k \) is the spring constant, and \( x \) is the displacement from equilibrium.

# Quantum Mechanical Description
- The Hamiltonian for the harmonic oscillator is given by:
  \[ H = \frac{p^2}{2m} + \frac{1}{2}m\omega^2x^2 \]
  where \( p \) is the momentum operator, \( \omega \) is the angular frequency, and \( x \) is the position operator.
- The time-independent Schrödinger equation for the harmonic oscillator is:
  \[ H\psi = E\psi \]
  where \( \psi \) is the wave function and \( E \) is the energy eigenvalue.

# Energy Levels of a Quantum Harmonic Oscillator
- The quantized energies are given by:
  \[ E_n = \left(n + \frac{1}{2}\right)\hbar\omega \]
  where \( n \) is a non-negative integer (quantum number), \( \hbar \) is the reduced Planck constant, and \( \omega \) is the angular frequency.
- The oscillator has zero-point energy, \( \frac{1}{2}\hbar\omega \), even in the ground state (\( n=0 \)).

# Wave Functions of the Quantum Harmonic Oscillator
- The wave functions are Hermite polynomials multiplied by a Gaussian factor:
  \[ \psi_n(x) = N_n H_n(\xi) e^{-\xi^2/2} \]
  where \( H_n(\xi) \) are the Hermite polynomials, \( \xi = \sqrt{\frac{m\omega}{\hbar}} x \), and \( N_n \) is a normalization factor.

# Generalization and Orthogonality
- The wave functions are orthogonal:
  \[ \int_{-\infty}^{\infty} \psi_m(x)\psi_n(x) dx = \delta_{nm} \]
  where \( \delta_{nm} \) is the Kronecker delta.

# Example of Hermite Polynomials
- The first few Hermite polynomials are:
  - \( H_0(\xi) = 1 \)
  - \( H_1(\xi) = 2\xi \)
  - \( H_2(\xi) = 4\xi^2 - 2 \)

# Coherent States
- Coherent states are eigenstates of the annihilation operator \( a \):
  \[ a|α\rangle = α|α\rangle \]
  where \( α \) is a complex number.
- The annihilation operator \( a \) is related to the position and momentum operators as follows:
  \[ a = \sqrt{\frac{m\omega}{2\hbar}} x + i\sqrt{\frac{1}{2m\omega\hbar}} p \]

# Properties of Coherent States
- Coherent states are normalized such that \( \langle α|α\rangle = 1 \).
- They satisfy the minimal uncertainty relation, making them as close to classical states as possible.
- The overlap between two coherent states is given by:
  \[ \langle β|α\rangle = \exp\left(-\frac{|β|^2}{2}\right) \exp\left(-\frac{|α|^2}{2}\right) \exp(β^*α) \]

# Visualization Using Circles in the Complex Plane
- A coherent state can be viewed as a point \( α \) in the complex plane, where the real and imaginary parts represent different aspects of the oscillatory motion.

# Time Evolution of Coherent States
- Under the evolution determined by the harmonic oscillator Hamiltonian:
  \[ |α\rangle \rightarrow e^{-iωt/2}e^{iθ}|α\rangle \]
  This means that the state rotates in the complex plane, but its shape and size remain unchanged.

# Physical Applications and Significance
- Quantum harmonic oscillators and coherent states have many applications in physics:
  - **Quantum Optics**: Coherent states model laser light, which exhibits properties similar to classical electromagnetic waves.
  - **Quantum Field Theory**: The concepts of harmonic oscillators are used to describe fundamental particles and fields.
  - **Molecular Physics**: Vibrational modes of molecules are analyzed using quantum harmonic oscillators.

# Example - Simple Pendulum Model
- A pendulum swinging back and forth with small amplitude can be approximated by a harmonic oscillator.
- Its behavior is described by quantized energy levels according to the quantum harmonic oscillator solution.

# Conclusion
- Quantum harmonic oscillators and coherent states provide important insights into quantum mechanics and bridge the gap between quantum and classical physics.
- Understanding these concepts is essential for advanced studies in quantum mechanics, making them important for theoretical physics and beyond.
