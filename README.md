# sypwgen
Password generator by syllables.

Perhaps someone will remember this password more easily.

build on 'go version go1.23.3 linux/amd64'

Usage of ./sypwgen:
```
  -c    Colored syllable
  -l int
        Password length (default 10)
  -n int
        Count of password (default 1)
  -u    Upper case
```

Examples:
```
./sypwgen -l 10
./sypwgen -l 10 -u
./sypwgen -l 10 -n 3 -u
./sypwgen -l 5 -c -n 3 -u
```