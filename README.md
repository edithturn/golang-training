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
```

E-->F{Environment files changed?}
F-->|No| G{Docs changed}
F-->|Yes| -->D
G-->|No|
G-->|Yes|