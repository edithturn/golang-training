```{mermaid}
flowchart TD
A(PR arrives)-->B[Selective Check]
B-->C{Direct push merge?}
C-->|Yes| N[Enable images]
N-->D(Run Full Test<br />+Quarantined<br />Run full static checks)
C-->|No| E[Retrieve changed files]
E-->F{Environment files changed?}
F-->|Yes| N
F-->|No| G{Docs changed}
G-->|Yes| O[Enable images building]
O-->I{Chart files changed?}
G-->|No| I
I-->|Yes| P[Enable helm tests]
P-->J{API files changed}
I-->|No| J
J-->|Yes| Q[Enable API tests]
Q-->H{Sources changed?}
J-->|No| H
H-->|Yes| R[Enable Pytests]
R-->K[Determine test type]
K-->S{Core files changed}
S-->|Yes| N
S-->|No| M(Run selected test+<br />Heisentest, Integration, Quarantined<br />Full static checks)
H-->|No| L[Skip running test<br />Run subset of static checks]
```
