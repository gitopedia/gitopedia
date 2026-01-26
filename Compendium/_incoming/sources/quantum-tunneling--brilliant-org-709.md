---
id: 01KFX5N71FDPB7BVD1KGW5KY10
slug: "quantum-tunneling--brilliant-org-709"
title: "Source: Quantum Tunneling | Brilliant Math & Science Wiki"
url: "https://brilliant.org/wiki/quantum-tunneling/"
type: source
related_article: "quantum-tunneling"
created: 2026-01-26T12:50:13Z
summary: "Summarized source material for Quantum Tunneling"
researcher_version: "0.3.29"
---

1. **Transmission Coefficient for a Rectangular Potential Barrier**  
   The transmission coefficient \( T \) for an electron encountering a one-dimensional barrier potential is given by:
   \[
   T = \left(1 + \frac{V_0^2}{4E(V_0 + E)} \sin^2\left(\frac{L}{\hbar} \sqrt{2m(V_0 + E)}\right)\right)^{-1}
   \]
   where:
   - \( V_0 = 3 \, \text{eV} \) (height of the barrier),
   - \( E = 1 \, \text{eV} \) (energy of the electron),
   - \( L = 1 \, \text{nm} \) (width of the barrier).

2. **Gamow Model of Radioactive Decay**  
   The transmission probability \( T \) for alpha decay is approximated by:
   \[
   T = e^{-2\gamma}
   \]
   where \( \gamma \) is defined as:
   \[
   \gamma = \frac{\sqrt{2mE}}{\hbar} \left(\frac{\pi}{2} r_1 - 2\sqrt{r_0 r_1}\right)
   \]
   Here, \( r_0 \) and \( r_1 \) are the radii where the energy \( E \) intersects the potential \( V(r) \).

3. **Applications of Quantum Tunneling**  
   - **Scanning Tunneling Microscopy (STM):** Uses quantum tunneling to map surface topography at atomic scales.
   - **Tunnel Diodes:** Utilize tunneling for unidirectional current flow in electronic devices.
   - **Josephson Junctions:** Enable superconducting currents through thin insulating barriers via Cooper pair tunneling.

4. **Impact of Potential Shift on Alpha Decay Probability**  
   If the equation for \( \gamma \) is shifted by \( \ln 2 / 2 \), the transmission coefficient (probability of alpha decay) changes exponentially:
   \[
   T \propto e^{-2(\gamma - \frac{\ln 2}{2})} = e^{-2\gamma + \ln 2}
   \]
   This implies a factor change in the transmission probability proportional to \( e^{\ln 2} = 2 \), effectively halving the decay probability.

**Final Answer:**  
The probability of alpha decay changes by a factor of \( \boxed{2} \).
