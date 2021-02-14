/**
 * (C) Copyright IBM Corp. 2021.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package ibm

import (
	"fmt"
	"log"

	"github.com/IBM/schematics-go-sdk/schematicsv1"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceIBMSchematicsWorkspace() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMSchematicsWorkspaceRead,

		Schema: map[string]*schema.Schema{
			"w_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The workspace ID for the workspace that you want to query.  You can run the GET /workspaces call if you need to look up the  workspace IDs in your IBM Cloud account.",
			},
			"applied_shareddata_ids": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of applied shared dataset id.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"catalog_ref": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Computed:    true,
				Description: "CatalogRef -.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dry_run": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Dry run.",
						},
						"item_icon_url": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Catalog item icon url.",
						},
						"item_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Catalog item id.",
						},
						"item_name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Catalog item name.",
						},
						"item_readme_url": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Catalog item readme url.",
						},
						"item_url": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Catalog item url.",
						},
						"launch_url": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Catalog item launch url.",
						},
						"offering_version": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Catalog item offering version.",
						},
					},
				},
			},
			"created_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Workspace created at.",
			},
			"created_by": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Workspace created by.",
			},
			"crn": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Workspace CRN.",
			},
			"description": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Workspace description.",
			},
			"last_health_check_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Last health checked at.",
			},
			"location": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Workspace location.",
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Workspace name.",
			},
			"resource_group": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Workspace resource group.",
			},
			"runtime_data": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Workspace runtime data.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"engine_cmd": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Engine command.",
						},
						"engine_name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Engine name.",
						},
						"engine_version": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Engine version.",
						},
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Template id.",
						},
						"log_store_url": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Log store url.",
						},
						"output_values": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "List of Output values.",
							Elem: &schema.Schema{
								Type: schema.TypeMap,
							},
						},
						"resources": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "List of resources.",
							Elem: &schema.Schema{
								Type: schema.TypeMap,
							},
						},
						"state_store_url": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "State store URL.",
						},
					},
				},
			},
			"shared_data": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Computed:    true,
				Description: "SharedTargetDataResponse -.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cluster_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Target cluster id.",
						},
						"cluster_name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Target cluster name.",
						},
						"entitlement_keys": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Entitlement keys.",
							Elem: &schema.Schema{
								Type: schema.TypeMap,
							},
						},
						"namespace": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Target namespace.",
						},
						"region": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Target region.",
						},
						"resource_group_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Target resource group id.",
						},
					},
				},
			},
			"status": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Workspace status type.",
			},
			"tags": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Workspace tags.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"template_env_settings": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "EnvVariableRequest ..",
				Elem:        &schema.Schema{Type: schema.TypeMap},
			},
			"template_git_folder": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Folder name.",
			},
			"template_init_state_file": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Init state file.",
			},
			"template_type": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Template type.",
			},
			"template_uninstall_script_name": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Uninstall script name.",
			},
			"template_values": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Value.",
			},
			"template_values_metadata": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Description: "List of values metadata.",
				Elem:        &schema.Schema{Type: schema.TypeMap},
			},
			"template_inputs": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "VariablesRequest -.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"description": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Variable description.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Variable name.",
						},
						"secure": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Variable is secure.",
						},
						"type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Variable type.",
						},
						"use_default": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Variable uses default value; and is not over-ridden.",
						},
						"value": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Value of the Variable.",
						},
					},
				},
			},
			"template_ref": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Workspace template ref.",
			},
			"template_git_branch": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Repo branch.",
			},
			"template_git_full_url": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Full repo URL.",
			},
			"template_git_has_uploadedgitrepotar": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "Has uploaded git repo tar",
			},
			"template_git_release": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Repo release.",
			},
			"template_git_repo_sha_value": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Repo SHA value.",
			},
			"template_git_repo_url": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Repo URL.",
			},
			"template_git_url": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Source URL.",
			},

			/*"template_type": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List of Workspace type.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},*/
			"updated_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Workspace updated at.",
			},
			"updated_by": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Workspace updated by.",
			},
			"frozen": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Frozen status.",
			},
			"frozen_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Frozen at.",
			},
			"frozen_by": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Frozen by.",
			},
			"locked": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Locked status.",
			},
			"locked_by": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Locked by.",
			},
			"locked_time": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Locked at.",
			},
			"status_code": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Status code.",
			},
			"status_msg": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Status message.",
			},
		},
	}
}

func dataSourceIBMSchematicsWorkspaceRead(d *schema.ResourceData, meta interface{}) error {
	schematicsClient, err := meta.(ClientSession).SchematicsV1()
	if err != nil {
		return err
	}

	getWorkspaceOptions := &schematicsv1.GetWorkspaceOptions{}

	getWorkspaceOptions.SetWID(d.Get("w_id").(string))

	workspaceResponse, response, err := schematicsClient.GetWorkspace(getWorkspaceOptions)
	if err != nil {
		log.Printf("[DEBUG] GetWorkspace failed %s\n%s", err, response)
		return err
	}

	d.SetId(*workspaceResponse.ID)
	if err = d.Set("applied_shareddata_ids", workspaceResponse.AppliedShareddataIds); err != nil {
		return fmt.Errorf("Error setting applied_shareddata_ids: %s", err)
	}

	if workspaceResponse.CatalogRef != nil {
		err = d.Set("catalog_ref", dataSourceWorkspaceResponseFlattenCatalogRef(*workspaceResponse.CatalogRef))
		if err != nil {
			return fmt.Errorf("Error setting catalog_ref %s", err)
		}
	}
	if err = d.Set("created_at", workspaceResponse.CreatedAt.String()); err != nil {
		return fmt.Errorf("Error setting created_at: %s", err)
	}
	if err = d.Set("created_by", workspaceResponse.CreatedBy); err != nil {
		return fmt.Errorf("Error setting created_by: %s", err)
	}
	if err = d.Set("crn", workspaceResponse.Crn); err != nil {
		return fmt.Errorf("Error setting crn: %s", err)
	}
	if err = d.Set("description", workspaceResponse.Description); err != nil {
		return fmt.Errorf("Error setting description: %s", err)
	}
	if err = d.Set("last_health_check_at", workspaceResponse.LastHealthCheckAt.String()); err != nil {
		return fmt.Errorf("Error setting last_health_check_at: %s", err)
	}
	if err = d.Set("location", workspaceResponse.Location); err != nil {
		return fmt.Errorf("Error setting location: %s", err)
	}
	if err = d.Set("name", workspaceResponse.Name); err != nil {
		return fmt.Errorf("Error setting name: %s", err)
	}
	if err = d.Set("resource_group", workspaceResponse.ResourceGroup); err != nil {
		return fmt.Errorf("Error setting resource_group: %s", err)
	}

	if workspaceResponse.RuntimeData != nil {
		err = d.Set("runtime_data", dataSourceWorkspaceResponseFlattenRuntimeData(workspaceResponse.RuntimeData))
		if err != nil {
			return fmt.Errorf("Error setting runtime_data %s", err)
		}
	}

	if workspaceResponse.SharedData != nil {
		err = d.Set("shared_data", dataSourceWorkspaceResponseFlattenSharedData(*workspaceResponse.SharedData))
		if err != nil {
			return fmt.Errorf("Error setting shared_data %s", err)
		}
	}
	if err = d.Set("status", workspaceResponse.Status); err != nil {
		return fmt.Errorf("Error setting status: %s", err)
	}
	if err = d.Set("tags", workspaceResponse.Tags); err != nil {
		return fmt.Errorf("Error setting tags: %s", err)
	}

	if workspaceResponse.TemplateData != nil {
		templateData := dataSourceWorkspaceResponseFlattenTemplateData(workspaceResponse.TemplateData)

		if err = d.Set("template_env_settings", templateData[0]["env_values"]); err != nil {
			return fmt.Errorf("Error reading env_values: %s", err)
		}
		if err = d.Set("template_git_folder", templateData[0]["folder"]); err != nil {
			return fmt.Errorf("Error reading folder: %s", err)
		}
		if err = d.Set("template_init_state_file", templateData[0]["init_state_file"]); err != nil {
			return fmt.Errorf("Error reading init_state_file: %s", err)
		}
		if err = d.Set("template_type", templateData[0]["type"]); err != nil {
			return fmt.Errorf("Error reading type: %s", err)
		}
		if err = d.Set("template_uninstall_script_name", templateData[0]["uninstall_script_name"]); err != nil {
			return fmt.Errorf("Error reading uninstall_script_name: %s", err)
		}
		if err = d.Set("template_values", templateData[0]["values"]); err != nil {
			return fmt.Errorf("Error reading values: %s", err)
		}
		if err = d.Set("template_values_metadata", templateData[0]["values_metadata"]); err != nil {
			return fmt.Errorf("Error reading values_metadata: %s", err)
		}
		if err = d.Set("template_inputs", templateData[0]["variablestore"]); err != nil {
			return fmt.Errorf("Error reading variablestore: %s", err)
		}
	}

	if err = d.Set("template_ref", workspaceResponse.TemplateRef); err != nil {
		return fmt.Errorf("Error setting template_ref: %s", err)
	}

	if workspaceResponse.TemplateRepo != nil {
		templateRepoMap := dataSourceWorkspaceResponseFlattenTemplateRepo(*workspaceResponse.TemplateRepo)
		if err = d.Set("template_git_branch", templateRepoMap[0]["branch"]); err != nil {
			return fmt.Errorf("Error reading branch: %s", err)
		}
		if err = d.Set("template_git_release", templateRepoMap[0]["release"]); err != nil {
			return fmt.Errorf("Error reading release: %s", err)
		}
		if err = d.Set("template_git_repo_sha_value", templateRepoMap[0]["repo_sha_value"]); err != nil {
			return fmt.Errorf("Error reading repo_sha_value: %s", err)
		}
		if err = d.Set("template_git_repo_url", templateRepoMap[0]["repo_url"]); err != nil {
			return fmt.Errorf("Error reading repo_url: %s", err)
		}
		if err = d.Set("template_git_url", templateRepoMap[0]["url"]); err != nil {
			return fmt.Errorf("Error reading url: %s", err)
		}
		if err = d.Set("template_git_has_uploadedgitrepotar", templateRepoMap[0]["has_uploadedgitrepotar"]); err != nil {
			return fmt.Errorf("Error reading has_uploadedgitrepotar: %s", err)
		}
	}
	/*if err = d.Set("type", workspaceResponse.Type); err != nil {
		return fmt.Errorf("Error setting type: %s", err)
	}*/
	if workspaceResponse.UpdatedAt != nil {
		if err = d.Set("updated_at", workspaceResponse.UpdatedAt.String()); err != nil {
			return fmt.Errorf("Error setting updated_at: %s", err)
		}
	}
	if err = d.Set("updated_by", workspaceResponse.UpdatedBy); err != nil {
		return fmt.Errorf("Error setting updated_by: %s", err)
	}

	if workspaceResponse.WorkspaceStatus != nil {
		workspaceStatusMap := dataSourceWorkspaceResponseFlattenWorkspaceStatus(*workspaceResponse.WorkspaceStatus)
		if err = d.Set("frozen", workspaceStatusMap[0]["frozen"]); err != nil {
			return fmt.Errorf("Error reading frozen: %s", err)
		}
		if err = d.Set("frozen_at", workspaceStatusMap[0]["frozen_at"]); err != nil {
			return fmt.Errorf("Error reading frozen_at: %s", err)
		}
		if err = d.Set("frozen_by", workspaceStatusMap[0]["frozen_by"]); err != nil {
			return fmt.Errorf("Error reading frozen_by: %s", err)
		}
		if err = d.Set("locked", workspaceStatusMap[0]["locked"]); err != nil {
			return fmt.Errorf("Error reading locked: %s", err)
		}
		if err = d.Set("locked_by", workspaceStatusMap[0]["locked_by"]); err != nil {
			return fmt.Errorf("Error reading locked_by: %s", err)
		}
		if err = d.Set("locked_time", workspaceStatusMap[0]["locked_time"]); err != nil {
			return fmt.Errorf("Error reading locked_time: %s", err)
		}
	}

	if workspaceResponse.WorkspaceStatusMsg != nil {
		workspaceStatusMsgMap := dataSourceWorkspaceResponseFlattenWorkspaceStatusMsg(*workspaceResponse.WorkspaceStatusMsg)
		if err = d.Set("status_code", workspaceStatusMsgMap[0]["status_code"]); err != nil {
			return fmt.Errorf("Error reading status_code: %s", err)
		}
		if err = d.Set("status_msg", workspaceStatusMsgMap[0]["status_msg"]); err != nil {
			return fmt.Errorf("Error reading status_msg: %s", err)
		}
	}

	return nil
}

func dataSourceWorkspaceResponseFlattenCatalogRef(result schematicsv1.CatalogRef) (finalList []map[string]interface{}) {
	finalList = []map[string]interface{}{}
	finalMap := dataSourceWorkspaceResponseCatalogRefToMap(result)
	finalList = append(finalList, finalMap)

	return finalList
}

func dataSourceWorkspaceResponseCatalogRefToMap(catalogRefItem schematicsv1.CatalogRef) (catalogRefMap map[string]interface{}) {
	catalogRefMap = map[string]interface{}{}

	if catalogRefItem.DryRun != nil {
		catalogRefMap["dry_run"] = catalogRefItem.DryRun
	}
	if catalogRefItem.ItemIconURL != nil {
		catalogRefMap["item_icon_url"] = catalogRefItem.ItemIconURL
	}
	if catalogRefItem.ItemID != nil {
		catalogRefMap["item_id"] = catalogRefItem.ItemID
	}
	if catalogRefItem.ItemName != nil {
		catalogRefMap["item_name"] = catalogRefItem.ItemName
	}
	if catalogRefItem.ItemReadmeURL != nil {
		catalogRefMap["item_readme_url"] = catalogRefItem.ItemReadmeURL
	}
	if catalogRefItem.ItemURL != nil {
		catalogRefMap["item_url"] = catalogRefItem.ItemURL
	}
	if catalogRefItem.LaunchURL != nil {
		catalogRefMap["launch_url"] = catalogRefItem.LaunchURL
	}
	if catalogRefItem.OfferingVersion != nil {
		catalogRefMap["offering_version"] = catalogRefItem.OfferingVersion
	}

	return catalogRefMap
}

func dataSourceWorkspaceResponseFlattenRuntimeData(result []schematicsv1.TemplateRunTimeDataResponse) (runtimeData []map[string]interface{}) {
	for _, runtimeDataItem := range result {
		runtimeData = append(runtimeData, dataSourceWorkspaceResponseRuntimeDataToMap(runtimeDataItem))
	}

	return runtimeData
}

func dataSourceWorkspaceResponseRuntimeDataToMap(runtimeDataItem schematicsv1.TemplateRunTimeDataResponse) (runtimeDataMap map[string]interface{}) {
	runtimeDataMap = map[string]interface{}{}

	if runtimeDataItem.EngineCmd != nil {
		runtimeDataMap["engine_cmd"] = runtimeDataItem.EngineCmd
	}
	if runtimeDataItem.EngineName != nil {
		runtimeDataMap["engine_name"] = runtimeDataItem.EngineName
	}
	if runtimeDataItem.EngineVersion != nil {
		runtimeDataMap["engine_version"] = runtimeDataItem.EngineVersion
	}
	if runtimeDataItem.ID != nil {
		runtimeDataMap["id"] = runtimeDataItem.ID
	}
	if runtimeDataItem.LogStoreURL != nil {
		runtimeDataMap["log_store_url"] = runtimeDataItem.LogStoreURL
	}
	if runtimeDataItem.OutputValues != nil {
		runtimeDataMap["output_values"] = runtimeDataItem.OutputValues
	}
	if runtimeDataItem.Resources != nil {
		runtimeDataMap["resources"] = runtimeDataItem.Resources
	}
	if runtimeDataItem.StateStoreURL != nil {
		runtimeDataMap["state_store_url"] = runtimeDataItem.StateStoreURL
	}

	return runtimeDataMap
}

func dataSourceWorkspaceResponseFlattenSharedData(result schematicsv1.SharedTargetDataResponse) (finalList []map[string]interface{}) {
	finalList = []map[string]interface{}{}
	finalMap := dataSourceWorkspaceResponseSharedDataToMap(result)
	finalList = append(finalList, finalMap)

	return finalList
}

func dataSourceWorkspaceResponseSharedDataToMap(sharedDataItem schematicsv1.SharedTargetDataResponse) (sharedDataMap map[string]interface{}) {
	sharedDataMap = map[string]interface{}{}

	if sharedDataItem.ClusterID != nil {
		sharedDataMap["cluster_id"] = sharedDataItem.ClusterID
	}
	if sharedDataItem.ClusterName != nil {
		sharedDataMap["cluster_name"] = sharedDataItem.ClusterName
	}
	if sharedDataItem.EntitlementKeys != nil {
		sharedDataMap["entitlement_keys"] = sharedDataItem.EntitlementKeys
	}
	if sharedDataItem.Namespace != nil {
		sharedDataMap["namespace"] = sharedDataItem.Namespace
	}
	if sharedDataItem.Region != nil {
		sharedDataMap["region"] = sharedDataItem.Region
	}
	if sharedDataItem.ResourceGroupID != nil {
		sharedDataMap["resource_group_id"] = sharedDataItem.ResourceGroupID
	}

	return sharedDataMap
}

func dataSourceWorkspaceResponseFlattenTemplateData(result []schematicsv1.TemplateSourceDataResponse) (templateData []map[string]interface{}) {
	for _, templateDataItem := range result {
		templateData = append(templateData, dataSourceWorkspaceResponseTemplateDataToMap(templateDataItem))
	}

	return templateData
}

func dataSourceWorkspaceResponseTemplateDataToMap(templateDataItem schematicsv1.TemplateSourceDataResponse) (templateDataMap map[string]interface{}) {
	templateDataMap = map[string]interface{}{}

	if templateDataItem.EnvValues != nil {
		envValuesList := []map[string]interface{}{}
		for _, envValuesItem := range templateDataItem.EnvValues {
			envValuesList = append(envValuesList, dataSourceWorkspaceResponseTemplateDataEnvValuesToMap(envValuesItem))
		}
		templateDataMap["env_values"] = envValuesList
	}
	if templateDataItem.Folder != nil {
		templateDataMap["folder"] = templateDataItem.Folder
	}
	if templateDataItem.HasGithubtoken != nil {
		templateDataMap["has_githubtoken"] = templateDataItem.HasGithubtoken
	}
	if templateDataItem.ID != nil {
		templateDataMap["id"] = templateDataItem.ID
	}
	if templateDataItem.Type != nil {
		templateDataMap["type"] = templateDataItem.Type
	}
	if templateDataItem.UninstallScriptName != nil {
		templateDataMap["uninstall_script_name"] = templateDataItem.UninstallScriptName
	}
	if templateDataItem.Values != nil {
		templateDataMap["values"] = templateDataItem.Values
	}
	if templateDataItem.ValuesMetadata != nil {
		templateDataMap["values_metadata"] = templateDataItem.ValuesMetadata
	}
	if templateDataItem.ValuesURL != nil {
		templateDataMap["values_url"] = templateDataItem.ValuesURL
	}
	if templateDataItem.Variablestore != nil {
		variablestoreList := []map[string]interface{}{}
		for _, variablestoreItem := range templateDataItem.Variablestore {
			variablestoreList = append(variablestoreList, dataSourceWorkspaceResponseTemplateDataVariablestoreToMap(variablestoreItem))
		}
		templateDataMap["variablestore"] = variablestoreList
	}

	return templateDataMap
}

func dataSourceWorkspaceResponseTemplateDataEnvValuesToMap(envValuesItem schematicsv1.EnvVariableResponse) (envValuesMap map[string]interface{}) {
	envValuesMap = map[string]interface{}{}

	if envValuesItem.Hidden != nil {
		envValuesMap["hidden"] = envValuesItem.Hidden
	}
	if envValuesItem.Name != nil {
		envValuesMap["name"] = envValuesItem.Name
	}
	if envValuesItem.Secure != nil {
		envValuesMap["secure"] = envValuesItem.Secure
	}
	if envValuesItem.Value != nil {
		envValuesMap["value"] = envValuesItem.Value
	}

	return envValuesMap
}

func dataSourceWorkspaceResponseTemplateDataVariablestoreToMap(variablestoreItem schematicsv1.WorkspaceVariableResponse) (variablestoreMap map[string]interface{}) {
	variablestoreMap = map[string]interface{}{}

	if variablestoreItem.Description != nil {
		variablestoreMap["description"] = variablestoreItem.Description
	}
	if variablestoreItem.Name != nil {
		variablestoreMap["name"] = variablestoreItem.Name
	}
	if variablestoreItem.Secure != nil {
		variablestoreMap["secure"] = variablestoreItem.Secure
	}
	if variablestoreItem.Type != nil {
		variablestoreMap["type"] = variablestoreItem.Type
	}
	if variablestoreItem.Value != nil {
		variablestoreMap["value"] = variablestoreItem.Value
	}

	return variablestoreMap
}

func dataSourceWorkspaceResponseFlattenTemplateRepo(result schematicsv1.TemplateRepoResponse) (finalList []map[string]interface{}) {
	finalList = []map[string]interface{}{}
	finalMap := dataSourceWorkspaceResponseTemplateRepoToMap(result)
	finalList = append(finalList, finalMap)

	return finalList
}

func dataSourceWorkspaceResponseTemplateRepoToMap(templateRepoItem schematicsv1.TemplateRepoResponse) (templateRepoMap map[string]interface{}) {
	templateRepoMap = map[string]interface{}{}

	if templateRepoItem.Branch != nil {
		templateRepoMap["branch"] = templateRepoItem.Branch
	}
	if templateRepoItem.FullURL != nil {
		templateRepoMap["full_url"] = templateRepoItem.FullURL
	}
	if templateRepoItem.HasUploadedgitrepotar != nil {
		templateRepoMap["has_uploadedgitrepotar"] = templateRepoItem.HasUploadedgitrepotar
	}
	if templateRepoItem.Release != nil {
		templateRepoMap["release"] = templateRepoItem.Release
	}
	if templateRepoItem.RepoShaValue != nil {
		templateRepoMap["repo_sha_value"] = templateRepoItem.RepoShaValue
	}
	if templateRepoItem.RepoURL != nil {
		templateRepoMap["repo_url"] = templateRepoItem.RepoURL
	}
	if templateRepoItem.URL != nil {
		templateRepoMap["url"] = templateRepoItem.URL
	}

	return templateRepoMap
}

func dataSourceWorkspaceResponseFlattenWorkspaceStatus(result schematicsv1.WorkspaceStatusResponse) (finalList []map[string]interface{}) {
	finalList = []map[string]interface{}{}
	finalMap := dataSourceWorkspaceResponseWorkspaceStatusToMap(result)
	finalList = append(finalList, finalMap)

	return finalList
}

func dataSourceWorkspaceResponseWorkspaceStatusToMap(workspaceStatusItem schematicsv1.WorkspaceStatusResponse) (workspaceStatusMap map[string]interface{}) {
	workspaceStatusMap = map[string]interface{}{}

	if workspaceStatusItem.Frozen != nil {
		workspaceStatusMap["frozen"] = workspaceStatusItem.Frozen
	}
	if workspaceStatusItem.FrozenAt != nil {
		workspaceStatusMap["frozen_at"] = workspaceStatusItem.FrozenAt
	}
	if workspaceStatusItem.FrozenBy != nil {
		workspaceStatusMap["frozen_by"] = workspaceStatusItem.FrozenBy
	}
	if workspaceStatusItem.Locked != nil {
		workspaceStatusMap["locked"] = workspaceStatusItem.Locked
	}
	if workspaceStatusItem.LockedBy != nil {
		workspaceStatusMap["locked_by"] = workspaceStatusItem.LockedBy
	}
	if workspaceStatusItem.LockedTime != nil {
		workspaceStatusMap["locked_time"] = workspaceStatusItem.LockedTime
	}

	return workspaceStatusMap
}

func dataSourceWorkspaceResponseFlattenWorkspaceStatusMsg(result schematicsv1.WorkspaceStatusMessage) (finalList []map[string]interface{}) {
	finalList = []map[string]interface{}{}
	finalMap := dataSourceWorkspaceResponseWorkspaceStatusMsgToMap(result)
	finalList = append(finalList, finalMap)

	return finalList
}

func dataSourceWorkspaceResponseWorkspaceStatusMsgToMap(workspaceStatusMsgItem schematicsv1.WorkspaceStatusMessage) (workspaceStatusMsgMap map[string]interface{}) {
	workspaceStatusMsgMap = map[string]interface{}{}

	if workspaceStatusMsgItem.StatusCode != nil {
		workspaceStatusMsgMap["status_code"] = workspaceStatusMsgItem.StatusCode
	}
	if workspaceStatusMsgItem.StatusMsg != nil {
		workspaceStatusMsgMap["status_msg"] = workspaceStatusMsgItem.StatusMsg
	}

	return workspaceStatusMsgMap
}
