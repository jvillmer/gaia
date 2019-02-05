# Model
model:
  rest_name: trigger
  resource_name: triggers
  entity_name: Trigger
  package: sephiroth
  group: integration/automation
  description: Trigger can be used to remotely trigger an automation.

# Attributes
attributes:
  v1:
  - name: payload
    description: Payload contains the eventual remote POST payload.
    type: string
    autogenerated: true
