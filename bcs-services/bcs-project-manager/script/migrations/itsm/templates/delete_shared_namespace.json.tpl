{
    "is_deleted": false,
    "name": "\u521b\u5efa\u5171\u4eab\u96c6\u7fa4\u547d\u540d\u7a7a\u95f4",
    "desc": "",
    "flow_type": "other",
    "is_enabled": true,
    "is_revocable": true,
    "revoke_config": {
        "type": 1,
        "state": 0
    },
    "is_draft": false,
    "is_builtin": false,
    "is_task_needed": false,
    "owners": "admin",
    "notify_rule": "ONCE",
    "notify_freq": 0,
    "is_biz_needed": false,
    "is_auto_approve": false,
    "is_iam_used": false,
    "is_supervise_needed": true,
    "supervise_type": "EMPTY",
    "supervisor": "",
    "engine_version": "PIPELINE_V1",
    "version_number": "20240308111613",
    "table": {
        "id": 36,
        "is_deleted": false,
        "name": "\u9ed8\u8ba4",
        "desc": "\u9ed8\u8ba4\u57fa\u7840\u6a21\u578b",
        "version": "EMPTY",
        "fields": [
            {
                "id": 1,
                "is_deleted": false,
                "is_builtin": true,
                "is_readonly": false,
                "is_valid": true,
                "display": true,
                "source_type": "CUSTOM",
                "source_uri": "",
                "api_instance_id": 0,
                "kv_relation": {},
                "type": "STRING",
                "key": "title",
                "name": "\u6807\u9898",
                "layout": "COL_12",
                "validate_type": "REQUIRE",
                "show_type": 1,
                "show_conditions": {},
                "regex": "EMPTY",
                "regex_config": {},
                "custom_regex": "",
                "desc": "\u8bf7\u8f93\u5165\u6807\u9898",
                "tips": "",
                "is_tips": false,
                "default": "",
                "choice": [],
                "related_fields": {},
                "meta": {},
                "flow_type": "DEFAULT",
                "project_key": "public",
                "source": "BASE-MODEL"
            },
            {
                "id": 2,
                "is_deleted": false,
                "is_builtin": true,
                "is_readonly": false,
                "is_valid": true,
                "display": true,
                "source_type": "DATADICT",
                "source_uri": "IMPACT",
                "api_instance_id": 0,
                "kv_relation": {},
                "type": "SELECT",
                "key": "impact",
                "name": "\u5f71\u54cd\u8303\u56f4",
                "layout": "COL_12",
                "validate_type": "REQUIRE",
                "show_type": 1,
                "show_conditions": {},
                "regex": "EMPTY",
                "regex_config": {},
                "custom_regex": "",
                "desc": "\u8bf7\u9009\u62e9\u5f71\u54cd\u8303\u56f4",
                "tips": "",
                "is_tips": false,
                "default": "",
                "choice": [],
                "related_fields": {},
                "meta": {},
                "flow_type": "DEFAULT",
                "project_key": "public",
                "source": "BASE-MODEL"
            },
            {
                "id": 3,
                "is_deleted": false,
                "is_builtin": true,
                "is_readonly": false,
                "is_valid": true,
                "display": true,
                "source_type": "DATADICT",
                "source_uri": "URGENCY",
                "api_instance_id": 0,
                "kv_relation": {},
                "type": "SELECT",
                "key": "urgency",
                "name": "\u7d27\u6025\u7a0b\u5ea6",
                "layout": "COL_12",
                "validate_type": "REQUIRE",
                "show_type": 1,
                "show_conditions": {},
                "regex": "EMPTY",
                "regex_config": {},
                "custom_regex": "",
                "desc": "\u8bf7\u9009\u62e9\u7d27\u6025\u7a0b\u5ea6",
                "tips": "",
                "is_tips": false,
                "default": "",
                "choice": [],
                "related_fields": {},
                "meta": {},
                "flow_type": "DEFAULT",
                "project_key": "public",
                "source": "BASE-MODEL"
            },
            {
                "id": 4,
                "is_deleted": false,
                "is_builtin": true,
                "is_readonly": true,
                "is_valid": true,
                "display": true,
                "source_type": "DATADICT",
                "source_uri": "PRIORITY",
                "api_instance_id": 0,
                "kv_relation": {},
                "type": "SELECT",
                "key": "priority",
                "name": "\u4f18\u5148\u7ea7",
                "layout": "COL_12",
                "validate_type": "REQUIRE",
                "show_type": 1,
                "show_conditions": {},
                "regex": "EMPTY",
                "regex_config": {},
                "custom_regex": "",
                "desc": "\u8bf7\u9009\u62e9\u4f18\u5148\u7ea7",
                "tips": "",
                "is_tips": false,
                "default": "",
                "choice": [],
                "related_fields": {
                    "rely_on": [
                        "urgency",
                        "impact"
                    ]
                },
                "meta": {},
                "flow_type": "DEFAULT",
                "project_key": "public",
                "source": "BASE-MODEL"
            },
            {
                "id": 5,
                "is_deleted": false,
                "is_builtin": true,
                "is_readonly": false,
                "is_valid": true,
                "display": true,
                "source_type": "RPC",
                "source_uri": "ticket_status",
                "api_instance_id": 0,
                "kv_relation": {},
                "type": "SELECT",
                "key": "current_status",
                "name": "\u5de5\u5355\u72b6\u6001",
                "layout": "COL_12",
                "validate_type": "REQUIRE",
                "show_type": 1,
                "show_conditions": {},
                "regex": "EMPTY",
                "regex_config": {},
                "custom_regex": "",
                "desc": "\u8bf7\u9009\u62e9\u5de5\u5355\u72b6\u6001",
                "tips": "",
                "is_tips": false,
                "default": "",
                "choice": [],
                "related_fields": {},
                "meta": {},
                "flow_type": "DEFAULT",
                "project_key": "public",
                "source": "BASE-MODEL"
            }
        ],
        "fields_order": [
            1,
            2,
            3,
            4,
            5
        ],
        "field_key_order": [
            "title",
            "impact",
            "urgency",
            "priority",
            "current_status"
        ]
    },
    "task_schemas": [],
    "creator": "",
    "updated_by": "",
    "workflow_id": 814,
    "version_message": "",
    "states": {
        "5557": {
            "workflow": 814,
            "id": 5557,
            "key": 5557,
            "name": "\u5f00\u59cb",
            "desc": "",
            "distribute_type": "PROCESS",
            "axis": {
                "x": 145,
                "y": 10
            },
            "is_builtin": true,
            "variables": {
                "inputs": [],
                "outputs": []
            },
            "tag": "DEFAULT",
            "processors_type": "OPEN",
            "processors": "",
            "assignors": "",
            "assignors_type": "EMPTY",
            "delivers": "",
            "delivers_type": "EMPTY",
            "can_deliver": false,
            "extras": {},
            "is_draft": false,
            "is_terminable": false,
            "fields": [],
            "type": "START",
            "api_instance_id": 0,
            "is_sequential": false,
            "finish_condition": {},
            "is_multi": false,
            "is_allow_skip": false,
            "creator": null,
            "create_at": "2024-03-05 21:15:29",
            "updated_by": "admin",
            "update_at": "2024-03-05 21:16:00",
            "end_at": null,
            "is_first_state": false
        },
        "5558": {
            "workflow": 814,
            "id": 5558,
            "key": 5558,
            "name": "\u63d0\u5355",
            "desc": "",
            "distribute_type": "PROCESS",
            "axis": {
                "x": 280,
                "y": 10
            },
            "is_builtin": true,
            "variables": {
                "inputs": [],
                "outputs": [
                    {
                        "key": "CLUSTER_TYPE",
                        "type": "SELECT",
                        "source": "field",
                        "state": 2948
                    },
                    {
                        "key": "CLUSTER_ID",
                        "type": "STRING",
                        "source": "field",
                        "state": 2582
                    }
                ]
            },
            "tag": "DEFAULT",
            "processors_type": "OPEN",
            "processors": "",
            "assignors": "",
            "assignors_type": "EMPTY",
            "delivers": "",
            "delivers_type": "EMPTY",
            "can_deliver": false,
            "extras": {},
            "is_draft": false,
            "is_terminable": false,
            "fields": [
                10802,
                10808,
                10803,
                10804
            ],
            "type": "NORMAL",
            "api_instance_id": 0,
            "is_sequential": false,
            "finish_condition": {},
            "is_multi": false,
            "is_allow_skip": false,
            "creator": null,
            "create_at": "2024-03-05 21:15:29",
            "updated_by": "admin",
            "update_at": "2024-03-05 21:15:54",
            "end_at": null,
            "is_first_state": true
        },
        "5559": {
            "workflow": 814,
            "id": 5559,
            "key": 5559,
            "name": "\u7ed3\u675f",
            "desc": "",
            "distribute_type": "PROCESS",
            "axis": {
                "x": 1215,
                "y": 5
            },
            "is_builtin": true,
            "variables": {
                "inputs": [],
                "outputs": []
            },
            "tag": "DEFAULT",
            "processors_type": "OPEN",
            "processors": "",
            "assignors": "",
            "assignors_type": "EMPTY",
            "delivers": "",
            "delivers_type": "EMPTY",
            "can_deliver": false,
            "extras": {},
            "is_draft": false,
            "is_terminable": false,
            "fields": [],
            "type": "END",
            "api_instance_id": 0,
            "is_sequential": false,
            "finish_condition": {},
            "is_multi": false,
            "is_allow_skip": false,
            "creator": null,
            "create_at": "2024-03-05 21:15:29",
            "updated_by": "admin",
            "update_at": "2024-03-05 21:16:18",
            "end_at": null,
            "is_first_state": false
        },
        "5560": {
            "workflow": 814,
            "id": 5560,
            "key": 5560,
            "name": "\u8d1f\u8d23\u4eba\u5ba1\u6279",
            "desc": "",
            "distribute_type": "PROCESS",
            "axis": {
                "x": 585,
                "y": 10
            },
            "is_builtin": false,
            "variables": {
                "inputs": [],
                "outputs": [
                    {
                        "source": "global",
                        "state": 2956,
                        "type": "STRING",
                        "key": "Fd6380d03621747689b9776224da468d",
                        "name": "\u5ba1\u6279\u7ed3\u679c",
                        "meta": {
                            "code": "NODE_APPROVE_RESULT",
                            "type": "SELECT",
                            "choice": [
                                {
                                    "key": "false",
                                    "name": "\u62d2\u7edd"
                                },
                                {
                                    "key": "true",
                                    "name": "\u901a\u8fc7"
                                }
                            ]
                        }
                    },
                    {
                        "source": "global",
                        "state": 2956,
                        "type": "STRING",
                        "key": "O1af1a6c7fceb2bbe9243d0cfd871028",
                        "name": "\u5ba1\u6279\u4eba",
                        "meta": {
                            "code": "NODE_APPROVER"
                        }
                    },
                    {
                        "source": "global",
                        "state": 2956,
                        "type": "INT",
                        "key": "dd93d6c0341ce48260408a2964448cb7",
                        "name": "\u5904\u7406\u4eba\u6570",
                        "meta": {
                            "code": "PROCESS_COUNT"
                        }
                    },
                    {
                        "source": "global",
                        "state": 2956,
                        "type": "INT",
                        "key": "c6619ac6399ebb6f4208406add9d971e",
                        "name": "\u901a\u8fc7\u4eba\u6570",
                        "meta": {
                            "code": "PASS_COUNT"
                        }
                    },
                    {
                        "source": "global",
                        "state": 2956,
                        "type": "INT",
                        "key": "l76a275fc8b01ceeb9a33f77ddb03679",
                        "name": "\u62d2\u7edd\u4eba\u6570",
                        "meta": {
                            "code": "REJECT_COUNT"
                        }
                    },
                    {
                        "source": "global",
                        "state": 2956,
                        "type": "INT",
                        "key": "f73b972755824685ca4cc7edd0a0bdab",
                        "name": "\u901a\u8fc7\u7387",
                        "meta": {
                            "code": "PASS_RATE",
                            "unit": "PERCENT"
                        }
                    },
                    {
                        "source": "global",
                        "state": 2956,
                        "type": "INT",
                        "key": "e987844249bd935b6e2b0b2609da593f",
                        "name": "\u62d2\u7edd\u7387",
                        "meta": {
                            "code": "REJECT_RATE",
                            "unit": "PERCENT"
                        }
                    }
                ]
            },
            "tag": "DEFAULT",
            "processors_type": "PERSON",
            "processors": "[[.Approvers]]",
            "assignors": "",
            "assignors_type": "EMPTY",
            "delivers": "",
            "delivers_type": "EMPTY",
            "can_deliver": false,
            "extras": {
                "ticket_status": {
                    "name": "",
                    "type": "keep"
                }
            },
            "is_draft": false,
            "is_terminable": false,
            "fields": [
                10805,
                10806,
                10807
            ],
            "type": "APPROVAL",
            "api_instance_id": 0,
            "is_sequential": false,
            "finish_condition": {
                "expressions": [],
                "type": "or"
            },
            "is_multi": false,
            "is_allow_skip": false,
            "creator": null,
            "create_at": "2024-03-05 21:15:29",
            "updated_by": "admin",
            "update_at": "2024-03-05 21:16:23",
            "end_at": null,
            "is_first_state": false
        },
        "5561": {
            "workflow": 814,
            "id": 5561,
            "key": 5561,
            "name": "\u6210\u529f\u56de\u8c03",
            "desc": "",
            "distribute_type": "PROCESS",
            "axis": {
                "x": 870,
                "y": -50
            },
            "is_builtin": false,
            "variables": {
                "outputs": [],
                "inputs": []
            },
            "tag": "DEFAULT",
            "processors_type": "PERSON",
            "processors": "admin",
            "assignors": "",
            "assignors_type": "EMPTY",
            "delivers": "",
            "delivers_type": "EMPTY",
            "can_deliver": false,
            "extras": {
                "webhook_info": {
                    "method": "POST",
                    "url": "[[.BCSGateway]]/bcsapi/v4/bcsproject/v1/projects/{{PROJECT_CODE}}/clusters/{{CLUSTER_ID}}/namespaces/{{NAMESPACE}}/callback/delete",
                    "query_params": [],
                    "auth": {
                        "auth_type": "bearer_token",
                        "auth_config": {
                            "token": "[[.BCSToken]]"
                        }
                    },
                    "headers": [],
                    "body": {
                        "type": "raw",
                        "raw_type": "JSON",
                        "content": "{\n    \"title\": \"{{ticket_title}}\",\n    \"currentStatus\": \"{{ticket_current_status}}\",\n    \"sn\": \"{{ticket_sn}}\",\n    \"ticketUrl\": \"{{ticket_ticket_url}}\",\n    \"applyInCluster\": true,\n    \"approveResult\": true\n}"
                    },
                    "settings": {
                        "timeout": 10
                    },
                    "success_exp": "resp.code==0"
                }
            },
            "is_draft": false,
            "is_terminable": false,
            "fields": [],
            "type": "WEBHOOK",
            "api_instance_id": 0,
            "is_sequential": false,
            "finish_condition": {},
            "is_multi": false,
            "is_allow_skip": false,
            "creator": null,
            "create_at": "2024-03-05 21:15:29",
            "updated_by": "admin",
            "update_at": "2024-03-05 21:16:09",
            "end_at": null,
            "is_first_state": false
        },
        "5562": {
            "workflow": 814,
            "id": 5562,
            "key": 5562,
            "name": "\u5931\u8d25\u56de\u8c03",
            "desc": "",
            "distribute_type": "PROCESS",
            "axis": {
                "x": 870,
                "y": 65
            },
            "is_builtin": false,
            "variables": {
                "outputs": [],
                "inputs": []
            },
            "tag": "DEFAULT",
            "processors_type": "PERSON",
            "processors": "admin",
            "assignors": "",
            "assignors_type": "EMPTY",
            "delivers": "",
            "delivers_type": "EMPTY",
            "can_deliver": false,
            "extras": {
                "webhook_info": {
                    "method": "POST",
                    "url": "[[.BCSGateway]]/bcsapi/v4/bcsproject/v1/projects/{{PROJECT_CODE}}/clusters/{{CLUSTER_ID}}/namespaces/{{NAMESPACE}}/callback/delete",
                    "query_params": [],
                    "auth": {
                        "auth_type": "bearer_token",
                        "auth_config": {
                            "token": "[[.BCSToken]]"
                        }
                    },
                    "headers": [],
                    "body": {
                        "type": "raw",
                        "raw_type": "JSON",
                        "content": "{\n    \"title\": \"{{ticket_title}}\",\n    \"currentStatus\": \"{{ticket_current_status}}\",\n    \"sn\": \"{{ticket_sn}}\",\n    \"ticketUrl\": \"{{ticket_ticket_url}}\",\n    \"applyInCluster\": false,\n    \"approveResult\": false\n}"
                    },
                    "settings": {
                        "timeout": 10
                    },
                    "success_exp": "resp.code==0"
                }
            },
            "is_draft": false,
            "is_terminable": false,
            "fields": [],
            "type": "WEBHOOK",
            "api_instance_id": 0,
            "is_sequential": false,
            "finish_condition": {},
            "is_multi": false,
            "is_allow_skip": false,
            "creator": null,
            "create_at": "2024-03-05 21:15:29",
            "updated_by": null,
            "update_at": "2024-03-05 21:15:45",
            "end_at": null,
            "is_first_state": false
        }
    },
    "transitions": {
        "6175": {
            "workflow": 814,
            "id": 6175,
            "from_state": 5557,
            "to_state": 5558,
            "name": "",
            "axis": {
                "start": "Right",
                "end": "Left"
            },
            "condition": {
                "expressions": [
                    {
                        "type": "and",
                        "expressions": [
                            {
                                "key": "G_INT_1",
                                "condition": "==",
                                "value": 1
                            }
                        ]
                    }
                ],
                "type": "and"
            },
            "condition_type": "default",
            "creator": null,
            "create_at": "2024-03-05 21:15:30",
            "updated_by": null,
            "update_at": "2024-03-05 21:15:30",
            "end_at": null
        },
        "6177": {
            "workflow": 814,
            "id": 6177,
            "from_state": 5560,
            "to_state": 5561,
            "name": "\u5ba1\u6279\u901a\u8fc7",
            "axis": {
                "start": "Right",
                "end": "Left"
            },
            "condition": {
                "expressions": [
                    {
                        "checkInfo": false,
                        "expressions": [
                            {
                                "choiceList": [],
                                "condition": "==",
                                "key": "Fd6380d03621747689b9776224da468d",
                                "source": "field",
                                "type": "SELECT",
                                "value": "true",
                                "meta": {
                                    "code": "NODE_APPROVE_RESULT",
                                    "type": "SELECT",
                                    "choice": [
                                        {
                                            "key": "false",
                                            "name": "\u62d2\u7edd"
                                        },
                                        {
                                            "key": "true",
                                            "name": "\u901a\u8fc7"
                                        }
                                    ]
                                }
                            }
                        ],
                        "type": "and"
                    }
                ],
                "type": "and"
            },
            "condition_type": "by_field",
            "creator": null,
            "create_at": "2024-03-05 21:15:30",
            "updated_by": null,
            "update_at": "2024-03-05 21:15:30",
            "end_at": null
        },
        "6178": {
            "workflow": 814,
            "id": 6178,
            "from_state": 5560,
            "to_state": 5562,
            "name": "\u5ba1\u6279\u9a73\u56de",
            "axis": {
                "start": "Right",
                "end": "Left"
            },
            "condition": {
                "expressions": [
                    {
                        "checkInfo": false,
                        "expressions": [
                            {
                                "choiceList": [],
                                "condition": "==",
                                "key": "Fd6380d03621747689b9776224da468d",
                                "source": "field",
                                "type": "SELECT",
                                "value": "false",
                                "meta": {
                                    "code": "NODE_APPROVE_RESULT",
                                    "type": "SELECT",
                                    "choice": [
                                        {
                                            "key": "false",
                                            "name": "\u62d2\u7edd"
                                        },
                                        {
                                            "key": "true",
                                            "name": "\u901a\u8fc7"
                                        }
                                    ]
                                }
                            }
                        ],
                        "type": "and"
                    }
                ],
                "type": "and"
            },
            "condition_type": "by_field",
            "creator": null,
            "create_at": "2024-03-05 21:15:30",
            "updated_by": null,
            "update_at": "2024-03-05 21:15:30",
            "end_at": null
        },
        "6179": {
            "workflow": 814,
            "id": 6179,
            "from_state": 5561,
            "to_state": 5559,
            "name": "\u6d41\u7a0b\u7ed3\u675f",
            "axis": {
                "start": "Right",
                "end": "Top"
            },
            "condition": {
                "expressions": [
                    {
                        "type": "and",
                        "expressions": [
                            {
                                "key": "G_INT_1",
                                "condition": "==",
                                "value": 1
                            }
                        ]
                    }
                ],
                "type": "and"
            },
            "condition_type": "default",
            "creator": null,
            "create_at": "2024-03-05 21:15:30",
            "updated_by": null,
            "update_at": "2024-03-05 21:15:30",
            "end_at": null
        },
        "6183": {
            "workflow": 814,
            "id": 6183,
            "from_state": 5562,
            "to_state": 5559,
            "name": "\u6d41\u7a0b\u7ed3\u675f",
            "axis": {
                "start": "Right",
                "end": "Bottom"
            },
            "condition": {
                "expressions": [
                    {
                        "type": "and",
                        "expressions": [
                            {
                                "key": "G_INT_1",
                                "condition": "==",
                                "value": 1
                            }
                        ]
                    }
                ],
                "type": "and"
            },
            "condition_type": "default",
            "creator": "admin",
            "create_at": "2024-03-05 21:15:45",
            "updated_by": "admin",
            "update_at": "2024-03-05 21:15:50",
            "end_at": null
        },
        "6184": {
            "workflow": 814,
            "id": 6184,
            "from_state": 5558,
            "to_state": 5560,
            "name": "\u9ed8\u8ba4",
            "axis": {
                "start": "Right",
                "end": "Left"
            },
            "condition": {
                "expressions": [
                    {
                        "type": "and",
                        "expressions": [
                            {
                                "key": "G_INT_1",
                                "condition": "==",
                                "value": 1
                            }
                        ]
                    }
                ],
                "type": "and"
            },
            "condition_type": "default",
            "creator": "admin",
            "create_at": "2024-03-05 21:16:26",
            "updated_by": "admin",
            "update_at": "2024-03-05 21:16:26",
            "end_at": null
        }
    },
    "triggers": [
        {
            "rules": [
                {
                    "name": "",
                    "condition": "",
                    "by_condition": false,
                    "action_schemas": [
                        {
                            "id": 6637,
                            "creator": "admin",
                            "updated_by": "admin",
                            "is_deleted": false,
                            "name": "",
                            "display_name": "",
                            "component_type": "automatic_announcement",
                            "operate_type": "BACKEND",
                            "delay_params": {
                                "type": "custom",
                                "value": 0
                            },
                            "can_repeat": false,
                            "params": [
                                {
                                    "key": "web_hook_id",
                                    "value": "BCS_CREATE_NAMESPACE_TICKET",
                                    "ref_type": "custom"
                                },
                                {
                                    "key": "chat_id",
                                    "value": "",
                                    "ref_type": "custom"
                                },
                                {
                                    "key": "content",
                                    "value": "\u60a8\u6709\u4e00\u6761\u5355\u636e\u5f85\u5904\u7406",
                                    "ref_type": "custom"
                                },
                                {
                                    "key": "mentioned_list",
                                    "value": "${ticket_current_processors}",
                                    "ref_type": "import"
                                }
                            ],
                            "inputs": {}
                        }
                    ]
                }
            ],
            "id": 6374,
            "creator": "admin",
            "updated_by": "admin",
            "is_deleted": false,
            "name": "\u4f01\u5fae\u901a\u77e5",
            "desc": "",
            "signal": "THROUGH_TRANSITION",
            "sender": "3002",
            "inputs": [],
            "source_type": "workflow",
            "source_id": 814,
            "source_table_id": 0,
            "is_draft": false,
            "is_enabled": true,
            "icon": "message",
            "project_key": "alkaid-test"
        },
        {
            "rules": [
                {
                    "name": "",
                    "condition": "",
                    "by_condition": false,
                    "action_schemas": [
                        {
                            "id": 6636,
                            "creator": "admin",
                            "updated_by": "admin",
                            "is_deleted": false,
                            "name": "",
                            "display_name": "",
                            "component_type": "send_message",
                            "operate_type": "BACKEND",
                            "delay_params": {
                                "type": "custom",
                                "value": 0
                            },
                            "can_repeat": false,
                            "params": [
                                {
                                    "key": "sub_message_component",
                                    "sub_components": [
                                        {
                                            "key": "send_rtx_message",
                                            "params": [
                                                {
                                                    "key": "title",
                                                    "value": "\u7b97\u529b\u96c6\u7fa4\u547d\u540d\u7a7a\u95f4\u652f\u6301\u5220\u9664\u64cd\u4f5c",
                                                    "ref_type": "custom"
                                                },
                                                {
                                                    "key": "receivers",
                                                    "value": [
                                                        {
                                                            "ref_type": "reference",
                                                            "value": {
                                                                "member_type": "VARIABLE",
                                                                "members": "ticket_creator"
                                                            }
                                                        }
                                                    ],
                                                    "ref_type": "custom"
                                                },
                                                {
                                                    "key": "content",
                                                    "value": "\u7b97\u529b\u96c6\u7fa4\u547d\u540d\u7a7a\u95f4\u652f\u6301\u5220\u9664\u64cd\u4f5c",
                                                    "ref_type": "custom"
                                                }
                                            ]
                                        }
                                    ]
                                }
                            ],
                            "inputs": {}
                        }
                    ]
                }
            ],
            "id": 6373,
            "creator": "admin",
            "updated_by": "admin",
            "is_deleted": false,
            "name": "\u901a\u77e5\u7528\u6237\u4e0d\u652f\u6301\u7b97\u529b\u96c6\u7fa4\u5220\u9664\u64cd\u4f5c",
            "desc": "xxx",
            "signal": "THROUGH_TRANSITION",
            "sender": "6182",
            "inputs": [],
            "source_type": "workflow",
            "source_id": 814,
            "source_table_id": 0,
            "is_draft": false,
            "is_enabled": true,
            "icon": "message",
            "project_key": "bcs-test"
        }
    ],
    "fields": {
        "10802": {
            "id": 10802,
            "is_deleted": false,
            "is_builtin": true,
            "is_readonly": false,
            "is_valid": true,
            "display": true,
            "source_type": "CUSTOM",
            "source_uri": "",
            "api_instance_id": 0,
            "kv_relation": {},
            "type": "STRING",
            "key": "title",
            "name": "\u6807\u9898",
            "layout": "COL_12",
            "validate_type": "REQUIRE",
            "show_type": 1,
            "show_conditions": {},
            "regex": "EMPTY",
            "regex_config": {},
            "custom_regex": "",
            "desc": "\u8bf7\u8f93\u5165\u6807\u9898",
            "tips": "",
            "is_tips": false,
            "default": "",
            "choice": [],
            "related_fields": {},
            "meta": {},
            "workflow_id": 814,
            "state_id": "",
            "source": "TABLE"
        },
        "10803": {
            "id": 10803,
            "is_deleted": false,
            "is_builtin": false,
            "is_readonly": false,
            "is_valid": true,
            "display": false,
            "source_type": "CUSTOM",
            "source_uri": "",
            "api_instance_id": 0,
            "kv_relation": {},
            "type": "STRING",
            "key": "CLUSTER_ID",
            "name": "\u96c6\u7fa4ID",
            "layout": "COL_12",
            "validate_type": "REQUIRE",
            "show_type": 1,
            "show_conditions": {},
            "regex": "CUSTOM",
            "regex_config": {
                "rule": {
                    "expressions": [
                        {
                            "condition": "",
                            "key": "",
                            "source": "field",
                            "type": "STRING",
                            "value": ""
                        }
                    ],
                    "type": "and"
                }
            },
            "custom_regex": "^BCS-K8S-[0-9]{5}$",
            "desc": "",
            "tips": "",
            "is_tips": false,
            "default": "",
            "choice": [],
            "related_fields": {},
            "meta": {},
            "workflow_id": 814,
            "state_id": 5558,
            "source": "CUSTOM"
        },
        "10804": {
            "id": 10804,
            "is_deleted": false,
            "is_builtin": false,
            "is_readonly": false,
            "is_valid": true,
            "display": false,
            "source_type": "CUSTOM",
            "source_uri": "",
            "api_instance_id": 0,
            "kv_relation": {},
            "type": "STRING",
            "key": "NAMESPACE",
            "name": "\u547d\u540d\u7a7a\u95f4",
            "layout": "COL_12",
            "validate_type": "REQUIRE",
            "show_type": 1,
            "show_conditions": {},
            "regex": "CUSTOM",
            "regex_config": {
                "rule": {
                    "expressions": [
                        {
                            "condition": "",
                            "key": "",
                            "source": "field",
                            "type": "STRING",
                            "value": ""
                        }
                    ],
                    "type": "and"
                }
            },
            "custom_regex": "",
            "desc": "",
            "tips": "",
            "is_tips": false,
            "default": "",
            "choice": [],
            "related_fields": {},
            "meta": {},
            "workflow_id": 814,
            "state_id": 5558,
            "source": "CUSTOM"
        },
        "10805": {
            "id": 10805,
            "is_deleted": false,
            "is_builtin": false,
            "is_readonly": false,
            "is_valid": true,
            "display": true,
            "source_type": "CUSTOM",
            "source_uri": "",
            "api_instance_id": 0,
            "kv_relation": {},
            "type": "RADIO",
            "key": "bfaba606fe9be5d6596270a00c87d428",
            "name": "\u5ba1\u6279\u610f\u89c1",
            "layout": "COL_6",
            "validate_type": "REQUIRE",
            "show_type": 1,
            "show_conditions": {},
            "regex": "EMPTY",
            "regex_config": {},
            "custom_regex": "",
            "desc": "",
            "tips": "",
            "is_tips": false,
            "default": "true",
            "choice": [
                {
                    "key": "true",
                    "name": "\u901a\u8fc7"
                },
                {
                    "key": "false",
                    "name": "\u62d2\u7edd"
                }
            ],
            "related_fields": {},
            "meta": {
                "code": "APPROVE_RESULT"
            },
            "workflow_id": 814,
            "state_id": 5560,
            "source": "CUSTOM"
        },
        "10806": {
            "id": 10806,
            "is_deleted": false,
            "is_builtin": false,
            "is_readonly": false,
            "is_valid": true,
            "display": false,
            "source_type": "CUSTOM",
            "source_uri": "",
            "api_instance_id": 0,
            "kv_relation": {},
            "type": "TEXT",
            "key": "ff9e6f2b83c5ea1c47f36e10310980c3",
            "name": "\u5907\u6ce8",
            "layout": "COL_12",
            "validate_type": "OPTION",
            "show_type": 0,
            "show_conditions": {
                "expressions": [
                    {
                        "value": "false",
                        "type": "RADIO",
                        "condition": "==",
                        "key": "bfaba606fe9be5d6596270a00c87d428"
                    }
                ],
                "type": "and"
            },
            "regex": "EMPTY",
            "regex_config": {},
            "custom_regex": "",
            "desc": "",
            "tips": "",
            "is_tips": false,
            "default": "",
            "choice": [],
            "related_fields": {},
            "meta": {},
            "workflow_id": 814,
            "state_id": 5560,
            "source": "CUSTOM"
        },
        "10807": {
            "id": 10807,
            "is_deleted": false,
            "is_builtin": false,
            "is_readonly": false,
            "is_valid": true,
            "display": false,
            "source_type": "CUSTOM",
            "source_uri": "",
            "api_instance_id": 0,
            "kv_relation": {},
            "type": "TEXT",
            "key": "I60e9046a05cdff0951ee0acf07d4db8",
            "name": "\u5907\u6ce8",
            "layout": "COL_12",
            "validate_type": "REQUIRE",
            "show_type": 0,
            "show_conditions": {
                "expressions": [
                    {
                        "value": "true",
                        "type": "RADIO",
                        "condition": "==",
                        "key": "bfaba606fe9be5d6596270a00c87d428"
                    }
                ],
                "type": "and"
            },
            "regex": "EMPTY",
            "regex_config": {},
            "custom_regex": "",
            "desc": "",
            "tips": "",
            "is_tips": false,
            "default": "",
            "choice": [],
            "related_fields": {},
            "meta": {},
            "workflow_id": 814,
            "state_id": 5560,
            "source": "CUSTOM"
        },
        "10808": {
            "id": 10808,
            "is_deleted": false,
            "is_builtin": false,
            "is_readonly": false,
            "is_valid": true,
            "display": false,
            "source_type": "CUSTOM",
            "source_uri": "",
            "api_instance_id": 0,
            "kv_relation": {},
            "type": "STRING",
            "key": "PROJECT_CODE",
            "name": "\u9879\u76ee\u82f1\u6587\u540d",
            "layout": "COL_12",
            "validate_type": "REQUIRE",
            "show_type": 1,
            "show_conditions": {},
            "regex": "EMPTY",
            "regex_config": {
                "rule": {
                    "expressions": [
                        {
                            "condition": "",
                            "key": "",
                            "source": "field",
                            "type": "STRING",
                            "value": ""
                        }
                    ],
                    "type": "and"
                }
            },
            "custom_regex": "",
            "desc": "",
            "tips": "",
            "is_tips": false,
            "default": "",
            "choice": [],
            "related_fields": {},
            "meta": {},
            "workflow_id": 814,
            "state_id": 5558,
            "source": "CUSTOM"
        }
    },
    "notify": [
        2
    ],
    "extras": {
        "task_settings": []
    }
}