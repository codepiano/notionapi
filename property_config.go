package notionapi

import (
	"encoding/json"
	"fmt"
	"github.com/iancoleman/orderedmap"
)

type PropertyConfigType string

type PropertyConfig interface {
	GetType() PropertyConfigType
}

type TitlePropertyConfig struct {
	ID    PropertyID         `json:"id,omitempty"`
	Type  PropertyConfigType `json:"type"`
	Title struct{}           `json:"title"`
}

func (p TitlePropertyConfig) GetType() PropertyConfigType {
	return p.Type
}

type RichTextPropertyConfig struct {
	ID       PropertyID         `json:"id,omitempty"`
	Type     PropertyConfigType `json:"type"`
	RichText struct{}           `json:"rich_text"`
}

func (p RichTextPropertyConfig) GetType() PropertyConfigType {
	return p.Type
}

type NumberPropertyConfig struct {
	ID     ObjectID           `json:"id,omitempty"`
	Type   PropertyConfigType `json:"type"`
	Number NumberFormat       `json:"number"`
}

type FormatType string

func (ft FormatType) String() string {
	return string(ft)
}

type NumberFormat struct {
	Format FormatType `json:"format"`
}

func (p NumberPropertyConfig) GetType() PropertyConfigType {
	return p.Type
}

type SelectPropertyConfig struct {
	ID     ObjectID           `json:"id,omitempty"`
	Type   PropertyConfigType `json:"type"`
	Select Select             `json:"select"`
}

func (p SelectPropertyConfig) GetType() PropertyConfigType {
	return p.Type
}

type MultiSelectPropertyConfig struct {
	ID          ObjectID           `json:"id,omitempty"`
	Type        PropertyConfigType `json:"type"`
	MultiSelect Select             `json:"multi_select"`
}

type Select struct {
	Options []Option `json:"options"`
}

func (p MultiSelectPropertyConfig) GetType() PropertyConfigType {
	return p.Type
}

type DatePropertyConfig struct {
	ID   ObjectID           `json:"id,omitempty"`
	Type PropertyConfigType `json:"type"`
	Date struct{}           `json:"date"`
}

func (p DatePropertyConfig) GetType() PropertyConfigType {
	return p.Type
}

type PeoplePropertyConfig struct {
	ID     ObjectID           `json:"id,omitempty"`
	Type   PropertyConfigType `json:"type"`
	People struct{}           `json:"people"`
}

func (p PeoplePropertyConfig) GetType() PropertyConfigType {
	return p.Type
}

type FilesPropertyConfig struct {
	ID    ObjectID           `json:"id,omitempty"`
	Type  PropertyConfigType `json:"type"`
	Files struct{}           `json:"files"`
}

func (p FilesPropertyConfig) GetType() PropertyConfigType {
	return p.Type
}

type CheckboxPropertyConfig struct {
	ID       ObjectID           `json:"id,omitempty"`
	Type     PropertyConfigType `json:"type"`
	Checkbox struct{}           `json:"checkbox"`
}

func (p CheckboxPropertyConfig) GetType() PropertyConfigType {
	return p.Type
}

type URLPropertyConfig struct {
	ID   ObjectID           `json:"id,omitempty"`
	Type PropertyConfigType `json:"type"`
	URL  struct{}           `json:"url"`
}

func (p URLPropertyConfig) GetType() PropertyConfigType {
	return p.Type
}

type EmailPropertyConfig struct {
	ID    PropertyID         `json:"id,omitempty"`
	Type  PropertyConfigType `json:"type"`
	Email struct{}           `json:"email"`
}

func (p EmailPropertyConfig) GetType() PropertyConfigType {
	return p.Type
}

type PhoneNumberPropertyConfig struct {
	ID          ObjectID           `json:"id,omitempty"`
	Type        PropertyConfigType `json:"type"`
	PhoneNumber struct{}           `json:"phone_number"`
}

func (p PhoneNumberPropertyConfig) GetType() PropertyConfigType {
	return p.Type
}

type FormulaPropertyConfig struct {
	ID      ObjectID           `json:"id,omitempty"`
	Type    PropertyConfigType `json:"type"`
	Formula FormulaConfig      `json:"formula"`
}

type FormulaConfig struct {
	Expression string `json:"expression"`
}

func (p FormulaPropertyConfig) GetType() PropertyConfigType {
	return p.Type
}

type RelationPropertyConfig struct {
	Type     PropertyConfigType `json:"type"`
	Relation RelationConfig     `json:"relation"`
}

type RelationConfigType string

func (rp RelationConfigType) String() string {
	return string(rp)
}

type SingleProperty struct{}

type DualProperty struct{}

type RelationConfig struct {
	DatabaseID         DatabaseID         `json:"database_id"`
	SyncedPropertyID   PropertyID         `json:"synced_property_id,omitempty"`
	SyncedPropertyName string             `json:"synced_property_name,omitempty"`
	Type               RelationConfigType `json:"type,omitempty"`
	SingleProperty     *SingleProperty    `json:"single_property,omitempty"`
	DualProperty       *DualProperty      `json:"dual_property,omitempty"`
}

func (p RelationPropertyConfig) GetType() PropertyConfigType {
	return p.Type
}

type RollupPropertyConfig struct {
	ID     ObjectID           `json:"id,omitempty"`
	Type   PropertyConfigType `json:"type"`
	Rollup RollupConfig       `json:"rollup"`
}

type RollupConfig struct {
	RelationPropertyName string       `json:"relation_property_name"`
	RelationPropertyID   PropertyID   `json:"relation_property_id"`
	RollupPropertyName   string       `json:"rollup_property_name"`
	RollupPropertyID     PropertyID   `json:"rollup_property_id"`
	Function             FunctionType `json:"function"`
}

func (p RollupPropertyConfig) GetType() PropertyConfigType {
	return p.Type
}

type CreatedTimePropertyConfig struct {
	ID          ObjectID           `json:"id,omitempty"`
	Type        PropertyConfigType `json:"type"`
	CreatedTime struct{}           `json:"created_time"`
}

func (p CreatedTimePropertyConfig) GetType() PropertyConfigType {
	return p.Type
}

type CreatedByPropertyConfig struct {
	ID        ObjectID           `json:"id"`
	Type      PropertyConfigType `json:"type"`
	CreatedBy struct{}           `json:"created_by"`
}

func (p CreatedByPropertyConfig) GetType() PropertyConfigType {
	return p.Type
}

type LastEditedTimePropertyConfig struct {
	ID             ObjectID           `json:"id"`
	Type           PropertyConfigType `json:"type"`
	LastEditedTime struct{}           `json:"last_edited_time"`
}

func (p LastEditedTimePropertyConfig) GetType() PropertyConfigType {
	return p.Type
}

type LastEditedByPropertyConfig struct {
	ID           ObjectID           `json:"id"`
	Type         PropertyConfigType `json:"type"`
	LastEditedBy struct{}           `json:"last_edited_by"`
}

func (p LastEditedByPropertyConfig) GetType() PropertyConfigType {
	return p.Type
}

type StatusPropertyConfig struct {
	ID     ObjectID           `json:"id"`
	Type   PropertyConfigType `json:"type"`
	Status StatusConfig       `json:"status"`
}

func (p StatusPropertyConfig) GetType() PropertyConfigType {
	return p.Type
}

type StatusConfig struct {
	Options []Option      `json:"options"`
	Groups  []GroupConfig `json:"groups"`
}

type GroupConfig struct {
	ID        ObjectID   `json:"id"`
	Name      string     `json:"name"`
	Color     string     `json:"color"`
	OptionIDs []ObjectID `json:"option_ids"`
}

// TODO: API docs don't have this listed yet, not sure of correct structure
type UniqueIDPropertyConfig struct{}

func (i UniqueIDPropertyConfig) GetType() PropertyConfigType {
	return ""
}

type PropertyConfigs orderedmap.OrderedMap

func NewPropertyConfigs() PropertyConfigs {
	return PropertyConfigs(*orderedmap.New())
}

func (p *PropertyConfigs) UnmarshalJSON(data []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	props, err := parsePropertyConfigs(raw)
	if err != nil {
		return err
	}

	*p = props
	return nil
}

func parsePropertyConfigs(raw map[string]interface{}) (PropertyConfigs, error) {
	om := orderedmap.New()
	var result PropertyConfigs
	for k, v := range raw {
		var p PropertyConfig
		switch rawProperty := v.(type) {
		case map[string]interface{}:
			switch PropertyConfigType(rawProperty["type"].(string)) {
			case PropertyConfigTypeTitle:
				p = &TitlePropertyConfig{}
			case PropertyConfigTypeRichText:
				p = &RichTextPropertyConfig{}
			case PropertyConfigTypeNumber:
				p = &NumberPropertyConfig{}
			case PropertyConfigTypeSelect:
				p = &SelectPropertyConfig{}
			case PropertyConfigTypeMultiSelect:
				p = &MultiSelectPropertyConfig{}
			case PropertyConfigTypeDate:
				p = &DatePropertyConfig{}
			case PropertyConfigTypePeople:
				p = &PeoplePropertyConfig{}
			case PropertyConfigTypeFiles:
				p = &FilesPropertyConfig{}
			case PropertyConfigTypeCheckbox:
				p = &CheckboxPropertyConfig{}
			case PropertyConfigTypeURL:
				p = &URLPropertyConfig{}
			case PropertyConfigTypeEmail:
				p = &EmailPropertyConfig{}
			case PropertyConfigTypePhoneNumber:
				p = &PhoneNumberPropertyConfig{}
			case PropertyConfigTypeFormula:
				p = &FormulaPropertyConfig{}
			case PropertyConfigTypeRelation:
				p = &RelationPropertyConfig{}
			case PropertyConfigTypeRollup:
				p = &RollupPropertyConfig{}
			case PropertyConfigCreatedTime:
				p = &CreatedTimePropertyConfig{}
			case PropertyConfigCreatedBy:
				p = &CreatedTimePropertyConfig{}
			case PropertyConfigLastEditedTime:
				p = &LastEditedTimePropertyConfig{}
			case PropertyConfigLastEditedBy:
				p = &LastEditedByPropertyConfig{}
			case PropertyConfigStatus:
				p = &StatusPropertyConfig{}
			case PropertyConfigUniqueID:
				p = &UniqueIDPropertyConfig{}
			default:

				return result, fmt.Errorf("unsupported property type: %s", rawProperty["type"].(string))
			}
			b, err := json.Marshal(rawProperty)
			if err != nil {
				return result, err
			}

			if err = json.Unmarshal(b, &p); err != nil {
				return result, err
			}
			om.Set(k, p)
		default:
			return result, fmt.Errorf("unsupported property format %T", v)
		}
	}
	result = PropertyConfigs(*om)
	return result, nil
}
