package database

type Doctor struct{
	Name string `json:"name"`
	First_name string `json:"fristName"`
	Last_name string  `json:"lastName"`
	Specialized_field string  `json:"specilizedField"`
	Working_hours string `json:"workingHours"`
}


type Hospital struct{
	Hospital_name string `json:"hospitalName"`
	Contact string `json:"contact"`
	Town string `json:"town"`
	Sector string  `json:"sector"`
	First_name string  `json:"firstName"`
	Last_name string  `json:"lastName"`
	Specialized_field string  `json:"specializedField"`
}

type HospitalAndDoctor struct{
	Hospital
	Working_hours string `json:"workingHours"`
	Working_days string `json:"workingDays"`
}

type Speciality struct{
	First_name string `json:"fristName"`
	Last_name string  `json:"lastName"`
	Hospital_name string `json:"hospitalName"`
}
