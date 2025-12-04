---
id: 01KBNK1J8QEKMV9N77P9MR8TDV
slug: "quantum-information-science--en-wikipedia-org-9"
title: "Source: Von Neumann entropy - Wikipedia"
url: "https://en.wikipedia.org/wiki/Von_Neumann_entropy"
type: source
related_article: "quantum-information-science"
created: 2025-12-04T21:13:02Z
tags: ["topic:quantum-information-theory", "topic:quantum-entropy", "topic:von-neumann-entropy", "topic:shannon-entropy", "topic:gibbs-entropy", "topic:density-matrix", "topic:quantum-entanglement", "topic:quantum-mechanics", "topic:probability", "topic:hilbert-space", "topic:qubit", "topic:pauli-matrices", "topic:pure-state", "topic:mixed-state", "topic:wavefunction", "topic:subadditivity", "topic:strong-subadditivity", "topic:holevos-theorem", "topic:bell-state", "topic:entanglement", "topic:quantum-channel", "topic:gibbs-states", "topic:boltzmann-constant", "topic:quantum-mutual-information", "topic:quantum-rnyi-entropies", "topic:principle-of-maximum-entropy", "topic:thermodynamic-potential", "topic:ferromagnetism", "topic:ising-model", "topic:potts-model", "topic:heisenberg-model", "topic:percolation", "topic:lennard-jones-potential", "topic:mean-field-theory", "topic:conformal-field-theory", "topic:critical-exponents", "topic:statistical-field-theory", "topic:superfluidity", "topic:condensed-matter-physics", "topic:complex-systems", "topic:chaos-theory", "topic:information-theory", "topic:boltzmann-machine", "topic:quantum-mechanical-entropy", "topic:quantum-information-science"]
people: ["person:john-von-neumann", "person:karl-kraus", "person:gerhart-lders", "person:von-neumann", "person:lev-landau", "person:albert-einstein", "person:leo-szilard", "person:max-delbrck", "person:gert-molire", "person:hisaharu-umegaki", "person:huzihiro-araki", "person:elliott-h.-lieb", "person:oscar-lanford", "person:derek-robinson", "person:mary-beth-ruskai", "person:bengtsson-ingemar", "person:yczkowski-karol", "person:alexander-s.-holevo", "person:michael-a.-nielsen", "person:isaac-l.-chuang", "person:asher-peres", "person:jochen-rau", "person:eleanor-rieffel", "person:wolfgang-polak", "person:mark-m.-wilde", "person:barton-zwiebach"]
orgs: ["org:wikipedia", "org:wikimedia-foundation", "org:cambridge-university-press", "org:springer", "org:mit-press", "org:oxford-university-press"]
summary: "Summarized source material for Quantum Information Science"
researcher_version: "0.3.7"
---

# The outline lacks a dedicated section on quantum information theory's mathematical foundations (e.g., quantum entropy, von Neumann entropy)

- Please don't skip this 1-minute read.
- It's Friday, 5 December, and this moment matters for Wikipedia because we haven't hit our goal yet.
- If you're like us, you've used Wikipedia countless times.
- To settle an argument with a friend.
- Whether it's 3 in the morning or afternoon, Wikipedia is useful in your life.

- After nearly 25 years, Wikipedia is still the internet we were promised—created by people, not by machines. It's not perfect, but it's not here to push a point of view. It's owned by a nonprofit, not a giant technology company or a billionaire.

- If you are undecided, remember that most readers don't donate, so any contribution helps—and monthly gifts are the best option for a lasting impact. So if Wikipedia still provides you with $2.75 of knowledge, kindly donate today. Thank you.

- 25 years of the internet at its best

- $25 Celebrate & Give 🎉

- Unlock tax benefits by directing your donation via your Donor-Advised Fund (DAF)

- Qualified Charitable Distributions from a tax efficient eligible IRA

- Involve your employer and increase the impact of your donation

- Editing Wikipedia is easier than you might think! You will get suggestions and help as you start.

- Nearly half of our budget goes toward supporting the technology that powers Wikipedia and other Wikimedia projects.
- We are constantly working to enhance the user experience for both contributors and readers, improve site security, and ensure reliable access to our websites globally.
- This infrastructure and product support sustain one of the top ten most visited websites in the world, all at a fraction of the cost of popular for-profit websites.

- The global reach of Wikimedia projects is made possible by the hard work of volunteers from across the globe.
- We provide grants, legal support, and other resources to help build vibrant volunteer communities.
- Additionally, we promote community engagement through outreach initiatives and advocate for the growth and protection of free knowledge.

- Operational costs are essential for the smooth management and governance of the Wikimedia Foundation. These expenses help us recruit top talent and support staff around the world, empowering them to carry out the mission of the Wikimedia Foundation.

- Donor support is crucial to sustaining Wikipedia and our other free knowledge endeavors. Our team is committed to efficient and effective fundraising throughout the year, ensuring that every contribution helps advance our mission.

- 2.2 Strong subadditivity

- 2.3 Minimum Shannon entropy

- 2.4 Holevo χ quantity

- 4 Thermodynamic meaning

- 5.1 Conditional entropy

- 5.3 Entanglement measures

- 5.4 Quantum Rényi entropies

- In physics, the von Neumann entropy, named after John von Neumann, is a measure of the statistical uncertainty within a description of a quantum system.
- It extends the concept of Gibbs entropy from classical statistical mechanics to quantum statistical mechanics, and it is the quantum counterpart of the Shannon entropy from classical information theory.
- For a quantum-mechanical system described by a density matrix ρ, the von Neumann entropy is[1] S = − tr ⁡ ( ρ ln ⁡ ρ ) , {\displaystyle S=-\operatorname {tr} (\rho \ln \rho ),} where tr {\displaystyle \operatorname {tr} } denotes the trace and ln {\displaystyle \operatorname {ln} } denotes the matrix version of the natural logarithm.
- If the density matrix ρ is written in a basis of its eigenvectors | 1 ⟩ , | 2 ⟩ , | 3 ⟩ , … {\displaystyle |1\rangle ,|2\rangle ,|3\rangle ,\dots } as ρ = ∑ j η j | j ⟩ ⟨ j | , {\displaystyle \rho =\sum _{j}\eta _{j}\left|j\right\rangle \left\langle j\right|,} then the von Neumann entropy is merely S = − ∑ j η j ln ⁡ η j . {\displaystyle S=-\sum _{j}\eta _{j}\ln \eta _{j}.} In this form, S can be seen as the Shannon entropy of the eigenvalues, reinterpreted as probabilities.[2]

- The von Neumann entropy and quantities based upon it are widely used in the study of quantum entanglement.[3]

- In quantum mechanics, probabilities for the outcomes of experiments made upon a system are calculated from the quantum state describing that system.
- Each physical system is associated with a vector space, or more specifically a Hilbert space.
- The dimension of the Hilbert space may be infinite, as it is for the space of square-integrable functions on a line, which is used to define the quantum physics of a continuous degree of freedom.
- Alternatively, the Hilbert space may be finite-dimensional, as occurs for spin degrees of freedom.
- A density operator, the mathematical representation of a quantum state, is a positive semi-definite, self-adjoint operator of trace one acting on the Hilbert space of the system.[4][5][6] A density operator that is a rank-1 projection is known as a pure quantum state, and all quantum states that are not pure are designated mixed.
- Pure states are also known as wavefunctions.
- Assigning a pure state to a quantum system implies certainty about the outcome of some measurement on that system (i.e., P ( x ) = 1 {\displaystyle P(x)=1} for some outcome x {\displaystyle x} ).
- The state space of a quantum system is the set of all states, pure and mixed, that can be assigned to it.
- For any system, the state space is a convex set: Any mixed state can be written as a convex combination of pure states, though not in a unique way.[7] The von Neumann entropy quantifies the extent to which a state is mixed.[8]

- The prototypical example of a finite-dimensional Hilbert space is a qubit, a quantum system whose Hilbert space is 2-dimensional.
- An arbitrary state for a qubit can be written as a linear combination of the Pauli matrices, which provide a basis for 2 × 2 {\displaystyle 2\times 2} self-adjoint matrices:[9] ρ = 1 2 ( I + r x σ x + r y σ y + r z σ z ) , {\displaystyle \rho ={\tfrac {1}{2}}\left(I+r_{x}\sigma _{x}+r_{y}\sigma _{y}+r_{z}\sigma _{z}\right),} where the real numbers ( r x , r y , r z ) {\displaystyle (r_{x},r_{y},r_{z})} are the coordinates of a point within the unit ball and σ x = ( 0 1 1 0 ) , σ y = ( 0 − i i 0 ) , σ z = ( 1 0 0 − 1 ) . {\displaystyle \sigma _{x}={\begin{pmatrix}0&1\\1&0\end{pmatrix}},\quad \sigma _{y}={\begin{pmatrix}0&-i\\i&0\end{pmatrix}},\quad \sigma _{z}={\begin{pmatrix}1&0\\0&-1\end{pmatrix}}.} The von Neumann entropy vanishes when ρ {\displaystyle \rho } is a pure state, i.e., when the point ( r x , r y , r z ) {\displaystyle (r_{x},r_{y},r_{z})} lies on the surface of the unit ball, and it attains its maximum value when ρ {\displaystyle \rho } is the maximally mixed state, which is given by r x = r y = r z = 0 {\displaystyle r_{x}=r_{y}=r_{z}=0} .[10]

- Some properties of the von Neumann entropy:

- S(ρ) is zero if and only if ρ represents a pure state.[11]

- S(ρ) is maximal and equal to ln ⁡ N {\displaystyle \ln N} for a maximally mixed state, N being the dimension of the Hilbert space.[12]

- S(ρ) is invariant under changes in the basis of ρ, that is, S(ρ) = S(UρU†), with U a unitary transformation.[13]

- S(ρ) is concave, that is, given a collection of positive numbers λi which sum to unity ( Σ i λ i = 1 {\displaystyle \Sigma _{i}\lambda _{i}=1} ) and density operators ρi, we have[14]

- S ( ∑ i = 1 k λ i ρ i ) ≥ ∑ i = 1 k λ i S ( ρ i ) . {\displaystyle S{\bigg (}\sum _{i=1}^{k}\lambda _{i}\rho _{i}{\bigg )}\geq \sum _{i=1}^{k}\lambda _{i}S(\rho _{i}).}

- S(ρ) is additive for independent systems. Given two density matrices ρA , ρB describing independent systems A and B, we have[15]

- S ( ρ A ⊗ ρ B ) = S ( ρ A ) + S ( ρ B ) . {\displaystyle S(\rho _{A}\otimes \rho _{B})=S(\rho _{A})+S(\rho _{B}).}

- S(ρ) is strongly subadditive. That is, for any three systems A, B, and C:[16]

- S ( ρ A B C ) + S ( ρ B ) ≤ S ( ρ A B ) + S ( ρ B C ) . {\displaystyle S(\rho _{ABC})+S(\rho _{B})\leq S(\rho _{AB})+S(\rho _{BC}).}

- S ( ρ A C ) ≤ S ( ρ A ) + S ( ρ C ) . {\displaystyle S(\rho _{AC})\leq S(\rho _{A})+S(\rho _{C}).}

- Below, the concept of subadditivity is discussed, followed by its generalization to strong subadditivity.

- If ρA, ρB are the reduced density matrices of the general state ρAB, then | S ( ρ A ) − S ( ρ B ) | ≤ S ( ρ A B ) ≤ S ( ρ A ) + S ( ρ B ) . {\displaystyle \left|S(\rho _{A})-S(\rho _{B})\right|\leq S(\rho _{AB})\leq S(\rho _{A})+S(\rho _{B}).}

- The right hand inequality is known as subadditivity, and the left is sometimes known as the triangle inequality.[17] While in Shannon's theory the entropy of a composite system can never be lower than the entropy of any of its parts, in quantum theory this is not the case; i.e., it is possible that S(ρAB) = 0, while S(ρA) = S(ρB) > 0.
- This is expressed by saying that the Shannon entropy is monotonic but the von Neumann entropy is not.[18] For example, take the Bell state of two spin-1/2 particles: | ψ ⟩ = | ↑↓ ⟩ + | ↓↑ ⟩ . {\displaystyle \left|\psi \right\rangle =\left|\uparrow \downarrow \right\rangle +\left|\downarrow \uparrow \right\rangle .} This is a pure state with zero entropy, but each spin has maximum entropy when considered individually, because its reduced density matrix is the maximally mixed state.
- This indicates that it is an entangled state;[19] the use of entropy as an entanglement measure is discussed further below.

- The von Neumann entropy is also strongly subadditive.[20] Given three Hilbert spaces, A, B, C, S ( ρ A B C ) + S ( ρ B ) ≤ S ( ρ A B ) + S ( ρ B C ) . {\displaystyle S(\rho _{ABC})+S(\rho _{B})\leq S(\rho _{AB})+S(\rho _{BC}).} By using the proof technique that establishes the left side of the triangle inequality above, one can show that the strong subadditivity inequality is equivalent to the following inequality: S ( ρ A ) + S ( ρ C ) ≤ S ( ρ A B ) + S ( ρ B C ) {\displaystyle S(\rho _{A})+S(\rho _{C})\leq S(\rho _{AB})+S(\rho _{BC})} where ρAB, etc. are the reduced density matrices of a density matrix ρABC.[21] If we apply ordinary subadditivity to the left side of this inequality, we then find S ( ρ A C ) ≤ S ( ρ A B ) + S ( ρ B C ) . {\displaystyle S(\rho _{AC})\leq S(\rho _{AB})+S(\rho _{BC}).} By symmetry, for any tripartite state ρABC, each of the three numbers S(ρAB), S(ρBC), S(ρAC) is less than or equal to the sum of the other two.[22]

- Minimum Shannon entropy

- Given a quantum state and a specification of a quantum measurement, we can calculate the probabilities for the different possible results of that measurement, and thus we can find the Shannon entropy of that probability distribution.
- A quantum measurement can be specified mathematically as a positive operator valued measure, or POVM.[23] In the simplest case, a system with a finite-dimensional Hilbert space and measurement with a finite number of outcomes, a POVM is a set of positive semi-definite matrices { F i } {\displaystyle \{F_{i}\}} on the Hilbert space that sum to the identity matrix,[24] ∑ i = 1 n F i = I . {\displaystyle \sum _{i=1}^{n}F_{i}=\operatorname {I} .} The POVM element F i {\displaystyle F_{i}} is associated with the measurement outcome i {\displaystyle i} , such that the probability of obtaining it when making a measurement on the quantum state ρ {\displaystyle \rho } is given by Prob ( i ) = tr ⁡ ( ρ F i ) . {\displaystyle {\text{Prob}}(i)=\operatorname {tr} (\rho F_{i}).} A POVM is rank-1 if all of the elements are proportional to rank-1 projection operators.
- The von Neumann entropy is the minimum achievable Shannon entropy, where the minimization is taken over all rank-1 POVMs.[25]

- If ρi are density operators and λi is a collection of positive numbers which sum to unity ( Σ i λ i = 1 {\displaystyle \Sigma _{i}\lambda _{i}=1} ), then ρ = ∑ i = 1 k λ i ρ i {\displaystyle \rho =\sum _{i=1}^{k}\lambda _{i}\rho _{i}} is a valid density operator, and the difference between its von Neumann entropy and the weighted average of the entropies of the ρi is bounded by the Shannon entropy of the λi: S ( ∑ i = 1 k λ i ρ i ) − ∑ i = 1 k λ i S ( ρ i ) ≤ − ∑ i = 1 k λ i log ⁡ λ i . {\displaystyle S{\bigg (}\sum _{i=1}^{k}\lambda _{i}\rho _{i}{\bigg )}-\sum _{i=1}^{k}\lambda _{i}S(\rho _{i})\leq -\sum _{i=1}^{k}\lambda _{i}\log \lambda _{i}.} Equality is attained when the supports of the ρi – the spaces spanned by their eigenvectors corresponding to nonzero eigenvalues – are orthogonal.
- The difference on the left-hand side of this inequality is known as the Holevo χ quantity and also appears in Holevo's theorem, an important result in quantum information theory.[26]

- Change under time evolution

- The time evolution of an isolated system is described by a unitary operator: ρ → U ρ U † . {\displaystyle \rho \to U\rho U^{\dagger }.} Unitary evolution takes pure states into pure states,[27] and it leaves the von Neumann entropy unchanged.
- This follows from the fact that the entropy of ρ {\displaystyle \rho } is a function of the eigenvalues of ρ {\displaystyle \rho } .[28]

- A measurement upon a quantum system will generally bring about a change of the quantum state of that system.
- Writing a POVM does not provide the complete information necessary to describe this state-change process.[29] To remedy this, further information is specified by decomposing each POVM element into a product: E i = A i † A i . {\displaystyle E_{i}=A_{i}^{\dagger }A_{i}.} The Kraus operators A i {\displaystyle A_{i}} , named for Karl Kraus, provide a specification of the state-change process.
- They are not necessarily self-adjoint, but the products A i † A i {\displaystyle A_{i}^{\dagger }A_{i}} are.
- If upon performing the measurement the outcome E i {\displaystyle E_{i}} is obtained, then the initial state ρ {\displaystyle \rho } is updated to ρ → ρ ′ = A i ρ A i † P r o b ( i ) = A i ρ A i † tr ⁡ ( ρ E i ) . {\displaystyle \rho \to \rho '={\frac {A_{i}\rho A_{i}^{\dagger }}{\mathrm {Prob} (i)}}={\frac {A_{i}\rho A_{i}^{\dagger }}{\operatorname {tr} (\rho E_{i})}}.} An important special case is the Lüders rule, named for Gerhart Lüders.[30][31] If the POVM elements are projection operators, then the Kraus operators can be taken to be the projectors themselves: ρ → ρ ′ = Π i ρ Π i tr ⁡ ( ρ Π i ) . {\displaystyle \rho \to \rho '={\frac {\Pi _{i}\rho \Pi _{i}}{\operatorname {tr} (\rho \Pi _{i})}}.} If the initial state ρ {\displaystyle \rho } is pure, and the projectors Π i {\displaystyle \Pi _{i}} have rank 1, they can be written as projectors onto the vectors | ψ ⟩ {\displaystyle |\psi \rangle } and | i ⟩ {\displaystyle |i\rangle } , respectively.
- The formula simplifies thus to ρ = | ψ ⟩ ⟨ ψ | → ρ ′ = | i ⟩ ⟨ i | ψ ⟩ ⟨ ψ | i ⟩ ⟨ i | | ⟨ i | ψ ⟩ | 2 = | i ⟩ ⟨ i | . {\displaystyle \rho =|\psi \rangle \langle \psi |\to \rho '={\frac {|i\rangle \langle i|\psi \rangle \langle \psi |i\rangle \langle i|}{|\langle i|\psi \rangle |^{2}}}=|i\rangle \langle i|.} We can define a linear, trace-preserving, completely positive map, by summing over all the possible post-measurement states of a POVM without the normalisation: ρ → ∑ i A i ρ A i † . {\displaystyle \rho \to \sum _{i}A_{i}\rho A_{i}^{\dagger }.} It is an example of a quantum channel,[32] and can be interpreted as expressing how a quantum state changes if a measurement is performed but the result of that measurement is lost.[33] Channels defined by projective measurements can never decrease the von Neumann entropy; they leave the entropy unchanged only if they do not change the density matrix.[34] A quantum channel will increase or leave constant the von Neumann entropy of every input state if and only if the channel is unital, i.e., if it leaves fixed the maximally mixed state.
- An example of a channel that decreases the von Neumann entropy is the amplitude damping channel for a qubit, which sends all mixed states towards a pure state.[35]

- Thermodynamic meaning

- The quantum version of the canonical distribution, the Gibbs states, are found by maximizing the von Neumann entropy under the constraint that the expected value of the Hamiltonian is fixed.
- A Gibbs state is a density operator with the same eigenvectors as the Hamiltonian, and its eigenvalues are λ i = 1 Z exp ⁡ ( − E i k B T ) , {\displaystyle \lambda _{i}={\frac {1}{Z}}\exp \left(-{\frac {E_{i}}{k_{B}T}}\right),} where T is the temperature, k B {\displaystyle k_{B}} is the Boltzmann constant, and Z is the partition function.[36][37] The von Neumann entropy of a Gibbs state is, up to a factor k B {\displaystyle k_{B}} , the thermodynamic entropy.[38]

- Generalizations and derived quantities

- Let ρ A B {\displaystyle \rho _{AB}} be a joint state for the bipartite quantum system AB.
- Then the conditional von Neumann entropy S ( A | B ) {\displaystyle S(A|B)} is the difference between the entropy of ρ A B {\displaystyle \rho _{AB}} and the entropy of the marginal state for subsystem B alone: S ( A | B ) = S ( ρ A B ) − S ( ρ B ) . {\displaystyle S(A|B)=S(\rho _{AB})-S(\rho _{B}).} This is bounded above by S ( ρ A ) {\displaystyle S(\rho _{A})} .
- In other words, conditioning the description of subsystem A upon subsystem B cannot increase the entropy associated with A.[39]

- Quantum mutual information can be defined as the difference between the entropy of the joint state and the total entropy of the marginals: S ( A : B ) = S ( ρ A ) + S ( ρ B ) − S ( ρ A B ) , {\displaystyle S(A:B)=S(\rho _{A})+S(\rho _{B})-S(\rho _{AB}),} which can also be expressed in terms of conditional entropy:[40] S ( A : B ) = S ( A ) − S ( A | B ) = S ( B ) − S ( B | A ) . {\displaystyle S(A:B)=S(A)-S(A|B)=S(B)-S(B|A).}

- Let ρ {\displaystyle \rho } and σ {\displaystyle \sigma } be two density operators in the same state space.
- The relative entropy is defined to be S ( σ | ρ ) = tr ⁡ [ ρ ( log ⁡ ρ − log ⁡ σ ) ] . {\displaystyle S(\sigma |\rho )=\operatorname {tr} [\rho (\log \rho -\log \sigma )].} The relative entropy is always greater than or equal to zero; it equals zero if and only if ρ = σ {\displaystyle \rho =\sigma } .[41] Unlike the von Neumann entropy itself, the relative entropy is monotonic, in that it decreases (or remains constant) when part of a system is traced over:[42] S ( σ A | ρ A ) ≤ S ( σ A B | ρ A B ) . {\displaystyle S(\sigma _{A}|\rho _{A})\leq S(\sigma _{AB}|\rho _{AB}).}

- Entanglement measures

- Just as energy is a resource that facilitates mechanical operations, entanglement is a resource that facilitates performing tasks that involve communication and computation.[43] The mathematical definition of entanglement can be paraphrased as saying that maximal knowledge about the whole of a system does not imply maximal knowledge about the individual parts of that system.[44] If the quantum state that describes a pair of particles is entangled, then the results of measurements upon one half of the pair can be strongly correlated with the results of measurements upon the other.
- However, entanglement is not the same as "correlation" as understood in classical probability theory and in daily life.
- Instead, entanglement can be thought of as potential correlation that can be used to generate actual correlation in an appropriate experiment.[45] The state of a composite system is always expressible as a sum, or superposition, of products of states of local constituents; it is entangled if this sum cannot be written as a single product term.[46] Entropy provides one tool that can be used to quantify entanglement.[47][48] If the overall system is described by a pure state, the entropy of one subsystem can be used to measure its degree of entanglement with the other subsystems.
- For bipartite pure states, the von Neumann entropy of reduced states is the unique measure of entanglement in the sense that it is the only function on the family of states that satisfies certain axioms required of an entanglement measure.[49][50] It is thus known as the entanglement entropy.[51]

- It is a classical result that the Shannon entropy achieves its maximum at, and only at, the uniform probability distribution {1/n, ..., 1/n}.[52] Therefore, a bipartite pure state ρ ∈ HA ⊗ HB is said to be a maximally entangled state if the reduced state of each subsystem of ρ is the diagonal matrix[53] ( 1 n ⋱ 1 n ) . {\displaystyle {\begin{pmatrix}{\frac {1}{n}}&&\\&\ddots &\\&&{\frac {1}{n}}\end{pmatrix}}.}

- For mixed states, the reduced von Neumann entropy is not the only reasonable entanglement measure.[54] Some of the other measures are also entropic in character.
- For example, the relative entropy of entanglement is given by minimizing the relative entropy between a given state ρ {\displaystyle \rho } and the set of nonentangled, or separable, states.[55] The entanglement of formation is defined by minimizing, over all possible ways of writing of ρ {\displaystyle \rho } as a convex combination of pure states, the average entanglement entropy of those pure states.[56] The squashed entanglement is based on the idea of extending a bipartite state ρ A B {\displaystyle \rho _{AB}} to a state describing a larger system, ρ A B E {\displaystyle \rho _{ABE}} , such that the partial trace of ρ A B E {\displaystyle \rho _{ABE}} over E yields ρ A B {\displaystyle \rho _{AB}} .
- One then finds the infimum of the quantity 1 2 [ S ( ρ A E ) + S ( ρ B E ) − S ( ρ E ) − S ( ρ A B E ) ] , {\displaystyle {\frac {1}{2}}[S(\rho _{AE})+S(\rho _{BE})-S(\rho _{E})-S(\rho _{ABE})],} over all possible choices of ρ A B E {\displaystyle \rho _{ABE}} .[57]

- Quantum Rényi entropies

- Just as the Shannon entropy function is one member of the broader family of classical Rényi entropies, so too can the von Neumann entropy be generalized to the quantum Rényi entropies: S α ( ρ ) = 1 1 − α ln ⁡ [ tr ⁡ ρ α ] = 1 1 − α ln ⁡ ∑ i = 1 N λ i α . {\displaystyle S_{\alpha }(\rho )={\frac {1}{1-\alpha }}\ln[\operatorname {tr} \rho ^{\alpha }]={\frac {1}{1-\alpha }}\ln \sum _{i=1}^{N}\lambda _{i}^{\alpha }.} In the limit that α → 1 {\displaystyle \alpha \to 1} , this recovers the von Neumann entropy.
- The quantum Rényi entropies are all additive for product states, and for any α {\displaystyle \alpha } , the Rényi entropy S α {\displaystyle S_{\alpha }} vanishes for pure states and is maximized by the maximally mixed state.
- For any given state ρ {\displaystyle \rho } , S α ( ρ ) {\displaystyle S_{\alpha }(\rho )} is a continuous, nonincreasing function of the parameter α {\displaystyle \alpha } .
- A weak version of subadditivity can be proven: S α ( ρ A ) − S 0 ( ρ B ) ≤ S α ( ρ A B ) ≤ S α ( ρ A ) + S 0 ( ρ B ) . {\displaystyle S_{\alpha }(\rho _{A})-S_{0}(\rho _{B})\leq S_{\alpha }(\rho _{AB})\leq S_{\alpha }(\rho _{A})+S_{0}(\rho _{B}).} Here, S 0 {\displaystyle S_{0}} is the quantum version of the Hartley entropy, i.e., the logarithm of the rank of the density matrix.[58]

- The density matrix was introduced, with different motivations, by von Neumann and by Lev Landau.
- The motivation that inspired Landau was the impossibility of describing a subsystem of a composite quantum system by a state vector.[59] On the other hand, von Neumann introduced the density matrix in order to develop both quantum statistical mechanics and a theory of quantum measurements.[60] He introduced the expression now known as von Neumann entropy by arguing that a probabilistic combination of pure states is analogous to a mixture of ideal gases.[61][62] Von Neumann first published on the topic in 1927.[63] His argument was built upon earlier work by Albert Einstein and Leo Szilard.[64][65][66]

- Max Delbrück and Gert Molière proved the concavity and subadditivity properties of the von Neumann entropy in 1936.
- Quantum relative entropy was introduced by Hisaharu Umegaki in 1962.[67][68] The subadditivity and triangle inequalities were proved in 1970 by Huzihiro Araki and Elliott H.
- Lieb.[69] Strong subadditivity is a more difficult theorem.
- It was conjectured by Oscar Lanford and Derek Robinson in 1968.[70] Lieb and Mary Beth Ruskai proved the theorem in 1973,[71][72] using a matrix inequality proved earlier by Lieb.[73][74]

- Bengtsson, Ingemar; Życzkowski, Karol (2017). Geometry of Quantum States: An Introduction to Quantum Entanglement (2nd ed.). Cambridge University Press. ISBN 978-1-107-02625-4.

- Holevo, Alexander S. (2001). Statistical Structure of Quantum Theory. Lecture Notes in Physics. Monographs. Springer. ISBN 3-540-42082-7.

- Nielsen, Michael A.; Chuang, Isaac L. (2010). Quantum Computation and Quantum Information (10th anniversary ed.). Cambridge: Cambridge Univ. Press. ISBN 978-0-521-63503-5.

- Peres, Asher (1993). Quantum Theory: Concepts and Methods. Kluwer. ISBN 0-7923-2549-4.

- Rau, Jochen (2017). Statistical Physics and Thermodynamics. Oxford University Press. ISBN 978-0-19-959506-8.

- Rau, Jochen (2021). Quantum Theory: An Information Processing Approach. Oxford University Press. ISBN 978-0-19-289630-8.

- Rieffel, Eleanor; Polak, Wolfgang (2011). Quantum Computing: A Gentle Introduction. Scientific and engineering computation. Cambridge, Mass: MIT Press. ISBN 978-0-262-01506-6.

- Wilde, Mark M. (2017). Quantum Information Theory (2nd ed.). Cambridge University Press. arXiv:1106.1445. doi:10.1017/9781316809976. ISBN 978-1-316-80997-6.

- Zwiebach, Barton (2022). Mastering Quantum Mechanics: Essentials, Theory, and Applications. MIT Press. ISBN 978-0-262-04613-8.

- Principle of maximum entropy

- thermodynamic potential: U H F G

- Ferromagnetism models Ising Potts Heisenberg percolation

- Particles with force field depletion force Lennard-Jones potential

- Lennard-Jones potential

- mean-field theory and conformal field theory

- Critical exponents correlation length size scaling

- Statistical field theory elementary particle superfluidity

- Condensed matter physics

- Complex system chaos information theory Boltzmann machine

- Quantum mechanical entropy

- CS1 errors: ISBN date

- Articles with short description

- Short description is different from Wikidata

- Edit preview settings
