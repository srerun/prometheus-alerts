---
title: PostgresqlCommitRateLow
description: Troubleshooting for alert PostgresqlCommitRateLow
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PostgresqlCommitRateLow

Postgresql seems to be processing very few transactions

<details>
  <summary>Alert Rule</summary>

{{% rule "postgresql/postgres-exporter.yml" "PostgresqlCommitRateLow" %}}

{{% comment %}}

```yaml
alert: PostgresqlCommitRateLow
expr: rate(pg_stat_database_xact_commit[1m]) < 10
for: 2m
labels:
    severity: critical
annotations:
    summary: Postgresql commit rate low (instance {{ $labels.instance }})
    description: |-
        Postgresql seems to be processing very few transactions
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/postgres-exporter/PostgresqlCommitRateLow.md

```

{{% /comment %}}

</details>

## **Meaning**
This alert triggers when the rate of committed transactions in a PostgreSQL database instance drops below 10 transactions per second, as averaged over a 1-minute period. This indicates that the database might not be processing transactions as expected, which could be a sign of reduced activity, resource contention, or a broader issue in the application stack.


## Impact
[//]: # "What could / will happen if the alert is not addressed"

A low commit rate can have several implications, including:
- Reduced application responsiveness or downtime, as critical transactions may not be completed.
- Potential user-facing errors if the database is unable to process transactions efficiently.
- Decreased system throughput, which could impact dependent services or systems.


## **Diagnosis**
### **Step 1: Confirm the Alert**
1. Review the alert details for the affected instance (e.g., `{{ $labels.instance }}`).
2. Check the `VALUE` of the commit rate to verify if it is indeed below the threshold of 10 commits/sec.

### **Step 2: Assess Database Health**
1. **Query Database Metrics:**
   - Use `pg_stat_database` to check for active and idle connections:
     ```sql
     SELECT datname, numbackends, xact_commit, xact_rollback, deadlocks
     FROM pg_stat_database;
     ```
   - Look for anomalies like a high number of idle connections or deadlocks.

2. **Check Logs:**
   - Examine PostgreSQL logs for errors or warnings related to transactions, connections, or resource contention.

3. **Resource Utilization:**
   - Monitor CPU, memory, and disk I/O usage on the database server.
   - Identify any contention or resource saturation.

4. **Network Issues:**
   - Verify network connectivity and latency between the database and application servers.

### **Step 3: Assess Application Activity**
1. Check if there is a sudden drop in incoming requests to the application that depend on the database.
2. Validate that application servers are operational and properly communicating with the database.
3. Ensure no recent configuration changes or deployments might have affected the database usage.

---

## **Mitigation**
### **Immediate Actions**
1. **Restart Connections:**
   - Restart problematic connections or the application to recover from transient issues.

2. **Resource Allocation:**
   - Allocate additional resources to the database server if resource constraints are identified.

3. **Throttle Traffic:**
   - Temporarily throttle incoming requests to reduce load on the database while investigating further.

### **Root Cause Resolution**
1. **Fix Deadlocks or Contention:**
   - Identify and resolve deadlocks or long-running transactions using tools like `pg_stat_activity`:
     ```sql
     SELECT * FROM pg_stat_activity WHERE state = 'active';
     ```

2. **Address Application Issues:**
   - Work with application teams to debug and resolve any issues causing reduced transaction activity.

3. **Database Tuning:**
   - Optimize queries, indexes, and configurations to ensure efficient transaction processing.

4. **Investigate External Factors:**
   - Check for upstream or downstream system issues that might be causing a drop in transaction rates.

### **Post-Mitigation**
1. Verify that the commit rate exceeds the threshold of 10 commits/sec.
2. Monitor for recurring alerts to ensure stability.
3. Document findings and any permanent fixes implemented.

---

## **Escalation**
If the issue persists or the root cause cannot be identified:
1. Escalate to the database administration team.
2. Engage application development teams to investigate potential application-level issues.
3. Consult infrastructure teams if resource or network-related issues are suspected.

---

## **References**
- PostgreSQL Documentation: https://www.postgresql.org/docs/
- `pg_stat_database` Reference: https://www.postgresql.org/docs/current/monitoring-stats.html
- Troubleshooting Deadlocks: https://www.postgresql.org/docs/current/runtime-config-locks.html

---

