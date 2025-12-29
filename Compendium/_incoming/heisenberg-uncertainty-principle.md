---
id: 01KDP6G127JNW1KB6RN9HMXSSR
title: "Heisenberg Uncertainty Principle"
slug: "heisenberg-uncertainty-principle"
created: 2025-12-29T23:18:34Z
tags: ["topic:heisenberg-uncertainty-principle", "topic:quantum-mechanics", "topic:planck-constant", "topic:heisenberg-uncertainty-principle"]
people: ["person:werner-heisenberg"]
researcher_version: "1"
model: "deepseek-r1:14b"
iterations: 0
summary: "Initial overview based on 7.3: The Heisenberg Uncertainty Principle - Physics LibreTexts"
---

# Heisenberg Uncertainty Principle

## Overview
The Heisenberg Uncertainty Principle is a fundamental concept in quantum mechanics that places inherent limits on the precision with which certain pairs of physical properties can be simultaneously measured. Formulated by Werner Heisenberg in 1927, it reveals that there are intrinsic limitations to knowledge in quantum systems due to the wave-particle duality of matter and energy.

## Key Concepts

### Momentum and Position Uncertainty
The principle asserts that the product of the uncertainties in position (Δx) and momentum (Δp) cannot be less than half the reduced Planck constant:
\[
\Delta x \Delta p \geq \frac{\hbar}{2}
\]
This relationship implies that a particle's position and momentum cannot both be precisely determined simultaneously. A localized wave packet, which minimizes uncertainty (Gaussian function), satisfies this condition with equality.

### Energy and Time Uncertainty
Similarly, the energy-time uncertainty principle states:
\[
\Delta E \Delta t \geq \frac{\hbar}{2}
\]
A quantum state with a short lifetime has an uncertain energy, leading to a broad spectral line. This principle applies to processes like atomic transitions.

## Applications

### Example Calculations
- **Electron**: With speed precision of \(1.0 \times 10^{-3}\) m/s:
  \[
  \Delta p = (9.1 \times 10^{-31} \text{ kg})(1.0 \times 10^{-3} \text{ m/s}) = 9.1 \times 10^{-34} \text{ kg·m/s}
  \]
  \[
  \Delta x \approx 5.8 \text{ cm}
  \]

- **Bowling Ball**: With mass 6.0 kg:
  \[
  \Delta p = (6.0 \text{ kg})(1.0 \times 10^{-3} \text{ m/s}) = 6.0 \times 10^{-3} \text{ kg·m/s}
  \]
  \[
  \Delta x \approx 8.8 \times 10^{-33} \text{ m}
  \]

### Ground-State Energy of a Hydrogen Atom
Using the uncertainty principle:
\[
E_0 = \frac{\hbar^2}{8mL^2}
\]
For \(L \approx 0.1\) nm, \(E_0 \approx 1 \text{ eV}\).

### Sodium Atom Transition
Transition emits a photon with energy 2.105 eV and wavelength 589.0 nm. With lifetime \(\Delta t = 1.6 \times 10^{-8} \text{ s}\):
\[
\Delta E \approx \frac{\hbar}{2\Delta t} \approx 4.1 \times 10^{-8} \text{ eV}
\]
Spectral line width:
\[
\Delta \lambda \approx 1.1 \times 10^{-5} \text{ nm}
\]

## Conclusion
The Heisenberg Uncertainty Principle imposes fundamental limits on measurement precision in quantum mechanics, reflecting the wave-like nature of matter and energy rather than experimental limitations.

## References

[^1]: [7.3: The Heisenberg Uncertainty Principle - Physics LibreTexts](https://phys.libretexts.org/Bookshelves/University_Physics/University_Physics_(OpenStax)/University_Physics_III_-_Optics_and_Modern_Physics_(OpenStax)/07%3A_Quantum_Mechanics/7.03%3A_The_Heisenberg_Uncertainty_Principle)