CREATE SCHEMA IF NOT EXISTS dental_clinic;
USE dental_clinic;
CREATE TABLE dentists (
    id INT NOT NULL AUTO_INCREMENT,
    surname VARCHAR(50) NOT NULL,
    name VARCHAR(25) NOT NULL,
    license_number VARCHAR(10) NOT NULL UNIQUE,

    PRIMARY KEY (id)
)ENGINE = INNODB;

CREATE TABLE patients (
    id INT NOT NULL AUTO_INCREMENT,
    surname VARCHAR(50) NOT NULL,
    name VARCHAR(25) NOT NULL,
    identity_number VARCHAR(10) NOT NULL UNIQUE,
    created_at DATETIME NOT NULL,

    PRIMARY KEY (id)
)ENGINE = INNODB;

CREATE TABLE appointments (
    id INT NOT NULL AUTO_INCREMENT,
    description VARCHAR(250) NOT NULL,
    date_and_time DATETIME NOT NULL,
    dentist_license VARCHAR(10) NOT NULL,
    patient_identity VARCHAR(10) NOT NULL,

    PRIMARY KEY (id),

    CONSTRAINT fk_dentist
                          FOREIGN KEY (dentist_license)
                          REFERENCES dentists(license_number),
    CONSTRAINT fk_patient
                          FOREIGN KEY (patient_identity)
                          REFERENCES patients(identity_number)
)ENGINE = INNODB;