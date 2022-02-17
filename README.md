# golang-training
Notes about Udemy course to learn Golang: https://www.udemy.com/course/go-the-complete-developers-guide/


## Selective Checks Mermaid

```mermaid
flowchart TD
A[PR arrives]-->B[Selective Check]
B-->C{Direct push merge?}
C-->|Yes: Enable images| D[Run Full Test, +Quarantined, run full static checks]
C-->|No| E[Retrieve changed files]
E-->F{Environment files changed?}
F-->|Yes: Enable images| D
F-->|No| G{Docs changed}
G-->|Yes: Enable images building| I{Chart files changed?}
G-->|No| I
I-->|Yes: Enable helm tests| J{API files changed}
I-->|No| J
J-->|Yes: Enable API tests| H{Sources changed?}
J-->|No| H
H-->|Yes: Enable Pytest| K{Determine test type}
H-->|No| L[Skip running test, Run subset of static checks]
```

F-->|No| G{Docs changed}
F-->|Yes| -->D
G-->|No|
G-->|Yes|