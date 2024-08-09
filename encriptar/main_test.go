package main

import (
	"testing"
)

// go test -run TestEncryptMsg -v
func TestEncryptMsg(t *testing.T) {
	// Test cases
	tests := []struct {
		name     string
		key      string
		msg      string
		expected string
	}{
		{
			name:     "Base case: Mensaje normal",
			key:      "dcj",
			msg:      "I love prOgrAmming!",
			expected: "dcjI ldcjovdcje prdcjOgrdcjAmmdcjing!",
		},
		{
			name:     "Clave vacía: Se usa la clave predeterminada",
			key:      "",
			msg:      "Go is awesome",
			expected: "GDCJo DCJis DCJawDCJesDCJomDCJe",
		},
		{
			name:     "Mensaje vacío: Retorna cadena vacía",
			key:      "abc",
			msg:      "",
			expected: "",
		},
		{
			name:     "Mensaje sin vocales: No cambia el mensaje",
			key:      "xyz",
			msg:      "Tstng Msg",
			expected: "Tstng Msg",
		},
		{
			name:     "Sin clave y mensaje con una vocal",
			key:      "",
			msg:      "Test Msg",
			expected: "TDCJest Msg",
		},
		{
			name:     "Mensaje con todas las vocales",
			key:      "xyz",
			msg:      "aeiou",
			expected: "xyzaxyzexyzixyzoxyzu",
		},
		{
			name:     "Clave no vacía y mensaje con vocales",
			key:      "123",
			msg:      "Gophers love Go!",
			expected: "G123oph123ers l123ov123e G123o!",
		},
		{
			name:     "Clave y mensaje sin vocales",
			key:      "abc",
			msg:      "Crypt Msg",
			expected: "Crypt Msg",
		},
		{
			name:     "Clave larga y mensaje con vocales",
			key:      "longkey",
			msg:      "Testing long keys",
			expected: "Tlongkeyestlongkeying llongkeyong klongkeyeys",
		},
		{
			name:     "Clave mayuscula y mensaje vocales con acento",
			key:      "RX",
			msg:      "Última pirámide",
			expected: "ÚltRXimRXa pRXirámRXidRXe",
		},
	}

	// Run table cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := encryptMsg(tt.key, tt.msg)
			if got != tt.expected {
				t.Errorf("encryptMsg(%q, %q) = %q; want %q", tt.key, tt.msg, got, tt.expected)
			}
		})
	}
}
