# ID_Generator

**Main purpose of this program is to train my skills with golang. I want to make casual program which takes
from user's input (or generate it randomly) date of birth and then create whole ID on it.
Everything should be as standarized and randomized as possible. On the other hand
i'll try to use Git to push every change and maybe even write changelog.**

I am going to use official names database from: https://dane.gov.pl/pl/dataset/1667,lista-imion-wystepujacych-w-rejestrze-pesel-osoby-zyjace

**PESEL** is unique ID number that each citizen of Poland has
It's string of 11 numbers based on date of birth, sex and control number

_YYMMDDSSSSC_ - **Y** stands for year, **M** for month, **D** for day, **S** for sex*, and **C** is control number**

_*Sex part of PESEL is described well, but if it's even it means that the owner is a woman, and if number is odd it stands for a man._

_**Control number is a result of mathematic function. Evey number in PESEL has its own wage, pattern looks like this: 1-3-7-9-1-3-7-9-1-3
to get Control number each number has to be multiplicated by corresponding wage. Then all of the results has te be added.
Finally we have to take ones digit and substract it from the 10. That's how we get it._

## Changelog

### v. 1.0 
#### 24.07.2022
* Whole program is working properly. 

### v. 1.1
#### 20.08.2022
* New identity now exists as a struct.
* *  Struct has a json tags (for future reasons)
* Code cleanup