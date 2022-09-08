Go Quiz 1

The program probably doesn't compile, the interface I made for sort.Sort() when used for an array was throwing errors

Once I fix it, the way to run the program will be `go run GoQuiz1.go {x}` where x is the input for problems 1 and 2

Regardless, for question 3, the O(x) for sort.Sort is O(n*log(n)), and the O(x) for sort.Stable is O(n * n*log(n) * n*log(n)*log(n))

<b>CORRECTIONS BELOW</b>

Additional corrections have been made. The corrections directory has been deleted, and GoQuiz1.go is now the most current version.

<b>Instructions to run</b>

    - `go run GoQuiz1.go {x}` 

where x is the integer input for the length of arrays, slices, and maps


<b>Expected outcome</b>

    - `Input from cmd line is 1000`
    - `Initializing an array of len 1000 took 83ns`
    - `Initializing an array of len 100 took 41ns`
    - `Initializing an array of len 10 took 42ns`
    - `Initializing an slice of len 1000 took 708ns`
    - `Initializing an map of len 0 took 83ns`
    - `The map is dynamic, and thus has not been assigned a length yet`
    - `incrementing an array by 1 took 42ns`
    - `incrementing a slice by 1 took 42ns`
    - `incrementing a map by 1 took 83ns`
    - `sorting a slice of length 1000 using sort.Sort took 66459ns`
    - `sorting a slice using sort.Stable took 4292ns`
    - `sorting an array of length 1000 using sort.Sort took 65667ns`
    - `sorting an array of length 1000 using sort.Stable took 3750ns`
    - `sorting an array of length 100 using sort.Sort took 4458ns`
    - `sorting an array of length 100 using sort.Stable took 583ns`


The Big-O notation for Sort and Stable from the go documentation seems to be accurate. 


Sorting using sort.Sort() results in a longer time the larger the slices get, approximately doubling every time x grows an order of magnitude, concurring with the assumption that sort.Sort is O(log(n)) for slices. For arrays, we must first convert the array to a slice to use the sort interface, so we add an additional step n that increases the time needed to sort. After conversion the sort.Sort() time is the same as slices.

When sorting slices (and therefore arrays after conversion), sort.Stable quadruples when inputs grow by an order of magnitude. These inputs align with the expected `O(n*n*log(n)*n*log(n)*log(n)) == O(3n * 3log(n))`
