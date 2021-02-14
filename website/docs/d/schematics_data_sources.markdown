---

copyright:
  years: 2021
lastupdated: "2021"

keywords: terraform

subcollection: terraform

---

# IBM Cloud Schematics API data sources
{: #schematics-data-sources}

Review the data sources that you can use to retrieve information about your IBM Cloud Schematics API resources.
All data sources are imported as read-only information. You can reference the output parameters for each data source by using Terraform interpolation syntax.

Before you start working with your data source, make sure to review the [required parameters](/docs/terraform?topic=terraform-provider-reference#required-parameters) 
that you need to specify in the `provider` block of your Terraform configuration file.
{: important}

## `ibm_schematics_output`
{: #schematics_output}

Retrieve information about Information about the Terraform output values that are defined in the Terraform template or IBM Cloud software template..
{: shortdesc}

### Sample Terraform code
{: #schematics_output-sample}

```
data "ibm_schematics_output" "schematics_output" {
  w_id = "w_id"
}
```

### Input parameters
{: #schematics_output-input}

Review the input parameters that you can specify for your data source. {: shortdesc}

|Name|Data type|Required/optional|Description|
|----|-----------|-------|----------|
|`w_id`|String|Required|The ID of the workspace for which you want to retrieve output values. To find the workspace ID, use the `GET /workspaces` API.|

### Output parameters
{: #schematics_output-output}

Review the output parameters that you can access after you retrieved your data source. {: shortdesc}

|Name|Data type|Description|
|----|-----------|---------|
|`output_values`|String|The output values exported as a map of key:value pairs

## `ibm_schematics_state`
{: #schematics_state}

Retrieve information about The content of the Terraform statefile (`terraform.tfstate`)..
{: shortdesc}

### Sample Terraform code
{: #schematics_state-sample}

```
data "ibm_schematics_state" "schematics_state" {
  w_id = "w_id"
  t_id = "t_id"
}
```

### Input parameters
{: #schematics_state-input}

Review the input parameters that you can specify for your data source. {: shortdesc}

|Name|Data type|Required/optional|Description|
|----|-----------|-------|----------|
|`w_id`|String|Required|The ID of the workspace for which you want to retrieve the Terraform statefile. To find the workspace ID, use the `GET /v1/workspaces` API.|
|`t_id`|String|Required|The ID of the Terraform template for which you want to retrieve the Terraform statefile. When you create a workspace, the Terraform template that your workspace points to is assigned a unique ID. To find this ID, use the `GET /v1/workspaces` API and review the `template_data.id` value.|

### Output parameters
{: #schematics_state-output}

Review the output parameters that you can access after you retrieved your data source. {: shortdesc}

|Name|Data type|Description|
|----|-----------|---------|
|`state_store_json`|String|The JSON representation of the state store data in string format.|


## `ibm_schematics_workspace`
{: #schematics_workspace}

Retrieve information about Overview of workspace details..
{: shortdesc}

### Sample Terraform code
{: #schematics_workspace-sample}

```
data "ibm_schematics_workspace" "schematics_workspace" {
  w_id = "w_id"
}
```

### Input parameters
{: #schematics_workspace-input}

Review the input parameters that you can specify for your data source. {: shortdesc}

|Name|Data type|Required/optional|Description|
|----|-----------|-------|----------|
|`w_id`|String|Required|The ID of the workspace for which you want to retrieve detailed information. To find the workspace ID, use the `GET /v1/workspaces` API.|

### Output parameters
{: #schematics_workspace-output}

Review the output parameters that you can access after you retrieved your data source. {: shortdesc}

|Name|Data type|Description|
|----|-----------|---------|
|`applied_shareddata_ids`|List|List of applied shared dataset id.|
|`catalog_ref`|List|Information about the software template that you chose from the IBM Cloud catalog. This information is returned for IBM Cloud catalog offerings only. This list contains only one item.|
|`catalog_ref.dry_run`|Boolean|Dry run.|
|`catalog_ref.item_icon_url`|String|The URL to the icon of the software template in the IBM Cloud catalog.|
|`catalog_ref.item_id`|String|The ID of the software template that you chose to install from the IBM Cloud catalog. This software is provisioned with Schematics.|
|`catalog_ref.item_name`|String|The name of the software that you chose to install from the IBM Cloud catalog.|
|`catalog_ref.item_readme_url`|String|The URL to the readme file of the software template in the IBM Cloud catalog.|
|`catalog_ref.item_url`|String|The URL to the software template in the IBM Cloud catalog.|
|`catalog_ref.launch_url`|String|The URL to the dashboard to access your software.|
|`catalog_ref.offering_version`|String|The version of the software template that you chose to install from the IBM Cloud catalog.|
|`created_at`|String|The timestamp when the workspace was created.|
|`created_by`|String|The user ID that created the workspace.|
|`crn`|String|Workspace CRN.|
|`description`|String|The description of the workspace.|
|`id`|String|The unique identifier of the workspace.|
|`last_health_check_at`|String|The timestamp when the last health check was performed by Schematics.|
|`location`|String|The IBM Cloud location where your workspace was provisioned.|
|`name`|String|The name of the workspace.|
|`resource_group`|String|The resource group the workspace was provisioned in.|
|`runtime_data`|List|Information about the provisioning engine, state file, and runtime logs.|
|`runtime_data.engine_cmd`|String|The command that was used to apply the Terraform template or IBM Cloud catalog software template.|
|`runtime_data.engine_name`|String|The provisioning engine that was used to apply the Terraform template or IBM Cloud catalog software template.|
|`runtime_data.engine_version`|String|The version of the provisioning engine that was used.|
|`runtime_data.id`|String|The ID that was assigned to your Terraform template or IBM Cloud catalog software template.|
|`runtime_data.log_store_url`|String|The URL to access the logs that were created during the creation, update, or deletion of your IBM Cloud resources.|
|`runtime_data.output_values`|List|List of Output values.|
|`runtime_data.resources`|List|List of resources.|
|`runtime_data.state_store_url`|String|The URL where the Terraform statefile (`terraform.tfstate`) is stored. You can use the statefile to find an overview of IBM Cloud resources that were created by Schematics. Schematics uses the statefile as an inventory list to determine future create, update, or deletion actions.|
|`shared_data`|List|Information that is shared across templates in IBM Cloud catalog offerings. This information is not provided when you create a workspace from your own Terraform template. This list contains only one item.|
|`shared_data.cluster_id`|String|The ID of the cluster where you want to provision the resources of all IBM Cloud catalog templates that are included in the catalog offering.|
|`shared_data.cluster_name`|String|Target cluster name.|
|`shared_data.entitlement_keys`|List|The entitlement key that you want to use to install IBM Cloud entitled software.|
|`shared_data.namespace`|String|The Kubernetes namespace or OpenShift project where the resources of all IBM Cloud catalog templates that are included in the catalog offering are deployed into.|
|`shared_data.region`|String|The IBM Cloud region that you want to use for the resources of all IBM Cloud catalog templates that are included in the catalog offering.|
|`shared_data.resource_group_id`|String|The ID of the resource group that you want to use for the resources of all IBM Cloud catalog templates that are included in the catalog offering.|
|`status`|String|The status of the workspace.  **Active**: After you successfully ran your infrastructure code by applying your Terraform execution plan, the state of your workspace changes to `Active`.  **Connecting**: Schematics tries to connect to the template in your source repo. If successfully connected, the template is downloaded and metadata, such as input parameters, is extracted. After the template is downloaded, the state of the workspace changes to `Scanning`.  **Draft**: The workspace is created without a reference to a GitHub or GitLab repository.  **Failed**: If errors occur during the execution of your infrastructure code in IBM Cloud Schematics, your workspace status is set to `Failed`.  **Inactive**: The Terraform template was scanned successfully and the workspace creation is complete. You can now start running Schematics plan and apply actions to provision the IBM Cloud resources that you specified in your template. If you have an `Active` workspace and decide to remove all your resources, your workspace is set to `Inactive` after all your resources are removed.  **In progress**: When you instruct IBM Cloud Schematics to run your infrastructure code by applying your Terraform execution plan, the status of our workspace changes to `In progress`.  **Scanning**: The download of the Terraform template is complete and vulnerability scanning started. If the scan is successful, the workspace state changes to `Inactive`. If errors in your template are found, the state changes to `Template Error`.  **Stopped**: The Schematics plan, apply, or destroy action was cancelled manually.  **Template Error**: The Schematics template contains errors and cannot be processed.|
|`tags`|List|A list of tags that are associated with the workspace.|
|`template_data`|List|Information about the Terraform or IBM Cloud software template that you want to use.|
|`template_data.env_values`|List|List of environment values.|
|`template_data.env_values.hidden`|Boolean|Environment variable is hidden.|
|`template_data.env_values.name`|String|Environment variable name.|
|`template_data.env_values.secure`|Boolean|Environment variable is secure.|
|`template_data.env_values.value`|String|Value for environment variable.|
|`template_data.folder`|String|The subfolder in your GitHub or GitLab repository where your Terraform template is stored. If your template is stored in the root directory, `.` is returned.|
|`template_data.has_githubtoken`|Boolean|Has github token.|
|`template_data.id`|String|The ID that was assigned to your Terraform template or IBM Cloud catalog software template.|
|`template_data.template_type`|String|The Terraform version that was used to run your Terraform code.|
|`template_data.uninstall_script_name`|String|Uninstall script name.|
|`template_data.values`|String|A list of variable values that you want to apply during the Helm chart installation. The list must be provided in JSON format, such as `""autoscaling:  enabled: true  minReplicas: 2"`. The values that you define here override the default Helm chart values. This field is supported only for IBM Cloud catalog offerings that are provisioned by using the Terraform Helm provider.|
|`template_data.values_metadata`|List|A list of input variables that are associated with the workspace.|
|`template_data.values_url`|String|The API endpoint to access the input variables that you defined for your template.|
|`template_data.variablestore`|List|Information about the input variables that your template uses.|
|`template_data.variablestore.description`|String|The description of your input variable.|
|`template_data.variablestore.name`|String|The name of the variable.|
|`template_data.variablestore.secure`|Boolean|If set to `true`, the value of your input variable is protected and not returned in your API response.|
|`template_data.variablestore.type`|String|`Terraform v0.11` supports `string`, `list`, `map` data type. For more information, about the syntax, see [Configuring input variables](https://www.terraform.io/docs/configuration-0-11/variables.html). <br> `Terraform v0.12` additionally, supports `bool`, `number` and complex data types such as `list(type)`, `map(type)`, `object({attribute name=type,..})`, `set(type)`, `tuple([type])`. For more information, about the syntax to use the complex data type, see [Configuring variables](https://www.terraform.io/docs/configuration/variables.html#type-constraints).|
|`template_data.variablestore.value`|String|Enter the value as a string for the primitive types such as `bool`, `number`, `string`, and `HCL` format for the complex variables, as you provide in a `.tfvars` file. **You need to enter escaped string of `HCL` format for the complex variable value**. For more information, about how to declare variables in a terraform configuration file and provide value to schematics, see [Providing values for the declared variables](/docs/schematics?topic=schematics-create-tf-config#declare-variable).|
|`template_ref`|String|Workspace template ref.|
|`template_repo`|List|Information about the Terraform template that your workspace points to. This list contains only one item.|
|`template_repo.branch`|String|The branch in GitHub where your Terraform template is stored.|
|`template_repo.full_url`|String|Full repo URL.|
|`template_repo.has_uploadedgitrepotar`|Boolean|Has uploaded git repo tar.|
|`template_repo.release`|String|The release tag in GitHub of your Terraform template.|
|`template_repo.repo_sha_value`|String|Repo SHA value.|
|`template_repo.repo_url`|String|The URL to the repository where the IBM Cloud catalog software template is stored.|
|`template_repo.url`|String|The URL to the GitHub or GitLab repository where your Terraform template is stored.|
|`type`|List|The Terraform version that was used to run your Terraform code.|
|`updated_at`|String|The timestamp when the workspace was last updated.|
|`updated_by`|String|The user ID that updated the workspace.|
|`workspace_status`|List|Response parameter that indicate if a workspace is frozen or locked. This list contains only one item.|
|`workspace_status.frozen`|Boolean|If set to true, the workspace is frozen and changes to the workspace are disabled.|
|`workspace_status.frozen_at`|String|The timestamp when the workspace was frozen.|
|`workspace_status.frozen_by`|String|The user ID that froze the workspace.|
|`workspace_status.locked`|Boolean|If set to true, the workspace is locked and disabled for changes.|
|`workspace_status.locked_by`|String|The user ID that initiated a resource-related action, such as applying or destroying resources, that locked the workspace.|
|`workspace_status.locked_time`|String|The timestamp when the workspace was locked.|
|`workspace_status_msg`|List|Information about the last action that ran against the workspace. This list contains only one item.|
|`workspace_status_msg.status_code`|String|The success or error code that was returned for the last plan, apply, or destroy action that ran against your workspace.|
|`workspace_status_msg.status_msg`|String|The success or error message that was returned for the last plan, apply, or destroy action that ran against your workspace.|

## `ibm_schematics_action`
{: #schematics_action}

Retrieve information about Complete action details with the user inputs and the system generated data..
{: shortdesc}

### Sample Terraform code
{: #schematics_action-sample}

```
data "ibm_schematics_action" "schematics_action" {
  action_id = "action_id"
}
```

### Input parameters
{: #schematics_action-input}

Review the input parameters that you can specify for your data source. {: shortdesc}

|Name|Data type|Required/optional|Description|
|----|-----------|-------|----------|
|`action_id`|String|Required|Use GET or actions API to look up the action IDs in your IBM Cloud account.|

### Output parameters
{: #schematics_action-output}

Review the output parameters that you can access after you retrieved your data source. {: shortdesc}

|Name|Data type|Description|
|----|-----------|---------|
|`name`|String|Action name (unique for an account).|
|`description`|String|Action description.|
|`location`|String|List of action locations supported by IBM Cloud Schematics service.  **Note** this does not limit the location of the resources provisioned using Schematics.|
|`resource_group`|String|Resource-group name for an action.  By default, action is created in default resource group.|
|`tags`|List|Action tags.|
|`user_state`|List|User defined status of the Schematics object. This list contains only one item.|
|`user_state.state`|String|User defined states  * `draft` Object can be modified, and can be used by jobs run by an author, during execution  * `live` Object can be modified, and can be used by jobs during execution  * `locked` Object cannot be modified, and can be used by jobs during execution  * `disable` Object can be modified, and cannot be used by Jobs during execution.|
|`user_state.set_by`|String|Name of the user who set the state of an Object.|
|`user_state.set_at`|String|When the user who set the state of an Object.|
|`source_readme_url`|String|URL of the `README` file, for the source.|
|`source`|List|Source of templates, playbooks, or controls. This list contains only one item.|
|`source.source_type`|String|Type of source for the Template.|
|`source.git`|List|Connection details to Git source. This list contains only one item.|
|`source.git.git_repo_url`|String|URL to the GIT Repo that can be used to clone the template.|
|`source.git.git_token`|String|Personal Access Token to connect to Git URLs.|
|`source.git.git_repo_folder`|String|Name of the folder in the Git Repo, that contains the template.|
|`source.git.git_release`|String|Name of the release tag, used to fetch the Git Repo.|
|`source.git.git_branch`|String|Name of the branch, used to fetch the Git Repo.|
|`source_type`|String|Type of source for the Template.|
|`command_parameter`|String|Schematics job command parameter (playbook-name, capsule-name or flow-name).|
|`bastion`|List|Complete target details with the user inputs and the system generated data. This list contains only one item.|
|`bastion.name`|String|Target name.|
|`bastion.type`|String|Target type (`cluster`, `vsi`, `icd`, `vpc`).|
|`bastion.description`|String|Target description.|
|`bastion.resource_query`|String|Resource selection query string.|
|`bastion.credential`|String|Override credential for each resource.  Reference to credentials values, used by all the resources.|
|`bastion.id`|String|Target ID.|
|`bastion.created_at`|String|Targets creation time.|
|`bastion.created_by`|String|E-mail address of the user who created the targets.|
|`bastion.updated_at`|String|Targets updation time.|
|`bastion.updated_by`|String|E-mail address of user who updated the targets.|
|`bastion.sys_lock`|List|System lock status. This list contains only one item.|
|`bastion.sys_lock.sys_locked`|Boolean|Is the Workspace locked by the Schematic action ?.|
|`bastion.sys_lock.sys_locked_by`|String|Name of the user who performed the action, that lead to lock the Workspace.|
|`bastion.sys_lock.sys_locked_at`|String|When the user performed the action that lead to lock the Workspace ?.|
|`bastion.resource_ids`|List|Array of the resource IDs.|
|`targets_ini`|String|Inventory of host and host group for the playbook in `INI` file format. For example, `"targets_ini": "[webserverhost]  172.22.192.6  [dbhost]  172.22.192.5"`. For more information, about an inventory host group syntax, see [Inventory host groups](/docs/schematics?topic=schematics-schematics-cli-reference#schematics-inventory-host-grps).|
|`inputs`|List|Input variables for an action.|
|`inputs.name`|String|Name of the variable.|
|`inputs.value`|String|Value for the variable or reference to the value.|
|`inputs.metadata`|List|User editable metadata for the variables. This list contains only one item.|
|`inputs.metadata.type`|String|Type of the variable.|
|`inputs.metadata.aliases`|List|List of aliases for the variable name.|
|`inputs.metadata.description`|String|Description of the meta data.|
|`inputs.metadata.default_value`|String|Default value for the variable, if the override value is not specified.|
|`inputs.metadata.secure`|Boolean|Is the variable secure or sensitive ?.|
|`inputs.metadata.immutable`|Boolean|Is the variable readonly ?.|
|`inputs.metadata.hidden`|Boolean|If true, the variable will not be displayed on UI or CLI.|
|`inputs.metadata.options`|List|List of possible values for this variable.  If type is integer or date, then the array of string will be  converted to array of integers or date during runtime.|
|`inputs.metadata.min_value`|Integer|Minimum value of the variable. Applicable for integer type.|
|`inputs.metadata.max_value`|Integer|Maximum value of the variable. Applicable for integer type.|
|`inputs.metadata.min_length`|Integer|Minimum length of the variable value. Applicable for string type.|
|`inputs.metadata.max_length`|Integer|Maximum length of the variable value. Applicable for string type.|
|`inputs.metadata.matches`|String|Regex for the variable value.|
|`inputs.metadata.position`|Integer|Relative position of this variable in a list.|
|`inputs.metadata.group_by`|String|Display name of the group this variable belongs to.|
|`inputs.metadata.source`|String|Source of this meta-data.|
|`inputs.link`|String|Reference link to the variable value By default the expression will point to self.value.|
|`outputs`|List|Output variables for an action.|
|`outputs.name`|String|Name of the variable.|
|`outputs.value`|String|Value for the variable or reference to the value.|
|`outputs.metadata`|List|User editable metadata for the variables. This list contains only one item.|
|`outputs.metadata.type`|String|Type of the variable.|
|`outputs.metadata.aliases`|List|List of aliases for the variable name.|
|`outputs.metadata.description`|String|Description of the meta data.|
|`outputs.metadata.default_value`|String|Default value for the variable, if the override value is not specified.|
|`outputs.metadata.secure`|Boolean|Is the variable secure or sensitive ?.|
|`outputs.metadata.immutable`|Boolean|Is the variable readonly ?.|
|`outputs.metadata.hidden`|Boolean|If true, the variable will not be displayed on UI or CLI.|
|`outputs.metadata.options`|List|List of possible values for this variable.  If type is integer or date, then the array of string will be  converted to array of integers or date during runtime.|
|`outputs.metadata.min_value`|Integer|Minimum value of the variable. Applicable for integer type.|
|`outputs.metadata.max_value`|Integer|Maximum value of the variable. Applicable for integer type.|
|`outputs.metadata.min_length`|Integer|Minimum length of the variable value. Applicable for string type.|
|`outputs.metadata.max_length`|Integer|Maximum length of the variable value. Applicable for string type.|
|`outputs.metadata.matches`|String|Regex for the variable value.|
|`outputs.metadata.position`|Integer|Relative position of this variable in a list.|
|`outputs.metadata.group_by`|String|Display name of the group this variable belongs to.|
|`outputs.metadata.source`|String|Source of this meta-data.|
|`outputs.link`|String|Reference link to the variable value By default the expression will point to self.value.|
|`settings`|List|Environment variables for an action.|
|`settings.name`|String|Name of the variable.|
|`settings.value`|String|Value for the variable or reference to the value.|
|`settings.metadata`|List|User editable metadata for the variables. This list contains only one item.|
|`settings.metadata.type`|String|Type of the variable.|
|`settings.metadata.aliases`|List|List of aliases for the variable name.|
|`settings.metadata.description`|String|Description of the meta data.|
|`settings.metadata.default_value`|String|Default value for the variable, if the override value is not specified.|
|`settings.metadata.secure`|Boolean|Is the variable secure or sensitive ?.|
|`settings.metadata.immutable`|Boolean|Is the variable readonly ?.|
|`settings.metadata.hidden`|Boolean|If true, the variable will not be displayed on UI or CLI.|
|`settings.metadata.options`|List|List of possible values for this variable.  If type is integer or date, then the array of string will be  converted to array of integers or date during runtime.|
|`settings.metadata.min_value`|Integer|Minimum value of the variable. Applicable for integer type.|
|`settings.metadata.max_value`|Integer|Maximum value of the variable. Applicable for integer type.|
|`settings.metadata.min_length`|Integer|Minimum length of the variable value. Applicable for string type.|
|`settings.metadata.max_length`|Integer|Maximum length of the variable value. Applicable for string type.|
|`settings.metadata.matches`|String|Regex for the variable value.|
|`settings.metadata.position`|Integer|Relative position of this variable in a list.|
|`settings.metadata.group_by`|String|Display name of the group this variable belongs to.|
|`settings.metadata.source`|String|Source of this meta-data.|
|`settings.link`|String|Reference link to the variable value By default the expression will point to self.value.|
|`trigger_record_id`|String|ID to the trigger.|
|`id`|String|Action ID.|
|`crn`|String|Action Cloud Resource Name.|
|`account`|String|Action account ID.|
|`source_created_at`|String|Action Playbook Source creation time.|
|`source_created_by`|String|E-mail address of user who created the Action Playbook Source.|
|`source_updated_at`|String|The action playbook updation time.|
|`source_updated_by`|String|E-mail address of user who updated the action playbook source.|
|`created_at`|String|Action creation time.|
|`created_by`|String|E-mail address of the user who created an action.|
|`updated_at`|String|Action updation time.|
|`updated_by`|String|E-mail address of the user who updated an action.|
|`namespace`|String|Name of the namespace.|
|`state`|List|Computed state of an action. This list contains only one item.|
|`state.status_code`|String|Status of automation (workspace or action).|
|`state.status_message`|String|Automation status message - to be displayed along with the status_code.|
|`playbook_names`|List|Playbook names retrieved from the respository.|
|`sys_lock`|List|System lock status. This list contains only one item.|
|`sys_lock.sys_locked`|Boolean|Is the Workspace locked by the Schematic action ?.|
|`sys_lock.sys_locked_by`|String|Name of the user who performed the action, that lead to lock the Workspace.|
|`sys_lock.sys_locked_at`|String|When the user performed the action that lead to lock the Workspace ?.|

## `ibm_schematics_job`
{: #schematics_job}

Retrieve information about Complete job with the user inputs and the system generated data..
{: shortdesc}

### Sample Terraform code
{: #schematics_job-sample}

```
data "ibm_schematics_job" "schematics_job" {
  job_id = "job_id"
}
```

### Input parameters
{: #schematics_job-input}

Review the input parameters that you can specify for your data source. {: shortdesc}

|Name|Data type|Required/optional|Description|
|----|-----------|-------|----------|
|`job_id`|String|Required|Use GET jobs API to look up the Job IDs in your IBM Cloud account.|

### Output parameters
{: #schematics_job-output}

Review the output parameters that you can access after you retrieved your data source. {: shortdesc}

|Name|Data type|Description|
|----|-----------|---------|
|`command_object`|String|Name of the Schematics automation resource.|
|`command_object_id`|String|Job command object ID (`workspace-id, action-id or control-id`).|
|`command_name`|String|Schematics job command name.|
|`command_parameter`|String|Schematics job command parameter (`playbook-name, capsule-name or flow-name`).|
|`command_options`|List|Command line options for the command.|
|`inputs`|List|Job inputs used by an action.|
|`inputs.name`|String|Name of the variable.|
|`inputs.value`|String|Value for the variable or reference to the value.|
|`inputs.metadata`|List|User editable metadata for the variables. This list contains only one item.|
|`inputs.metadata.type`|String|Type of the variable.|
|`inputs.metadata.aliases`|List|List of aliases for the variable name.|
|`inputs.metadata.description`|String|Description of the meta data.|
|`inputs.metadata.default_value`|String|Default value for the variable, if the override value is not specified.|
|`inputs.metadata.secure`|Boolean|Is the variable secure or sensitive ?.|
|`inputs.metadata.immutable`|Boolean|Is the variable readonly ?.|
|`inputs.metadata.hidden`|Boolean|If true, the variable will not be displayed on UI or CLI.|
|`inputs.metadata.options`|List|List of possible values for this variable.  If type is integer or date, then the array of string will be  converted to array of integers or date during runtime.|
|`inputs.metadata.min_value`|Integer|Minimum value of the variable. Applicable for integer type.|
|`inputs.metadata.max_value`|Integer|Maximum value of the variable. Applicable for integer type.|
|`inputs.metadata.min_length`|Integer|Minimum length of the variable value. Applicable for string type.|
|`inputs.metadata.max_length`|Integer|Maximum length of the variable value. Applicable for string type.|
|`inputs.metadata.matches`|String|Regex for the variable value.|
|`inputs.metadata.position`|Integer|Relative position of this variable in a list.|
|`inputs.metadata.group_by`|String|Display name of the group this variable belongs to.|
|`inputs.metadata.source`|String|Source of this meta-data.|
|`inputs.link`|String|Reference link to the variable value By default the expression will point to self.value.|
|`settings`|List|Environment variables used by the job while performing an action.|
|`settings.name`|String|Name of the variable.|
|`settings.value`|String|Value for the variable or reference to the value.|
|`settings.metadata`|List|User editable metadata for the variables. This list contains only one item.|
|`settings.metadata.type`|String|Type of the variable.|
|`settings.metadata.aliases`|List|List of aliases for the variable name.|
|`settings.metadata.description`|String|Description of the meta data.|
|`settings.metadata.default_value`|String|Default value for the variable, if the override value is not specified.|
|`settings.metadata.secure`|Boolean|Is the variable secure or sensitive ?.|
|`settings.metadata.immutable`|Boolean|Is the variable readonly ?.|
|`settings.metadata.hidden`|Boolean|If true, the variable will not be displayed on UI or CLI.|
|`settings.metadata.options`|List|List of possible values for this variable.  If type is integer or date, then the array of string will be  converted to array of integers or date during runtime.|
|`settings.metadata.min_value`|Integer|Minimum value of the variable. Applicable for integer type.|
|`settings.metadata.max_value`|Integer|Maximum value of the variable. Applicable for integer type.|
|`settings.metadata.min_length`|Integer|Minimum length of the variable value. Applicable for string type.|
|`settings.metadata.max_length`|Integer|Maximum length of the variable value. Applicable for string type.|
|`settings.metadata.matches`|String|Regex for the variable value.|
|`settings.metadata.position`|Integer|Relative position of this variable in a list.|
|`settings.metadata.group_by`|String|Display name of the group this variable belongs to.|
|`settings.metadata.source`|String|Source of this meta-data.|
|`settings.link`|String|Reference link to the variable value By default the expression will point to self.value.|
|`tags`|List|User defined tags, while running the job.|
|`id`|String|Job ID.|
|`name`|String|Job name, uniquely derived from the related action.|
|`description`|String|Job description derived from the related action.|
|`location`|String|List of action locations supported by IBM Cloud Schematics service.  **Note** this does not limit the location of the resources provisioned using Schematics.|
|`resource_group`|String|Resource group name derived from the related action.|
|`submitted_at`|String|Job submission time.|
|`submitted_by`|String|E-mail address of the user who submitted the job.|
|`start_at`|String|Job start time.|
|`end_at`|String|Job end time.|
|`duration`|String|Duration of job execution, for example, `40 sec`.|
|`status`|List|Job Status. This list contains only one item.|
|`status.action_job_status`|List|Action Job Status. This list contains only one item.|
|`status.action_job_status.action_name`|String|Action name.|
|`status.action_job_status.status_code`|String|Status of the jobs.|
|`status.action_job_status.status_message`|String|Action job status message to be displayed along with the `action_status_code`.|
|`status.action_job_status.bastion_status_code`|String|Status of the resources.|
|`status.action_job_status.bastion_status_message`|String|Bastion status message to be displayed along with the `bastion_status_code`.|
|`status.action_job_status.targets_status_code`|String|Status of the resources.|
|`status.action_job_status.targets_status_message`|String|Aggregated status message for all target resources, to be displayed along with the `targets_status_code`.|
|`status.action_job_status.updated_at`|String|Job status updation timestamp.|
|`data`|List|Job data. This list contains only one item.|
|`data.job_type`|String|Type of the job.|
|`data.action_job_data`|List|Action Job data. This list contains only one item.|
|`data.action_job_data.action_name`|String|Flow name.|
|`data.action_job_data.inputs`|List|Input variables data used by an action job.|
|`data.action_job_data.inputs.name`|String|Name of the variable.|
|`data.action_job_data.inputs.value`|String|Value for the variable or reference to the value.|
|`data.action_job_data.inputs.metadata`|List|User editable metadata for the variables. This list contains only one item.|
|`data.action_job_data.inputs.metadata.type`|String|Type of the variable.|
|`data.action_job_data.inputs.metadata.aliases`|List|List of aliases for the variable name.|
|`data.action_job_data.inputs.metadata.description`|String|Description of the meta data.|
|`data.action_job_data.inputs.metadata.default_value`|String|Default value for the variable, if the override value is not specified.|
|`data.action_job_data.inputs.metadata.secure`|Boolean|Is the variable secure or sensitive ?.|
|`data.action_job_data.inputs.metadata.immutable`|Boolean|Is the variable readonly ?.|
|`data.action_job_data.inputs.metadata.hidden`|Boolean|If true, the variable will not be displayed on UI or CLI.|
|`data.action_job_data.inputs.metadata.options`|List|List of possible values for this variable.  If type is integer or date, then the array of string will be  converted to array of integers or date during runtime.|
|`data.action_job_data.inputs.metadata.min_value`|Integer|Minimum value of the variable. Applicable for integer type.|
|`data.action_job_data.inputs.metadata.max_value`|Integer|Maximum value of the variable. Applicable for integer type.|
|`data.action_job_data.inputs.metadata.min_length`|Integer|Minimum length of the variable value. Applicable for string type.|
|`data.action_job_data.inputs.metadata.max_length`|Integer|Maximum length of the variable value. Applicable for string type.|
|`data.action_job_data.inputs.metadata.matches`|String|Regex for the variable value.|
|`data.action_job_data.inputs.metadata.position`|Integer|Relative position of this variable in a list.|
|`data.action_job_data.inputs.metadata.group_by`|String|Display name of the group this variable belongs to.|
|`data.action_job_data.inputs.metadata.source`|String|Source of this meta-data.|
|`data.action_job_data.inputs.link`|String|Reference link to the variable value By default the expression will point to self.value.|
|`data.action_job_data.outputs`|List|Output variables data from an action job.|
|`data.action_job_data.outputs.name`|String|Name of the variable.|
|`data.action_job_data.outputs.value`|String|Value for the variable or reference to the value.|
|`data.action_job_data.outputs.metadata`|List|User editable metadata for the variables. This list contains only one item.|
|`data.action_job_data.outputs.metadata.type`|String|Type of the variable.|
|`data.action_job_data.outputs.metadata.aliases`|List|List of aliases for the variable name.|
|`data.action_job_data.outputs.metadata.description`|String|Description of the meta data.|
|`data.action_job_data.outputs.metadata.default_value`|String|Default value for the variable, if the override value is not specified.|
|`data.action_job_data.outputs.metadata.secure`|Boolean|Is the variable secure or sensitive ?.|
|`data.action_job_data.outputs.metadata.immutable`|Boolean|Is the variable readonly ?.|
|`data.action_job_data.outputs.metadata.hidden`|Boolean|If true, the variable will not be displayed on UI or CLI.|
|`data.action_job_data.outputs.metadata.options`|List|List of possible values for this variable.  If type is integer or date, then the array of string will be  converted to array of integers or date during runtime.|
|`data.action_job_data.outputs.metadata.min_value`|Integer|Minimum value of the variable. Applicable for integer type.|
|`data.action_job_data.outputs.metadata.max_value`|Integer|Maximum value of the variable. Applicable for integer type.|
|`data.action_job_data.outputs.metadata.min_length`|Integer|Minimum length of the variable value. Applicable for string type.|
|`data.action_job_data.outputs.metadata.max_length`|Integer|Maximum length of the variable value. Applicable for string type.|
|`data.action_job_data.outputs.metadata.matches`|String|Regex for the variable value.|
|`data.action_job_data.outputs.metadata.position`|Integer|Relative position of this variable in a list.|
|`data.action_job_data.outputs.metadata.group_by`|String|Display name of the group this variable belongs to.|
|`data.action_job_data.outputs.metadata.source`|String|Source of this meta-data.|
|`data.action_job_data.outputs.link`|String|Reference link to the variable value By default the expression will point to self.value.|
|`data.action_job_data.settings`|List|Environment variables used by all the templates in an action.|
|`data.action_job_data.settings.name`|String|Name of the variable.|
|`data.action_job_data.settings.value`|String|Value for the variable or reference to the value.|
|`data.action_job_data.settings.metadata`|List|User editable metadata for the variables. This list contains only one item.|
|`data.action_job_data.settings.metadata.type`|String|Type of the variable.|
|`data.action_job_data.settings.metadata.aliases`|List|List of aliases for the variable name.|
|`data.action_job_data.settings.metadata.description`|String|Description of the meta data.|
|`data.action_job_data.settings.metadata.default_value`|String|Default value for the variable, if the override value is not specified.|
|`data.action_job_data.settings.metadata.secure`|Boolean|Is the variable secure or sensitive ?.|
|`data.action_job_data.settings.metadata.immutable`|Boolean|Is the variable readonly ?.|
|`data.action_job_data.settings.metadata.hidden`|Boolean|If true, the variable will not be displayed on UI or CLI.|
|`data.action_job_data.settings.metadata.options`|List|List of possible values for this variable.  If type is integer or date, then the array of string will be  converted to array of integers or date during runtime.|
|`data.action_job_data.settings.metadata.min_value`|Integer|Minimum value of the variable. Applicable for integer type.|
|`data.action_job_data.settings.metadata.max_value`|Integer|Maximum value of the variable. Applicable for integer type.|
|`data.action_job_data.settings.metadata.min_length`|Integer|Minimum length of the variable value. Applicable for string type.|
|`data.action_job_data.settings.metadata.max_length`|Integer|Maximum length of the variable value. Applicable for string type.|
|`data.action_job_data.settings.metadata.matches`|String|Regex for the variable value.|
|`data.action_job_data.settings.metadata.position`|Integer|Relative position of this variable in a list.|
|`data.action_job_data.settings.metadata.group_by`|String|Display name of the group this variable belongs to.|
|`data.action_job_data.settings.metadata.source`|String|Source of this meta-data.|
|`data.action_job_data.settings.link`|String|Reference link to the variable value By default the expression will point to self.value.|
|`data.action_job_data.updated_at`|String|Job status updation timestamp.|
|`targets_ini`|String|Inventory of host and host group for the playbook in `INI` file format. For example, `"targets_ini": "[webserverhost]  172.22.192.6  [dbhost]  172.22.192.5"`. For more information, about an inventory host group syntax, see [Inventory host groups](/docs/schematics?topic=schematics-schematics-cli-reference#schematics-inventory-host-grps).|
|`bastion`|List|Complete target details with the user inputs and the system generated data. This list contains only one item.|
|`bastion.name`|String|Target name.|
|`bastion.type`|String|Target type (`cluster`, `vsi`, `icd`, `vpc`).|
|`bastion.description`|String|Target description.|
|`bastion.resource_query`|String|Resource selection query string.|
|`bastion.credential`|String|Override credential for each resource.  Reference to credentials values, used by all the resources.|
|`bastion.id`|String|Target ID.|
|`bastion.created_at`|String|Targets creation time.|
|`bastion.created_by`|String|E-mail address of the user who created the targets.|
|`bastion.updated_at`|String|Targets updation time.|
|`bastion.updated_by`|String|E-mail address of user who updated the targets.|
|`bastion.sys_lock`|List|System lock status. This list contains only one item.|
|`bastion.sys_lock.sys_locked`|Boolean|Is the Workspace locked by the Schematic action ?.|
|`bastion.sys_lock.sys_locked_by`|String|Name of the user who performed the action, that lead to lock the Workspace.|
|`bastion.sys_lock.sys_locked_at`|String|When the user performed the action that lead to lock the Workspace ?.|
|`bastion.resource_ids`|List|Array of the resource IDs.|
|`log_summary`|List|Job log summary record. This list contains only one item.|
|`log_summary.job_id`|String|Workspace ID.|
|`log_summary.job_type`|String|Type of Job.|
|`log_summary.log_start_at`|String|Job log start timestamp.|
|`log_summary.log_analyzed_till`|String|Job log update timestamp.|
|`log_summary.elapsed_time`|Float|Job log elapsed time (`log_analyzed_till - log_start_at`).|
|`log_summary.log_errors`|List|Job log errors.|
|`log_summary.log_errors.error_code`|String|Error code in the Log.|
|`log_summary.log_errors.error_msg`|String|Summary error message in the log.|
|`log_summary.log_errors.error_count`|Float|Number of occurrence.|
|`log_summary.repo_download_job`|List|Repo download Job log summary. This list contains only one item.|
|`log_summary.repo_download_job.scanned_file_count`|Float|Number of files scanned.|
|`log_summary.repo_download_job.quarantined_file_count`|Float|Number of files quarantined.|
|`log_summary.repo_download_job.detected_filetype`|String|Detected template or data file type.|
|`log_summary.repo_download_job.inputs_count`|String|Number of inputs detected.|
|`log_summary.repo_download_job.outputs_count`|String|Number of outputs detected.|
|`log_summary.action_job`|List|Flow Job log summary. This list contains only one item.|
|`log_summary.action_job.target_count`|Float|number of targets or hosts.|
|`log_summary.action_job.task_count`|Float|number of tasks in playbook.|
|`log_summary.action_job.play_count`|Float|number of plays in playbook.|
|`log_summary.action_job.recap`|List|Recap records. This list contains only one item.|
|`log_summary.action_job.recap.target`|List|List of target or host name.|
|`log_summary.action_job.recap.ok`|Float|Number of OK.|
|`log_summary.action_job.recap.changed`|Float|Number of changed.|
|`log_summary.action_job.recap.failed`|Float|Number of failed.|
|`log_summary.action_job.recap.skipped`|Float|Number of skipped.|
|`log_summary.action_job.recap.unreachable`|Float|Number of unreachable.|
|`log_store_url`|String|Job log store URL.|
|`state_store_url`|String|Job state store URL.|
|`results_url`|String|Job results store URL.|
|`updated_at`|String|Job status updation timestamp.|

