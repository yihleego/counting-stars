# Counting Stars

A utility for querying user and organization repositories.

## Install

```shell
go install github.com/yihleego/counting-stars
```

## Usage

### User Repositories

```go
counting-stars -u octocat
```

```text
================================================================================
Spoon-Knife         HTML            Stars 11010     Forks 126302    Watchers 11010     
Hello-World                         Stars 1856      Forks 1734      Watchers 1856
octocat.github.io   CSS             Stars 154       Forks 169       Watchers 154
hello-worId                         Stars 127       Forks 101       Watchers 127
git-consortium                      Stars 93        Forks 62        Watchers 93
test-repo1                          Stars 54        Forks 12        Watchers 54
================================================================================
Total               Repos 6         Stars 13294     Forks 128380    Watchers 13294
```

### Organization Repositories

```go
counting-stars -o golang
```

```text
================================================================================
go                  Go              Stars 98579     Forks 14697     Watchers 98579     
dep                 Go              Stars 13057     Forks 1102      Watchers 13057     
groupcache          Go              Stars 11360     Forks 1279      Watchers 11360     
protobuf            Go              Stars 8406      Forks 1524      Watchers 8406      
mock                Go              Stars 7110      Forks 514       Watchers 7110      
tools               Go              Stars 6069      Forks 1996      Watchers 6069
mobile              Go              Stars 4957      Forks 651       Watchers 4957
oauth2              Go              Stars 4142      Forks 839       Watchers 4142
lint                Go              Stars 3946      Forks 528       Watchers 3946
glog                Go              Stars 3160      Forks 862       Watchers 3160
proposal            Go              Stars 2882      Forks 382       Watchers 2882
vscode-go           TypeScript      Stars 2576      Forks 493       Watchers 2576
net                 Go              Stars 2540      Forks 1034      Watchers 2540
crypto              Go              Stars 2408      Forks 1296      Watchers 2408
example             Go              Stars 2125      Forks 697       Watchers 2125
vgo                 Go              Stars 1532      Forks 65        Watchers 1532
tour                Go              Stars 1428      Forks 514       Watchers 1428
geo                 Go              Stars 1364      Forks 157       Watchers 1364
snappy              Go              Stars 1262      Forks 153       Watchers 1262
gddo                Go              Stars 1098      Forks 293       Watchers 1098
leveldb             Go              Stars 1050      Forks 141       Watchers 1050
sys                 Go              Stars 1007      Forks 507       Watchers 1007
pkgsite             Go              Stars 899       Forks 126       Watchers 899
talks                               Stars 697       Forks 97        Watchers 697
sync                Go              Stars 674       Forks 125       Watchers 674
exp                 Go              Stars 671       Forks 162       Watchers 671
text                Go              Stars 661       Forks 268       Watchers 661
freetype            Go              Stars 649       Forks 160       Watchers 649
gofrontend          Go              Stars 648       Forks 109       Watchers 648
appengine           Go              Stars 611       Forks 196       Watchers 611
playground          Go              Stars 593       Forks 171       Watchers 593
build               Go              Stars 512       Forks 131       Watchers 512
image               Go              Stars 463       Forks 148       Watchers 463
blog                                Stars 389       Forks 144       Watchers 389
vulndb              Go              Stars 356       Forks 32        Watchers 356
time                Go              Stars 335       Forks 113       Watchers 335
sublime-build       Python          Stars 322       Forks 47        Watchers 322
perf                Go              Stars 302       Forks 47        Watchers 302
website             Go              Stars 252       Forks 216       Watchers 252
xerrors             Go              Stars 250       Forks 48        Watchers 250
debug               Go              Stars 189       Forks 42        Watchers 189
term                Go              Stars 168       Forks 41        Watchers 168
mod                 Go              Stars 138       Forks 57        Watchers 138
review              Go              Stars 137       Forks 42        Watchers 137
cwg                                 Stars 131       Forks 25        Watchers 131
arch                Go              Stars 116       Forks 40        Watchers 116
benchmarks          Go              Stars 116       Forks 36        Watchers 116
dl                  Go              Stars 112       Forks 27        Watchers 112
sublime-config      Python          Stars 84        Forks 22        Watchers 84
winstrap            Go              Stars 48        Forks 24        Watchers 48
vuln                Go              Stars 41        Forks 10        Watchers 41
scratch             Go              Stars 25        Forks 26        Watchers 25
go-get-issue-15410  Go              Stars 2         Forks 6         Watchers 2
================================================================================
Total               Repos 53        Stars 192649    Forks 32462     Watchers 192649
```

## License

This project is under the Apache 2.0 license. See the [LICENSE](LICENSE) file for details.
