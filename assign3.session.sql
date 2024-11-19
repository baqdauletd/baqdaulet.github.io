-- DROP TABLE IF EXISTS Specialize;
-- DROP TABLE IF EXISTS Doctor;
-- DROP TABLE IF EXISTS Record;
-- -- DROP TABLE IF EXISTS Patients;
-- DROP TABLE IF EXISTS Discover;
-- -- DROP TABLE IF EXISTS Disease;
-- -- DROP TABLE IF EXISTS DiseaseType;
-- -- DROP TABLE IF EXISTS Country;
-- -- DROP TABLE IF EXISTS Users;
-- DROP TABLE IF EXISTS PublicServant;
-- -- DROP TABLE IF EXISTS PatientDisease;

-- drop view patient_disease_view


-- Deletion in reverse order of foreign key dependencies
-- DELETE FROM Specialize;
-- DELETE FROM Doctor;
-- DELETE FROM PublicServant;
-- DELETE FROM Record;
-- DELETE FROM PatientDisease;
-- DELETE FROM Patients;
-- DELETE FROM Discover;
-- DELETE FROM Disease;
-- DELETE FROM DiseaseType;
-- DELETE FROM Country;
-- DELETE FROM Users;


-- -- Insert records into Country
-- INSERT INTO Country (cname, population) VALUES
--     ('Brazil', 211000000),
--     ('Germany', 83000000),
--     ('China', 1393000000),
--     ('France', 67000000),
--     ('Kazakhstan', 24500000),
--     ('Italy', 60000000),
--     ('Mexico', 126000000),
--     ('Canada', 38000000),
--     ('Japan', 126000000),
--     ('Australia', 25000000);

-- -- Insert records into DiseaseType
-- INSERT INTO DiseaseType (id, description) VALUES
--     (1, 'Bacterial Diseases'),
--     (2, 'Viral Diseases'),
--     (3, 'Fungal Diseases'),
--     (4, 'Parasitic Diseases'),
--     (5, 'Prion Diseases'),
--     (6, 'Genetic Disorders'),
--     (7, 'Autoimmune Diseases'),
--     (8, 'Degenerative Diseases'),
--     (9, 'Nutritional Deficiencies'),
--     (10, 'Psychiatric Disorders'),
--     (11, 'Infectious Diseases');


-- Insert records into Users
-- INSERT INTO Users (email, name, surname, salary, phone, cname) VALUES
--     ('bakdaulet.dauletov@example.com', 'Bakdaulet', 'Dauletov', 55000, '123-456-7890', 'Brazil'),
--     ('damir.alipbayev@example.com', 'Damir', 'Alipbayev', 60000, '123-456-7891', 'Germany'),
--     ('zhanara.bekbergenova@example.com', 'Zhanara', 'Bekbergenova', 62000, '123-456-7892', 'China'),
--     ('adele.dauletova@example.com', 'Adele', 'Dauletova', 57000, '123-456-7893', 'France'),
--     ('zhanel.dauletova@example.com', 'Zhanel', 'Dauletova', 59000, '123-456-7894', 'Kazakhstan'),
--     ('gulshat.abdikaimova@example.com', 'Gulshat', 'Abdikaimova', 61000, '123-456-7895', 'Italy'),
--     ('guldana.zhubandyk@example.com', 'Guldana', 'Zhubandyk', 58000, '123-456-7896', 'Mexico'),
--     ('aibek.narik@example.com', 'Aibek', 'Narik', 53000, '123-456-7897', 'Canada'),
--     ('nurbek.nysanbayev@example.com', 'Nurbek', 'Nysanbayev', 54000, '123-456-7898', 'Japan'),
--     ('temirlan.suleimenov@example.com', 'Temirlan', 'Suleimenov', 60000, '123-456-7899', 'Australia'),
--     ('batyrbek.aliev@example.com', 'Batyrbek', 'Aliev', 56000, '123-456-7800', 'Brazil'),
--     ('murad.kassym@example.com', 'Murad', 'Kassym', 59000, '123-456-7801', 'Germany'),
--     ('karina.ongar@example.com', 'Karina', 'Ongar', 62000, '123-456-7802', 'China'),
--     ('diana.nurly@example.com', 'Diana', 'Nurly', 54000, '123-456-7803', 'France'),
--     ('yerassyl.sagymbay@example.com', 'Yerassyl', 'Sagymbay', 61000, '123-456-7804', 'Kazakhstan'),
--     ('marzhan.talassova@example.com', 'Marzhan', 'Talassova', 58000, '123-456-7805', 'Italy'),
--     ('omirbek.askar@example.com', 'Omirbek', 'Askar', 62000, '123-456-7806', 'Mexico'),
--     ('rauan.kazhim@example.com', 'Rauan', 'Kazhim', 57000, '123-456-7807', 'Canada'),
--     ('meirzhan.nurtas@example.com', 'Meirzhan', 'Nurtas', 55000, '123-456-7808', 'Japan'),
--     ('saule.yermakhan@example.com', 'Saule', 'Yermakhan', 60000, '123-456-7809', 'Australia');

-- -- Insert records into Doctor
-- INSERT INTO Doctor (email, degree) VALUES
--     ('bakdaulet.dauletov@example.com', 'MD'),
--     ('damir.alipbayev@example.com', 'PhD'),
--     ('zhanara.bekbergenova@example.com', 'MD'),
--     ('adele.dauletova@example.com', 'MD'),
--     ('zhanel.dauletova@example.com', 'PhD'),
--     ('gulshat.abdikaimova@example.com', 'MD'),
--     ('guldana.zhubandyk@example.com', 'PhD'),
--     ('aibek.narik@example.com', 'MD'),
--     ('nurbek.nysanbayev@example.com', 'MD'),
--     ('temirlan.suleimenov@example.com', 'PhD');

-- -- Insert records into PublicServant
-- INSERT INTO PublicServant (email, department) VALUES
--     ('bakdaulet.dauletov@example.com', 'Health'),
--     ('damir.alipbayev@example.com', 'Research'),
--     ('zhanara.bekbergenova@example.com', 'Health'),
--     ('adele.dauletova@example.com', 'Admin'),
--     ('zhanel.dauletova@example.com', 'Health'),
--     ('gulshat.abdikaimova@example.com', 'Admin'),
--     ('guldana.zhubandyk@example.com', 'Research'),
--     ('aibek.narik@example.com', 'Health'),
--     ('nurbek.nysanbayev@example.com', 'Health'),
--     ('temirlan.suleimenov@example.com', 'Research'),
--     ('batyrbek.aliev@example.com', 'Health'),
--     ('murad.kassym@example.com', 'Research'),
--     ('karina.ongar@example.com', 'Health'),
--     ('diana.nurly@example.com', 'Health'),
--     ('yerassyl.sagymbay@example.com', 'Research');

-- -- Insert records into Patients
-- INSERT INTO Patients (email) VALUES
--     ('marzhan.talassova@example.com'),
--     ('omirbek.askar@example.com'),
--     ('rauan.kazhim@example.com'),
--     ('meirzhan.nurtas@example.com'),
--     ('saule.yermakhan@example.com');

-- Insert records into Disease
-- INSERT INTO Disease (disease_code, pathogen, description, id) VALUES
--     ('COVID-19', 'Virus', 'Coronavirus disease 2019', 11),
--     ('Flu', 'Virus', 'Seasonal influenza', 2),
--     ('Tuberculosis', 'Bacteria', 'Bacterial infection', 1),
--     ('Malaria', 'Parasite', 'Parasitic infection', 4),
--     ('Celiac', 'Autoimmune', 'Gluten intolerance', 7),
--     ('Depression', 'Psychiatric', 'Mental health disorder', 10),
--     ('Scurvy', 'Nutritional', 'Vitamin C deficiency', 9),
--     ('BSE', 'Prion', 'Mad cow disease', 5),
--     ('ALS', 'Degenerative', 'Motor neuron disease', 8),
--     ('Herpes', 'Virus', 'Herpes simplex', 2);

-- -- Insert records into Discover
-- INSERT INTO Discover (cname, disease_code, first_enc_date) VALUES
--     ('Brazil', 'COVID-19', '2020-01-15'),
--     ('Germany', 'Flu', '1900-01-01'),
--     ('China', 'Tuberculosis', '1882-03-24'),
--     ('France', 'Malaria', '1902-04-02'),
--     ('Kazakhstan', 'Celiac', '2015-05-06');


-- -- Insert records into PatientDisease
-- INSERT INTO PatientDisease (email, disease_code) VALUES
--     ('marzhan.talassova@example.com', 'COVID-19'),
--     ('omirbek.askar@example.com', 'Tuberculosis'),
--     ('rauan.kazhim@example.com', 'Malaria'),
--     ('meirzhan.nurtas@example.com', 'Celiac'),
--     ('saule.yermakhan@example.com', 'Scurvy');

-- -- Insert records into Record
-- INSERT INTO Record (email, cname, disease_code, total_deaths, total_patients) VALUES
--     ('bakdaulet.dauletov@example.com', 'Brazil', 'COVID-19', 5000, 100000),
--     ('damir.alipbayev@example.com', 'Germany', 'Flu', 1000, 20000),
--     ('zhanara.bekbergenova@example.com', 'China', 'Tuberculosis', 3000, 150000),
--     ('zhanel.dauletova@example.com', 'France', 'Malaria', 1500, 50000),
--     ('gulshat.abdikaimova@example.com', 'Kazakhstan', 'Celiac', 200, 5000),
--     ('guldana.zhubandyk@example.com', 'Italy', 'BSE', 200, 10000),
--     ('aibek.narik@example.com', 'Canada', 'Herpes', 500, 15000),
--     ('nurbek.nysanbayev@example.com', 'Japan', 'Scurvy', 200, 10000);


-- Insert specializations, ensuring some doctors have more than two specializations
-- INSERT INTO Specialize (id, email) VALUES
--     (1, 'bakdaulet.dauletov@example.com'),   -- Bacterial Diseases
--     (2, 'bakdaulet.dauletov@example.com'),   -- Viral Diseases
--     (3, 'bakdaulet.dauletov@example.com'),   -- Fungal Diseases
    
--     (2, 'damir.alipbayev@example.com'),      -- Viral Diseases
--     (4, 'damir.alipbayev@example.com'),      -- Parasitic Diseases
--     (5, 'damir.alipbayev@example.com'),      -- Prion Diseases

--     (6, 'zhanara.bekbergenova@example.com'), -- Genetic Disorders
--     (7, 'zhanara.bekbergenova@example.com'), -- Autoimmune Diseases
    
--     (8, 'adele.dauletova@example.com'),      -- Degenerative Diseases
--     (10, 'adele.dauletova@example.com'),     -- Psychiatric Disorders
--     (11, 'adele.dauletova@example.com'),     -- Infectious Diseases
    
--     (1, 'zhanel.dauletova@example.com'),     -- Bacterial Diseases
--     (3, 'zhanel.dauletova@example.com'),     -- Fungal Diseases
    
--     (9, 'gulshat.abdikaimova@example.com'),  -- Nutritional Deficiencies
--     (2, 'gulshat.abdikaimova@example.com'),  -- Viral Diseases
    
--     (11, 'guldana.zhubandyk@example.com'),   -- Infectious Diseases
--     (6, 'guldana.zhubandyk@example.com'),    -- Genetic Disorders
    
--     (4, 'aibek.narik@example.com'),          -- Parasitic Diseases
--     (9, 'aibek.narik@example.com');          -- Nutritional Deficiencies

-- -- Alter the Doctor table to cascade deletes
-- ALTER TABLE Doctor
-- DROP CONSTRAINT doctor_email_fkey, -- Replace with the actual foreign key constraint name if needed
-- ADD CONSTRAINT doctor_email_fkey FOREIGN KEY (email)
-- REFERENCES Users(email) ON DELETE CASCADE;

-- -- Alter the Patients table to cascade deletes
-- ALTER TABLE Patients
-- DROP CONSTRAINT patients_email_fkey, -- Replace with the actual foreign key constraint name if needed
-- ADD CONSTRAINT patients_email_fkey FOREIGN KEY (email)
-- REFERENCES Users(email) ON DELETE CASCADE;

-- -- Alter the PublicServant table to cascade deletes
-- ALTER TABLE PublicServant
-- DROP CONSTRAINT publicservant_email_fkey, -- Replace with the actual foreign key constraint name if needed
-- ADD CONSTRAINT publicservant_email_fkey FOREIGN KEY (email)
-- REFERENCES Users(email) ON DELETE CASCADE;

-- -- Alter the PatientDisease table to cascade deletes when a user is deleted
-- ALTER TABLE PatientDisease
-- DROP CONSTRAINT patientdisease_email_fkey, -- Replace with the actual foreign key constraint name for `email` if needed
-- ADD CONSTRAINT patientdisease_email_fkey FOREIGN KEY (email)
-- REFERENCES Users(email) ON DELETE CASCADE;


-- Alter the Specialize table to cascade deletes when a doctor is deleted
-- ALTER TABLE Specialize
-- DROP CONSTRAINT specialize_email_fkey, -- Replace with actual foreign key name if needed
-- ADD CONSTRAINT specialize_email_fkey FOREIGN KEY (email)
-- REFERENCES Doctor(email) ON DELETE CASCADE;


-- -- Alter the Record table to cascade deletes when a public servant is deleted
-- ALTER TABLE Record
-- DROP CONSTRAINT record_email_fkey, -- Replace with actual foreign key name if needed
-- ADD CONSTRAINT record_email_fkey FOREIGN KEY (email)
-- REFERENCES PublicServant(email) ON DELETE CASCADE;

-- ALTER TABLE Disease
-- DROP CONSTRAINT disease_id_fkey, -- Replace with actual foreign key name if needed
-- ADD CONSTRAINT disease_id_fkey FOREIGN KEY (id)
-- REFERENCES DiseaseType(id) ON DELETE CASCADE;

-- ALTER TABLE Specialize
-- DROP CONSTRAINT specialize_id_fkey, -- Replace with actual foreign key name if needed
-- ADD CONSTRAINT specialize_id_fkey FOREIGN KEY (id)
-- REFERENCES DiseaseType(id) ON DELETE CASCADE;

-- ALTER TABLE Discover
-- DROP CONSTRAINT discover_disease_code_fkey, -- Replace with actual foreign key name if needed
-- ADD CONSTRAINT discover_disease_code_fkey FOREIGN KEY (disease_code)
-- REFERENCES disease(disease_code) ON DELETE CASCADE;

-- ALTER TABLE Record
-- DROP CONSTRAINT record_disease_code_fkey, -- Replace with actual foreign key name if needed
-- ADD CONSTRAINT record_disease_code_fkey FOREIGN KEY (disease_code)
-- REFERENCES disease(disease_code) ON DELETE CASCADE;

-- ALTER TABLE Users
-- DROP CONSTRAINT users_cname_fkey, -- Replace with actual foreign key name if needed
-- ADD CONSTRAINT users_cname_fkey FOREIGN KEY (cname)
-- REFERENCES Country(cname) ON DELETE CASCADE;

-- ALTER TABLE PatientDisease
-- DROP CONSTRAINT patientdisease_disease_code_fkey, -- Replace with actual foreign key name if needed
-- ADD CONSTRAINT patientdisease_disease_code_fkey FOREIGN KEY (disease_code)
-- REFERENCES disease(disease_code) ON DELETE CASCADE;

-- ALTER TABLE Discover
-- DROP CONSTRAINT discover_cname_fkey, -- Replace with actual foreign key name if needed
-- ADD CONSTRAINT discover_cname_fkey FOREIGN KEY (cname)
-- REFERENCES Country(cname) ON DELETE CASCADE;

-- ALTER TABLE Record
-- DROP CONSTRAINT record_cname_fkey, -- Replace with actual foreign key name if needed
-- ADD CONSTRAINT record_cname_fkey FOREIGN KEY (cname)
-- REFERENCES Country(cname) ON DELETE CASCADE;

-- Alter table patientdisease
-- drop constraint patientdisease_email_fkey,
-- add constraint patientdisease_email_fkey FOREIGN KEY (email)
-- REFERENCES Patients(email) ON DELETE CASCADE;

-- INSERT INTO Country (cname, population) VALUES
--     ('United States', 331000000),
--     ('India', 1380000000),
--     ('Brazil', 213000000),
--     ('Kazakhstan', 20000000),
--     ('United Kingdom', 67000000),
--     ('Germany', 83000000),
--     ('China', 1440000000),
--     ('France', 65000000),
--     ('South Africa', 60000000),
--     ('Canada', 38000000),
--     ('Australia', 25000000),
--     ('Japan', 125000000),
--     ('Italy', 60000000),
--     ('Mexico', 126000000),
--     ('Turkey', 85000000);

    -- Populate DiseaseType table with 15 entries (adding other types for diversity)
-- INSERT INTO DiseaseType (id, description) VALUES
--     (16, 'Infectous Diseases');

-- -- Populate Disease table with 15 entries, including additional bacterial diseases
-- INSERT INTO Disease (disease_code, pathogen, description, id) VALUES
--     ('TB001', 'bacteria', 'Tuberculosis', 1),
--     ('STP002', 'not bacteria', 'Strep Throat', 1),
--     ('CHL003', 'bacteria', 'Cholera', 1),
--     ('PNE006', 'not bacteria', 'Pneumonia', 1),
--     ('DIP007', 'bacteria', 'Diphtheria', 1),
--     ('LPN008', 'not bacteria', 'Legionnaires Disease', 1),
--     ('PLG009', 'bacteria', 'Plague', 1),
--     ('TYP010', 'bacteria', 'Typhoid Fever', 1),
--     ('SHG011', 'not bacteria', 'Shigellosis', 1),
--     ('MNG012', 'bacteria', 'Meningitis', 1),
--     ('BTS013', 'not bacteria', 'Botulism', 1),
--     ('ANT014', 'not bacteria', 'Anthrax', 1),
--     ('SYP015', 'not bacteria', 'Syphilis', 1),
--     ('LYM016', 'not bacteria', 'Lyme Disease', 1),
--     ('GON017', 'bacteria', 'Gonorrhea', 1);

-- Populate Users table with doctors, public servants, and other users
-- INSERT INTO Users (email, name, surname, salary, phone, cname) VALUES
--     ('bakdaulet.dauletov@example.com', 'Bakdaulet', 'Dauletov', 55000, '123-456-7890', 'Brazil'),
--     ('damir.alipbayev@example.com', 'Damir', 'Alipbayev', 60000, '123-456-7891', 'Germany'),
--     ('zhanara.bekbergenova@example.com', 'Zhanara', 'Bekbergenova', 62000, '123-456-7892', 'China'),
--     ('adele.dauletova@example.com', 'Adele', 'Dauletova', 57000, '123-456-7893', 'France'),
--     ('zhanel.dauletova@example.com', 'Zhanel', 'Dauletova', 59000, '123-456-7894', 'Kazakhstan'),
--     ('gulshat.abdikaimova@example.com', 'Gulshat', 'Abdikaimova', 61000, '123-456-7895', 'Italy'),
--     ('guldana.zhubandyk@example.com', 'Guldana', 'Zhubandyk', 58000, '123-456-7896', 'Mexico'),
--     ('aibek.narik@example.com', 'Aibek', 'Narik', 53000, '123-456-7897', 'Canada'),
--     ('nurbek.nysanbayev@example.com', 'Nurbek', 'Nysanbayev', 54000, '123-456-7898', 'Japan'),
--     ('temirlan.suleimenov@example.com', 'Temirlan', 'Suleimenov', 60000, '123-456-7899', 'Australia');

-- INSERT INTO Users (email, name, surname, salary, phone, cname) VALUES
--     ('aliya.batyrbekova@example.com', 'Aliya', 'Batyrbekova', 56000, '123-456-7800', 'Brazil'),
--     ('murad.kassym@example.com', 'Murad', 'Kassym', 59000, '123-456-7801', 'Germany'),
--     ('karina.ongar@example.com', 'Karina', 'Ongar', 62000, '123-456-7802', 'China'),
--     ('diana.nurly@example.com', 'Diana', 'Nurly', 54000, '123-456-7803', 'France'),
--     ('yerassyl.sagymbay@example.com', 'Yerassyl', 'Sagymbay', 61000, '123-456-7804', 'Kazakhstan'),
--     ('marzhan.talassova@example.com', 'Marzhan', 'Talassova', 58000, '123-456-7805', 'Italy'),
--     ('askar.omirbek@example.com', 'Askar', 'Omirbek', 62000, '123-456-7806', 'Mexico'),
--     ('rauan.kazhim@example.com', 'Rauan', 'Kazhim', 57000, '123-456-7807', 'Canada'),
--     ('meirzhan.nurtas@example.com', 'Meirzhan', 'Nurtas', 55000, '123-456-7808', 'Japan'),
--     ('saule.yermakhan@example.com', 'Saule', 'Yermakhan', 60000, '123-456-7809', 'Australia');


-- -- Populate Patients table with a subset of users
-- INSERT INTO Patients (email) VALUES
--     ('carol@example.com'),
--     ('eve@example.com'),
--     ('grace@example.com'),
--     ('jack@example.com'),
--     ('kate@example.com'),
--     ('leo@example.com'),
--     ('noah@example.com'),
--     ('olivia@example.com'),
--     ('hank@example.com'),
--     ('frank@example.com'),
--     ('dave@example.com'),
--     ('bob@example.com'),
--     ('alice@example.com'),
--     ('irene@example.com'),
--     ('mia@example.com');

-- Populate Discover table with bacterial diseases and discovery dates before 2024
-- INSERT INTO Discover (cname, disease_code, first_enc_date) VALUES
--     ('United States', 'TB001', '1882-03-24'),
--     ('India', 'CHL003', '1817-01-01'),
--     ('Brazil', 'STP002', '1905-06-10'),
--     ('Kazakhstan', 'PNE006', '1892-12-01'),
--     ('United Kingdom', 'DIP007', '1923-06-01'),
--     ('France', 'LPN008', '1976-01-01'),
--     ('China', 'PLG009', '1347-10-01'),
--     ('Germany', 'TYP010', '1880-01-01'),
--     ('South Africa', 'SHG011', '1935-04-15'),
--     ('Canada', 'MNG012', '1887-02-18'),
--     ('Australia', 'BTS013', '1897-09-10'),
--     ('Japan', 'ANT014', '1950-06-30'),
--     ('Italy', 'SYP015', '1495-03-20'),
--     ('Mexico', 'LYM016', '1975-08-30'),
--     ('Turkey', 'GON017', '1879-02-15');

-- INSERT INTO Doctor (email, degree) VALUES
--     ('bakdaulet.dauletov@example.com', 'MD'),
--     ('damir.alipbayev@example.com', 'PhD'),
--     ('zhanara.bekbergenova@example.com', 'MD'),
--     ('adele.dauletova@example.com', 'DO'),
--     ('zhanel.dauletova@example.com', 'PhD'),
--     ('gulshat.abdikaimova@example.com', 'MD'),
--     ('guldana.zhubandyk@example.com', 'DDS'),
--     ('aibek.narik@example.com', 'MD'),
--     ('nurbek.nysanbayev@example.com', 'DO'),
--     ('temirlan.suleimenov@example.com', 'PhD');

-- Update 'DO' degrees to 'MD'
-- UPDATE Doctor
-- SET degree = 'MD'
-- WHERE degree = 'DO';

-- -- Update 'DDS' degrees to 'PhD'
-- UPDATE Doctor
-- SET degree = 'PhD'
-- WHERE degree = 'DDS';



-- INSERT INTO Specialize (id, email) VALUES
--     (1, 'bakdaulet.dauletov@example.com'),
--     (2, 'bakdaulet.dauletov@example.com'),
--     (3, 'damir.alipbayev@example.com'),
--     (4, 'damir.alipbayev@example.com'),
--     (5, 'zhanara.bekbergenova@example.com'),
--     (11, 'zhanara.bekbergenova@example.com'),
--     (1, 'adele.dauletova@example.com'),
--     (2, 'adele.dauletova@example.com'),
--     (8, 'zhanel.dauletova@example.com'),
--     (10, 'zhanel.dauletova@example.com'),
--     (9, 'gulshat.abdikaimova@example.com'),
--     (6, 'gulshat.abdikaimova@example.com'),
--     (7, 'guldana.zhubandyk@example.com'),
--     (11, 'guldana.zhubandyk@example.com'),
--     (4, 'aibek.narik@example.com'),
--     (2, 'aibek.narik@example.com'),
--     (3, 'nurbek.nysanbayev@example.com'),
--     (5, 'nurbek.nysanbayev@example.com'),
--     (1, 'temirlan.suleimenov@example.com'),
--     (8, 'temirlan.suleimenov@example.com');

-- INSERT INTO Specialize (id, email) VALUES
--     (16, 'zhanel.dauletova@example.com'),
--     (16, 'gulshat.abdikaimova@example.com');

-- CREATE TABLE Country (
--     cname VARCHAR(50) PRIMARY KEY,
--     population BIGINT
-- );

-- CREATE TABLE DiseaseType (
--     id INTEGER PRIMARY KEY,
--     description VARCHAR(140)
-- );

-- CREATE TABLE Disease (
--     disease_code VARCHAR(50) PRIMARY KEY,
--     pathogen VARCHAR(20),
--     description VARCHAR(140),
--     id INTEGER,
--     FOREIGN KEY (id) REFERENCES DiseaseType(id)
-- );

-- CREATE TABLE Users (
--     email VARCHAR(60) PRIMARY KEY,
--     name VARCHAR(30),
--     surname VARCHAR(40),
--     salary INTEGER,
--     phone VARCHAR(20),
--     cname VARCHAR(50),
--     FOREIGN KEY (cname) REFERENCES Country(cname)
-- );

-- CREATE TABLE Patients (
--     email VARCHAR(60) PRIMARY KEY,
--     FOREIGN KEY (email) REFERENCES Users(email)
-- );

-- CREATE TABLE Discover (
--     cname VARCHAR(50),
--     disease_code VARCHAR(50),
--     first_enc_date DATE,
--     PRIMARY KEY (cname, disease_code),
--     FOREIGN KEY (cname) REFERENCES Country(cname),
--     FOREIGN KEY (disease_code) REFERENCES Disease(disease_code)
-- );

-- CREATE TABLE PatientDisease (
--     email VARCHAR(60),
--     disease_code VARCHAR(50),
--     PRIMARY KEY (email, disease_code),
--     FOREIGN KEY (email) REFERENCES Users(email),
--     FOREIGN KEY (disease_code) REFERENCES Disease(disease_code)
-- );

-- CREATE TABLE PublicServant (
--     email VARCHAR(60) PRIMARY KEY,
--     department VARCHAR(50),
--     FOREIGN KEY (email) REFERENCES Users(email)
-- );

-- CREATE TABLE Doctor (
--     email VARCHAR(60) PRIMARY KEY,
--     degree VARCHAR(20),
--     FOREIGN KEY (email) REFERENCES Users(email)
-- );

-- CREATE TABLE Specialize (
--     id INTEGER,
--     email VARCHAR(60),
--     PRIMARY KEY (id, email),
--     FOREIGN KEY (id) REFERENCES DiseaseType(id),
--     FOREIGN KEY (email) REFERENCES Doctor(email)
-- );

-- CREATE TABLE Record (
--     email VARCHAR(60),
--     cname VARCHAR(50),
--     disease_code VARCHAR(50),
--     total_deaths INTEGER,
--     total_patients INTEGER,
--     PRIMARY KEY (email, cname, disease_code),
--     FOREIGN KEY (email) REFERENCES PublicServant(email),
--     FOREIGN KEY (cname) REFERENCES Country(cname),
--     FOREIGN KEY (disease_code) REFERENCES Disease(disease_code)
-- );
