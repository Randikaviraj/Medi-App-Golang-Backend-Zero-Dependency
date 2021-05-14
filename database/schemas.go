package database

type Doctor struct{
	Name string `json:"name"`
	First_name string `json:"fristName"`
	Last_name string  `json:"lastName"`
	Specialized_field string  `json:"specilizedField"`
	Working_hours string `json:"workingHours"`
}


type Hospital struct{
	Hospital_name string
	Contact string
	Town string
	Sector string
	First_name string 
	Last_name string  
	Specialized_field string  
}

type HospitalAndDoctor struct{
	Hospital
	Working_hours string
	Working_days string
}

type Speciality struct{
	First_name string
	Last_name string
	Hospital_name string
}
