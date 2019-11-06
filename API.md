# API Documentation

## Config
```
GET /config
```
Get the config in JSON format.


```
PUT /config
```
Update the config with the specified JSON document in the request body.

## Channels
```
PUT /select/{channel}
```
Select the specified channel.


## Volume
```
PUT /volume/{volume}
```
Set the volume to the specified value, which must be in the 0 - 100 range.


```
PUT /mute/{true|false}
```
Mute or unmute the radio. Does not change the value for the volume.

## Logs

```
GET /logs
```
Get the logs in JSON format.

```
GET /logs?simple=1
```
Get the logs in human-readable format.
