#!python
from os import system


SIZE = 256
for D in range(4, 7):
    print(f"Start {D}")
    system(f"mkdir data{D}")
    B = 3 ** D - 1 // 3
    system(f"mkdir data{D}")
    for b in range(-D, D + 1):
        b_rule = B + b
        system(f"mkdir data{D}/testForRule{B}")
        s = [b]
        for k in range(D - 1):
            if len(s) % 2 == 0:
                s.append(s[i] + 1)
            else:
                s.insert(0, s[0] - 1)
        b_rule = str(b_rule)
        s_rule = ",".join(s)
        for _ in range(100):
            system(f"./gameoflife -g 10 -b {b_rule} -s {s_rule} -l {D} -o data{D}/testForRule{B}/result{_}.life -d {D} -S {SIZE}")
    print(f"End {D}")

