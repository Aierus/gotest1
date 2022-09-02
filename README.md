Go Quiz 1

The program probably doesn't compile, the interface I made for sort.Sort() when used for an array was throwing errors

Once I fix it, the way to run the program will be `go run GoQuiz1.go {x}` where x is the input for problems 1 and 2

Regardless, for question 3, the O(x) for sort.Sort is O(n*log(n)), and the O(x) for sort.Stable is O(n * n*log(n) * n*log(n)*log(n))

<b>CORRECTIONS BELOW</b>

Instructions to run
    - `go run corrections/corrections.go {x}` where x is the integer input for the length of arrays, slices, and maps

Expected outcome

    - `Input from cmd line is 100`

    - `Initializing an array of len 1000 took 84ns miliseconds`

    - `Initializing an slice of len 100 took 708ns miliseconds`

    - `Initializing an map of len 0 took 84ns miliseconds`

    - `incrementing an array by 1 took 42 miliseconds`

    - `incrementing a slice by 1 took 42 miliseconds`

    - `sorting a slice using sort.Sort took 5042 miliseconds`

    - `sorting a slice using sort.Stable took 875 miliseconds`

    - `sorting an array using sort.Sort took 25875 miliseconds`

    - `sorting an array using sort.Stable took 30250 miliseconds`


The Big-O notation for Sort and Stable from the go documentation seems to be accurate. 
    - Sorting using sort.Sort() results in a longer time the larger the slices get, approximately doubling every time x grows an order of magnitude, concurring with the assumption that sort.Sort is O(log(n)) for slices. For arrays, sort.Sort() actually sorts faster when the arrays are larger (this may be a bug). 
    - Sorting using sort.Stable() results in a slower rate of growth when sorting an array, staying the nearly same when inputs grow by an order of magnitude. When sorting slices, sort.Stable quadruples when inputs grow by an order of magnitude. These inputs align with the expected O(n*n*log(n)*n*log(n)*log(n)) == O(3n * 3log(n))
