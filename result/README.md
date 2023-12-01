## hello world
- hyperfine 'curl localhost:3000'
### bun
Benchmark 1: curl localhost:3000
  Time (mean ± σ):       5.6 ms ±   1.0 ms    [User: 3.8 ms, System: 1.0 ms]
  Range (min … max):     4.5 ms …  10.1 ms    312 runs

### tsnode_fastify
Benchmark 1: curl localhost:3000
  Time (mean ± σ):       5.4 ms ±   0.4 ms    [User: 4.0 ms, System: 1.0 ms]
  Range (min … max):     4.8 ms …   8.0 ms    433 runs

### go
Benchmark 1: curl localhost:3000
  Time (mean ± σ):       5.2 ms ±   0.4 ms    [User: 4.0 ms, System: 1.0 ms]
  Range (min … max):     4.7 ms …   7.8 ms    495 runs

## calc (10m)
- (ループ回数を10000000に)
- hyperfine 'curl localhost:3000/calc'
### bun
Benchmark 1: curl localhost:3000/calc
  Time (mean ± σ):      13.3 ms ±   1.3 ms    [User: 3.9 ms, System: 1.0 ms]
  Range (min … max):    11.5 ms …  20.3 ms    143 runs
 
### tsnode_fastify
Benchmark 1: curl localhost:3000/calc
  Time (mean ± σ):      10.2 ms ±   1.3 ms    [User: 4.2 ms, System: 0.8 ms]
  Range (min … max):     9.4 ms …  18.7 ms    156 runs

### go
Benchmark 1: curl localhost:3000/calc
  Time (mean ± σ):       7.3 ms ±   0.6 ms    [User: 3.8 ms, System: 1.0 ms]
  Range (min … max):     6.8 ms …  15.0 ms    369 runs

## calc (100m)
- (ループ回数を100000000に)
- hyperfine 'curl localhost:3000/calc'
### bun
Benchmark 1: curl localhost:3000/calc
  Time (mean ± σ):      76.2 ms ±   6.5 ms    [User: 4.5 ms, System: 0.8 ms]
  Range (min … max):    72.8 ms … 101.1 ms    29 runs
 
### tsnode_fastify
Benchmark 1: curl localhost:3000/calc
  Time (mean ± σ):      52.5 ms ±   6.8 ms    [User: 4.5 ms, System: 0.8 ms]
  Range (min … max):    49.7 ms …  83.7 ms    35 runs
 
### go
Benchmark 1: curl localhost:3000/calc
  Time (mean ± σ):      28.1 ms ±   0.7 ms    [User: 4.1 ms, System: 1.0 ms]
  Range (min … max):    27.1 ms …  30.8 ms    101 runs
