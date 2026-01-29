---
id: TEST123
domain: "Science"
domain-slug: "science"
category: "Physics"
category-slug: "physics"
topic: "Quantum Mechanics"
topic-slug: "quantum-mechanics"
article: "Math Rendering Test"
article-slug: "math-test"
github_issue_ids: []
github_pr_ids: []
created: 2026-01-29T00:00:00Z
model: "test"
---

## Math Rendering Test

This page tests different math formats to verify KaTeX rendering.

---

### Format 1: LaTeX Display Math \[...\]

The Dirac equation:
\[
(i \gamma^\mu \partial_\mu - m) \psi = 0
\]

Schrödinger equation:
\[
i\hbar \frac{\partial}{\partial t}\Psi = \hat{H}\Psi
\]

---

### Format 2: LaTeX Inline Math \(...\)

The wavefunction \( \psi \) describes the quantum state. The gamma matrices \( \gamma^\mu \) satisfy \( \{\gamma^\mu, \gamma^\nu\} = 2\eta^{\mu\nu} \).

Einstein's equation \( E = mc^2 \) relates mass and energy.

---

### Format 3: Dollar Sign Display Math $$...$$

Maxwell's equations in differential form:
$$
\nabla \cdot \mathbf{E} = \frac{\rho}{\epsilon_0}
$$

$$
\nabla \times \mathbf{B} = \mu_0\mathbf{J} + \mu_0\epsilon_0\frac{\partial\mathbf{E}}{\partial t}
$$

---

### Format 4: Dollar Sign Inline Math $...$

Planck's constant is $h = 6.626 \times 10^{-34}$ joule-seconds.

The reduced Planck constant is $\hbar = h / 2\pi$.

The speed of light is $c = 3 \times 10^8$ m/s.

---

### Format 5: Mixed Formats

Using both inline \( \alpha \) and $\beta$ in the same paragraph.

A display equation using brackets:
\[
\sum_{n=1}^{\infty} \frac{1}{n^2} = \frac{\pi^2}{6}
\]

And one using dollars:
$$
\int_{-\infty}^{\infty} e^{-x^2} dx = \sqrt{\pi}
$$

---

### Summary

If this page renders correctly:
- Display math should be centered on its own line
- Inline math should flow with the text
- Greek letters like α, β, γ should render as symbols
- Fractions, integrals, and summations should display properly
