# Music 504

Pushing Last.fm stats for 504 into a telegram channel.

## Cron configuration

Cron config to run the job:
```
0 12 * * FRI /home/uladbohdan/music504 --prod --time_period=twoWeeks --enable_tips
0 0 1 * * /home/uladbohdan/music504 --prod --time_period=oneMonth
0 0 1 APR,JUL,OCT,JAN * /home/uladbohdan/music504 --prod --time_period=threeMonths
0 0 1 JUL,JAN * /home/uladbohdan/music504 --prod --time_period=sixMonths
0 0 1 JAN * /home/uladbohdan/music504 --prod --time_period=oneYear
```

Example of cron config to run debug:
```
7 3,18 * * * /home/uladbohdan/music504 --time_period=oneWeek --enable_tips
```

Run `crontab -e` to edit this on your machine.
