# Sample API Using Go

Sample API project using go lang
<br\>
 using static sample data without Database connection
<br\>
https://golang.org/

<br>
+--------+----------------------+---------------------+------------------------------------------+------------------+------------+
| Verb   | Path                 |  Action             |  Body Data
+--------+----------------------+---------------------+------------------------------------------+------------------+------------+
| GET    | /person/{id}         | getPerson           |                                                                          |
| GET    | /people              | getPeople           |                                                                          |
| POST   | /person              | createPerson        |data{"id":"5","firstname":"john","lastname":"doe"}                        |
| DELETE | /person/{id}         | deletePerson        |                                                                          |
+--------+----------------------+---------------------+------------------------------------------+------------------+------------+
