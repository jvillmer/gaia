# Model
model:
  rest_name: export
  resource_name: export
  entity_name: Export
  package: yuna
  group: core
  description: Export the policies and related objects in a given namespace.

# Attributes
attributes:
  v1:
  - name: APIVersion
    description: APIVersion of the api used for the exported data.
    type: integer
    exposed: true
    read_only: true
    autogenerated: true

  - name: data
    description: List of all exported data.
    type: external
    exposed: true
    subtype: map_of_string_of_lists_of_maps_of_string_of_objects
    autogenerated: true

  - name: identities
    description: The list of identities to export.
    type: list
    exposed: true
    subtype: string
    stored: true

  - name: label
    description: |-
      Label allows to define a unique label for this export. When importing the
      content of the export, this label will be added as a tag that will be used to
      recognize imported object in a later import.
    type: string
    exposed: true
    read_only: true
    autogenerated: true
