package movie
func FindName(imbID string) string {
	switch imbID {
	case "tt4323314":
		return "Avengers: Endgame!!!"
	case "tt4154796":
		return "Black Panther"
	}

	return "Unknown Movie."
}