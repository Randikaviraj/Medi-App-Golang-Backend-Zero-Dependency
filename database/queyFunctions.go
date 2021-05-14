package database

import (
	"errors"
	"log"
)


func QueryDoctorByDate(hospitalName,date string,doctors *[] Doctor) error{

	result, err := Db.Query(`SELECT hospital.name,doctor.first_name,doctor.last_name,doctor.specialized_field,works_at.working_hours FROM((works_at INNER JOIN doctor ON works_at.MCID=doctor.MCID) INNER JOIN hospital ON works_at.Hos_reg_id=hospital.Hos_reg_id) WHERE hospital.name LIKE CONCAT(?,'%') AND works_at.working_days=?;`,hospitalName,date)

	if err != nil {
		log.Print(err.Error())
		return errors.New("Something Wrong")
	}

	defer result.Close()
	
	for result.Next(){

		var row Doctor

		result.Scan(&row.Name,&row.First_name,&row.Last_name,&row.Specialized_field,&row.Working_hours)
		*doctors=append(*doctors, row)
	}
	return nil
}


func QueryByHospitalName(hospitalName string,doctors *[] Hospital) error{

	result, err := Db.Query(`SELECT hospital.name,hospital.contact,hospital.town,hospital.sector,doctor.first_name,doctor.last_name,doctor.specialized_field FROM ((works_at INNER JOIN doctor ON works_at.MCID=doctor.MCID) INNER JOIN hospital ON works_at.hos_reg_id=hospital.hos_reg_id) WHERE hospital.name LIKE CONCAT(?,'%');`,hospitalName)

	if err != nil {
		log.Print(err.Error())
		return errors.New("Something Wrong")
	}

	defer result.Close()
	
	for result.Next(){

		var row Hospital

		result.Scan(&row.Hospital_name,&row.Contact,&row.Town,&row.Sector,&row.First_name,&row.Last_name,&row.Specialized_field)
		*doctors=append(*doctors, row)
	}

	return nil
}


func QueryByDoctorName(doctorName string,doctors *[] HospitalAndDoctor) error{
	

	result, err := Db.Query(`SELECT doctor.first_name,doctor.last_name,doctor.specialized_field,hospital.name,hospital.contact,hospital.town,works_at.Working_hours,works_at.Working_days FROM ((works_at INNER JOIN doctor on works_at.MCID=doctor.MCID) INNER JOIN hospital ON works_at.Hos_reg_id=hospital.Hos_reg_id) WHERE doctor.first_name LIKE CONCAT(?,'%') || doctor.last_name LIKE CONCAT(?,'%');`,doctorName,doctorName)

	if err != nil {
		log.Print(err.Error())
		return errors.New("Something Wrong")
	}

	defer result.Close()
	
	for result.Next(){

		var row HospitalAndDoctor

		result.Scan(&row.First_name,&row.Last_name,&row.Specialized_field,&row.Hospital_name,&row.Contact,&row.Town,&row.Working_hours,&row.Working_days)
		*doctors=append(*doctors, row)
	}

	return nil
}

func QueryBySpecialityName(specialityName string,doctors *[] Speciality) error{

	result, err := Db.Query(`SELECT doctor.first_name,doctor.last_name,hospital.name FROM(( works_at INNER JOIN doctor ON works_at.MCID=doctor.MCID) 
	INNER JOIN hospital ON works_at.Hos_reg_id=hospital.Hos_reg_id)  WHERE doctor.specialized_field=?;`,specialityName)

	if err != nil {
		log.Print(err.Error())
		return errors.New("Something Wrong")
	}

	defer result.Close()
	
	for result.Next(){

		var row Speciality

		result.Scan(&row.First_name,&row.Last_name,&row.Hospital_name)
		*doctors=append(*doctors, row)
	}

	return nil
}



