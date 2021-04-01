# Game Of Life
Game of Life in N-dimension world in Moore's neighborhood 1 degree
## Usage
```sh
Usage of ./gameoflifemodel:
  -a, --attempt int       Count attempts (default 100)
  -b, --b-rule string     Rules for birth (default "5")
  -g, --count int         count generations. (default 100)
  -d, --dimension int     dimension of world (default 3)
  -h, --help              Show help message
  -3, --model3d           Use 3D model
  -o, --out string        Database name (default "output.db")
  -p, --probability int   probability in % (default 50)
  -s, --s-rule string     Rules for save (default "4,5")
  -S, --size int          side size (default 128)
Use "," to split different numbers on rule.
Use "{start}.{end}" to set range [start, end] (end and start includes)
```
