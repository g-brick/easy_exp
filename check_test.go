package easy_exp

import (
	"testing"
)

func TestCompileNormalTxt(t *testing.T) {
	expRule := "John" // If some txt includes John, it will be true
	t.Log("expRule: ", expRule)
	check, err := Compile(expRule)
	if err != nil {
		t.Error(err)
		return
	}
	target1 := "My college teacher is Mr.John"
	t.Log("Matched by: ", target1)
	t.Log("True is expected: ", check(target1))
	target2 := "My best friend is Jay"
	t.Log("Matched by: ", target2)
	t.Log("False is expected: ", check(target2))
	target3 := "Yesterday, John got sick."
	t.Log("Matched by: ", target3)
	t.Log("True is expected: ", check(target3))
}

func TestCompileNotRule(t *testing.T) {
	expRule := "!John" // If some txt includes John, it will be false
	t.Log("expRule: ", expRule)
	check, err := Compile(expRule)
	if err != nil {
		t.Error(err)
		return
	}
	target1 := "My college teacher is Mr.John"
	t.Log("Matched by: ", target1)
	t.Log("False is expected: ", check(target1))
	target2 := "My best friend is Jay"
	t.Log("Matched by: ", target2)
	t.Log("True is expected: ", check(target2))
	target3 := "Yesterday, John and Joy got sick."
	t.Log("Matched by: ", target3)
	t.Log("False is expected: ", check(target3))
}

func TestCompileAndRule(t *testing.T) {
	expRule := "\"golang1.1\"&&\"golang1.2\""
	t.Log("Expression rule: ", expRule)
	check, err := Compile(expRule)
	if err != nil {
		t.Error(err)
		return
	}
	target1 := "I has made my golang version upgrade from golang1.1 to golang1.2"
	t.Log("Matched by: ", target1)
	t.Log("True is expected: ", check(target1))
	target2 := "I has made my golang version downgrade from golang1.2 to golang1.1"
	t.Log("Matched by: ", target2)
	t.Log("True is expected: ", check(target2))
	target3 := "The version of go in this program is golang1.1" // No golang1.2 in the txt, so result is false
	t.Log("Matched by: ", target3)
	t.Log("False is expected: ", check(target3))
}

func TestCompileOrRule(t *testing.T) {
	expRule := "\"golang1.1\"||\"golang1.2\""
	t.Log("Expression rule: ", expRule)
	check, err := Compile(expRule)
	if err != nil {
		t.Error(err)
		return
	}
	target1 := "I has made my golang version upgrade from golang1.1 to golang1.2"
	t.Log("Matched by: ", target1)
	t.Log("True is expected: ", check(target1))
	target2 := "I has made my golang version downgrade from golang1.2 to golang1.1"
	t.Log("Matched by: ", target2)
	t.Log("True is expected: ", check(target2))
	target3 := "The version of go in this program is golang1.1"
	t.Log("Matched by: ", target3)
	t.Log("True is expected: ", check(target3))
	target4 := "The version of go in this program is golang1.2"
	t.Log("Matched by: ", target4)
	t.Log("True is expected: ", check(target4))
	target5 := "The version of go in this program is golang1.0" // Not matched the rule:  golang1.1 || golang1.2
	t.Log("Matched by: ", target5)
	t.Log("False is expected: ", check(target5))
}

func TestCompileMixedRule1(t *testing.T) {
	expRule := "golang1.2&&!golang1.1" // if some txt includes the golang1.2 and excludes the golang1.1, it will be true
	t.Log("Expression rule: ", expRule)

	check, err := Compile(expRule)
	if err != nil {
		t.Error(err)
		return
	}
	target1 := "I has made my golang version upgrade from golang1.1 to golang1.2"
	t.Log("Matched by: ", target1)
	t.Log("False is expected: ", check(target1))
	target2 := "I has made my golang version upgrade from golang1.2 to golang1.3"
	t.Log("Matched by: ", target2)
	t.Log("True is expected: ", check(target2))
	target3 := "The version of go in this program is golang1.1"
	t.Log("Matched by: ", target3)
	t.Log("False is expected: ", check(target3))
	target4 := "The version of go in this program is golang1.2"
	t.Log("Matched by: ", target4)
	t.Log("True is expected: ", check(target4))
	target5 := "The version of go in this program is golang1.0"
	t.Log("Matched by: ", target5)
	t.Log("False is expected: ", check(target5))
}

func TestCompileMixedRule2(t *testing.T) {
	expRule := "(\"China\"&&\"USA\")||(\"Japan\"&&\"Korea\")||(\"Asia\"&&\"North American\")"
	// expRule equals expRule2
	// expRule2 := "(China && USA) || (Japan && Korea) || (Asia && North American)"
	t.Log("Expression rule: ", expRule)
	check, err := Compile(expRule)
	if err != nil {
		t.Error(err)
		return
	}
	target1 := "I am from USA, but I love China."
	t.Log("Matched by: ", target1)
	t.Log("True is expected: ", check(target1))
	target2 := "Japan is not in Africa, and neither is South Korea."
	t.Log("Matched by: ", target2)
	t.Log("True is expected: ", check(target2))
	target3 := "The North American is smaller than Asia."
	t.Log("Matched by: ", target3)
	t.Log("True is expected: ", check(target3))
	target4 := "The USA is a part of the North American."
	t.Log("Matched by: ", target4)
	t.Log("False is expected: ", check(target4))
}

func TestCompileMixedRule3(t *testing.T) {
	expRule := "\"China\"&&\"USA\"||\"Japan\"&&\"Korea\"" // If no parenthesis in the expression, it also works.
	// expRule equals expRule2
	// expRule2 := "China && USA || Japan && Korea
	t.Log("Expression rule: ", expRule)
	check, err := Compile(expRule)
	if err != nil {
		t.Error(err)
		return
	}
	target1 := "I am from USA, but I love China."
	t.Log("Matched by: ", target1)
	t.Log("True is expected: ", check(target1))
	target2 := "I am from China, and I want to travel to Japan or Korea."
	t.Log("Matched by: ", target2)
	t.Log("True is expected: ", check(target2))
	target3 := "I am from USA, and I want to travel to Japan."
	t.Log("Matched by: ", target3)
	t.Log("False is expected: ", check(target3))
	target4 := "I am from USA, and I want to travel to Korea."
	t.Log("Matched by: ", target4)
	t.Log("False is expected: ", check(target4))
}

func TestCompileMixedRule4(t *testing.T) {
	expRule := "(!\"JAVA\"||\"Golang\")&&\"PHP\""
	// expRule equals expRule2
	// expRule2 := "(!JAVA || Golang) && PHP"
	t.Log("Expression rule: ", expRule)
	check, err := Compile(expRule)
	if err != nil {
		t.Error(err)
		return
	}
	target1 := "I am good at JAVA, not PHP."
	t.Log("Matched by: ", target1)
	t.Log("False is expected: ", check(target1))
	target2 := "I am good at Golang, not PHP."
	t.Log("Matched by: ", target2)
	t.Log("True is expected: ", check(target2))
	target3 := "I am good at PHP." // Match !JAVA && PHP
	t.Log("Matched by: ", target3)
	t.Log("True is expected: ", check(target3))
	target4 := "I am good at Golang."
	t.Log("Matched by: ", target4)
	t.Log("False is expected: ", check(target4))
}
