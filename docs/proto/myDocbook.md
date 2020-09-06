zeigarnik.proto
===============

CreateReminderReq {#reminder.v1alpha1.CreateReminderReq}
-----------------

  Field      Type                                      Label   Description
  ---------- ----------------------------------------- ------- -------------
  reminder   [Reminder](#reminder.v1alpha1.Reminder)           

  : `CreateReminderReq` Fields

CreateReminderRes {#reminder.v1alpha1.CreateReminderRes}
-----------------

  Field      Type                                      Label   Description
  ---------- ----------------------------------------- ------- -------------
  reminder   [Reminder](#reminder.v1alpha1.Reminder)           

  : `CreateReminderRes` Fields

DeleteReminderReq {#reminder.v1alpha1.DeleteReminderReq}
-----------------

  Field   Type                Label   Description
  ------- ------------------- ------- -------------
  id      [string](#string)           

  : `DeleteReminderReq` Fields

DeleteReminderRes {#reminder.v1alpha1.DeleteReminderRes}
-----------------

  Field     Type            Label   Description
  --------- --------------- ------- -------------
  deleted   [bool](#bool)           

  : `DeleteReminderRes` Fields

GetReminderByIDReq {#reminder.v1alpha1.GetReminderByIDReq}
------------------

  Field   Type                Label   Description
  ------- ------------------- ------- -------------
  id      [string](#string)           

  : `GetReminderByIDReq` Fields

GetReminderByIDRes {#reminder.v1alpha1.GetReminderByIDRes}
------------------

  Field      Type                                      Label   Description
  ---------- ----------------------------------------- ------- -------------
  reminder   [Reminder](#reminder.v1alpha1.Reminder)           

  : `GetReminderByIDRes` Fields

ListRemindersReq {#reminder.v1alpha1.ListRemindersReq}
----------------

ListRemindersRes {#reminder.v1alpha1.ListRemindersRes}
----------------

  Field       Type                Label      Description
  ----------- ------------------- ---------- -------------
  reminders   [string](#string)   repeated   

  : `ListRemindersRes` Fields

Reminder {#reminder.v1alpha1.Reminder}
--------

  Field     Type                                                  Label      Description
  --------- ----------------------------------------------------- ---------- -------------
  id        [string](#string)                                                
  created   [int64](#int64)                                                  
  message   [string](#string)                                                
  to        [string](#string)                                                
  status    [ReminderStatus](#reminder.v1alpha1.ReminderStatus)              
  when      [int64](#int64)                                                  
  type      [ReminderType](#reminder.v1alpha1.ReminderType)                  
  warnAt    [int64](#int64)                                       repeated   

  : `Reminder` Fields

UpdateReminderReq {#reminder.v1alpha1.UpdateReminderReq}
-----------------

  Field      Type                                      Label   Description
  ---------- ----------------------------------------- ------- -------------
  reminder   [Reminder](#reminder.v1alpha1.Reminder)           

  : `UpdateReminderReq` Fields

UpdateReminderRes {#reminder.v1alpha1.UpdateReminderRes}
-----------------

  Field      Type                                      Label   Description
  ---------- ----------------------------------------- ------- -------------
  reminder   [Reminder](#reminder.v1alpha1.Reminder)           

  : `UpdateReminderRes` Fields

ReminderStatus {#reminder.v1alpha1.ReminderStatus}
--------------

  Name      Number   Description
  --------- -------- -------------
  UNKNOWN   0        
  CREATED   1        
  QUEUED    2        
  FIRED     3        
  MISSED    4        

  : `ReminderStatus` Values

ReminderType {#reminder.v1alpha1.ReminderType}
------------

  Name      Number   Description
  --------- -------- -------------
  INVALID   0        
  AT        1        
  AFTER     2        

  : `ReminderType` Values

ReminderService {#reminder.v1alpha1.ReminderService}
---------------

  Method Name      Request Type                                                  Response Type                                                 Description
  ---------------- ------------------------------------------------------------- ------------------------------------------------------------- -------------
  CreateReminder   [CreateReminderReq](#reminder.v1alpha1.CreateReminderReq)     [CreateReminderRes](#reminder.v1alpha1.CreateReminderRes)     
  GetReminder      [GetReminderByIDReq](#reminder.v1alpha1.GetReminderByIDReq)   [GetReminderByIDRes](#reminder.v1alpha1.GetReminderByIDRes)   
  UpdateReminder   [UpdateReminderReq](#reminder.v1alpha1.UpdateReminderReq)     [UpdateReminderRes](#reminder.v1alpha1.UpdateReminderRes)     
  DeleteReminder   [DeleteReminderReq](#reminder.v1alpha1.DeleteReminderReq)     [DeleteReminderRes](#reminder.v1alpha1.DeleteReminderRes)     

  : `ReminderService` Methods

Scalar Value Types
==================

  .proto Type   Notes                                                                                                                                              C++ Type   Java Type    Python Type
  ------------- -------------------------------------------------------------------------------------------------------------------------------------------------- ---------- ------------ -------------
  double                                                                                                                                                           double     double       float
  float                                                                                                                                                            float      float        float
  int32         Uses variable-length encoding. Inefficient for encoding negative numbers -- if your field is likely to have negative values, use sint32 instead.   int32      int          int
  int64         Uses variable-length encoding. Inefficient for encoding negative numbers -- if your field is likely to have negative values, use sint64 instead.   int64      long         int/long
  uint32        Uses variable-length encoding.                                                                                                                     uint32     int          int/long
  uint64        Uses variable-length encoding.                                                                                                                     uint64     long         int/long
  sint32        Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s.                               int32      int          int
  sint64        Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s.                               int64      long         int/long
  fixed32       Always four bytes. More efficient than uint32 if values are often greater than 2\^28.                                                              uint32     int          int
  fixed64       Always eight bytes. More efficient than uint64 if values are often greater than 2\^56.                                                             uint64     long         int/long
  sfixed32      Always four bytes.                                                                                                                                 int32      int          int
  sfixed64      Always eight bytes.                                                                                                                                int64      long         int/long
  bool                                                                                                                                                             bool       boolean      boolean
  string        A string must always contain UTF-8 encoded or 7-bit ASCII text.                                                                                    string     String       str/unicode
  bytes         May contain any arbitrary sequence of bytes.                                                                                                       string     ByteString   str
