---
title: CassandraBadHacker
description: Troubleshooting for alert CassandraBadHacker
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# CassandraBadHacker

Increase of Cassandra authentication failures

<details>
  <summary>Alert Rule</summary>

{{% rule "cassandra/criteo-cassandra-exporter.yml" "CassandraBadHacker" %}}

{{% comment %}}

```yaml
alert: CassandraBadHacker
expr: rate(cassandra_stats{name="org:apache:cassandra:metrics:client:authfailure:count"}[1m]) > 5
for: 2m
labels:
    severity: warning
annotations:
    summary: Cassandra bad hacker (instance {{ $labels.instance }})
    description: |-
        Increase of Cassandra authentication failures
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/criteo-cassandra-exporter/CassandraBadHacker.md

```

{{% /comment %}}

</details>


## Meaning

The CassandraBadHacker alert is triggered when the rate of Cassandra authentication failures exceeds 5 attempts per minute, sustained for 2 minutes. This alert indicates a potential security threat, as it may indicate an attacker attempting to brute-force access to the Cassandra cluster.

## Impact

If left unchecked, a successful brute-force attack could lead to unauthorized access to sensitive data stored in Cassandra. This could result in:

* Data breaches or exfiltration
* Disruption of business operations
* Compliance and regulatory issues
* Reputation damage

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Cassandra logs for authentication failure events to identify the source IP address(es) and username(s) involved.
2. Verify that the Cassandra cluster is properly configured with strong passwords and secure authentication mechanisms.
3. Investigate the network traffic to identify any suspicious patterns or activities.
4. Check if there are any other security-related alerts or incidents reported around the same time.

## Mitigation

To mitigate the issue, take the following steps:

1. Immediately block the IP address(es) involved in the authentication attempts from accessing the Cassandra cluster.
2. Rotate the passwords for all Cassandra nodes and update the configuration accordingly.
3. Enable additional security measures, such as IP whitelisting or two-factor authentication.
4. Perform a thorough security audit of the Cassandra cluster and related systems to identify and address any vulnerabilities.
5. Consider implementing a Web Application Firewall (WAF) or an Intrusion Detection System (IDS) to detect and prevent similar attacks in the future.

Remember to update the runbook with any additional details or procedures specific to your environment.