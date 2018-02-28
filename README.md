# CmusNotify
Notifications from cmus music player.

This is a simple Golang project that creates notifications on your system when a song in cmus changes.

### How to compile

Just type
```go build```
in the same folder where you downloaded the file and the executable should appear.
You must have Go installed on your system.

### How to use
In cmus, type:
```
:set status_display_program=/path/CmusNotify
```
and press enter. CmusNotify is the executable that was created in the previous step.

If you have more than one status display program, add a comma to separate them.

#### Example notification

![Notification](https://github.com/sebojanko/CmusNotify/blob/master/cmusNotify-example.png)
