# Music 504

Pushing Last.fm stats for 504 into a telegram channel.

## Cron configuration

Cron config to run the job:
```
0 12 * * FRI /home/uladbohdan/music504 --time_period=twoWeeks
0 0 1 * * /home/uladbohdan/music504 --time_period=oneMonth
0 0 1 APR,JUL,OCT,JAN * /home/uladbohdan/music504 --time_period=threeMonths
0 0 1 JAN * /home/uladbohdan/music504 --time_period=oneYear
```

Example of cron config to run debug:
```
7 3,6,9,12,15,18,21 * * * /home/uladbohdan/music504 --debug --time_period=oneWeek
```

Run `crontab -e` to edit this on your machine.
