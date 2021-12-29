# ccl
A simple calculator to return the final price with the input included percentage or vice versa

## Usage

### To get percentage
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
