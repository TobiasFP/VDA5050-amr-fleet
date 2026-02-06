# MQTT Topics (VDA5050)

BotNana uses the standard VDA5050 topic hierarchy:

```
<prefix>/<manufacturer>/<serialNumber>/<topic>
```

## Prefix
The default prefix is `vda5050`. You can override this in `config/*.yaml`:

```
mqttTopicPrefix: vda5050
```

## Topics
Inbound (subscribed):
- `vda5050/+/+/state`
- `vda5050/+/+/connection`

Outbound (published):
- `vda5050/<manufacturer>/<serialNumber>/order`
- `vda5050/<manufacturer>/<serialNumber>/instantActions`

## Payload requirements
For outbound publishing, `manufacturer` and `serialNumber` must be present in the payload so BotNana can build the topic.
