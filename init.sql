DROP TABLE IF EXISTS File;
DROP TABLE IF EXISTS Submission;
DROP TABLE IF EXISTS Link;

create table Link (
    HELinkUuid uuid primary key,
    created TIMESTAMP DEFAULT now()
);

create table Submission (
	LinkUuid uuid references Link,
	SubmissionUuid uuid primary key,
	fname text,
    lname text,
    created TIMESTAMP DEFAULT now()
);

create table File (
    SubmissionUuid uuid primary key REFERENCES Submission (SubmissionUuid),
    Text text,
    status text,
    created TIMESTAMP DEFAULT now()
);


-- inserts
/*
insert into HELink (HELinkUuid) values ('cfc7103d-84e8-4190-898c-47b61f50296f');
insert into HELink (HELinkUuid) values ('8f426c72-88ab-4efc-866d-2f947dd6c1ad');

insert into HE (HELinkUuid, HeUuid, fname, lname, status) values ('cfc7103d-84e8-4190-898c-47b61f50296f', '891c0d40-de5e-4cae-bc6b-24a98e059502', 'Hansi', 'Hansman', 'Submitted');
insert into HE (HELinkUuid, HeUuid, fname, lname, status) values ('cfc7103d-84e8-4190-898c-47b61f50296f', '97c1629e-7053-4a58-89c2-7221dbf0a7ff', 'Heinrich', 'Meier', 'Submitted');
insert into HE (HELinkUuid, HeUuid, fname, lname, status) values ('cfc7103d-84e8-4190-898c-47b61f50296f', '78fb1632-a42b-49a0-a371-95238a6e0761', 'Axel', 'Schweiß', 'Submitted');

insert into HE (HELinkUuid, HeUuid, fname, lname, status) values ('8f426c72-88ab-4efc-866d-2f947dd6c1ad', '387977ad-ca6d-4df6-aaf2-ab4a357d6b83', 'Hansi', 'Hansman', 'Corrected');
insert into HE (HELinkUuid, HeUuid, fname, lname, status) values ('8f426c72-88ab-4efc-866d-2f947dd6c1ad', 'b050065d-6056-425c-93ef-37b22908254c', 'Hansi', 'Hansman', 'Corrected');

*/
