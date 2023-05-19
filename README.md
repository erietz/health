The server in my closes sometimes goes off the rails. I'd like a cronjob to
email me once a day with a health check to ensure things are running smoothly.

I want something like this

```
  System information as of Thu May 18 06:48:58 PM CDT 2023

  System load:                      1.28662109375
  Usage of /:                       8.1% of 456.34GB
  Memory usage:                     7%
  Swap usage:                       0%
  Temperature:                      70.0 C
  Processes:                        278
  Users logged in:                  0
```

but without using a bunch of shell scripts to get there. It would be cool if
the tool had a flag to generate html or json output. Here we go...
