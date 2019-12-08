Human readable byte count
=========================

This package provides two functions to convert a byte count value into a 
human readable byte count value. A human readable byte count is a value 
with a unit rounded away from zero with the first decimal digit. 

The SI() function converting byte counts using the international system units :

|  count  |    result    |
|---------|--------------|
| 0       | "0 B"        |
| 100     | "100 B"      |
| 999     | "999 B"      |
| 1000    | "1.0 kB"     |
| 999949  | "999.9 kB"   |
| 999950  | "1.0 MB"     |
| 1000000 | "1.0 MB"     |

The Bin() function converts byte counts using the 
[binary prefix](https://en.wikipedia.org/wiki/Binary_prefix) units, 
sometimes also referred as the gibi units :

|   count   |    result    |
|-----------|--------------|
| 0         | "0 B"        |
| 100       | "100 B"      |
| 1023      | "1023 B"     |
| 1024      | "1.0 kiB"    |
| 1048524   | "1023.9 kiB" |
| 1048525   | "1.0 MiB"    |



Performance
-----------

The computation is implemented to be efficient in term of CPU and memory allocation.
The SI() and Bin() functions are compared with the straightforward functions SI1() 
and Bin1() using Printf and float operations.

Here are the benchmark results:

|  Function       |     N    |   ns/op  |  B/op  |  allocs/op  |
|---------------: |--------: |--------: |------: |-----------: |
| BenchmarkSI-4   |  357614  |    3286  |   800  |        100  |
| BenchmarkSI1-4  |   30014  |   39577  |  1600  |        200  |
| BenchmarkBin-4  |  375980  |    3097  |  1280  |         80  |
| BenchmarkBin1-4 |   36217  |   32936  |  1920  |        160  |


We can see that SI1() and Bin1() are 12 times slower and allocate two times more blocks than 
the functions SI() and Bin() provided in this package. 