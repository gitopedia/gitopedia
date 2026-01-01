---
id: 01KDWDX7V66GF0G8Z6C7EYYS31
slug: "wave-function--en-wikipedia-org-4"
title: "Source: Wave function - Wikipedia"
url: "https://en.wikipedia.org/wiki/Wave_function"
type: source
related_article: "wave-function"
created: 2026-01-01T09:23:50Z
tags: ["topic:quantum-mechanics", "topic:wave-function", "topic:hilbert-space", "topic:probability-interpretation", "topic:schrdinger-equation", "topic:dirac-equation", "topic:fock-space", "topic:wave-function"]
people: ["person:albert-einstein", "person:erwin-schrdinger", "person:niels-bohr", "person:david-bohm"]
summary: "Summarized source material for Wave Function"
researcher_version: "0.3.29"
used_status: "integrated"
used_at: 2026-01-01T22:27:50+13:00
---

### EXTRACTED FACTS:

1. **Wave Function in Quantum Mechanics**  
   - The wave function is a mathematical description of the quantum state of a system, represented as an abstract vector in Hilbert space.
   - It can be expressed as:  
     \[
     | \Psi \rangle = \sum_{\boldsymbol{\alpha}} \int d^m \omega \, \Psi_t({\boldsymbol{\alpha}}, {\boldsymbol{\omega}}) | {\boldsymbol{\alpha}}, {\boldsymbol{\omega}} \rangle
     \]  
     where:
     - \(|{\boldsymbol{\alpha}}, {\boldsymbol{\omega}}\rangle\) are basis vectors of the chosen representation.
     - \(d^m \omega = d\omega_1 d\omega_2 \dots d\omega_m\) is a differential volume element in continuous degrees of freedom.
     - \(\Psi_t({\boldsymbol{\alpha}}, {\boldsymbol{\omega}})\) are the components of the vector \(| \Psi \rangle\).

2. **Hilbert Space Structure**  
   - The set of all possible wave functions constitutes an infinite-dimensional Hilbert space, which is not unique due to multiple possible choices of representation and basis.
   - Each state is represented as an abstract vector in this Hilbert space.

3. **Probability Interpretation**  
   - The probability density of finding the system at time \(t\) in state \(|{\boldsymbol{\alpha}}, {\boldsymbol{\omega}}\rangle\) is:  
     \[
     \rho_{\alpha, \omega}(t) = |\Psi({\boldsymbol{\alpha}}, {\boldsymbol{\omega}}, t)|^2
     \]
   - The probability of finding the system in a specific configuration is given by summing and integrating over the density.

4. **Normalization Condition**  
   - The normalization condition ensures that the total probability sums to 1:  
     \[
     1 = \sum_{\boldsymbol{\alpha} \in A} \int_{\Omega} d^m \omega \, \rho_{\alpha, \omega}(t)
     \]

5. **Interpretations of the Wave Function**  
   - Different interpretations exist regarding the physical meaning of the wave function:
     - **Copenhagen Interpretation**: The wave function represents probabilities in the mind of the observer.
     - **Objective Reality**: Some argue that the wave function has an objective, physical existence (e.g., David Bohm).
     - **Mind Projection**: Others consider it a measure of our knowledge of reality.

6. **Units and Dimensions**  
   - The wave function must have units such that \(\rho dm\omega\) is dimensionless, implying \(\Psi\) has the same units as \((\omega_1 \omega_2 \dots \omega_m)^{-1/2}\).

7. **Historical Perspectives**  
   - Prominent physicists like Erwin Schrödinger, Albert Einstein, and Niels Bohr debated the nature of the wave function.
   - Einstein believed a complete description should refer directly to physical space and time, distinct from the abstract mathematical space of the wave function.

8. **Applications and Variations**  
   - The wave function is used in various formulations, such as the Schrödinger equation for non-relativistic systems and the Dirac equation for relativistic systems.
   - In quantum field theory, the underlying Hilbert space is Fock space, built from free single-particle states.

9. **Finite-Dimensional Cases**  
   - Finite-dimensional Hilbert spaces (e.g., \(C^n\)) are used to describe spin states and other subsystems within larger quantum systems.

10. **Relaxing Conditions**  
    - While wave functions are typically required to be square-integrable, continuous, and differentiable, these conditions can sometimes be relaxed for special purposes or specific potentials.

---

### SUMMARY:
The wave function in quantum mechanics is a fundamental concept that describes the state of a system in an abstract Hilbert space. It has various interpretations and applications across different physical scenarios, from single-particle systems to quantum field theory. The mathematical structure and physical meaning of the wave function have been extensively debated and remain central to understanding quantum mechanics.
