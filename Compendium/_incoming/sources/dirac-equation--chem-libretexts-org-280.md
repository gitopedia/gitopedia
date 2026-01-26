---
id: 01KFWXY39F32EYVHGCYXWW8HPH
slug: "dirac-equation--chem-libretexts-org-280"
title: "Source: 2.5: Introduction to the Dirac Equation - Chemistry LibreTexts"
url: "https://chem.libretexts.org/Courses/New_York_University/G25.2666%3A_Quantum_Chemistry_and_Dynamics/2%3A_General_theory_of_spin/2.5%3A_Introduction_to_the_Dirac_Equation"
type: source
related_article: "dirac-equation"
created: 2026-01-26T10:35:17Z
summary: "Summarized source material for Dirac Equation"
researcher_version: "0.3.29"
---

# Introduction to the Dirac Equation
- P.A.M. Dirac proposed a relativistic formulation of quantum mechanics for the electron in 1928, from which spin emerges naturally.
- Dirac's theory is essential for understanding low-lying states of heavy atoms where electron speeds approach the speed of light due to large Coulomb forces.
- The Dirac equation is the foundation of modern quantum electrodynamics (QED), one of the most accurate quantum theories.

# Problem with Relativistic Energy
- The relativistic energy of a free particle is given by \( E = \sqrt{p^2c^2 + m^2c^4} \).
- When \( p = 0 \), the energy reduces to the rest mass energy \( mc^2 \).
- In the non-relativistic limit, the kinetic energy \( p^2/(2m) \) is added to the rest mass energy.

# Attempts to Formulate a Relativistic Schrödinger Equation
- Promoting classical variables to quantum operators led to issues with interpreting square roots of operators.
- Squaring the Hamiltonian resulted in the Klein-Gordon equation, which has two solutions corresponding to forward and backward propagating waves.
- The Klein-Gordon equation does not incorporate spin and is only valid for spinless particles.

# Dirac's Approach
- Dirac sought a linear Hamiltonian that would square to the required relativistic energy expression.
- He proposed a general form of the Hamiltonian involving matrices \( \alpha \) and \( \beta \):
  \[
  H = c\, \boldsymbol{\alpha} \cdot \mathbf{P} + \beta mc^2
  \]
- The conditions for the matrices \( \alpha \) and \( \beta \) to satisfy anticommutation relations led to the conclusion that they must be traceless and anticommuting.

# Matrices in the Dirac Equation
- Four traceless, anticommuting matrices are required.
- A possible representation uses Pauli matrices and the identity matrix in a 4x4 block structure:
  \[
  \boldsymbol{\alpha} = 
  \begin{pmatrix}
    0 & \sigma \\
    \sigma & 0
  \end{pmatrix}, \quad
  \beta = 
  \begin{pmatrix}
    I & 0 \\
    0 & -I
  \end{pmatrix}
  \]
- These matrices ensure the correct anticommutation relations and traceless property.

# Solution of the Dirac Equation for a Free Particle
- The explicit form of the spin-\( 1/2 \) rotation operator is given by:
  \[
  R(\theta, \mathbf{n}) = e^{-i\theta \boldsymbol{\alpha} \cdot \mathbf{n}}
  \]
- This operator is used to solve the Dirac equation for a free particle, leading to solutions that describe both positive and negative energy states.

# Conclusion
- The Dirac equation successfully incorporates spin into relativistic quantum mechanics.
- It predicts the existence of antimatter and provides a framework for understanding high-energy phenomena in QED.
