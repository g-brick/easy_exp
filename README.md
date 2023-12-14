# Easy_exp

**Easy_exp** is a simplified version of regular expression matching rules.
It has a simple syntax tree, and it was originally intended to make it easier for non-IT people
to fill in certain text matching rules, such as those in operations.
You can determine whether the current string exists by easy_exp generating your own matching rules, 
so it only returns a boolean.
It supports the following logical operators:

| Logical Operator | Priority  |
|------------------|-----------|
| "!"              | Highest   |
| "&&"             | Higher    |
| "\|\|"           | Normal    |


In addition to the above rules, easy_exp also supports parentheses to be forced to be promoted to the highest priority. 
```
You can set the parentheses in part of rule to raise priority:
("Go" || "JAVA") && (!PHP || JS)
```
## Getting started
Simply add the following import in your code file,
and then `go [build|run|test]` will automatically fetch the necessary dependencies.
```
import "github.com/g-brick/easy_exp"
```
Otherwise, run the following Go command to install the `easy_exp` package:

```sh
$ go get -u github.com/g-brick/easy_exp
```
## How to use
### 1.Individual text matching
First you need to import easy_exp package for using easy_exp, one example :
```
 Your rule is a single text, such as "John", so your code snippets in golang could be like this:
```

```
        expRule := "John" // If some txt includes "John", it will be true
	fmt.Printf("expRule: ", expRule)
	check, err := Compile(expRule) // It will be build a matching rule
	if err != nil {
		return
	}
	target1 := "My college teacher is Mr.John"
	fmt.Printf("Matched by: ", target1)
	fmt.Printf("True is expected: ", check(target1))
	target2 := "My best friend is Jay"
	fmt.Printf("Matched by: ", target2)
	fmt.Printf("False is expected: ", check(target2))
	target3 := "Yesterday, John got sick."
	fmt.Printf("Matched by: ", target3)
	fmt.Printf("True is expected: ", check(target3))
```
### 2.Common usage
If you need a matching rule to support logical operators, you can set the rule up like this：
```
        expRule := "(China && USA) || (Japan && Korea) || (Asia && North American)"
	fmt.Printf("Expression rule: ", expRule)
	check, err := Compile(expRule)
	if err != nil {
		return
	}
	target1 := "I am from USA, but I want to travel to China."
	fmt.Printf("Matched by: ", target1)
	fmt.Printf("True is expected: ", check(target1))
	target2 := "Japan is not in Africa, and neither is South Korea."
	fmt.Printf("Matched by: ", target2)
	fmt.Printf("True is expected: ", check(target2))
	target3 := "The North American is smaller than Asia."
	fmt.Printf("Matched by: ", target3)
	fmt.Printf("True is expected: ", check(target3))
	target4 := "The USA is a part of the North American."
	fmt.Printf("Matched by: ", target4)
	fmt.Printf("False is expected: ", check(target4))
```
### 3.Other usage
```
        expRule := "golang1.2&&!golang1.1" // if some txt includes the golang1.2 and excludes the golang1.1, it will be true
	fmt.Printf("Expression rule: ", expRule)
	check, err := Compile(expRule)
	if err != nil {
		return
	}
	target1 := "I has made my golang version upgrade from golang1.1 to golang1.2"
	fmt.Printf("Matched by: ", target1)
	fmt.Printf("False is expected: ", check(target1))
	target2 := "I has made my golang version upgrade from golang1.2 to golang1.3"
	fmt.Printf("Matched by: ", target2)
	fmt.Printf("True is expected: ", check(target2))
	target3 := "The version of go in this program is golang1.1"
	fmt.Printf("Matched by: ", target3)
	fmt.Printf("False is expected: ", check(target3))
	target4 := "The version of go in this program is golang1.2"
	fmt.Printf("Matched by: ", target4)
	fmt.Printf("True is expected: ", check(target4))
	target5 := "The version of go in this program is golang1.0"
	fmt.Printf("Matched by: ", target5)
	fmt.Printf("False is expected: ", check(target5))
```
You can check the test file for more details.