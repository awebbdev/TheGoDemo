package GoPower


func CalcFTP(watts float64)float64{
	return watts * 0.95
}

func WattsperKG(watts, kg float64)float64{
	return watts / kg
}
