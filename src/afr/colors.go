package afr

func getDeepColor(dc string) string {
	switch dc {
	case "8":
		return "0"
	case "10":
		return "1"
	case "12":
		return "2"
	default:
		return dc
	}
}

func getColorSpace(cs string) string {
	switch cs {
	case "422":
		return "1"
	case "444":
		return "2"
	case "420":
		return "3"
	default:
		return cs
	}
}
