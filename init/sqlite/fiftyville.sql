-- Keep a log of any SQL queries you execute as you solve the mystery.

--The THIEF is:
--The city the thief ESCAPED TO:
--The ACCOMPLICE is:

.schema
.tables

CREATE TABLE crime_scene_reports (
    id INTEGER,
    year INTEGER,
    month INTEGER,
    day INTEGER,
    street TEXT,
    description TEXT,
    PRIMARY KEY(id)
);
CREATE TABLE interviews (
    id INTEGER,
    name TEXT,
    year INTEGER,
    month INTEGER,
    day INTEGER,
    transcript TEXT,
    PRIMARY KEY(id)
);
CREATE TABLE atm_transactions (
    id INTEGER,
    account_number INTEGER,
    year INTEGER,
    month INTEGER,
    day INTEGER,
    atm_location TEXT,
    transaction_type TEXT,
    amount INTEGER,
    PRIMARY KEY(id)
);
CREATE TABLE bank_accounts (
    account_number INTEGER,
    person_id INTEGER,
    creation_year INTEGER,
    FOREIGN KEY(person_id) REFERENCES people(id)
);
CREATE TABLE airports (
    id INTEGER,
    abbreviation TEXT,
    full_name TEXT,
    city TEXT,
    PRIMARY KEY(id)
);
CREATE TABLE flights (
    id INTEGER,
    origin_airport_id INTEGER,
    destination_airport_id INTEGER,
    year INTEGER,
    month INTEGER,
    day INTEGER,
    hour INTEGER,
    minute INTEGER,
    PRIMARY KEY(id),
    FOREIGN KEY(origin_airport_id) REFERENCES airports(id),
    FOREIGN KEY(destination_airport_id) REFERENCES airports(id)
);
CREATE TABLE passengers (
    flight_id INTEGER,
    passport_number INTEGER,
    seat TEXT,
    FOREIGN KEY(flight_id) REFERENCES flights(id)
);
CREATE TABLE phone_calls (
    id INTEGER,
    caller TEXT,
    receiver TEXT,
    year INTEGER,
    month INTEGER,
    day INTEGER,
    duration INTEGER,
    PRIMARY KEY(id)
);
CREATE TABLE people (
    id INTEGER,
    name TEXT,
    phone_number TEXT,
    passport_number INTEGER,
    license_plate TEXT,
    PRIMARY KEY(id)
);
CREATE TABLE bakery_security_logs (
    id INTEGER,
    year INTEGER,
    month INTEGER,
    day INTEGER,
    hour INTEGER,
    minute INTEGER,
    activity TEXT,
    license_plate TEXT,
    PRIMARY KEY(id)
);

-- to see who sitnessed theft
SELECT * FROM crime_scene_reports WHERE description LIKE "%Theft%";
-- three witness at humphrey street 10:15am we'll see in interview
-- interview mentions bakery
-- id 295
SELECT * FROM crime_scene_reports WHERE month=7 AND day=28 AND street="Humphrey Street";


SELECT * FROM interviews WHERE month=7 AND day=28;
-- people mentioned bakery
-- eugene id162  -> recognizes thief
-- raymond id163
    -- heards -> thief to take earliest flight tomorrow (month= 7 day=29)
-- ruth id161  thief went in car , watch security footage

SELECT * FROM people where name="Eugene";
 280744 | Eugene | (666) 555-5774 | 9584465633      | 47592FJ


-- final earch by password number gotten from flights
SELECT name
  FROM people
 WHERE passport_number IN
        (SELECT passport_number
           FROM passengers
          WHERE flight_id IN
                (SELECT id
                   FROM flights
                  WHERE year = 2021
                    AND month = 7
                    AND day = 29
                    AND origin_airport_id IN
                        (SELECT id
                           FROM airports
                          WHERE city = "Fiftyville")
                  ORDER BY hour
                  LIMIT 1))
 ORDER BY name;


 SELECT name
  FROM people
 WHERE phone_number IN
       (SELECT caller
          FROM phone_calls
         WHERE year = 2021
           AND month = 7
           AND day = 28
           AND duration < 60)
 ORDER BY name;