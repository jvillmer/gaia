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
            "description": "Object is the list of tags to use to find a enforcer profile.",
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
            "secret": false,
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
            "secret": false,
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
        "aliases": [
            "srvpol",
            "srvpols"
        ],
        "create": null,
        "delete": true,
        "description": "A Enforcer Profile Mapping Policy will tell what Enforcer Profile should be used by and Aporeto Agent based on the Enforcer that have been used during the registration. The policy can also be propagated down to the child namespace.",
        "entity_name": "EnforcerProfileMappingPolicy",
        "extends": [
            "@base",
            "@described",
            "@disabled",
            "@identifiable-nopk-nostored",
            "@named",
            "@propagated"
        ],
        "get": true,
        "package": "System",
        "resource_name": "enforcerprofilemappingpolicies",
        "rest_name": "enforcerprofilemappingpolicy",
        "root": null,
        "update": true
    }
}