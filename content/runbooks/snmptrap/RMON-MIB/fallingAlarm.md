---
title: fallingAlarm
description: Troubleshooting for SNMP trap fallingAlarm
#published: true
tags:
  - LGTM
  - generated
editor: markdown
---

# RMON-MIB::fallingAlarm 

The SNMP trap that is generated when an alarm
entry crosses its falling threshold and generates
an event that is configured for sending SNMP
traps. 


## Variables


  - alarmIndex
  - alarmVariable
  - alarmSampleType
  - alarmValue
  - alarmFallingThreshold 

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

**alarmFallingThreshold** 
: A threshold for the sampled statistic.  When the current
sampled value is less than or equal to this threshold,
and the value at the last sampling interval was greater than
this threshold, a single event will be generated.
A single event will also be generated if the first
sample after this entry becomes valid is less than or
equal to this threshold and the associated
alarmStartupAlarm is equal to fallingAlarm(2) or
risingOrFallingAlarm(3).
After a falling event is generated, another such event
will not be generated until the sampled value
rises above this threshold and reaches the
alarmRisingThreshold.
This object may not be modified if the associated
alarmStatus object is equal to valid(1). 


## Meaning

The RMON-MIB::fallingAlarm SNMP trap is generated when an alarm entry crosses its falling threshold and generates an event that is configured for sending SNMP traps. This trap indicates that a specific variable has fallen below a predetermined threshold, triggering an alarm event.

## Impact

The impact of this trap depends on the specific alarm configuration and the variable being monitored. However, in general, it may indicate a potential issue or anomaly in the system or network being monitored. The alarm event may trigger notifications, automated actions, or manual interventions to investigate and resolve the issue.

## Diagnosis

To diagnose the cause of this trap, follow these steps:

1. Identify the alarmIndex, which uniquely identifies the alarm entry.
2. Determine the alarmVariable, which is the object identifier of the variable being monitored.
3. Check the alarmSampleType to understand how the variable is being sampled (absoluteValue or deltaValue).
4. Review the alarmValue to see the current sampled value that triggered the falling alarm.
5. Verify the alarmFallingThreshold to ensure it is set correctly.
6. Check the system or network being monitored to identify the root cause of the alarm event.
7. Analyze the alarm history to determine if this is a one-time event or a recurring issue.

## Mitigation

To mitigate the issue, follow these steps:

1. Investigate and resolve the underlying cause of the alarm event.
2. Adjust the alarmFallingThreshold if necessary to reduce false positives or improve sensitivity.
3. Verify that the alarmVariable is correct and that the sampling method (alarmSampleType) is appropriate for the monitored variable.
4. Consider implementing automatic actions or scripts to respond to falling alarm events, such as sending notifications or executing remediation scripts.
5. Review and update the alarm configuration as needed to ensure it is aligned with the desired monitoring and notification policies.
---




