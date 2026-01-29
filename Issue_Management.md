# Issue Management

## Hierarchy Structure

Gitopedia uses GitHub Issues to organize research topics in a four-level hierarchy:

```
Encyclopedia Index
└── Research Domain
    └── Research Category (label: "research category")
        └── Research Topic (label: "research topic")
            └── Articles (task list checkboxes)
```

### Issue Types
- **Encyclopedia Index** — Root issue linking all domains
- **Research Domain** — Top-level knowledge domains
- **Research Category** — Subject areas within a domain
- **Research Topic** — Specific topics containing article lists

### Issue Labels
- `research category` — Applied to category-level issues
- `research topic` — Applied to topic-level issues containing article lists

## Issue Titles
- Issue titles follow pattern: `Domain > Category > Topic`
  - Example: `Science > Physics > Quantum Mechanics`

### Parent-Child Linking
Issues link to parents via markdown in the body:
```
**Parent:** [Parent Name](https://github.com/gitopedia/gitopedia/issues/123)
```
Or equivalently: `**Domain:**`, `**Category:**`

### Article Task Lists
Topic issues contain articles as markdown checkboxes:
```
- [ ] Article Name        (pending)
- [x] Completed Article   (done)
```

The researcher agent:
1. Picks unassigned topic issues with unchecked articles
2. Claims via GitHub assignee (distributed locking)
3. Creates/improves articles, checks them off
4. Creates PR and adds `pending review` label



### Image Style Configuration

Domains and categories have artistic styles defined in `researcher/config/artistic_styles.yaml`. The style hierarchy:

```
default           → fallback styles
└── domain        → domain-level styles (e.g., science)
    └── category  → category-level styles (e.g., physics)
```

When adding new domains/categories, add corresponding entries to the style config.

---

## Domains (Priority Order)

Status: `[ ]` not created, `[x]` created

- [x] Science
- [x] Mathematics
- [x] Philosophy
- [x] History
- [x] Technology
- [x] Engineering
- [x] Medicine
- [x] Computer Science
- [ ] Economics
- [ ] Psychology
- [ ] Sociology
- [ ] Political Science
- [ ] Anthropology
- [ ] Literature
- [ ] Art
- [ ] Music
- [ ] Linguistics
- [ ] Religion
- [ ] Geography
- [ ] Law
- [ ] Business
- [ ] Education
- [ ] Agriculture
- [ ] Sports
- [ ] Military
- [ ] Media & Communication
- [ ] People

---

## Current Structure

- **Encyclopedia Root** — #126

### Science (Domain) — #127
- Physics (#124)
  - Quantum Mechanics (#121)
  - Thermodynamics (#122)
- Biology (#140)
  - Genetics (#143)
  - Cell Biology (#144)
  - Ecology (#145)
- Chemistry (#141)
  - Organic Chemistry (#146)
  - Inorganic Chemistry (#147)
  - Physical Chemistry (#148)
- Astronomy (#142)
  - Stellar Astrophysics (#149)
  - Planetary Science (#150)
  - Cosmology (#151)
- Geography (#125)
  - Plate Tectonics (#123)

### Mathematics (Domain) — #152
- Algebra (#153)
  - Linear Algebra (#158)
  - Abstract Algebra (#159)
- Geometry (#154)
  - Euclidean Geometry (#160)
  - Differential Geometry (#161)
  - Topology (#162)
- Calculus (#155)
  - Differential Calculus (#164)
  - Integral Calculus (#165)
  - Multivariable Calculus (#166)
- Statistics (#156)
  - Probability Theory (#167)
  - Statistical Inference (#168)
- Number Theory (#157)
  - Elementary Number Theory (#169)
  - Analytic Number Theory (#170)

### Philosophy (Domain) — #172
- Ethics (#178)
  - Moral Philosophy (#202)
  - Bioethics (#203)
- Metaphysics (#179)
  - Ontology (#204)
  - Philosophy of Mind (#205)
- Epistemology (#180)
  - Theory of Knowledge (#206)
- Logic (#181)
  - Formal Logic (#207)

### History (Domain) — #173
- Ancient History (#182)
  - Ancient Civilizations (#208)
- Medieval History (#183)
  - Medieval Europe (#209)
- Modern History (#184)
  - Industrial Revolution (#210)
- World History (#185)
  - World Wars (#211)

### Technology (Domain) — #174
- Computing (#186)
  - Hardware (#212)
- Electronics (#187)
  - Semiconductors (#213)
- Energy (#188)
  - Renewable Energy (#214)
- Transportation (#189)
  - Automotive Technology (#215)

### Engineering (Domain) — #175
- Civil Engineering (#190)
  - Structural Engineering (#216)
- Mechanical Engineering (#191)
  - Thermodynamics (#217)
- Electrical Engineering (#192)
  - Power Systems (#218)
- Chemical Engineering (#193)
  - Process Engineering (#219)

### Medicine (Domain) — #176
- Anatomy (#194)
  - Human Body Systems (#220)
- Pharmacology (#195)
  - Drug Classes (#221)
- Pathology (#196)
  - Disease Mechanisms (#222)
- Public Health (#197)
  - Epidemiology (#223)

### Computer Science (Domain) — #177
- Algorithms (#198)
  - Algorithm Design (#224)
- Artificial Intelligence (#199)
  - Machine Learning (#225)
- Networking (#200)
  - Network Protocols (#226)
- Databases (#201)
  - Database Systems (#227)

---
