# jscollect

A simple tool to collect all Javascript files from a domain.

### Installation
```
go get -u github.com/anthdm/jscollect
```

### Usage

```
jscollect -d https://target.com
```

Pipe the response into other tools
```
jscollect -d https://google.com | wc -l
jscollect -d https://google.com | httprobe
jscollect -d https://google.com > js.txt
```