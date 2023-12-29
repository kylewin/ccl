# ccl
A simple calculator in Golang and Cobra to return the final price

![Release](https://github.com/kyledakid/ccl/actions/workflows/release.yml/badge.svg)
![CI](https://github.com/kyledakid/ccl/actions/workflows/ci.yml/badge.svg)

## Installation
```
go get -u github.com/kyledakid/ccl
or
go install github.com/kyledakid/ccl
or download binaries in Github Release page
```

## Usage

### To get Gold price in VND/L
```
ccl au -d [usd] -v [1 usd in vnd] <- if v is not toggled, use default of 24,275 VND

ccl au  au -d 2069.26
1USD: 24,275.00 VND
1L: 66,560,892.01 VND

ccl au -d 2069.26 -v 24680
1USD: 24,680.00 VND
1L: 67,671,382.69 VND
```

### To get gain/loss percentage
```
ccl pc -i [entry price] -o [exit price]

ccl pc -i 6 -o 8
Entry: 6.00
Exit: 8.00
Change: 33.33%
```

### To get price
```
ccl tg -i [entry price] -d [delta]

ccl tg -i 6 -d 8
Entry: 6.00
Delta: 8.00%
Exit: 6.48

ccl tg -i 6 -d -8
Entry: 6.00
Delta: -8.00%
Exit: 5.52
```
