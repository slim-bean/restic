// +build go1.9

// Copyright 2017 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder
// commit ID: dab57ee609fffdc578f48519d5cdc980efd8cc00

package insights

import original "github.com/Azure/azure-sdk-for-go/services/monitor/mgmt/2017-05-01-preview/insights"

type LogProfilesClient = original.LogProfilesClient
type CategoryType = original.CategoryType

const (
	Logs    CategoryType = original.Logs
	Metrics CategoryType = original.Metrics
)

type ComparisonOperationType = original.ComparisonOperationType

const (
	Equals             ComparisonOperationType = original.Equals
	GreaterThan        ComparisonOperationType = original.GreaterThan
	GreaterThanOrEqual ComparisonOperationType = original.GreaterThanOrEqual
	LessThan           ComparisonOperationType = original.LessThan
	LessThanOrEqual    ComparisonOperationType = original.LessThanOrEqual
	NotEquals          ComparisonOperationType = original.NotEquals
)

type ConditionOperator = original.ConditionOperator

const (
	ConditionOperatorGreaterThan        ConditionOperator = original.ConditionOperatorGreaterThan
	ConditionOperatorGreaterThanOrEqual ConditionOperator = original.ConditionOperatorGreaterThanOrEqual
	ConditionOperatorLessThan           ConditionOperator = original.ConditionOperatorLessThan
	ConditionOperatorLessThanOrEqual    ConditionOperator = original.ConditionOperatorLessThanOrEqual
)

type MetricStatisticType = original.MetricStatisticType

const (
	Average MetricStatisticType = original.Average
	Max     MetricStatisticType = original.Max
	Min     MetricStatisticType = original.Min
	Sum     MetricStatisticType = original.Sum
)

type OdataType = original.OdataType

const (
	OdataTypeMicrosoftAzureManagementInsightsModelsRuleManagementEventDataSource OdataType = original.OdataTypeMicrosoftAzureManagementInsightsModelsRuleManagementEventDataSource
	OdataTypeMicrosoftAzureManagementInsightsModelsRuleMetricDataSource          OdataType = original.OdataTypeMicrosoftAzureManagementInsightsModelsRuleMetricDataSource
)

type OdataType1 = original.OdataType1

const (
	OdataTypeMicrosoftAzureManagementInsightsModelsLocationThresholdRuleCondition OdataType1 = original.OdataTypeMicrosoftAzureManagementInsightsModelsLocationThresholdRuleCondition
	OdataTypeMicrosoftAzureManagementInsightsModelsManagementEventRuleCondition   OdataType1 = original.OdataTypeMicrosoftAzureManagementInsightsModelsManagementEventRuleCondition
	OdataTypeMicrosoftAzureManagementInsightsModelsThresholdRuleCondition         OdataType1 = original.OdataTypeMicrosoftAzureManagementInsightsModelsThresholdRuleCondition
)

type OdataType2 = original.OdataType2

const (
	OdataTypeMicrosoftAzureManagementInsightsModelsRuleEmailAction   OdataType2 = original.OdataTypeMicrosoftAzureManagementInsightsModelsRuleEmailAction
	OdataTypeMicrosoftAzureManagementInsightsModelsRuleWebhookAction OdataType2 = original.OdataTypeMicrosoftAzureManagementInsightsModelsRuleWebhookAction
)

type ReceiverStatus = original.ReceiverStatus

const (
	Disabled     ReceiverStatus = original.Disabled
	Enabled      ReceiverStatus = original.Enabled
	NotSpecified ReceiverStatus = original.NotSpecified
)

type RecurrenceFrequency = original.RecurrenceFrequency

const (
	Day    RecurrenceFrequency = original.Day
	Hour   RecurrenceFrequency = original.Hour
	Minute RecurrenceFrequency = original.Minute
	Month  RecurrenceFrequency = original.Month
	None   RecurrenceFrequency = original.None
	Second RecurrenceFrequency = original.Second
	Week   RecurrenceFrequency = original.Week
	Year   RecurrenceFrequency = original.Year
)

type ScaleDirection = original.ScaleDirection

const (
	ScaleDirectionDecrease ScaleDirection = original.ScaleDirectionDecrease
	ScaleDirectionIncrease ScaleDirection = original.ScaleDirectionIncrease
	ScaleDirectionNone     ScaleDirection = original.ScaleDirectionNone
)

type ScaleType = original.ScaleType

const (
	ChangeCount        ScaleType = original.ChangeCount
	ExactCount         ScaleType = original.ExactCount
	PercentChangeCount ScaleType = original.PercentChangeCount
)

type TimeAggregationOperator = original.TimeAggregationOperator

const (
	TimeAggregationOperatorAverage TimeAggregationOperator = original.TimeAggregationOperatorAverage
	TimeAggregationOperatorLast    TimeAggregationOperator = original.TimeAggregationOperatorLast
	TimeAggregationOperatorMaximum TimeAggregationOperator = original.TimeAggregationOperatorMaximum
	TimeAggregationOperatorMinimum TimeAggregationOperator = original.TimeAggregationOperatorMinimum
	TimeAggregationOperatorTotal   TimeAggregationOperator = original.TimeAggregationOperatorTotal
)

type TimeAggregationType = original.TimeAggregationType

const (
	TimeAggregationTypeAverage TimeAggregationType = original.TimeAggregationTypeAverage
	TimeAggregationTypeCount   TimeAggregationType = original.TimeAggregationTypeCount
	TimeAggregationTypeMaximum TimeAggregationType = original.TimeAggregationTypeMaximum
	TimeAggregationTypeMinimum TimeAggregationType = original.TimeAggregationTypeMinimum
	TimeAggregationTypeTotal   TimeAggregationType = original.TimeAggregationTypeTotal
)

type ActionGroup = original.ActionGroup
type ActionGroupList = original.ActionGroupList
type ActionGroupResource = original.ActionGroupResource
type ActivityLogAlert = original.ActivityLogAlert
type ActivityLogAlertActionGroup = original.ActivityLogAlertActionGroup
type ActivityLogAlertActionList = original.ActivityLogAlertActionList
type ActivityLogAlertAllOfCondition = original.ActivityLogAlertAllOfCondition
type ActivityLogAlertLeafCondition = original.ActivityLogAlertLeafCondition
type ActivityLogAlertList = original.ActivityLogAlertList
type ActivityLogAlertPatch = original.ActivityLogAlertPatch
type ActivityLogAlertPatchBody = original.ActivityLogAlertPatchBody
type ActivityLogAlertResource = original.ActivityLogAlertResource
type AlertRule = original.AlertRule
type AlertRuleResource = original.AlertRuleResource
type AlertRuleResourceCollection = original.AlertRuleResourceCollection
type AlertRuleResourcePatch = original.AlertRuleResourcePatch
type AutoscaleNotification = original.AutoscaleNotification
type AutoscaleProfile = original.AutoscaleProfile
type AutoscaleSetting = original.AutoscaleSetting
type AutoscaleSettingResource = original.AutoscaleSettingResource
type AutoscaleSettingResourceCollection = original.AutoscaleSettingResourceCollection
type AutoscaleSettingResourcePatch = original.AutoscaleSettingResourcePatch
type DiagnosticSettings = original.DiagnosticSettings
type DiagnosticSettingsCategory = original.DiagnosticSettingsCategory
type DiagnosticSettingsCategoryResource = original.DiagnosticSettingsCategoryResource
type DiagnosticSettingsCategoryResourceCollection = original.DiagnosticSettingsCategoryResourceCollection
type DiagnosticSettingsResource = original.DiagnosticSettingsResource
type DiagnosticSettingsResourceCollection = original.DiagnosticSettingsResourceCollection
type EmailNotification = original.EmailNotification
type EmailReceiver = original.EmailReceiver
type EnableRequest = original.EnableRequest
type ErrorResponse = original.ErrorResponse
type Incident = original.Incident
type IncidentListResult = original.IncidentListResult
type LocationThresholdRuleCondition = original.LocationThresholdRuleCondition
type LogProfileCollection = original.LogProfileCollection
type LogProfileProperties = original.LogProfileProperties
type LogProfileResource = original.LogProfileResource
type LogProfileResourcePatch = original.LogProfileResourcePatch
type LogSettings = original.LogSettings
type ManagementEventAggregationCondition = original.ManagementEventAggregationCondition
type ManagementEventRuleCondition = original.ManagementEventRuleCondition
type MetricSettings = original.MetricSettings
type MetricTrigger = original.MetricTrigger
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type ProxyOnlyResource = original.ProxyOnlyResource
type Recurrence = original.Recurrence
type RecurrentSchedule = original.RecurrentSchedule
type Resource = original.Resource
type RetentionPolicy = original.RetentionPolicy
type RuleAction = original.RuleAction
type RuleCondition = original.RuleCondition
type RuleDataSource = original.RuleDataSource
type RuleEmailAction = original.RuleEmailAction
type RuleManagementEventClaimsDataSource = original.RuleManagementEventClaimsDataSource
type RuleManagementEventDataSource = original.RuleManagementEventDataSource
type RuleMetricDataSource = original.RuleMetricDataSource
type RuleWebhookAction = original.RuleWebhookAction
type ScaleAction = original.ScaleAction
type ScaleCapacity = original.ScaleCapacity
type ScaleRule = original.ScaleRule
type SmsReceiver = original.SmsReceiver
type ThresholdRuleCondition = original.ThresholdRuleCondition
type TimeWindow = original.TimeWindow
type WebhookNotification = original.WebhookNotification
type WebhookReceiver = original.WebhookReceiver
type ActivityLogAlertsClient = original.ActivityLogAlertsClient

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ManagementClient = original.ManagementClient
type DiagnosticSettingsCategoryClient = original.DiagnosticSettingsCategoryClient
type AutoscaleSettingsClient = original.AutoscaleSettingsClient
type DiagnosticSettingsClient = original.DiagnosticSettingsClient
type OperationsClient = original.OperationsClient
type ActionGroupsClient = original.ActionGroupsClient
type AlertRuleIncidentsClient = original.AlertRuleIncidentsClient
type AlertRulesClient = original.AlertRulesClient

func NewLogProfilesClient(subscriptionID string) LogProfilesClient {
	return original.NewLogProfilesClient(subscriptionID)
}
func NewLogProfilesClientWithBaseURI(baseURI string, subscriptionID string) LogProfilesClient {
	return original.NewLogProfilesClientWithBaseURI(baseURI, subscriptionID)
}
func NewActivityLogAlertsClient(subscriptionID string) ActivityLogAlertsClient {
	return original.NewActivityLogAlertsClient(subscriptionID)
}
func NewActivityLogAlertsClientWithBaseURI(baseURI string, subscriptionID string) ActivityLogAlertsClient {
	return original.NewActivityLogAlertsClientWithBaseURI(baseURI, subscriptionID)
}
func New(subscriptionID string) ManagementClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) ManagementClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func NewDiagnosticSettingsCategoryClient(subscriptionID string) DiagnosticSettingsCategoryClient {
	return original.NewDiagnosticSettingsCategoryClient(subscriptionID)
}
func NewDiagnosticSettingsCategoryClientWithBaseURI(baseURI string, subscriptionID string) DiagnosticSettingsCategoryClient {
	return original.NewDiagnosticSettingsCategoryClientWithBaseURI(baseURI, subscriptionID)
}
func NewAutoscaleSettingsClient(subscriptionID string) AutoscaleSettingsClient {
	return original.NewAutoscaleSettingsClient(subscriptionID)
}
func NewAutoscaleSettingsClientWithBaseURI(baseURI string, subscriptionID string) AutoscaleSettingsClient {
	return original.NewAutoscaleSettingsClientWithBaseURI(baseURI, subscriptionID)
}
func NewDiagnosticSettingsClient(subscriptionID string) DiagnosticSettingsClient {
	return original.NewDiagnosticSettingsClient(subscriptionID)
}
func NewDiagnosticSettingsClientWithBaseURI(baseURI string, subscriptionID string) DiagnosticSettingsClient {
	return original.NewDiagnosticSettingsClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
func NewActionGroupsClient(subscriptionID string) ActionGroupsClient {
	return original.NewActionGroupsClient(subscriptionID)
}
func NewActionGroupsClientWithBaseURI(baseURI string, subscriptionID string) ActionGroupsClient {
	return original.NewActionGroupsClientWithBaseURI(baseURI, subscriptionID)
}
func NewAlertRuleIncidentsClient(subscriptionID string) AlertRuleIncidentsClient {
	return original.NewAlertRuleIncidentsClient(subscriptionID)
}
func NewAlertRuleIncidentsClientWithBaseURI(baseURI string, subscriptionID string) AlertRuleIncidentsClient {
	return original.NewAlertRuleIncidentsClientWithBaseURI(baseURI, subscriptionID)
}
func NewAlertRulesClient(subscriptionID string) AlertRulesClient {
	return original.NewAlertRulesClient(subscriptionID)
}
func NewAlertRulesClientWithBaseURI(baseURI string, subscriptionID string) AlertRulesClient {
	return original.NewAlertRulesClientWithBaseURI(baseURI, subscriptionID)
}
