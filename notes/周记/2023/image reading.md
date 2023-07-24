First, ask bing or bard to read the image, tell me what the image is:
![](https://pbs.twimg.com/media/F1xYLoCaMAUMav-?format=jpg&name=medium)
Then, let them give the mermaid code or ascii images.
Here is the result:
## Bing
``` ascii
+----------------+       +----------------+       +----------------------+       +-------------------+       +---------------+       +------------+
| DATA PULLER    |       | DATA VALIDATOR |       | DATA TRANSFORMATION |       | CALCULATION ENGINE|       | REPORT WRITER |       | CONTROLLER |
+----------------+       +----------------+       +----------------------+       +-------------------+       +---------------+       +------------+
  |                 \         /               \         /                 \         /               \         /                 \         /             |
  |                  \       /                 \       /                   \       /                 \       /                   \       /               |
  |                   \     /                   \     /                     \     /                   \     /                     \     /                 |
  |                    \   /                     \   /                       \   /                     \   /                       \   /                   |
  v                     v v                       v v                         v v                       v v                         v v                     v
+----------------+    +----------------+        +----------------+          +----------------------+  +---------------+          +------------+
| Extract data   |--->| Check data     |--->| Flag errors      |--->| Transform data      |--->| Perform        |--->| Create reports |--->| Manage system |
| from sources   |    | quality and    |    | or anomalies     |    | into standard format|    | calculations   |    | and dashboards |    | and execution |
+----------------+    | integrity      |    +----------------+    | and structure      |    | and analysis   |    +---------------+    +------------+
                      +----------------+                          +----------------------+    +-------------------+
```

## Bard
``` mermaid
sequenceDiagram
  participant ANSICAL SECURITY
  participant DATA PILL SOVICE
  participant DATA REDE
  participant PORTO SATTE TRAP
  participant STOP
  participant EVENT
  participant STRE
  participant ·STOPPED EVENT
  participant DATA VAL
  participant DATA REMBER
  participant DATA TRANSPORTATION INTEGRITY VALIDATION
  participant FAIL EVENT
  participant COMPUTE EVENT
  participant rejects
  participant F
  participant CALCULATION
  participant смт
  participant ENGINE
  participant EVONT
  participant 07
  participant COMPLETE EVENT
  participant FAILURE REPORT
  participant REPORT LIRITER

  ANSICAL SECURITY->>DATA PILL SOVICE: Start
  DATA PILL SOVICE->>DATA REDE: Send data
  DATA REDE->>PORTO SATTE TRAP: Validate data
  PORTO SATTE TRAP->>STOP: If data is invalid, stop
  STOP->>EVENT: Generate event
  EVENT->>STRE: Store event
  STRE->>·STOPPED EVENT: If event is a stop event, store as stopped event
  ·STOPPED EVENT->>DATA VAL: Validate data
  DATA VAL->>DATA REMBER: Remember data
  DATA REMBER->>DATA TRANSPORTATION INTEGRITY VALIDATION: Validate data transportation
  DATA TRANSPORTATION INTEGRITY VALIDATION->>FAIL EVENT: If data is not valid, generate fail event
  FAIL EVENT->>COMPUTE EVENT: Compute data
  COMPUTE EVENT->>rejects: If data is invalid, reject
  rejects->>F: Send reject flag
  F->>CALCULATION: Calculate data
  CALCULATION->>смт: Store data in смт
  смт->>ENGINE: Send data to engine
  ENGINE->>EVONT: Generate event
  EVONT->>07: Add 7 to event
  07->>COMPLETE EVENT: Generate complete event
  COMPLETE EVENT->>FAILURE REPORT: Generate failure report
  FAILURE REPORT->>REPORT LIRITER: Send report to report literator

```

``` mermaid
sequenceDiagram
  participant DataPillService
  participant DataRedundancy
  participant DataIntegrityValidation
  participant FailureReport

  DataPillService ->> DataRedundancy: Collect data from user or other sources
  DataRedundancy ->> DataIntegrityValidation: Store data in multiple places
  DataIntegrityValidation ->> DataPillService: Check data for accuracy and completeness
  DataPillService ->> FailureReport: Generate report if data processing fails
  FailureReport ->> DataPillService: Include information about the error
  FailureReport ->> DataPillService: Help to identify and fix the problem

```





