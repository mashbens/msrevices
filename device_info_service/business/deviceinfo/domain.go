package deviceinfo

type Result struct {
	// TotalAllData     int
	// CurrentPage      int
	// PerPage          int
	// TotalPage        int
	// TotalRunningTime int
	Domain Domain
}

type Domain struct {
	Imei            string
	GpsType         int
	StrGpsType      string
	DeviceName      string
	VehicleType     int
	StrVehicleType  string
	Palte           string
	GsmNumber       string
	Icon            int
	Stnk            string
	AdminBranch     string
	OwnerName       string
	InstalationTech string
	Year            int
	MachineNumber   string
	ChassisNumber   string
	SpeedLimit      int
	Norangka        string
	Nomesin         string
}
