package main

import "testing"

func TestAdd(t *testing.T) {
	expected := 6
	expression := "1 + 5"
	formedExp := formExpression(expression)
	result := int(count(formedExp))
	if result != expected {
		t.Errorf("Test add failed, %v %v", expected, result)
	}

	expected = 13
	expression = "(1 + 5) + 7"
	formedExp = formExpression(expression)
	result = int(count(formedExp))
	if result != expected {
		t.Errorf("Test add failed, %v %v", expected, result)
	}
}

func TestMinus(t *testing.T) {
	expected := -4
	expression := "9 - 13"
	formedExp := formExpression(expression)
	result := int(count(formedExp))
	if result != expected {
		t.Errorf("Test add failed, %v %v", expected, result)
	}

	expected = -43
	expression = "( 7 - 10 ) - 40"
	formedExp = formExpression(expression)
	result = int(count(formedExp))
	if result != expected {
		t.Errorf("Test add failed, %v %v", expected, result)
	}
}

func TestDel(t *testing.T) {
	expected := 3
	expression := "9 / 3"
	formedExp := formExpression(expression)
	result := int(count(formedExp))
	if result != expected {
		t.Errorf("Test add failed, %v %v", expected, result)
	}

	expected = 5
	expression = "(50 / 5) / 2"
	formedExp = formExpression(expression)
	result = int(count(formedExp))
	if result != expected {
		t.Errorf("Test add failed, %v %v", expected, result)
	}
}

func TestMult(t *testing.T) {
	expected := 27
	expression := "9 * 3"
	formedExp := formExpression(expression)
	result := int(count(formedExp))
	if result != expected {
		t.Errorf("Test add failed, %v %v", expected, result)
	}

	expected = 500
	expression = "(50 * 5) * 2"
	formedExp = formExpression(expression)
	result = int(count(formedExp))
	if result != expected {
		t.Errorf("Test add failed, %v %v", expected, result)
	}
}

func TestDifferent(t *testing.T) {
	expected := 32
	expression := "(9 * 3) + (40 / 4) - 5"
	formedExp := formExpression(expression)
	result := int(count(formedExp))
	if result != expected {
		t.Errorf("Test add failed, %v %v", expected, result)
	}

	expected = 17
	expression = "3 + 7 * (5 - 4) + 7"
	formedExp = formExpression(expression)
	result = int(count(formedExp))
	if result != expected {
		t.Errorf("Test add failed, %v %v", expected, result)
	}
}
