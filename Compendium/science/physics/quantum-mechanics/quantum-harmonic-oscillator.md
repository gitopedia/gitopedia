---
id: 01KFWQDFKEJ7T9HSZF41RTKE9V
domain: "Science"
domain-slug: "science"
category: "Physics"
category-slug: "physics"
topic: "Quantum Mechanics"
topic-slug: "quantum-mechanics"
article: "Quantum Harmonic Oscillator"
article-slug: "quantum-harmonic-oscillator"
github_issue_ids: [126, 127, 124, 121]
github_pr_ids: [139]
created: 2026-01-26T08:41:07Z
model: "deepseek-r1:14b"
iterations: 5
---

![Header](img/quantum-harmonic-oscillator_header.png)

## Overview

The quantum harmonic oscillator is a fundamental concept in quantum mechanics that describes the quantized motion of a particle undergoing harmonic oscillation. It is analogous to the classical harmonic oscillator, which consists of a mass attached to a spring that experiences a linear restoring force. However, the quantum version introduces discrete energy levels and wavefunctions that reflect the probabilistic nature of quantum states.

## Key Concepts

The energy levels of a quantum harmonic oscillator are quantized and given by the formula:

\[ E_n = \left(n + \frac{1}{2}\right)\hbar\omega \]

where \( n \) is the principal quantum number (a non-negative integer), \( \hbar \) is the reduced Planck's constant, and \( \omega \) is the angular frequency of the oscillator. The energy spacing between adjacent levels is:

\[ \Delta E = \hbar\omega \]

While this formula for energy spacing holds true theoretically, in real systems, it is only accurate for the lowest vibrational levels where the potential remains a good approximation of the 'mass on a spring' type harmonic potential. This principle plays a crucial role in determining the ground state energy of the quantum harmonic oscillator.

The natural frequency formula for a diatomic molecule's vibration relates the force constant (\( k \)) and reduced mass (\( \mu \)) as follows:

\[ \omega = \sqrt{\frac{k}{\mu}} \]

This phenomenon has significant implications for molecular behavior. Even at absolute zero, molecules modeled as quantum harmonic oscillators exhibit residual vibrational energy due to the zero-point energy:

\[ E_0 = \frac{1}{2}\hbar\omega \]

Additionally, the ground state is unique in that it represents the only eigenstate achieving the minimum uncertainty product in position and momentum space, satisfying Heisenberg’s Uncertainty Principle with equality.

The statistical distribution of these energy levels at thermal equilibrium is described by the harmonic oscillator partition function:

\[ q = \frac{e^{-\hbar\omega/2kT}}{1 - e^{-\hbar\omega/kT}} \]

This partition function, derived from summing over all energy levels, provides insight into how energy levels are distributed statistically under thermal conditions.

This surprising property highlights the fundamental differences between classical and quantum systems, where quantum effects ensure that molecular vibrations persist regardless of temperature.

The wavefunctions corresponding to these energy levels are expressed as:

\[ \psi_n(x) = N_n e^{-\beta^2 x^2 / 2} H_n(\beta x) \]

where \( N_n \) is the normalization constant, \( \beta = \frac{m\omega}{\hbar} \), and \( H_n(y) \) represents the nth Hermite polynomial. The normalization constant \( N_n \) ensures that the total probability of finding the particle within all possible positions integrates to 1, satisfying a fundamental requirement of quantum mechanics.

These wavefunctions describe the probability distribution of finding the particle in a given position.

The first four Hermite polynomials (\( H_0 \) to \( H_3 \)) are as follows:

- \( H_0(y) = 1 \)
- \( H_1(y) = 2y \)
- \( H_2(y) = 4y^2 - 2 \)
- \( H_3(y) = 8y^3 - 12y \)

These polynomials directly shape the probability density distributions of the quantum states, with each subsequent polynomial corresponding to higher energy levels and more complex wavefunction shapes.

## Wavefunction Derivation from Schrödinger Equation

The wavefunctions of the quantum harmonic oscillator are derived by solving the time-independent Schrödinger equation (TISE) for the potential energy of a harmonic oscillator. The TISE is given by:

\[
-\frac{\hbar^2}{2m} \frac{d^2 \psi(x)}{dx^2} + V(x)\psi(x) = E_n \psi(x)
\]

where \( V(x) = \frac{1}{2} m \omega^2 x^2 \) is the potential energy, \( m \) is the mass of the particle, and \( \omega \) is the angular frequency of the oscillator. Substituting the potential into the TISE leads to:

\[
-\frac{\hbar^2}{2m} \frac{d^2 \psi(x)}{dx^2} + \frac{1}{2} m \omega^2 x^2 \psi(x) = E_n \psi(x)
\]

To solve this differential equation, a dimensionless variable \( y = \alpha x \), where \( \alpha = \sqrt{\frac{m \omega}{\hbar}} \), is introduced. This substitution transforms the equation into:

\[
\frac{d^2 \psi(y)}{dy^2} + \left( y^2 - 2n - 1 \right) \psi(y) = 0
\]

The solutions to this equation are well-known and involve Hermite polynomials \( H_n(y) \), which are orthogonal and normalizable functions. The resulting wavefunctions take the form:

\[
\psi_n(x) = \left( \frac{\alpha}{\sqrt{\pi}} \right)^{1/2} e^{-y^2 / 2} H_n(y)
\]

where \( y = \alpha x \). These solutions are orthonormal and satisfy the normalization condition:

\[
\int_{-\infty}^{\infty} |\psi_n(x)|^2 dx = 1
\]

The Hermite polynomials arise naturally from the requirement of normalizability and orthogonality, ensuring that each wavefunction corresponds to a specific energy eigenvalue \( E_n \). This derivation establishes the mathematical foundation for the quantum harmonic oscillator's wavefunctions, which are essential in understanding quantized systems.

## Differences Between Classical and Quantum Harmonic Oscillators

The classical harmonic oscillator's total energy is the sum of its kinetic and potential energies. The system has defined turning points where the kinetic energy becomes zero, marking the extremes of displacement.

One key difference between classical and quantum harmonic oscillators lies in their ground state energy. While the classical oscillator has zero minimum energy, the quantum version has a nonzero ground state energy:

\[ E_0 = \frac{\hbar\omega}{2} \]

Another distinction is that both the classical and quantum harmonic oscillators share the same frequency formula. The frequency of the quantum harmonic oscillator is given by:

\[ f = \frac{1}{2\pi} \sqrt{\frac{k}{m}} \]

This is identical to the frequency of a classical simple harmonic oscillator, where \( k \) is the spring constant and \( m \) is the mass of the particle.

Furthermore, there are differences in the probability distribution of the particle's position. In the classical case, the probability density is highest near the turning points of the oscillation. However, for the quantum oscillator, especially in the ground state, the probability density peaks at the equilibrium position \( x = 0 \). Additionally, a quantum particle can be found with nonzero probability outside the classically forbidden region, which extends beyond the turning points. This probability is quantified to be approximately 16% for the ground state.

The classical oscillator has zero minimum energy, while the quantum version has a nonzero ground state energy:

\[ E_0 = \frac{\hbar\omega}{2} \]

Another distinction is in the probability distribution of the particle's position. In the classical case, the probability density is highest near the turning points of the oscillation. However, for the quantum oscillator, especially in the ground state, the probability density peaks at the equilibrium position \( x = 0 \). Additionally, a quantum particle can be found with nonzero probability outside the classically forbidden region, which extends beyond the turning points.

## Applications and Examples  

The quantum harmonic oscillator is fundamental in modeling molecular vibrations, particularly in diatomic molecules. Vibrational energy levels of these systems are described by the equation:

\[ \varepsilon_n = \hbar \omega \left( n + \frac{1}{2} \right) \]

where \( \omega \) is determined by the force constant and reduced mass of the molecule, given by \( \omega = \sqrt{\frac{k}{\mu}} \). Knowing the vibrational frequency \( \nu \), one can calculate the force constant \( k \) using \( k = m \omega^2 \), where \( \omega = 2\pi\nu \) and \( m \) is the reduced mass.

For example, in hydrogen chloride (HCl), the vibrational frequency is \( f = 8.88 \times 10^{13} \, \text{Hz} \). Using this, the force constant \( k \) can be calculated as approximately \( 520 \, \text{N/m} \). Similarly, for hydrogen iodide (HI), the vibrational frequency is \( f = 6.69 \times 10^{13} \, \text{Hz} \).

The reduced mass \( \mu \) for a diatomic molecule is given by:

\[ \mu = \frac{m_1 m_2}{m_1 + m_2} \]

In addition to its role in molecular vibrations, the quantum harmonic oscillator models wave packets in quantum optics. These wave packets are localized oscillatory solutions that describe light-matter interactions, providing a framework for understanding phenomena such as coherent states and Fock state superpositions.

This model allows determination of vibrational frequencies and energy levels in molecular systems, providing insight into the quantized nature of molecular vibrations.

## Thermodynamic Properties from Partition Function

The partition function of the quantum harmonic oscillator provides a pathway to calculate key thermodynamic quantities such as vibrational energy and heat capacity. By leveraging statistical mechanics, these properties can be derived directly from the partition function, which encapsulates the system's quantum mechanical behavior.

The internal energy \( U \) for a system of \( N \) quantum harmonic oscillators can be calculated using:

\[
U = Nk \left( \frac{\hbar\omega}{2} + \frac{\hbar\omega}{kT} \cdot \frac{e^{-\hbar\omega/kT}}{1 - e^{-\hbar\omega/kT}}} \right)
\]

where \( k \) is Boltzmann's constant, \( T \) is the temperature, and \( \hbar\omega \) represents the quantum energy spacing. This expression combines contributions from both the zero-point energy (the minimum energy at \( n=0 \)) and the thermal excitations.

The heat capacity \( C_v \), which measures the system's ability to absorb or release heat, is derived as:

\[
C_v = Nk \left( \frac{\hbar\omega}{kT} \right)^2 \cdot \frac{e^{\hbar\omega/kT}}{(1 - e^{-\hbar\omega/kT})^2}
\]

This formula reveals how the heat capacity depends on temperature and quantum energy spacing. At high temperatures (\( kT \gg \hbar\omega \)), \( C_v \) approaches the classical limit of \( Nk \), demonstrating the correspondence principle's validity in this regime.

These derivations highlight the power of statistical mechanics in connecting microscopic quantum mechanical models to macroscopic thermodynamic properties, enabling predictions about material behavior under various conditions.

## Correspondence Principle

The quantum description converges to the classical one in highly excited states (\( n \to \infty \)). In these cases, the probability density distribution of the quantum oscillator resembles the classical probability distribution. This convergence underscores the correspondence principle, which bridges quantum and classical mechanics at high energy levels.

## Classical Limit of Harmonic Oscillator

The classical limit of the quantum harmonic oscillator is observed at high temperatures where \( kT \gg \hbar\omega \). In this regime, the quantum system's behavior converges to classical predictions. The energy of the system approximates \( E \approx NkT \), aligning with the equipartition theorem in classical physics, which states that each degree of freedom contributes an average energy of \( \frac{1}{2}kT \). Consequently, for a harmonic oscillator with two degrees of freedom (kinetic and potential), the total energy becomes \( E = NkT \).

Similarly, the heat capacity \( C_v \) approaches \( Nk \) at high temperatures, consistent with classical results. This transition underscores the correspondence principle, which bridges quantum and classical mechanics at high energy levels, demonstrating the validity of classical physics in macroscopic or high-temperature regimes.

## Symmetry and Expectation Values

The symmetry properties of the wavefunctions play a crucial role in determining expectation values. For example, the expectation value of the position \( \langle x \rangle \) for a particle in the ground state is zero due to the even nature of the wavefunction, which ensures that the probability density is symmetric about the equilibrium position.

The quantum harmonic oscillator serves as a cornerstone in understanding quantized systems and has applications in various areas of physics, chemistry, and engineering. Its mathematical framework provides insights into more complex quantum systems and phenomena.

## References

[^1]: [7.6: The Quantum Harmonic Oscillator - Physics LibreTexts](https://phys.libretexts.org/Bookshelves/University_Physics/University_Physics_(OpenStax)/University_Physics_III_-_Optics_and_Modern_Physics_(OpenStax)/07%3A_Quantum_Mechanics/7.06%3A_The_Quantum_Harmonic_Oscillator)
[^2]: [7.5 The Quantum Harmonic Oscillator - University Physics ... - OpenStax](https://openstax.org/books/university-physics-volume-3/pages/7-5-the-quantum-harmonic-oscillator)
[^3]: [Quantum Harmonic Oscillator - HyperPhysics](http://www.hyperphysics.phy-astr.gsu.edu/hbase/quantum/hosc.html)
[^4]: [Quantum Harmonic Oscillator | Brilliant Math & Science Wiki](https://brilliant.org/wiki/quantum-harmonic-oscillator/)
[^5]: [The Quantum Harmonic Oscillator - University Physics Volume 3](https://pressbooks.online.ucf.edu/osuniversityphysics3/chapter/the-quantum-harmonic-oscillator/)

## Classical Harmonic Oscillator Mechanics

The classical harmonic oscillator operates under Hooke's law, which states that the restoring force is directly proportional to the displacement and acts in the opposite direction: \( F = -kx \). This leads to oscillatory motion where the system moves between two extreme points, known as turning points. The position of these points depends on the amplitude \( A \) of the oscillation, given by \( x = \pm A \).

The total mechanical energy of a classical harmonic oscillator is the sum of its kinetic and potential energies:

\[
E = \frac{1}{2}mv^2 + \frac{1}{2}kx^2
\]

This energy remains constant throughout the motion, with kinetic energy reaching its maximum when the particle passes through the equilibrium position (\( x = 0 \)) and potential energy peaking at the turning points. The system's behavior is described by simple harmonic motion (SHM), where displacement \( x(t) \) as a function of time follows sinusoidal functions.

The classical oscillator exhibits periodic motion with a well-defined frequency, given by:

\[
\omega = \sqrt{\frac{k}{m}}
\]

where \( k \) is the spring constant and \( m \) is the mass. This angular frequency determines the period \( T \) and the time it takes for the system to complete one full oscillation cycle.

In contrast to the quantum harmonic oscillator, which has a nonzero ground state energy (zero-point energy), the classical counterpart can theoretically have zero minimum energy. The classical model also lacks wavefunction characteristics and does not exhibit probability distributions outside the classically allowed region, as is possible in quantum mechanics.

The classical harmonic oscillator's simplicity makes it an essential tool for understanding more complex systems and serves as a foundation for analyzing phenomena such as resonance, damping, and coupled oscillations. Its mathematical framework provides insights into various physical processes, from mechanical vibrations to electrical circuits modeled using springs and masses.
