# goscores

A simple utility for fetching sports scores from the CLI.

## Development
### Requirements
- `task`: `go install github.com/go-task/task/v3/cmd/task@latest`
- `go` v1.24+


### Building
```bash
$ task build
```

### Running
```bash
$ goscores                                                                                                            main 
2025/03/30 15:32:20 INFO starting goscores with defaults gamedate=2025330 sport=hockey league=nhl
2025/03/30 15:32:20 DEBUG fetching scores provider=ESPN
2025/03/30 15:32:21 DEBUG finished fetching scores
+------+------+-----------+----------------+-------+-----------------+
| HOME | AWAY |    STATUS | PERIOD / START | CLOCK | SCORE           |
+------+------+-----------+----------------+-------+-----------------+
| FLA  | MTL  |   Ongoing | 3              | 4:25  |  FLA 2 -  MTL 3 |
| WSH  | BUF  |   Ongoing | 1              | 7:46  |  WSH 1 -  BUF 0 |
| WPG  | VAN  |   Ongoing | 1              | 6:02  |  WPG 0 -  VAN 0 |
| CHI  | UTAH | Scheduled | 4:00PM         | 0:00  |  CHI 0 - UTAH 0 |
| PIT  | OTT  | Scheduled | 5:00PM         | 0:00  |  PIT 0 -  OTT 0 |
| CAR  | NYI  | Scheduled | 5:00PM         | 0:00  |  CAR 0 -  NYI 0 |
| ANA  | TOR  | Scheduled | 8:00PM         | 0:00  |  ANA 0 -  TOR 0 |
| LA   | SJ   | Scheduled | 10:00PM        | 0:00  |   LA 0 -   SJ 0 |
+------+------+-----------+----------------+-------+-----------------+
2025/03/30 15:32:21 DEBUG exiting
```

## Links
- [ESPN API description](https://gist.github.com/akeaswaran/b48b02f1c94f873c6655e7129910fc3b) from [akeaswaran](https://github.com/akeaswaran)
- [json-to-go](https://mholt.github.io/json-to-go/): Helpful tool to emit a Go struct from a JSON document
