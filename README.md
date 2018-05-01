
# [Supported]

 - [x] windows 
 - [ ] linux

# [Usage]

```
1.cronPush will watch the filepath
2.if the filepath has some change, auto run the git command 'git add . && git commit && git push'

input the path you want to watch as the flag

Usage:
  cronPush [flags]

Flags:
  -h, --help            help for cronPush
  -p, --path string     input the path you want to watch as the flag
  -c, --pushCycle int   git push once time each %n seconds; the default is 5s (default 5)
```

# [Tips]

the best way: the filepath should not be the git repo.

For example:
 
 `/the/path/to/a`: this is the git repo, should not watch tish filepath

 `/the/path/to/a/b`: should watch this filepath


