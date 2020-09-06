# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [zeigarnik.proto](#zeigarnik.proto)
    - [CreateReminderReq](#reminder.v1alpha1.CreateReminderReq)
    - [CreateReminderRes](#reminder.v1alpha1.CreateReminderRes)
    - [DeleteReminderReq](#reminder.v1alpha1.DeleteReminderReq)
    - [DeleteReminderRes](#reminder.v1alpha1.DeleteReminderRes)
    - [GetReminderByIDReq](#reminder.v1alpha1.GetReminderByIDReq)
    - [GetReminderByIDRes](#reminder.v1alpha1.GetReminderByIDRes)
    - [ListRemindersReq](#reminder.v1alpha1.ListRemindersReq)
    - [ListRemindersRes](#reminder.v1alpha1.ListRemindersRes)
    - [Reminder](#reminder.v1alpha1.Reminder)
    - [UpdateReminderReq](#reminder.v1alpha1.UpdateReminderReq)
    - [UpdateReminderRes](#reminder.v1alpha1.UpdateReminderRes)
  
    - [ReminderStatus](#reminder.v1alpha1.ReminderStatus)
    - [ReminderType](#reminder.v1alpha1.ReminderType)
  
  
    - [ReminderService](#reminder.v1alpha1.ReminderService)
  

- [Scalar Value Types](#scalar-value-types)



<a name="zeigarnik.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## zeigarnik.proto



<a name="reminder.v1alpha1.CreateReminderReq"></a>

### CreateReminderReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| reminder | [Reminder](#reminder.v1alpha1.Reminder) |  |  |






<a name="reminder.v1alpha1.CreateReminderRes"></a>

### CreateReminderRes



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| reminder | [Reminder](#reminder.v1alpha1.Reminder) |  |  |






<a name="reminder.v1alpha1.DeleteReminderReq"></a>

### DeleteReminderReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="reminder.v1alpha1.DeleteReminderRes"></a>

### DeleteReminderRes



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| deleted | [bool](#bool) |  |  |






<a name="reminder.v1alpha1.GetReminderByIDReq"></a>

### GetReminderByIDReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="reminder.v1alpha1.GetReminderByIDRes"></a>

### GetReminderByIDRes



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| reminder | [Reminder](#reminder.v1alpha1.Reminder) |  |  |






<a name="reminder.v1alpha1.ListRemindersReq"></a>

### ListRemindersReq







<a name="reminder.v1alpha1.ListRemindersRes"></a>

### ListRemindersRes



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| reminders | [string](#string) | repeated |  |






<a name="reminder.v1alpha1.Reminder"></a>

### Reminder



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| created | [int64](#int64) |  |  |
| message | [string](#string) |  |  |
| to | [string](#string) |  |  |
| status | [ReminderStatus](#reminder.v1alpha1.ReminderStatus) |  |  |
| when | [int64](#int64) |  |  |
| type | [ReminderType](#reminder.v1alpha1.ReminderType) |  |  |
| warnAt | [int64](#int64) | repeated |  |






<a name="reminder.v1alpha1.UpdateReminderReq"></a>

### UpdateReminderReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| reminder | [Reminder](#reminder.v1alpha1.Reminder) |  |  |






<a name="reminder.v1alpha1.UpdateReminderRes"></a>

### UpdateReminderRes



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| reminder | [Reminder](#reminder.v1alpha1.Reminder) |  |  |





 


<a name="reminder.v1alpha1.ReminderStatus"></a>

### ReminderStatus


| Name | Number | Description |
| ---- | ------ | ----------- |
| UNKNOWN | 0 |  |
| CREATED | 1 |  |
| QUEUED | 2 |  |
| FIRED | 3 |  |
| MISSED | 4 |  |



<a name="reminder.v1alpha1.ReminderType"></a>

### ReminderType


| Name | Number | Description |
| ---- | ------ | ----------- |
| INVALID | 0 |  |
| AT | 1 |  |
| AFTER | 2 |  |


 

 


<a name="reminder.v1alpha1.ReminderService"></a>

### ReminderService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateReminder | [CreateReminderReq](#reminder.v1alpha1.CreateReminderReq) | [CreateReminderRes](#reminder.v1alpha1.CreateReminderRes) |  |
| GetReminder | [GetReminderByIDReq](#reminder.v1alpha1.GetReminderByIDReq) | [GetReminderByIDRes](#reminder.v1alpha1.GetReminderByIDRes) |  |
| UpdateReminder | [UpdateReminderReq](#reminder.v1alpha1.UpdateReminderReq) | [UpdateReminderRes](#reminder.v1alpha1.UpdateReminderRes) |  |
| DeleteReminder | [DeleteReminderReq](#reminder.v1alpha1.DeleteReminderReq) | [DeleteReminderRes](#reminder.v1alpha1.DeleteReminderRes) |  |

 



## Scalar Value Types

| .proto Type | Notes | C++ Type | Java Type | Python Type |
| ----------- | ----- | -------- | --------- | ----------- |
| <a name="double" /> double |  | double | double | float |
| <a name="float" /> float |  | float | float | float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long |
| <a name="bool" /> bool |  | bool | boolean | boolean |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str |

