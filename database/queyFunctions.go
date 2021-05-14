package database

import (
	"log"
)
func QueryDoctorByDate(hospitalName,date string,doctors *[] Doctor){

	result, err := Db.Query(`SELECT hospital.name,doctor.first_name,doctor.last_name,doctor.specialized_field,works_at.working_hours FROM((works_at INNER JOIN doctor ON works_at.MCID=doctor.MCID) INNER JOIN hospital ON works_at.Hos_reg_id=hospital.Hos_reg_id) WHERE hospital.name LIKE CONCAT(?,'%') AND works_at.working_days=?;`,hospitalName,date)

	if err != nil {
        log.Print(err.Error())
	}

	defer result.Close()
	
	for result.Next(){

		var row Doctor

		result.Scan(&row.Name,&row.First_name,&row.Last_name,&row.Specialized_field,&row.Working_hours)
		*doctors=append(*doctors, row)
	}
}