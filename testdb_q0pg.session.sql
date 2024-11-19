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
--     FOREIGN KEY (id) REFERENCES DiseaseType(id) ON DELETE CASCADE ON UPDATE CASCADE
-- );

-- CREATE TABLE Users (
--     email VARCHAR(60) PRIMARY KEY,
--     name VARCHAR(30),
--     surname VARCHAR(40),
--     salary INTEGER,
--     phone VARCHAR(20),
--     cname VARCHAR(50),
--     FOREIGN KEY (cname) REFERENCES Country(cname) ON DELETE CASCADE ON UPDATE CASCADE
-- );

-- CREATE TABLE Patients (
--     email VARCHAR(60) PRIMARY KEY,
--     FOREIGN KEY (email) REFERENCES Users(email) ON DELETE CASCADE ON UPDATE CASCADE
-- );

-- CREATE TABLE Discover (
--     cname VARCHAR(50),
--     disease_code VARCHAR(50),
--     first_enc_date DATE,
--     PRIMARY KEY (cname, disease_code),
--     FOREIGN KEY (cname) REFERENCES Country(cname)  ON DELETE CASCADE ON UPDATE CASCADE,
--     FOREIGN KEY (disease_code) REFERENCES Disease(disease_code) ON DELETE CASCADE ON UPDATE CASCADE
-- );

-- CREATE TABLE PatientDisease (
--     email VARCHAR(60),
--     disease_code VARCHAR(50),
--     PRIMARY KEY (email, disease_code),
--     FOREIGN KEY (email) REFERENCES Users(email) ON DELETE CASCADE ON UPDATE CASCADE,
--     FOREIGN KEY (disease_code) REFERENCES Disease(disease_code) ON DELETE CASCADE ON UPDATE CASCADE
-- );

-- CREATE TABLE PublicServant (
--     email VARCHAR(60) PRIMARY KEY,
--     department VARCHAR(50),
--     FOREIGN KEY (email) REFERENCES Users(email) ON DELETE CASCADE ON UPDATE CASCADE
-- );

-- CREATE TABLE Doctor (
--     email VARCHAR(60) PRIMARY KEY,
--     degree VARCHAR(20),
--     FOREIGN KEY (email) REFERENCES Users(email) ON DELETE CASCADE ON UPDATE CASCADE
-- );

-- CREATE TABLE Specialize (
--     id INTEGER,
--     email VARCHAR(60),
--     PRIMARY KEY (id, email),
--     FOREIGN KEY (id) REFERENCES DiseaseType(id) ON DELETE CASCADE ON UPDATE CASCADE,
--     FOREIGN KEY (email) REFERENCES Doctor(email) ON DELETE CASCADE ON UPDATE CASCADE
-- );

-- CREATE TABLE Record (
--     email VARCHAR(60),
--     cname VARCHAR(50),
--     disease_code VARCHAR(50),
--     total_deaths INTEGER,
--     total_patients INTEGER,
--     PRIMARY KEY (email, cname, disease_code),
--     FOREIGN KEY (email) REFERENCES PublicServant(email) ON DELETE CASCADE ON UPDATE CASCADE,
--     FOREIGN KEY (cname) REFERENCES Country(cname) ON DELETE CASCADE ON UPDATE CASCADE,
--     FOREIGN KEY (disease_code) REFERENCES Disease(disease_code) ON DELETE CASCADE ON UPDATE CASCADE
-- );



-- Insert records into Country
INSERT INTO Country (cname, population) VALUES
    ('Brazil', 211000000),
    ('Germany', 83000000),
    ('China', 1393000000),
    ('France', 67000000),
    ('Kazakhstan', 24500000),
    ('Italy', 60000000),
    ('Mexico', 126000000),
    ('Canada', 38000000),
    ('Japan', 126000000),
    ('Australia', 25000000);

-- Insert records into DiseaseType
INSERT INTO DiseaseType (id, description) VALUES
    (1, 'Bacterial Diseases'),
    (2, 'Viral Diseases'),
    (3, 'Fungal Diseases'),
    (4, 'Parasitic Diseases'),
    (5, 'Prion Diseases'),
    (6, 'Genetic Disorders'),
    (7, 'Autoimmune Diseases'),
    (8, 'Degenerative Diseases'),
    (9, 'Nutritional Deficiencies'),
    (10, 'Psychiatric Disorders'),
    (11, 'Infectious Diseases');


-- Insert records into Users
INSERT INTO Users (email, name, surname, salary, phone, cname) VALUES
    ('bakdaulet.dauletov@example.com', 'Bakdaulet', 'Dauletov', 55000, '123-456-7890', 'Brazil'),
    ('damir.alipbayev@example.com', 'Damir', 'Alipbayev', 60000, '123-456-7891', 'Germany'),
    ('zhanara.bekbergenova@example.com', 'Zhanara', 'Bekbergenova', 62000, '123-456-7892', 'China'),
    ('adele.dauletova@example.com', 'Adele', 'Dauletova', 57000, '123-456-7893', 'France'),
    ('zhanel.dauletova@example.com', 'Zhanel', 'Dauletova', 59000, '123-456-7894', 'Kazakhstan'),
    ('gulshat.abdikaimova@example.com', 'Gulshat', 'Abdikaimova', 61000, '123-456-7895', 'Italy'),
    ('guldana.zhubandyk@example.com', 'Guldana', 'Zhubandyk', 58000, '123-456-7896', 'Mexico'),
    ('aibek.narik@example.com', 'Aibek', 'Narik', 53000, '123-456-7897', 'Canada'),
    ('nurbek.nysanbayev@example.com', 'Nurbek', 'Nysanbayev', 54000, '123-456-7898', 'Japan'),
    ('temirlan.suleimenov@example.com', 'Temirlan', 'Suleimenov', 60000, '123-456-7899', 'Australia'),
    ('batyrbek.aliev@example.com', 'Batyrbek', 'Aliev', 56000, '123-456-7800', 'Brazil'),
    ('murad.kassym@example.com', 'Murad', 'Kassym', 59000, '123-456-7801', 'Germany'),
    ('karina.ongar@example.com', 'Karina', 'Ongar', 62000, '123-456-7802', 'China'),
    ('diana.nurly@example.com', 'Diana', 'Nurly', 54000, '123-456-7803', 'France'),
    ('yerassyl.sagymbay@example.com', 'Yerassyl', 'Sagymbay', 61000, '123-456-7804', 'Kazakhstan'),
    ('marzhan.talassova@example.com', 'Marzhan', 'Talassova', 58000, '123-456-7805', 'Italy'),
    ('omirbek.askar@example.com', 'Omirbek', 'Askar', 62000, '123-456-7806', 'Mexico'),
    ('rauan.kazhim@example.com', 'Rauan', 'Kazhim', 57000, '123-456-7807', 'Canada'),
    ('meirzhan.nurtas@example.com', 'Meirzhan', 'Nurtas', 55000, '123-456-7808', 'Japan'),
    ('saule.yermakhan@example.com', 'Saule', 'Yermakhan', 60000, '123-456-7809', 'Australia');

-- Insert records into Doctor
INSERT INTO Doctor (email, degree) VALUES
    ('bakdaulet.dauletov@example.com', 'MD'),
    ('damir.alipbayev@example.com', 'PhD'),
    ('zhanara.bekbergenova@example.com', 'MD'),
    ('adele.dauletova@example.com', 'MD'),
    ('zhanel.dauletova@example.com', 'PhD'),
    ('gulshat.abdikaimova@example.com', 'MD'),
    ('guldana.zhubandyk@example.com', 'PhD'),
    ('aibek.narik@example.com', 'MD'),
    ('nurbek.nysanbayev@example.com', 'MD'),
    ('temirlan.suleimenov@example.com', 'PhD');

-- Insert records into PublicServant
INSERT INTO PublicServant (email, department) VALUES
    ('bakdaulet.dauletov@example.com', 'Health'),
    ('damir.alipbayev@example.com', 'Research'),
    ('zhanara.bekbergenova@example.com', 'Health'),
    ('adele.dauletova@example.com', 'Admin'),
    ('zhanel.dauletova@example.com', 'Health'),
    ('gulshat.abdikaimova@example.com', 'Admin'),
    ('guldana.zhubandyk@example.com', 'Research'),
    ('aibek.narik@example.com', 'Health'),
    ('nurbek.nysanbayev@example.com', 'Health'),
    ('temirlan.suleimenov@example.com', 'Research'),
    ('batyrbek.aliev@example.com', 'Health'),
    ('murad.kassym@example.com', 'Research'),
    ('karina.ongar@example.com', 'Health'),
    ('diana.nurly@example.com', 'Health'),
    ('yerassyl.sagymbay@example.com', 'Research');

-- Insert records into Patients
INSERT INTO Patients (email) VALUES
    ('marzhan.talassova@example.com'),
    ('omirbek.askar@example.com'),
    ('rauan.kazhim@example.com'),
    ('meirzhan.nurtas@example.com'),
    ('saule.yermakhan@example.com');

-- Insert records into Disease
INSERT INTO Disease (disease_code, pathogen, description, id) VALUES
    ('COVID-19', 'Virus', 'Coronavirus disease 2019', 11),
    ('Flu', 'Virus', 'Seasonal influenza', 2),
    ('Tuberculosis', 'Bacteria', 'Bacterial infection', 1),
    ('Malaria', 'Parasite', 'Parasitic infection', 4),
    ('Celiac', 'Autoimmune', 'Gluten intolerance', 7),
    ('Depression', 'Psychiatric', 'Mental health disorder', 10),
    ('Scurvy', 'Nutritional', 'Vitamin C deficiency', 9),
    ('BSE', 'Prion', 'Mad cow disease', 5),
    ('ALS', 'Degenerative', 'Motor neuron disease', 8),
    ('Herpes', 'Virus', 'Herpes simplex', 2);

-- Insert records into Discover
INSERT INTO Discover (cname, disease_code, first_enc_date) VALUES
    ('Brazil', 'COVID-19', '2020-01-15'),
    ('Germany', 'Flu', '1900-01-01'),
    ('China', 'Tuberculosis', '1882-03-24'),
    ('France', 'Malaria', '1902-04-02'),
    ('Kazakhstan', 'Celiac', '2015-05-06');


-- Insert records into PatientDisease
INSERT INTO PatientDisease (email, disease_code) VALUES
    ('marzhan.talassova@example.com', 'COVID-19'),
    ('omirbek.askar@example.com', 'Tuberculosis'),
    ('rauan.kazhim@example.com', 'Malaria'),
    ('meirzhan.nurtas@example.com', 'Celiac'),
    ('saule.yermakhan@example.com', 'Scurvy');

-- Insert records into Record
INSERT INTO Record (email, cname, disease_code, total_deaths, total_patients) VALUES
    ('bakdaulet.dauletov@example.com', 'Brazil', 'COVID-19', 5000, 100000),
    ('damir.alipbayev@example.com', 'Germany', 'Flu', 1000, 20000),
    ('zhanara.bekbergenova@example.com', 'China', 'Tuberculosis', 3000, 150000),
    ('zhanel.dauletova@example.com', 'France', 'Malaria', 1500, 50000),
    ('gulshat.abdikaimova@example.com', 'Kazakhstan', 'Celiac', 200, 5000),
    ('guldana.zhubandyk@example.com', 'Italy', 'BSE', 200, 10000),
    ('aibek.narik@example.com', 'Canada', 'Herpes', 500, 15000),
    ('nurbek.nysanbayev@example.com', 'Japan', 'Scurvy', 200, 10000);


-- Insert specializations, ensuring some doctors have more than two specializations
INSERT INTO Specialize (id, email) VALUES
    (1, 'bakdaulet.dauletov@example.com'),   -- Bacterial Diseases
    (2, 'bakdaulet.dauletov@example.com'),   -- Viral Diseases
    (3, 'bakdaulet.dauletov@example.com'),   -- Fungal Diseases
    
    (2, 'damir.alipbayev@example.com'),      -- Viral Diseases
    (4, 'damir.alipbayev@example.com'),      -- Parasitic Diseases
    (5, 'damir.alipbayev@example.com'),      -- Prion Diseases

    (6, 'zhanara.bekbergenova@example.com'), -- Genetic Disorders
    (7, 'zhanara.bekbergenova@example.com'), -- Autoimmune Diseases
    
    (8, 'adele.dauletova@example.com'),      -- Degenerative Diseases
    (10, 'adele.dauletova@example.com'),     -- Psychiatric Disorders
    (11, 'adele.dauletova@example.com'),     -- Infectious Diseases
    
    (1, 'zhanel.dauletova@example.com'),     -- Bacterial Diseases
    (3, 'zhanel.dauletova@example.com'),     -- Fungal Diseases
    
    (9, 'gulshat.abdikaimova@example.com'),  -- Nutritional Deficiencies
    (2, 'gulshat.abdikaimova@example.com'),  -- Viral Diseases
    
    (11, 'guldana.zhubandyk@example.com'),   -- Infectious Diseases
    (6, 'guldana.zhubandyk@example.com'),    -- Genetic Disorders
    
    (4, 'aibek.narik@example.com'),          -- Parasitic Diseases
    (9, 'aibek.narik@example.com');          -- Nutritional Deficiencies