
# Dint

A port of a minimalistic library to handle dates as easy as you handle simple integers [Dint](https://github.com/konmik/dint/).

Dint is an integer which contains date in a human-readable integer format: YYYYMMDD.

Example: 20230912 = September 12, 2023

Benefits are obvious:
* Dint does not have timezone information, it does not have time information,
  all operations are extremely simple. Your user will never get a wrong date because
  his timezone changed.
* Storage in a database is trivial.
* Easy to debug (in a human readable format).
* The entire library size is 2Kb. There is just a couple of math functions!

### Usage

```java
int dint = Dint.create(2014, 9, 12);  // creates a dint
Integer.toString(dint);               // "20140912" - human-readable

Dint.addDays(dint, 199);              // adds 199 days to a dint

if (dint1 < dint2)                    // Simple comparison

int daysBetween = Dint.diff(20150912, 20140912); // = 365

Date asDate = Dint.toDate(dint);      // Simple conversion

Dint.year(20140912);  // 2014
Dint.month(20140912); // 9
Dint.day(20140912);   // 12

Dint.compose(2014, 13, 12);       // 20150112!
Dint.compose(2013, 9, 12 + 365);  // 20140912!

Dint.addMonths(20140930, 1);       // 20141030
Dint.addMonthsExtend(20140930, 1); // 20141031 keeps the last day of month!
```


