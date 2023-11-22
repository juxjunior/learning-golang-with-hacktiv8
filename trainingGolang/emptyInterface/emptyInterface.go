package main

func main() {
	mainUntukEmptyInterface()
	mainUntukEmptyInterfaceTypeAssertion()
	mainUntukMapDanSLiceDariEmptyInterface()
}

// Contoh sederhana empty interface

func mainUntukEmptyInterface() {
	var randomValues interface{}

	_ = randomValues

	randomValues = "Jalan Sudirman"
	randomValues = 20
	randomValues = true
	randomValues = []string{"Airell", "Nanda"}
}

// Contoh empty interface (type assertion)

func mainUntukEmptyInterfaceTypeAssertion() {
	var v interface{}
	v = 20

	if value, ok := v.(int); ok == true {
		v = value * 9
	}
}

// Contoh map & slice dari empty interface

func mainUntukMapDanSLiceDariEmptyInterface() {
	rs := []interface{}{1, "Airell", true, 2, "Ananda", true}
	rm := map[string]interface{}{
		"Name":   "Airell",
		"Status": true,
		"Age":    23,
	}
	_, _ = rs, rm
}
