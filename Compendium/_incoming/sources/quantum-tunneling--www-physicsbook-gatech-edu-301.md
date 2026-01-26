---
id: 01KFX5BDBJ0SANN5NTVE5BMPJW
slug: "quantum-tunneling--www-physicsbook-gatech-edu-301"
title: "Source: Quantum Tunneling through Potential Barriers - Physics Book"
url: "https://www.physicsbook.gatech.edu/Quantum_Tunneling_through_Potential_Barriers"
type: source
related_article: "quantum-tunneling"
created: 2026-01-26T12:44:46Z
summary: "Summarized source material for Quantum Tunneling"
researcher_version: "0.3.29"
---

# Quantum Tunneling through Potential Barriers
- Quantum tunneling is a phenomenon where a particle has a nonzero probability of appearing on the far side of a potential barrier, even when its classical energy is insufficient to cross the barrier.
- This behavior arises from the wave-like nature of matter and is described by quantum mechanics principles such as the Schrödinger equation and the Heisenberg uncertainty principle.

# Classical Expectations
- In classical physics, if a particle's energy (E) is less than the potential barrier height (V₀), it cannot cross the barrier.
- Transmission probability is strictly zero in classical mechanics for E < V₀.

# Wave Behavior of Particles
- Quantum mechanics models particles as wavefunctions that satisfy the time-independent Schrödinger equation.
- These wavefunctions can extend into and beyond regions forbidden in classical mechanics, allowing for tunneling.
- The wavefunction decays exponentially within the barrier but continues to "leak through" on the far side.

# Heisenberg Uncertainty Principle
- The uncertainty principle provides intuitive support for tunneling by suggesting that a particle's energy cannot be known exactly over extremely short time intervals.
- This uncertainty allows temporary fluctuations that permit penetration into forbidden regions.

# The Schrödinger Equation
- The time-independent Schrödinger equation is given by:
  \[
  -\frac{\hbar^2}{2m}\frac{d^2\psi(x)}{dx^2} + V(x)\psi(x) = E\psi(x)
  \]
- For a finite potential barrier (V₀), the equation changes in regions where the particle's energy is less than the barrier height.

# Boundary Conditions and Wave Functions
- The wave functions for each region are:
  - Region I: \( \psi(x)_I = A e^{ikx} + B e^{-ikx} \)
  - Region II: \( \psi(x)_{II} = C e^{-\alpha x} + D e^{\alpha x} \)
  - Region III: \( \psi(x)_{III} = F e^{ikx} \)
- Constants (A, B, C, D, F) represent the amplitudes of incident, reflected, and transmitted waves.

# Tunneling Probability
- The tunneling probability is the ratio of the transmitted wave intensity (\( |F|^2 \)) to the incident wave intensity (\( |A|^2 \)).
- The formula for tunneling probability is:
  \[
  T \approx \frac{16 E V_0 (1 - \frac{E}{V_0})}{e^{2kL}}
  \]
  where \( k = \frac{2m(V₀ - E)}{\hbar^2} \) and L is the barrier width.

# Example Calculation
- For a particle with 5.0 eV of energy and a potential barrier of 10.0 eV with a width of 1.00 nm:
  - \( k = 1.145 \times 10^{10} m^{-1} \)
  - Tunneling probability (\( T \)) ≈ \( 4.52 \times 10^{-10} \)

# Applications of Quantum Tunneling
- **Nuclear Fusion in Stars**: Enables protons to overcome electrostatic repulsion, facilitating fusion reactions.
- **Scanning Tunneling Microscopy (STM)**: Uses tunneling currents between a sharp tip and surface for atomic resolution imaging.
- **Semiconductor Devices**: Exploits controlled tunneling currents in devices like tunnel diodes and modern transistors.

# Historical Background
- The concept of barrier penetration was introduced by physicist Friedrich Hund in the late 1920s.
- George Gamow applied tunneling to explain alpha decay, establishing it as a central quantum phenomenon.

# Related Concepts
- **Heisenberg Uncertainty Principle**: Fundamental to understanding tunneling.
- **Wave-Particle Duality**: Underpins the wave-like behavior of particles.
- **Quantum Wells and Bound States**: Relevant for studying tunneling in confined systems.
