# MATH-SKILLS

This project is a simple statistical calculator that reads a dataset from a file and calculates the following statistics:


1. **Calculate Average**: Sums all the values and divides by the number of values.
2. **Calculate Median**: Sorts the values and finds the middle value.
3. **Calculate Variance**: Computes the average of the squared differences from the Mean.
4. **Calculate Standard Deviation**: Takes the square root of the variance.


## Installation

To clone and navigate this program to your local machine, follow the steps below :

```
git clone https://learn.zone01kisumu.ke/git/bernaotieno/math-skills.git
cd math-skills
```

## Usage

If you want to run the program, do the following:

1. Go to the data.txt and write the number you want the operations to be done with.

example:
```
12
23
45
67
4
```

2. go run the program now with the command below.
```
go run . data.txt
```
### Output
The output will be as follow:

```
Average: 30
Median: 23
Variance: 529
Standard Deviation: 23
```

## Testing
Run the test in your program by following these commands:
1. Move into the folder that contains the test:
```
cd calculation/
```
2. Run the test:
```
go test -v
```


## Contributing

If you would like to contribute to this project, please fork the repository and submit a pull request.

## Author

[Bernad Okumu](https://github.com/bernotieno)

