# tempus-fugit

Commandline utility to see differences in dates

## Build

```bash
$ go build
```

## Usage

Accepts dates in YYYY-MM-DD format.

```bash
$ tempus-fugit --from '2000-01-01' --to '2010-10-10'
from: 2000-01-01 00:00:00 -0800 PST
  to: 2010-10-10 00:00:00 -0700 PDT
94439h0m0s
339980400 seconds
5666340 minutes
94439 hours
3934 days
562 weeks
129 months
10 years
```

Default `--from` date is 1970-01-01.
Default `--to` date is now.

```bash
$ tempus-fugit --help
Usage of tempus-fugit:
  -from string
        the from input date (default "1970-01-01")
  -to string
        the to input date (default "now")
```
