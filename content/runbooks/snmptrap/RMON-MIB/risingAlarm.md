---
title: risingAlarm
description: Troubleshooting for SNMP trap risingAlarm
#published: true
tags:
  - LGTM
  - generated
editor: markdown
---

# RMON-MIB::risingAlarm 

The SNMP trap that is generated when an alarm
entry crosses its rising threshold and generates
an event that is configured for sending SNMP
traps. 


## Variables


  - alarmIndex
  - alarmVariable
  - alarmSampleType
  - alarmValue
  - alarmRisingThreshold 

### Definitions 


**alarmIndex** 
: An index that uniquely identifies an entry in the
alarm table.  Each such entry defines a
diagnostic sample at a particular interval
for an object on the device. 

**alarmVariable** 
: The object identifier of the particular variable to be
sampled.  Only variables that resolve to an ASN.1 primitive
type of INTEGER (INTEGER, Integer32, Counter32, Counter64,
Gauge, or TimeTicks) may be sampled.
Because SNMP access control is articulated entirely
in terms of the contents of MIB views, no access
control mechanism exists that can restrict the value of
this object to identify only those objects that exist
in a particular MIB view.  Because there is thus no
acceptable means of restricting the read access that
could be obtained through the alarm mechanism, the
probe must only grant write access to this object in
those views that have read access to all objects on
the probe.
During a set operation, if the supplied variable name is
not available in the selected MIB view, a badValue error
must be returned.  If at any time the variable name of
an established alarmEntry is no longer available in the
selected MIB view, the probe must change the status of
this alarmEntry to invalid(4).
This object may not be modified if the associated
alarmStatus object is equal to valid(1). 

**alarmSampleType** 
: The method of sampling the selected variable and
calculating the value to be compared against the
thresholds.  If the value of this object is
absoluteValue(1), the value of the selected variable
will be compared directly with the thresholds at the
end of the sampling interval.  If the value of this
object is deltaValue(2), the value of the selected
variable at the last sample will be subtracted from
the current value, and the difference compared with
the thresholds.
This object may not be modified if the associated
alarmStatus object is equal to valid(1). 

**alarmValue** 
: The value of the statistic during the last sampling
period.  For example, if the sample type is deltaValue,
this value will be the difference between the samples
at the beginning and end of the period.  If the sample
type is absoluteValue, this value will be the sampled
value at the end of the period.
This is the value that is compared with the rising and
falling thresholds.
The value during the current sampling period is not
made available until the period is completed and will
remain available until the next period completes. 

**alarmRisingThreshold** 
: A threshold for the sampled statistic.  When the current
sampled value is greater than or equal to this threshold,
and the value at the last sampling interval was less than
this threshold, a single event will be generated.
A single event will also be generated if the first
sample after this entry becomes valid is greater than or
equal to this threshold and the associated
alarmStartupAlarm is equal to risingAlarm(1) or
risingOrFallingAlarm(3).
After a rising event is generated, another such event
will not be generated until the sampled value
falls below this threshold and reaches the
alarmFallingThreshold.
This object may not be modified if the associated
alarmStatus object is equal to valid(1). 


## Meaning

The RMON-MIB::risingAlarm SNMP trap is generated when an alarm entry crosses its rising threshold and meets the conditions for sending SNMP traps. This trap is triggered when the sampled value of a monitored object exceeds the rising threshold, indicating a significant change in the object's value.

## Impact

Receiving this trap may indicate a potential issue with the monitored object or the system being monitored. The rising threshold crossing may signal an anomaly, error, or unexpected behavior that requires attention from network administrators or operators. Ignoring this trap may lead to undetected problems, performance degradation, or even system crashes.

## Diagnosis

To diagnose the issue, follow these steps:

1. Identify the alarmIndex, which uniquely identifies the entry in the alarm table.
2. Determine the object being monitored (alarmVariable) and its current value (alarmValue).
3. Verify the sampling method (alarmSampleType) and the rising threshold (alarmRisingThreshold) that was exceeded.
4. Analyze the historical data and recent trends to understand the cause of the rising threshold crossing.
5. Check the system logs and other monitoring tools for correlated events or errors.

## Mitigation

To mitigate the issue, consider the following steps:

1. Investigate and address the root cause of the rising threshold crossing.
2. Adjust the alarmRisingThreshold to a more suitable value, if necessary.
3. Implement additional monitoring or logging to provide more insights into the system behavior.
4. Consider setting up additional alarm entries for similar variables to provide early warnings for potential issues.
5. Develop and implement a corrective action plan to prevent similar events from occurring in the future.
---




