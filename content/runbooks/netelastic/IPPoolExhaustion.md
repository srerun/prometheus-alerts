---
title: IPPoolExhaustion
description: Troubleshooting for alert IPPoolExhaustion
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# IPPoolExhaustion

IP pool usage for group {{ $labels.vbng_groupName }} is above 90% for more than 10 minutes (current usage: {{ $value }}%). This may prevent new users from connecting.

<details>
  <summary>Alert Rule</summary>

{{% rule "netelastic/netelastic.yml" "IPPoolExhaustion" %}}

{{% comment %}}

```yaml
alert: IPPoolExhaustion
expr: (vbng_pool_ip_used_total / vbng_pool_ip_max) * 100 > 90
for: 10m
labels:
    severity: critical
annotations:
    summary: IP pool nearing exhaustion
    description: 'IP pool usage for group {{ $labels.vbng_groupName }} is above 90% for more than 10 minutes (current usage: {{ $value }}%). This may prevent new users from connecting.'
    runbook: https://srerun.github.io/prometheus-alerts/runbooks/netelastic/ippoolexhaustion/

```

{{% /comment %}}

</details>


## Meaning

The IPPoolExhaustion alert is triggered when the IP pool usage for a specific group exceeds 90% for more than 10 minutes. This alert indicates that the IP pool is nearing exhaustion, which may prevent new users from connecting.

## Impact

If this alert is not addressed, it can lead to:

* New users unable to connect to the network
* Existing users experiencing connectivity issues
* Increased latency and packet loss
* Potential revenue loss due to service unavailability

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the `vbng_pool_ip_used_total` and `vbng_pool_ip_max` metrics to understand the current IP pool usage.
2. Identify the specific VBNG group that is experiencing high IP pool usage using the `vbng_groupName` label.
3. Verify that the IP pool usage has been above 90% for more than 10 minutes using the Prometheus console or another monitoring tool.
4. Check the network infrastructure for any issues that may be contributing to the high IP pool usage, such as misconfigured DHCP settings or high traffic volumes.

## Mitigation

To mitigate the issue, follow these steps:

1. Investigate the root cause of the high IP pool usage and address any underlying issues.
2. Consider increasing the size of the IP pool to accommodate the current demand.
3. Implement rate limiting or traffic policing to reduce the load on the network and prevent further IP pool exhaustion.
4. Consider migrating users to a different VBNG group or network segment to reduce the load on the affected group.
5. Monitor the IP pool usage closely and adjust the mitigation strategy as needed to prevent further exhaustion.

Note: Refer to the provided runbook URL for more detailed steps and procedures for mitigating the IPPoolExhaustion alert.