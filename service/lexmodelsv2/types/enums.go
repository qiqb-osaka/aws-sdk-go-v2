// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type BotAliasStatus string

// Enum values for BotAliasStatus
const (
	BotAliasStatusCreating  BotAliasStatus = "Creating"
	BotAliasStatusAvailable BotAliasStatus = "Available"
	BotAliasStatusDeleting  BotAliasStatus = "Deleting"
	BotAliasStatusFailed    BotAliasStatus = "Failed"
)

// Values returns all known values for BotAliasStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (BotAliasStatus) Values() []BotAliasStatus {
	return []BotAliasStatus{
		"Creating",
		"Available",
		"Deleting",
		"Failed",
	}
}

type BotFilterName string

// Enum values for BotFilterName
const (
	BotFilterNameBotName BotFilterName = "BotName"
)

// Values returns all known values for BotFilterName. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (BotFilterName) Values() []BotFilterName {
	return []BotFilterName{
		"BotName",
	}
}

type BotFilterOperator string

// Enum values for BotFilterOperator
const (
	BotFilterOperatorContains BotFilterOperator = "CO"
	BotFilterOperatorEquals   BotFilterOperator = "EQ"
)

// Values returns all known values for BotFilterOperator. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (BotFilterOperator) Values() []BotFilterOperator {
	return []BotFilterOperator{
		"CO",
		"EQ",
	}
}

type BotLocaleFilterName string

// Enum values for BotLocaleFilterName
const (
	BotLocaleFilterNameBotLocaleName BotLocaleFilterName = "BotLocaleName"
)

// Values returns all known values for BotLocaleFilterName. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (BotLocaleFilterName) Values() []BotLocaleFilterName {
	return []BotLocaleFilterName{
		"BotLocaleName",
	}
}

type BotLocaleFilterOperator string

// Enum values for BotLocaleFilterOperator
const (
	BotLocaleFilterOperatorContains BotLocaleFilterOperator = "CO"
	BotLocaleFilterOperatorEquals   BotLocaleFilterOperator = "EQ"
)

// Values returns all known values for BotLocaleFilterOperator. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (BotLocaleFilterOperator) Values() []BotLocaleFilterOperator {
	return []BotLocaleFilterOperator{
		"CO",
		"EQ",
	}
}

type BotLocaleSortAttribute string

// Enum values for BotLocaleSortAttribute
const (
	BotLocaleSortAttributeBotLocaleName BotLocaleSortAttribute = "BotLocaleName"
)

// Values returns all known values for BotLocaleSortAttribute. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (BotLocaleSortAttribute) Values() []BotLocaleSortAttribute {
	return []BotLocaleSortAttribute{
		"BotLocaleName",
	}
}

type BotLocaleStatus string

// Enum values for BotLocaleStatus
const (
	BotLocaleStatusCreating            BotLocaleStatus = "Creating"
	BotLocaleStatusBuilding            BotLocaleStatus = "Building"
	BotLocaleStatusBuilt               BotLocaleStatus = "Built"
	BotLocaleStatusReadyExpressTesting BotLocaleStatus = "ReadyExpressTesting"
	BotLocaleStatusFailed              BotLocaleStatus = "Failed"
	BotLocaleStatusDeleting            BotLocaleStatus = "Deleting"
	BotLocaleStatusNotBuilt            BotLocaleStatus = "NotBuilt"
)

// Values returns all known values for BotLocaleStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (BotLocaleStatus) Values() []BotLocaleStatus {
	return []BotLocaleStatus{
		"Creating",
		"Building",
		"Built",
		"ReadyExpressTesting",
		"Failed",
		"Deleting",
		"NotBuilt",
	}
}

type BotSortAttribute string

// Enum values for BotSortAttribute
const (
	BotSortAttributeBotName BotSortAttribute = "BotName"
)

// Values returns all known values for BotSortAttribute. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (BotSortAttribute) Values() []BotSortAttribute {
	return []BotSortAttribute{
		"BotName",
	}
}

type BotStatus string

// Enum values for BotStatus
const (
	BotStatusCreating   BotStatus = "Creating"
	BotStatusAvailable  BotStatus = "Available"
	BotStatusInactive   BotStatus = "Inactive"
	BotStatusDeleting   BotStatus = "Deleting"
	BotStatusFailed     BotStatus = "Failed"
	BotStatusVersioning BotStatus = "Versioning"
)

// Values returns all known values for BotStatus. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (BotStatus) Values() []BotStatus {
	return []BotStatus{
		"Creating",
		"Available",
		"Inactive",
		"Deleting",
		"Failed",
		"Versioning",
	}
}

type BotVersionSortAttribute string

// Enum values for BotVersionSortAttribute
const (
	BotVersionSortAttributeBotVersion BotVersionSortAttribute = "BotVersion"
)

// Values returns all known values for BotVersionSortAttribute. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (BotVersionSortAttribute) Values() []BotVersionSortAttribute {
	return []BotVersionSortAttribute{
		"BotVersion",
	}
}

type BuiltInIntentSortAttribute string

// Enum values for BuiltInIntentSortAttribute
const (
	BuiltInIntentSortAttributeIntentSignature BuiltInIntentSortAttribute = "IntentSignature"
)

// Values returns all known values for BuiltInIntentSortAttribute. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (BuiltInIntentSortAttribute) Values() []BuiltInIntentSortAttribute {
	return []BuiltInIntentSortAttribute{
		"IntentSignature",
	}
}

type BuiltInSlotTypeSortAttribute string

// Enum values for BuiltInSlotTypeSortAttribute
const (
	BuiltInSlotTypeSortAttributeSlotTypeSignature BuiltInSlotTypeSortAttribute = "SlotTypeSignature"
)

// Values returns all known values for BuiltInSlotTypeSortAttribute. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (BuiltInSlotTypeSortAttribute) Values() []BuiltInSlotTypeSortAttribute {
	return []BuiltInSlotTypeSortAttribute{
		"SlotTypeSignature",
	}
}

type IntentFilterName string

// Enum values for IntentFilterName
const (
	IntentFilterNameIntentName IntentFilterName = "IntentName"
)

// Values returns all known values for IntentFilterName. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (IntentFilterName) Values() []IntentFilterName {
	return []IntentFilterName{
		"IntentName",
	}
}

type IntentFilterOperator string

// Enum values for IntentFilterOperator
const (
	IntentFilterOperatorContains IntentFilterOperator = "CO"
	IntentFilterOperatorEquals   IntentFilterOperator = "EQ"
)

// Values returns all known values for IntentFilterOperator. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (IntentFilterOperator) Values() []IntentFilterOperator {
	return []IntentFilterOperator{
		"CO",
		"EQ",
	}
}

type IntentSortAttribute string

// Enum values for IntentSortAttribute
const (
	IntentSortAttributeIntentName          IntentSortAttribute = "IntentName"
	IntentSortAttributeLastUpdatedDateTime IntentSortAttribute = "LastUpdatedDateTime"
)

// Values returns all known values for IntentSortAttribute. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (IntentSortAttribute) Values() []IntentSortAttribute {
	return []IntentSortAttribute{
		"IntentName",
		"LastUpdatedDateTime",
	}
}

type ObfuscationSettingType string

// Enum values for ObfuscationSettingType
const (
	ObfuscationSettingTypeNone               ObfuscationSettingType = "None"
	ObfuscationSettingTypeDefaultObfuscation ObfuscationSettingType = "DefaultObfuscation"
)

// Values returns all known values for ObfuscationSettingType. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ObfuscationSettingType) Values() []ObfuscationSettingType {
	return []ObfuscationSettingType{
		"None",
		"DefaultObfuscation",
	}
}

type SlotConstraint string

// Enum values for SlotConstraint
const (
	SlotConstraintRequired SlotConstraint = "Required"
	SlotConstraintOptional SlotConstraint = "Optional"
)

// Values returns all known values for SlotConstraint. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SlotConstraint) Values() []SlotConstraint {
	return []SlotConstraint{
		"Required",
		"Optional",
	}
}

type SlotFilterName string

// Enum values for SlotFilterName
const (
	SlotFilterNameSlotName SlotFilterName = "SlotName"
)

// Values returns all known values for SlotFilterName. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SlotFilterName) Values() []SlotFilterName {
	return []SlotFilterName{
		"SlotName",
	}
}

type SlotFilterOperator string

// Enum values for SlotFilterOperator
const (
	SlotFilterOperatorContains SlotFilterOperator = "CO"
	SlotFilterOperatorEquals   SlotFilterOperator = "EQ"
)

// Values returns all known values for SlotFilterOperator. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SlotFilterOperator) Values() []SlotFilterOperator {
	return []SlotFilterOperator{
		"CO",
		"EQ",
	}
}

type SlotSortAttribute string

// Enum values for SlotSortAttribute
const (
	SlotSortAttributeSlotName            SlotSortAttribute = "SlotName"
	SlotSortAttributeLastUpdatedDateTime SlotSortAttribute = "LastUpdatedDateTime"
)

// Values returns all known values for SlotSortAttribute. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SlotSortAttribute) Values() []SlotSortAttribute {
	return []SlotSortAttribute{
		"SlotName",
		"LastUpdatedDateTime",
	}
}

type SlotTypeFilterName string

// Enum values for SlotTypeFilterName
const (
	SlotTypeFilterNameSlotTypeName SlotTypeFilterName = "SlotTypeName"
)

// Values returns all known values for SlotTypeFilterName. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SlotTypeFilterName) Values() []SlotTypeFilterName {
	return []SlotTypeFilterName{
		"SlotTypeName",
	}
}

type SlotTypeFilterOperator string

// Enum values for SlotTypeFilterOperator
const (
	SlotTypeFilterOperatorContains SlotTypeFilterOperator = "CO"
	SlotTypeFilterOperatorEquals   SlotTypeFilterOperator = "EQ"
)

// Values returns all known values for SlotTypeFilterOperator. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SlotTypeFilterOperator) Values() []SlotTypeFilterOperator {
	return []SlotTypeFilterOperator{
		"CO",
		"EQ",
	}
}

type SlotTypeSortAttribute string

// Enum values for SlotTypeSortAttribute
const (
	SlotTypeSortAttributeSlotTypeName        SlotTypeSortAttribute = "SlotTypeName"
	SlotTypeSortAttributeLastUpdatedDateTime SlotTypeSortAttribute = "LastUpdatedDateTime"
)

// Values returns all known values for SlotTypeSortAttribute. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SlotTypeSortAttribute) Values() []SlotTypeSortAttribute {
	return []SlotTypeSortAttribute{
		"SlotTypeName",
		"LastUpdatedDateTime",
	}
}

type SlotValueResolutionStrategy string

// Enum values for SlotValueResolutionStrategy
const (
	SlotValueResolutionStrategyOriginalValue SlotValueResolutionStrategy = "OriginalValue"
	SlotValueResolutionStrategyTopResolution SlotValueResolutionStrategy = "TopResolution"
)

// Values returns all known values for SlotValueResolutionStrategy. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (SlotValueResolutionStrategy) Values() []SlotValueResolutionStrategy {
	return []SlotValueResolutionStrategy{
		"OriginalValue",
		"TopResolution",
	}
}

type SortOrder string

// Enum values for SortOrder
const (
	SortOrderAscending  SortOrder = "Ascending"
	SortOrderDescending SortOrder = "Descending"
)

// Values returns all known values for SortOrder. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (SortOrder) Values() []SortOrder {
	return []SortOrder{
		"Ascending",
		"Descending",
	}
}
