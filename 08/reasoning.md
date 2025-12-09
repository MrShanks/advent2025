# Day 08

## Merge two circuits

    Size:6       Size:2
       A            G
      /\           /
     B  C         H
         \
          D
         / \
        E   F

        A   B   C   D   E   F   G   H
id:     1   2   3   4   5   6   7   8
parent: 1   1   1   3   4   4   7   7
size:   6   1   4   3   1   1   2   1

---

## Result

    Size: 8
       A
      /|\
     B C G
        \ \
        D  H
       / \
      E   F  

        A   B   C   D   E   F   G   H
id:     1   2   3   4   5   6   7   8
parent: 1   1   1   3   4   4   1   7
size:   8   1   4   3   1   1   2   1
