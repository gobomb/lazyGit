
# [Supported]

 - [x] windows 
 - [x] macos 
 - [x] linux

# [Usage]

```
1.lazyGit will watch the filepath
2.if the filepath has some change, auto run the git command 'git add . && git commit && git push'

input the path you want to watch as the flag

Usage:
  lazyGit [flags]

Flags:
  -h, --help            help for lazyGit
  -p, --path string     input the path you want to watch as the flag
  -c, --pushCycle int   git push once time each %n seconds; the default is 5s (default 5)
```

# [Tips]

The best way: the filepath should be the git repo.

And the .git directory will be ignored.


