from os import system


for D in range(4, 10):
    SIZE = 2 ** (D * 2)
    system(f"mkdir data{D}")
    for i in range(18, 24):
        for j in range(100):
            s_rule = f"{i - 1},{i}"
            system(f"./gameoflife -g 10 -b {i} -s {s_rule} -l {D} -o data{D}/result{j}for{i}_{s_rule}.life -d {D} -size {SIZE}")

