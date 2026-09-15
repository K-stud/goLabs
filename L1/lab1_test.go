package main

import "testing"


func TestTask4(t *testing.T) {
	var res [5]int = task4(11, 7)

	if res[0] != 18 {
		t.Errorf("Сумма 11 и 7 должна быть равна 18: Получено %d", res[0])
	}

	if res[1] != 4 {
		t.Errorf("Разность 11 и 7 должна быть равна 4: Получено %d", res[1])
	}

	if res[2] != 77 {
		t.Errorf("Произведение 11 и 7 должно быть равна 77: Получено %d", res[2])
	}

	if res[3] != 1 {
		t.Errorf("Деление 11 и 7 должно быть равно 1: Получено %d", res[3])
	}

	if res[4] != 4 {
		t.Errorf("Остаток деления 11 и 7 должен быть равна 4: Получено %d", res[4])
	}
}

func TestTask5(t *testing.T) {
	var res [2]float64 = task5(10.23, 5.12)

	if res[0] - 15.35 > 0.00001 {
		t.Errorf("Сумма 10.23 и 5.12 должна быть 15.35. Получено %f", res[0])
	}
	if res[1] - 5.11 > 0.00001 {
		t.Errorf("Разность 10.23 и 5.12 должна быть 5.11. Получено %f", res[1])
	}
}