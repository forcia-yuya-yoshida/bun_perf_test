# hello_world
- hyperfine 'curl localhost:3000'

# calc_10m
- (ループ回数を10000000に)
- hyperfine 'curl localhost:3000/calc'

# calc_100m
- (ループ回数を100000000に)
- hyperfine 'curl localhost:3000/calc'
