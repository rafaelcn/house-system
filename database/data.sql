-- File containing data to be tested on the database

INSERT INTO TB_USER_TYPE VALUES(1, 'admin');
INSERT INTO TB_USER_TYPE VALUES(2, 'user');

INSERT INTO TB_USER VALUES(default, 'Rafael Campos Nunes',
                           'rcamposnunes@outlook.com',
                           'rafael.nunes',
                           'different-password',
                           '45999215452',
                           '1996-07-15',
                           1);

INSERT INTO TB_USER VALUES(default, 'Rodrigo Beloni',
						   'rodrigo@gmail.com',
						   'rodrigo.beloni',
                           'password',
                           '61992901458',
                           '1996-07-15',
                           2);

INSERT INTO TB_HOUSE VALUES(default, 'Rafaels House',
                            7);

INSERT INTO TB_OBJECT_TYPE VALUES(1, 'Light');
INSERT INTO TB_OBJECT_TYPE VALUES(2, 'Sound');
INSERT INTO TB_OBJECT_TYPE VALUES(3, 'Sensor');
INSERT INTO TB_OBJECT_TYPE VALUES(4, 'AirConditioner');

-- Types of objects:
-- Light = 1
-- Sound = 2
-- Sensor = 3
-- Air Conditioner = 4

INSERT INTO TB_OBJECT (OBJECT_ID,
                       OBJECT_NAME,
                       OBJECT_STATUS,
                       OBJECT_TYPE,
                       OBJECT_HOUSE,
                       OBJECT_ATTR_INTENSITY)
VALUES('1abc212fd212',
       'Light Room',
       false,
       1,
       2,
       100);

INSERT INTO TB_OBJECT (OBJECT_ID,
                       OBJECT_NAME,
                       OBJECT_STATUS,
                       OBJECT_TYPE,
                       OBJECT_HOUSE,
                       OBJECT_ATTR_INTENSITY)
VALUES('1abc212fd21123',
       'Light Bathroom',
       false,
       1,
       2,
       100);

INSERT INTO TB_OBJECT (OBJECT_ID,
                       OBJECT_NAME,
                       OBJECT_STATUS,
                       OBJECT_TYPE,
                       OBJECT_HOUSE,
                       OBJECT_ATTR_DISTANCE)
VALUES('1abc212f1643212',
       'Door sensor',
       false,
       3,
       2,
       3);

INSERT INTO TB_OBJECT (OBJECT_ID,
                       OBJECT_NAME,
                       OBJECT_STATUS,
                       OBJECT_TYPE,
                       OBJECT_HOUSE,
                       OBJECT_ATTR_VOLUME)
VALUES('1abc212f1643432',
       'Stereo sound',
       false,
       2,
       2,
       0.0);