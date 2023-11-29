package validator

import (
	"testing"
)

func TestIsCPFCNPJInvalid(t *testing.T) {
	cpfCNPJList := []string{
		"",
		"028982-892-29999999",
		"0289",
		" ",
		"           ",
		"              ",
		"111.111.111-11",
		"222.222.222-22",
		"333.333.333-33",
		"44444444444",
		"527.636.260-50",
		"52763626050",
		"52276326503",
		"90.198.810/0601-97",
		"55.555.555/5555-55",
		"26163008000158",
		"11111111111111",
		"63008000158",
	}

	for _, cpfCNPJ := range cpfCNPJList {
		if IsCPFCNPJ(cpfCNPJ) != false {
			t.Errorf("Expected invalid cpf/cnpj but receive valid")
		}
	}
}

func TestIsCPFCNPJValid(t *testing.T) {
	cpfCNPJList := []string{"987.654.321-00", "98765432100", "714.602.380-01", "71460238001", "313.030.210-72", "31303021072", "144.796.170-60", "14479617060", "35.040.714/0001-08", "96950131000109"}

	for _, cpfCNPJ := range cpfCNPJList {
		if IsCPFCNPJ(cpfCNPJ) != true {
			t.Errorf("Expected valid cpf but receive invalid")
		}
	}
}

func TestIsCPFInvalid(t *testing.T) {
	cpfList := []string{
		"",
		"028982-892-29999999",
		"0289",
		" ",
		"           ",
		"              ",
		"111.111.111-11",
		"222.222.222-22",
		"333.333.333-33",
		"44444444444",
		"527.636.260-50",
		"52763626050",
		"52276326503",
	}

	for _, cpfNumber := range cpfList {
		if IsCPF(cpfNumber) != false {
			t.Errorf("Expected invalid cpf but receive valid")
		}
	}
}

func TestIsCPFValid(t *testing.T) {
	cpfList := []string{"987.654.321-00", "98765432100", "714.602.380-01", "71460238001", "313.030.210-72", "31303021072", "144.796.170-60", "14479617060"}

	for _, cpfNumber := range cpfList {
		if IsCPF(cpfNumber) != true {
			t.Errorf("Expected valid cpf but receive invalid")
		}
	}
}

func TestIsCNPJValid(t *testing.T) {
	cnpjList := []string{"35.040.714/0001-08", "96950131000109"}

	for _, cnpj := range cnpjList {
		if IsCNPJ(cnpj) != true {
			t.Errorf("Expected valid cnpj but receive invalid")
		}
	}
}

func TestIsCNPJInvalid(t *testing.T) {
	cnpjList := []string{"90.198.810/0601-97", "55.555.555/5555-55", "26163008000158", "11111111111111", "63008000158", "              ", " "}

	for _, cnpj := range cnpjList {
		if IsCNPJ(cnpj) != false {
			t.Errorf("Expected invalid cnpj but receive valid")
		}
	}
}
