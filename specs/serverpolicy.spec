{
    "attributes": [
        {
            "allowed_chars": null,
            "allowed_choices": null,
            "autogenerated": null,
            "channel": null,
            "creation_only": null,
            "default_order": null,
            "default_value": null,
            "deprecated": null,
            "description": "Object is the list of tags to use to find a server profile.",
            "exposed": true,
            "filterable": false,
            "foreign_key": null,
            "format": null,
            "getter": null,
            "identifier": null,
            "index": null,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "object",
            "orderable": false,
            "primary_key": null,
            "read_only": null,
            "required": true,
            "setter": null,
            "stored": true,
            "subtype": "policies_list",
            "transient": false,
            "type": "external",
            "unique": null,
            "uniqueScope": null
        },
        {
            "allowed_chars": null,
            "allowed_choices": null,
            "autogenerated": null,
            "channel": null,
            "creation_only": null,
            "default_order": null,
            "default_value": null,
            "deprecated": null,
            "description": "Subject is the subject of the policy.",
            "exposed": true,
            "filterable": false,
            "foreign_key": null,
            "format": null,
            "getter": null,
            "identifier": null,
            "index": null,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "subject",
            "orderable": false,
            "primary_key": null,
            "read_only": null,
            "required": true,
            "setter": null,
            "stored": true,
            "subtype": "policies_list",
            "transient": false,
            "type": "external",
            "unique": null,
            "uniqueScope": null
        }
    ],
    "children": [],
    "model": {
        "create": null,
        "delete": true,
        "description": null,
        "entity_name": "ServerPolicy",
        "extends": [
            "@base",
            "@described",
            "@identifiable-nopk-nostored",
            "@named"
        ],
        "get": true,
        "package": null,
        "resource_name": "serverpolicies",
        "rest_name": "serverpolicy",
        "root": null,
        "update": true
    }
}