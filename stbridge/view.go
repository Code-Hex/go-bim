package stbridge

// View returns a small typed wrapper around the document.
//
// The raw node API is still available, and it stays useful when you already
// know the XML names. View exists for the opposite case: you want gopls to
// suggest the obvious section names and common attributes while you write.
func (d *Document) View() DocumentView {
	return DocumentView{doc: d}
}

// SectionNames returns the top-level element names present under ST_BRIDGE.
func (d *Document) SectionNames() []SectionName {
	if d == nil || d.Root == nil {
		return nil
	}

	names := make([]SectionName, 0, len(d.Root.Children))
	for _, child := range d.Root.ElementChildren() {
		names = append(names, child.Name.Local)
	}
	return names
}

// DocumentView is a completion-friendly view over a parsed document.
type DocumentView struct {
	doc *Document
}

// Raw returns the underlying parsed document.
func (v DocumentView) Raw() *Document {
	return v.doc
}

// Common wraps the StbCommon section.
func (v DocumentView) Common() CommonSection {
	return CommonSection{node: sectionNode(v.doc, Section_StbCommon)}
}

// Model wraps the StbModel section.
func (v DocumentView) Model() ModelSection {
	return ModelSection{node: sectionNode(v.doc, Section_StbModel)}
}

// Extensions wraps the optional StbExtensions section.
func (v DocumentView) Extensions() RawSection {
	return RawSection{node: sectionNode(v.doc, Section_StbExtensions)}
}

// ExportInformation wraps the optional StbExportInformation section.
//
// This section was added in ST-Bridge 2.1.0.
func (v DocumentView) ExportInformation() ExportInformationSection {
	return ExportInformationSection{node: sectionNode(v.doc, Section_StbExportInformation)}
}

// CalData wraps the optional StbCalData section.
//
// This section was added in ST-Bridge 2.0.2.
func (v DocumentView) CalData() CalDataSection {
	return CalDataSection{node: sectionNode(v.doc, Section_StbCalData)}
}

// AnaModels wraps the optional StbAnaModels section.
//
// This section was added in ST-Bridge 2.0.2.
func (v DocumentView) AnaModels() AnaModelsSection {
	return AnaModelsSection{node: sectionNode(v.doc, Section_StbAnaModels)}
}

// RawSection is a thin wrapper for sections that do not need extra sugar yet.
type RawSection struct {
	node *Node
}

// Present reports whether the section exists in the parsed document.
func (s RawSection) Present() bool {
	return s.node != nil
}

// Raw returns the underlying XML node.
func (s RawSection) Raw() *Node {
	return s.node
}

// CommonSection wraps StbCommon.
//
// These getters cover the attributes and child nodes people usually reach for
// first. The raw node stays available for everything else.
type CommonSection struct {
	node *Node
}

// Present reports whether StbCommon exists.
func (s CommonSection) Present() bool {
	return s.node != nil
}

// Raw returns the underlying StbCommon node.
func (s CommonSection) Raw() *Node {
	return s.node
}

// GUID returns the guid attribute.
func (s CommonSection) GUID() (string, bool) {
	return stringAttr(s.node, Attr_guid)
}

// ProjectName returns the project_name attribute.
func (s CommonSection) ProjectName() (string, bool) {
	return stringAttr(s.node, Attr_project_name)
}

// AppName returns the app_name attribute.
func (s CommonSection) AppName() (string, bool) {
	return stringAttr(s.node, Attr_app_name)
}

// AppVersion returns the app_version attribute when it exists.
func (s CommonSection) AppVersion() (string, bool) {
	return stringAttr(s.node, Attr_app_version)
}

// ConvertAppName returns the convert_app_name attribute when it exists.
func (s CommonSection) ConvertAppName() (string, bool) {
	return stringAttr(s.node, Attr_convert_app_name)
}

// ConvertAppVersion returns the convert_app_version attribute when it exists.
func (s CommonSection) ConvertAppVersion() (string, bool) {
	return stringAttr(s.node, Attr_convert_app_version)
}

// StrengthConcrete returns the strength_concrete attribute when it exists.
func (s CommonSection) StrengthConcrete() (string, bool) {
	return stringAttr(s.node, Attr_strength_concrete)
}

// GlobalOffsetX returns the global_offset_X attribute when it exists.
func (s CommonSection) GlobalOffsetX() (float64, bool, error) {
	return float64Attr(s.node, Attr_global_offset_X)
}

// GlobalOffsetY returns the global_offset_Y attribute when it exists.
func (s CommonSection) GlobalOffsetY() (float64, bool, error) {
	return float64Attr(s.node, Attr_global_offset_Y)
}

// GlobalOffsetZ returns the global_offset_Z attribute when it exists.
func (s CommonSection) GlobalOffsetZ() (float64, bool, error) {
	return float64Attr(s.node, Attr_global_offset_Z)
}

// GlobalRotation returns the global_rotation attribute when it exists.
func (s CommonSection) GlobalRotation() (float64, bool, error) {
	return float64Attr(s.node, Attr_global_rotation)
}

// ReinforcementStrengthList returns the StbReinforcementStrengthList child.
func (s CommonSection) ReinforcementStrengthList() *Node {
	return childNode(s.node, Element_StbReinforcementStrengthList)
}

// ReinforcementStrengthListPile returns the StbReinforcementStrengthListPile child.
//
// This child was added in ST-Bridge 2.1.0.
func (s CommonSection) ReinforcementStrengthListPile() *Node {
	return childNode(s.node, Element_StbReinforcementStrengthListPile)
}

// ApplyConditionsList returns the StbApplyConditionsList child.
func (s CommonSection) ApplyConditionsList() *Node {
	return childNode(s.node, Element_StbApplyConditionsList)
}

// StandardPlateThicknessList returns the StbStandardPlateThicknessList child.
//
// This child was added in ST-Bridge 2.1.0.
func (s CommonSection) StandardPlateThicknessList() *Node {
	return childNode(s.node, Element_StbStandardPlateThicknessList)
}

// ConnectionSpecs returns the StbConnectionSpecs child.
//
// This child was added in ST-Bridge 2.1.0.
func (s CommonSection) ConnectionSpecs() *Node {
	return childNode(s.node, Element_StbConnectionSpecs)
}

// WeldCommon returns the StbWeldCommon child.
//
// This child was added in ST-Bridge 2.1.0.
func (s CommonSection) WeldCommon() *Node {
	return childNode(s.node, Element_StbWeldCommon)
}

// AdditionalInformation returns the StbAdditionalInformation child.
//
// This child was added in ST-Bridge 2.1.0.
func (s CommonSection) AdditionalInformation() *Node {
	return childNode(s.node, Element_StbAdditionalInformation)
}

// ModelSection wraps StbModel.
type ModelSection struct {
	node *Node
}

// Present reports whether StbModel exists.
func (s ModelSection) Present() bool {
	return s.node != nil
}

// Raw returns the underlying StbModel node.
func (s ModelSection) Raw() *Node {
	return s.node
}

// Nodes returns the StbNodes child.
func (s ModelSection) Nodes() *Node {
	return childNode(s.node, Element_StbNodes)
}

// Axes returns the StbAxes child.
func (s ModelSection) Axes() *Node {
	return childNode(s.node, Element_StbAxes)
}

// Stories returns the StbStories child.
func (s ModelSection) Stories() *Node {
	return childNode(s.node, Element_StbStories)
}

// Members returns the StbMembers child.
func (s ModelSection) Members() *Node {
	return childNode(s.node, Element_StbMembers)
}

// Sections returns the StbSections child.
func (s ModelSection) Sections() *Node {
	return childNode(s.node, Element_StbSections)
}

// Joints returns the StbJoints child.
func (s ModelSection) Joints() *Node {
	return childNode(s.node, Element_StbJoints)
}

// Connections returns the StbConnections child.
//
// This child was added in ST-Bridge 2.1.0.
func (s ModelSection) Connections() *Node {
	return childNode(s.node, Element_StbConnections)
}

// Weld returns the StbWeld child.
//
// This child was added in ST-Bridge 2.1.0.
func (s ModelSection) Weld() *Node {
	return childNode(s.node, Element_StbWeld)
}

// ExportInformationSection wraps StbExportInformation.
//
// StbExportInformation was added in ST-Bridge 2.1.0. Its child collections are
// repeatable in the schema, so this wrapper returns slices rather than a single
// first match.
type ExportInformationSection struct {
	node *Node
}

// Present reports whether StbExportInformation exists.
func (s ExportInformationSection) Present() bool {
	return s.node != nil
}

// Raw returns the underlying StbExportInformation node.
func (s ExportInformationSection) Raw() *Node {
	return s.node
}

// ExportPolicies returns every StbExportPolicy child in order.
func (s ExportInformationSection) ExportPolicies() []*Node {
	return childrenByName(s.node, Element_StbExportPolicy)
}

// ExportLogs returns every StbExportLog child in order.
func (s ExportInformationSection) ExportLogs() []*Node {
	return childrenByName(s.node, Element_StbExportLog)
}

// CalDataSection wraps StbCalData.
type CalDataSection struct {
	node *Node
}

// Present reports whether StbCalData exists.
func (s CalDataSection) Present() bool {
	return s.node != nil
}

// Raw returns the underlying StbCalData node.
func (s CalDataSection) Raw() *Node {
	return s.node
}

// CalCommon returns the StbCalCommon child.
func (s CalDataSection) CalCommon() *Node {
	return childNode(s.node, Element_StbCalCommon)
}

// CalLoad returns the StbCalLoad child.
func (s CalDataSection) CalLoad() *Node {
	return childNode(s.node, Element_StbCalLoad)
}

// CalCondition returns the StbCalCondition child.
func (s CalDataSection) CalCondition() *Node {
	return childNode(s.node, Element_StbCalCondition)
}

// CalLoadArrangements returns the StbCalLoadArrangements child.
func (s CalDataSection) CalLoadArrangements() *Node {
	return childNode(s.node, Element_StbCalLoadArrangements)
}

// CalConditionArrangements returns the StbCalConditionArrangements child.
func (s CalDataSection) CalConditionArrangements() *Node {
	return childNode(s.node, Element_StbCalConditionArrangements)
}

// AnaModelsSection wraps StbAnaModels.
type AnaModelsSection struct {
	node *Node
}

// Present reports whether StbAnaModels exists.
func (s AnaModelsSection) Present() bool {
	return s.node != nil
}

// Raw returns the underlying StbAnaModels node.
func (s AnaModelsSection) Raw() *Node {
	return s.node
}

// Models returns every StbAnaModel child in order.
func (s AnaModelsSection) Models() []*Node {
	if s.node == nil {
		return nil
	}
	return s.node.ChildrenByName(Element_StbAnaModel)
}

func sectionNode(doc *Document, name SectionName) *Node {
	if doc == nil {
		return nil
	}
	return doc.Section(name)
}

func childNode(node *Node, name ElementName) *Node {
	if node == nil {
		return nil
	}
	return node.Child(name)
}

func childrenByName(node *Node, name ElementName) []*Node {
	if node == nil {
		return nil
	}
	return node.ChildrenByName(name)
}

func stringAttr(node *Node, name AttributeName) (string, bool) {
	if node == nil {
		return "", false
	}
	return node.Attr(name)
}
